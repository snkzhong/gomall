package forms

type CategoryForm struct {
	ParentID int    `form:"parent_id" json:"parent_id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
}
