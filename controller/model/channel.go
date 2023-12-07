package model

type Channel struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type Channels []*Channel
