package requestParams

//添加博客
type BlogAddParams struct {
	Name    string `json:"name" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

//更新博客
type BlogUpdateParams struct {
	Name    string `json:"name" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
