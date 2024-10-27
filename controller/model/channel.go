package model

// Swaggoのコメントを記述するためのファイル

type Channel struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type Channels []*Channel

type PostChannelPayload struct {
	Name string `json:"name"`
}
