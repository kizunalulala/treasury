package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
	"treasury/models"
	"treasury/services"
)

func main() {
	models.Init()

	ctx := context.Background()
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				services.ConsumeApprovedTasks(ctx)
				services.ConsumeWaitingTasks(ctx)
			}
		}
	}()

	// router
	r := gin.Default()
	r.POST("/withdraw/claim/create", claimHandler)
	r.POST("/withdraw/approve/create", approveHandler)
	r.Run()

}
