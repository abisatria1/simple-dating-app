package model

type GetFeedsForUserParam struct {
	PaginationQueryParams
}

type DislikeRequest struct {
	TargetUserID int64 `json:"target_user_id" valid:"required"`
}
