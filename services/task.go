package services

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"treasury/mock"
	"treasury/models"
)

func ConsumeApprovedTasks(ctx context.Context) {
	approvedList, err := models.FilterByStatus(models.StatusApproved)
	if err != nil {
		log.Error().Err(err).Msg("ApprovedList err")
		return
	}
	if len(approvedList) == 0 {
		return
	}

	client, err := getEthClient(rpcURL)
	if err != nil {
		log.Error().Err(err).Msg("getEthClient err")
		return
	}
	for _, data := range approvedList {
		user, ok := mock.UserData[data.UserID]
		if !ok {
			log.Error().Msg("user not found")
			continue
		}
		txHash, err := sendTransaction(ctx, client, data.SignTransaction, user.Address, data.Value)
		if err != nil {
			log.Error().Err(err).Msg("sendTransaction err")
			continue
		}
		err = models.UpdateToWaiting(data.ID, txHash)
		if err != nil {
			log.Error().Err(err).Msg("UpdateToWaiting err")
		}
	}
}

func ConsumeWaitingTasks(ctx context.Context) {
	waitingList, err := models.FilterByStatus(models.StatusWaiting)
	if err != nil {
		log.Error().Err(err).Msg("FilterByStatus err")
		return
	}
	if len(waitingList) == 0 {
		return
	}

	ethClient, err := getEthClient(rpcURL)
	if err != nil {
		log.Error().Err(err).Msg("getEthClient err")
		return
	}
	for _, data := range waitingList {
		receipt, err := ethClient.TransactionReceipt(ctx, common.HexToHash(data.TxHash))
		if err != nil {
			log.Error().Err(err).Msg("ConsumePendingClaims failed, get transactionReceipt err")
			return
		}
		var claimStatus int
		if receipt.Status == 1 {
			claimStatus = models.StatusDone
		} else {
			claimStatus = models.StatusApproved
		}
		err = models.UpdateStatus(data.ID, claimStatus)
		if err != nil {
			log.Error().Err(err).Msg("ConsumePendingClaims failed, update claim status err")
			return
		}
	}
}
