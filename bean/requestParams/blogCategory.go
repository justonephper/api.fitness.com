package requestParams

type BlogCategoryAddParams struct {
	Name string `json:"name" binding:"required,max=64"`
	Desc string `json:"desc" binding:"required,max=128"`
}

type BlogCategoryUpdateParams struct {
	Name string `json:"name" binding:"required,max=64"`
	Desc string `json:"desc" binding:"required,max=128"`
}
