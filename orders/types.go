package orders

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			ImageLinks struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo"`
	} `json:"items"`
}
