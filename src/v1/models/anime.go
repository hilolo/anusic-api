package models

// Anime type
type Anime struct {
	ID       uint16   `json:"id"`
	Name     string   `json:"name"`
	AltNames []string `json:"altNames"`
	Year     uint16   `json:"year"`
	Themes   []Theme  `json:"themes"`
	LinkInfo string
}
