package request

type CreatePost struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePost struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
