package model

type MovieModel struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Genre       string  `json:"genre"`
	Rating      float64 `json:"rating"`
	ReleaseDate string  `json:"release_date"`
	UpdatedAt   string  `json:"updatedAt"`
	CreatedAt   string  `json:"createdAt"`
	DeletedAt   string  `json:"deletedAt,omitempty"`
	Plot        string  `json:"plot"`
	Released    bool    `json:"released"`
}
