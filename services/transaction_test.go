package services

import (
	"context"
	"fmt"
	"testing"
	"treasury/models"

	"github.com/glycerine/goconvey/convey"
)

func TestTransaction(t *testing.T) {
	models.Init()
	convey.Convey("Test_Transaction", t, func(convCtx convey.C) {
		convCtx.Convey("When everything goes positive", func(convCtx convey.C) {
			WithdrawalTask(context.Background())
		})
	})
}

func TestSignTx(t *testing.T) {
	convey.Convey("Test_SignTx", t, func(convCtx convey.C) {
		models.Init()
		convCtx.Convey("When everything goes positive", func(convCtx convey.C) {
			res, err := signTx(context.Background(), "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", 500000000)
			fmt.Println("res:", res)
			fmt.Println("err:", err)
		})
	})
}

func TestSendTransaction(t *testing.T) {
	convey.Convey("Test_SendTransaction", t, func(convCtx convey.C) {
		convCtx.Convey("When everything goes positive", func(convCtx convey.C) {
			client, _ := getEthClient("http://127.0.0.1:8545")
			res, err := sendTransaction(context.Background(), client, "0xb8b302f8b0827a6904831e84808447868c00828e78945fbdb2315678afecb367f032d93f642f64180aa380b84440c10f1900000000000000000000000070997970c51812dc3a010c7d01b50e0d17dc79c8000000000000000000000000000000000000000000000000000000012a05f200c001a08c21d280199c79ed072165ba963d3981c0c6402fabde8895b6dd5f26286aac9ca0775da5bbaf4e2e6059451b1059ecc7fac891b90a91be6b6724d96a7115edfbde", "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", 5000000000)
			fmt.Println("res:", res)
			fmt.Println("err:", err)
		})
	})
}

func TestConsumePendingClaims(t *testing.T) {
	convey.Convey("Test_ConsumePendingClaims", t, func(convCtx convey.C) {
		models.Init()
		convCtx.Convey("When everything goes positive", func(convCtx convey.C) {
			ConsumeWaitingClaims(context.Background())
		})
	})
}
