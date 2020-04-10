package twitterjson

type UpdateResponse struct {
	kind string
	path string
}

func NewUpdateResponse(s string) *UpdateCmdResponse {
	return &UpdateResponse{}
}

// DestroyTweetCmd destroys a tweet given the tweet identification number, id
type DestroyTweetResponse struct {
}

func NewDestroyTweetResponse(i string) *DestroyTweetResponse {
	return &DestroyTweetResponse{}
}

type ShowTweetResponse struct {
}

func NewShowTweetResponse(i string) *ShowTweetResponse {
	return &ShowTweetResponse{}
}

/*
// TO DO FINISH THIS:
// Returns a tweet in oembed form given the tweets
type EmbedCmd struct {
	kind string
	path string
}

// insert id here.
func NewEmbedCmd() *EmbedCmd {
	return &EmbedCmd{
		kind: get,
		path: "statuses/oembed",
	}
}
*/

// StatusLookUpCmd gets a fully-hydrated tweet objects
//
// It is suggested to use
type StatusLookUpResponse struct {
}

func NewStatusLookUpResponse() *StatusLookUpResponse {
	return &StatusLookUpResponse{}
}

type RetweetResponse struct {
}

func NewRetweetResponse() *RetweetResponse {
	return &RetweetResponse{}
}

type UnRetweetResponse struct {
}

func NewUnRetweetResponse() *UnRetweetResponse {
	return &UnRetweetResponse{}
}

type RetweetResponse struct {
}

func NewRetweetResponse() *RetweetResponse {
	&RetweetResponse{}
}

type RetweetsOfMeResponse struct {
}

func NewRetweetsOfMeResponse() *RetweetsOfMeResponse {
	return &RetweetsOfMeResponse{}
}

// Likes (formerly favorites) - 3

// 1
type CreateLikeResponse struct {
}

func NewCreateLikeResponse() *CreateLikeResponse {
	return &CreateLikeResponse{}
}

// 2
type DestroyLikeResponse struct {
}

func NewDestroyLikeResponse() *DestroyLikeResponse {
	return &DestroyLikeResponse{}
}

// 3
type GetLikesResponse struct {
}

func NewGetLikesResponse() *GetLikesResponse {
	return &GetLikesResponse{}
}
