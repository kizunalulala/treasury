package models

type ClaimRequest struct {
	UserID uint64 `json:"userID" form:"userID"`
	Value  int64  `json:"value" form:"value"`
}

type ApproveRequest struct {
	ApproverID uint64 `json:"approverID" form:"approverID"`
	ClaimID    uint64 `json:"claimID" form:"claimID"`
}
