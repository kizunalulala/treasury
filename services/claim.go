package services

import (
	"context"
	"errors"
	"treasury/mock"
	"treasury/models"
)

func ClaimInfo(c context.Context, id int64) (*models.WithdrawClaim, error) {
	claimInfo := &models.WithdrawClaim{ID: uint64(id)}
	err := claimInfo.GetByID()
	return claimInfo, err
}

func ClaimCreate(c context.Context, req *models.ClaimRequest) error {
	_, ok := mock.UserData[req.UserID]
	if !ok {
		return errors.New("invalid user")
	}

	withdrawalClaim := &models.WithdrawClaim{
		UserID: req.UserID,
		Value:  req.Value,
		Status: models.StatusPending,
	}

	return withdrawalClaim.Create()
}
