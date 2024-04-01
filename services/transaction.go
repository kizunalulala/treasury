package services

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strings"
	"treasury/services/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sendTransaction(ctx context.Context, client *ethclient.Client, signedTx, toMintAddress string, mintValue int64) (string, error) {
	if signedTx == "" {
		var err error
		signedTx, err = signTx(ctx, toMintAddress, mintValue)
		if err != nil {
			return "", err
		}
	}
	if strings.HasPrefix(signedTx, "0x") {
		signedTx = signedTx[2:]
	}
	signedTxBytes, err := hex.DecodeString(signedTx)
	if err != nil {
		return "", err
	}

	var tx types.Transaction
	if err = rlp.DecodeBytes(signedTxBytes, &tx); err != nil {
		return "", err
	}

	err = client.SendTransaction(ctx, &tx)
	if err != nil {
		if err.Error() == "nonce too low" {
			return tx.Hash().String(), nil
		}
		return "", err

	}
	return tx.Hash().String(), nil
}

func signTx(ctx context.Context, toMintAddrStr string, claimValue int64) (string, error) {
	client, err := getEthClient(rpcURL)
	if err != nil {
		return "", err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(signerPk)
	if err != nil {
		return "", err
	}

	signerAddress := common.HexToAddress(signerAddrStr)
	contractAddress := common.HexToAddress(contractAddrStr)

	nonce, err := client.PendingNonceAt(ctx, signerAddress)
	if err != nil {
		return "", err
	}

	myTokenABI, _ := contracts.MyTokenMetaData.GetAbi()
	toMintAddress := common.HexToAddress(toMintAddrStr)
	data, err := myTokenABI.Pack("mint", toMintAddress, big.NewInt(claimValue))
	if err != nil {
		return "", errors.New("fail to pack data for mint")
	}

	estimateGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:      signerAddress,
		To:        &contractAddress,
		Gas:       0,
		GasPrice:  nil,
		GasTipCap: big.NewInt(2000000),
		GasFeeCap: big.NewInt(1200000000),
		Value:     big.NewInt(0),
		Data:      data,
	})
	if err != nil {
		return "", errors.New("fail to pack data for mint")
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: big.NewInt(2000000),
		GasFeeCap: big.NewInt(1200000000),
		Gas:       estimateGas,
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return "", errors.New("get signedTx failed")
	}
	txBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", errors.New("EncodeToBytes failed")
	}
	return fmt.Sprintf("0x%x", txBytes), nil
}

func getEthClient(endpoint string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	return client, nil
}
