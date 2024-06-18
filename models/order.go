package models

import "time"

type Order struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"author"`
	Author      Author
	ReleaseDate time.Time `json:"release_date" gorm:"column:release_date"`
	ISBN13      uint      `json:"isbn_13" gorm:"column:isbn_13"`
	FormatID    uint      `json:"format"`
	Format      Format
}

type CreateOrderInput struct {
	ID          uint      `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Author      uint      `json:"author" binding:"required"`
	Format      uint      `json:"format" binding:"required"`
	ReleaseDate time.Time `json:"release_date"`
	ISBN13      uint      `json:"isbn_13" binding:"required"`
}

type UpdateOrderInput struct {
	Title string `json:"title"`
}
