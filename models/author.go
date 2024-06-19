package models

type Author struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	FullName string `json:"full_name"`
}

type CreateAuthorInput struct {
	ID       uint   `json:"id" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

type UpdateAuthorInput struct {
	FullName string `json:"full_name"`
}

func NewAuthor(id uint, full_name string) Author {
	author := Author{
		ID:       id,
		FullName: full_name,
	}
	return author
}
