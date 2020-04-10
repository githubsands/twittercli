package methods

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
