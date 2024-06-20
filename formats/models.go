package formats

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

func NewFormat(id uint, format_name string) Format {
	format := Format{
		ID:     id,
		Format: format_name,
	}
	return format
}
