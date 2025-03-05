package domain

type News struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Content         string `json:"content"`
	Topic           string `json:"topic"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication"`
}
