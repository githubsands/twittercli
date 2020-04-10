package main

// Get Tweet timelines - 3
//
// A timeline is simply a list, or an aggregate stream of tweets.

// 1
type GetHomeTimeLineResponse struct {
}

func NewGetUsersTimeLineResponse() *GetHomeTimeLineResponse {
	return &GetHomeTimeLineResponse{}
}
