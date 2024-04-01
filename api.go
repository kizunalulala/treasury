package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"treasury/models"
	"treasury/services"
)

func claimHandler(c *gin.Context) {
	req := new(models.ClaimRequest)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.UserID == 0 || req.Value == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request params invalid, should be greater than zero"})
		return
	}
	ctx := c.Request.Context()
	err := services.Claim(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Claim received"})
}

func approveHandler(c *gin.Context) {
	req := new(models.ApproveRequest)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.ApproverID == 0 || req.ClaimID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request params invalid, should be greater than zero"})
		return
	}
	ctx := c.Request.Context()
	err := services.Approve(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Confirm received"})
}
