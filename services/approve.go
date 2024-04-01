package services

import (
	"context"
	"errors"
	"fmt"
	"treasury/mock"
	"treasury/models"
)

const (
	signerPk      = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	signerAddrStr = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	//rpcURL          = "http://127.0.0.1:8545"
	rpcURL          = "http://host.docker.internal:8545"
	contractAddrStr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
)

func Approve(ctx context.Context, req *models.ApproveRequest) error {
	manager, ok := mock.ManagerData[req.ApproverID]
	if !ok {
		return errors.New("invalid approver")
	}

	claimInfo := &models.WithdrawClaim{ID: req.ClaimID}
	err := claimInfo.GetByID()
	if err != nil {
		return err
	}
	if claimInfo.Status != models.StatusPending {
		return errors.New("claim status invalid")
	}

	confirm := &models.WithdrawApprove{ClaimID: req.ClaimID, ApproverID: manager.ID}
	err = confirm.FirstOrCreate()
	if err != nil {
		return err
	}

	approvalCounts, err := models.CountApprovals(req.ClaimID)
	if approvalCounts < 2 {
		return nil
	}

	user, ok := mock.UserData[claimInfo.UserID]
	if !ok {
		return errors.New("invalid user")
	}

	signedTx, err := signTx(ctx, user.Address, claimInfo.Value)
	if err != nil {
		fmt.Println("signTx err:", err)
		return err
	}

	return models.UpdateToApproved(claimInfo.ID, signedTx)

}
