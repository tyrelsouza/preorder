package orders

import (
	"encoding/json"
	"io"
	"net/http"
	"preorder/authors"
	"preorder/formats"
	"strconv"
	"time"
)

const googleBooksAPI = "https://www.googleapis.com/books/v1/volumes?q=isbn:"

type Order struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"author"`
	Author      authors.Author
	ReleaseDate time.Time `json:"release_date" gorm:"column:release_date"`
	ISBN13      uint      `json:"isbn_13" gorm:"column:isbn_13"`
	FormatID    uint      `json:"format"`
	Format      formats.Format
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

func NewOrder(id uint, title string, author uint, format uint, isbn13 uint, releaseDate time.Time) Order {
	order := Order{
		ID:          id,
		Title:       title,
		AuthorID:    author,
		FormatID:    format,
		ISBN13:      isbn13,
		ReleaseDate: releaseDate,
	}
	return order
}

func (o *Order) fetchCoverImageURL() (string, error) {
	resp, err := http.Get(googleBooksAPI + strconv.Itoa(int(o.ISBN13)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var googleBooksResponse GoogleBooksResponse
	if err := json.Unmarshal(body, &googleBooksResponse); err != nil {
		return "", err
	}

	if len(googleBooksResponse.Items) > 0 {
		return googleBooksResponse.Items[0].VolumeInfo.ImageLinks.Thumbnail, nil
	}

	return "", nil
}
