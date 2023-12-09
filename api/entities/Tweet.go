package entities

import "github.com/pborman/uuid"

type Tweet struct {
	ID          string `json:"id"`
	User        string `json:"user"`
	Description string `json:"description"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}

	return &tweet
}
