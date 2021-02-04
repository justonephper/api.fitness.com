package requestParams

type BlogCategoryAddParams struct {
	Name string `json:"name" binding:"required,Maximum=64"`
	Desc string `json:"desc" binding:"required,Maximum=128"`
}

type BlogCategoryUpdateParams struct {
	Name string `json:"name" binding:"required,Maximum=64"`
	Desc string `json:"desc" binding:"required,Maximum=128"`
}
