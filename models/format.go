package models

type Format struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Format string `json:"format"`
}

type CreateFormatInput struct {
	ID     uint   `json:"id" binding:"required"`
	Format string `json:"format" binding:"required"`
}

type UpdateFormatInput struct {
	Format string `json:"format"`
}
