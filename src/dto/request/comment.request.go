package request

type CreateComment struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateComment struct {
	ID      uint   `json:"id"`
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
