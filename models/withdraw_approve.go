package models

import (
	"gorm.io/gorm"
	"time"
)

type WithdrawApprove struct {
	ID         uint64         `gorm:"primaryKey;column:id;autoIncrement;comment:primary key" json:"id"`
	ApproverID uint64         `gorm:"column:approver_id;comment:withdrawal approver id" json:"approverId"`
	ClaimID    uint64         `gorm:"column:claim_id;index:idx_claim_id;comment:withdrawal claim id" json:"claimId"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}

func (WithdrawApprove) TableName() string {
	return "withdraw_approve"
}

func (confirm *WithdrawApprove) FirstOrCreate() error {
	result := self.db.FirstOrCreate(&WithdrawApprove{}, &WithdrawApprove{
		ApproverID: confirm.ApproverID,
		ClaimID:    confirm.ClaimID,
	})
	return result.Error
}

func CountApprovals(claimID uint64) (int64, error) {
	var count int64
	err := self.db.Model(&WithdrawApprove{}).Where("claim_id = ?", claimID).Count(&count).Error
	return count, err
}
