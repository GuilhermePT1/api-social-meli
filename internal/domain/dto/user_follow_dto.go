package dto

type FollowRequestDTO struct {
	UserID     uint `json:"user_id" binding:"required"`
	FollowerID uint `json:"follower_id" binding:"required"`
}

type FollowResponseDTO struct {
	Message string `json:"message"`
}
