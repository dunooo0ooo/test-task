package models

type Post struct {
	ID       string     `json:"id"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	IsLocked bool       `json:"isLocked"`
	Comments []*Comment `json:"comments"`
}
