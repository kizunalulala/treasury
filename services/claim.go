package services

import (
	"context"
	"errors"
	"treasury/mock"
	"treasury/models"
)

func Claim(c context.Context, req *models.ClaimRequest) error {
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
