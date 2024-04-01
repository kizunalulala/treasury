package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"treasury/models"
	"treasury/services"
)

func claimGetHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	res, err := services.ClaimInfo(ctx, idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"claim": res})
}

func claimCreateHandler(c *gin.Context) {
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
	err := services.ClaimCreate(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Claim received"})
}

func approveCreateHandler(c *gin.Context) {
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
	err := services.ApproveCreate(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Confirm received"})
}
