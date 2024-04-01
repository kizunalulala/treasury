package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	StatusPending  = 1
	StatusApproved = 2
	StatusWaiting  = 3
	StatusDone     = 4
)

type WithdrawClaim struct {
	ID              uint64         `json:"id" gorm:"primaryKey;column:id;autoIncrement;comment:primary key"`
	UserID          uint64         `json:"userID" gorm:"column:user_id;index:idx_user_id;comment:withdrawal claimer id"`
	Value           int64          `json:"value" gorm:"column:value;comment:withdrawal value, unit: gwei"`
	Status          int            `json:"status" gorm:"column:status;index:idx_status;comment:job status, 1-pending 2-approved 3-WIP 4-waitingForReceipt 5-done"`
	TxHash          string         `json:"txHash" gorm:"column:tx_hash"`
	SignTransaction string         `json:"signTransaction" gorm:"sign_transaction"`
	CreatedAt       time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type WithdrawalClaimUser struct {
	WithdrawClaim
}

func (WithdrawClaim) TableName() string {
	return "withdraw_claim"
}

func (claimInfo *WithdrawClaim) Create() error {
	return self.db.Create(claimInfo).Error
}

func (claimInfo *WithdrawClaim) GetByID() error {
	return self.db.Where("`id` = ?", claimInfo.ID).First(claimInfo).Error
}

func FilterByStatus(status int) ([]*WithdrawClaim, error) {
	list := make([]*WithdrawClaim, 0)
	err := self.db.Model(&WithdrawClaim{}).Where("status = ?", status).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateToApproved(id uint64, signedTx string) error {
	return self.db.Model(&WithdrawClaim{}).Where("`id` = ?", id).Updates(map[string]interface{}{
		"status":           StatusApproved,
		"sign_transaction": signedTx,
	}).Error
}

func UpdateToWaiting(id uint64, txHash string) error {
	return self.db.Model(&WithdrawClaim{}).Where("`id` = ?", id).Updates(map[string]interface{}{
		"status":  StatusWaiting,
		"tx_hash": txHash,
	}).Error
}

func UpdateStatus(id uint64, status int) error {
	return self.db.Model(&WithdrawClaim{}).Where("`id` = ?", id).Updates(map[string]interface{}{
		"status": status,
	}).Error
}
