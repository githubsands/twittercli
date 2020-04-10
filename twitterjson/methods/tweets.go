package methods

import "errors"

type UpdateCmd struct {
	kind string
	path string
}

func NewUpdateCmd() *UpdateCmd {
	return &UpdateCmd{
		kind: post,
		path: "statuses/update",
	}
}

// DestroyTweetCmd destroys a tweet given the tweet identification number, id
type DestroyTweetCmd struct {
	kind string
	path string
}

func NewDestroyTweetCmd(i string) *DestroyTweetCmd {
	return &DestroyTweetCmd{
		kind: post,
		path: "statuses/destroy/:id" + i,
	}
}

// ShowTweetCmd shows a tweet given the tweet identification number, id
type ShowTweetCmd struct {
	kind string
	path string
}

func NewShowCmd(i string) *ShowTweetCmd {
	return &ShowTweetCmd{
		kind: get,
		path: "statuses/show/:" + i,
	}
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
type StatusLookUpCmd struct {
	kind string
	path string
}

func NewStatusLookUpCmd(k string) (*StatusLookUpCmd, error) {
	if k != get || k != post {
		return nil, errors.New("Post request must be GET or POST")
	}

	return &StatusLookUpCmd{
		kind: k,
		path: "statuses/lookup",
	}, nil
}

type RetweetCmd struct {
	kind string
	path string
}

// insert id here.
func NewRetweetCmd() *RetweetCmd {
	return &RetweetCmd{
		kind: post,
		path: "statuses/retweet/:id",
	}
}

type UnRetweetCmd struct {
	kind string
	path string
}

// insert id here.
func NewUnRetweetCmd() *UnRetweetCmd {
	return &UnRetweetCmd{
		kind: post,
		path: "statuses/unretweet/:id",
	}
}

type RetweetCmdGet struct {
	kind string
	path string
}

type RetweetsOfMeCmd struct {
	kind string
	path string
}

// insert id here.
func NewUnRetweetCmdGet() *RetweetsOfMeCmd {
	return &RetweetsOfMeCmd{
		kind: get,
		path: "statuses/retweets_of_me",
	}
}

// Likes (formerly favorites) - 3

// 1
type CreateLikeCmd struct {
	kind string
	path string
}

func NewCreateLikeCmd(i string) *CreateLikeCmd {
	return &CreateLikeCmd{
		kind: post,
		path: "favorites/create/:" + i,
	}
}

// 2
type DestroyLikeCmd struct {
	kind string
	path string
}

func NewDestroyLikeCmd(i string) *DestroyLikeCmd {
	return &DestroyLikeCmd{
		kind: post,
		path: "favorites/destroy/:" + i,
	}
}

// 3
type GetLikesCmd struct {
	kind string
	path string
}

func NewGetLikesCmd() *GetLikesCmd {
	return &GetLikesCmd{
		kind: get,
		path: "favorites/likes",
	}
}
