package model

// Swaggoのコメントを記述するためのファイル

type User struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type Users []*Channel
