package twitterjson

const (
	post     = "POST"
	get      = "GET"
)

type command struct {
	commandType string
	cmd         interface{}
	path        string
}

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

func NewDestroyTweetCmd(i string) *DestroyCmd {
	return &DestroyCmd{
		kind: post,
		path: "statuses/destroy/:id"+i,
	}
}

// ShowTweetCmd shows a tweet given the tweet identification number, id
type ShowTweetCmd struct {
	kind string
	path string
}

func NewShowCmd(i string) *ShowTweetCmd {
	return &ShowCmd{
		kind: get,
		path: "statuses/show/:"+i, 
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

func NewStatusLookUpCmd(k string) *StatusCmd, error {
	if k != get || k != post {
		return nil, errors.New("Post request must be GET or POST")
	}

	return &StatusCmd{
		kind: k,
		path: "statuses/lookup",
	}
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
		path: "favorites/create/:"+i,
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
		path: "favorites/destroy/:"+i, 
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

// Get Tweet timelines - 3
//
// A timeline is simply a list, or an aggregate stream of tweets.

// 1
type GetHomeTimeLineCmd struct {
	kind         string
	homeTimeLine string
	path         string
	user_id      string
	screen_name  string
	since_id     int
	count        int
	max_id       int
	trim_user    bool
	er           bool
	include_rts  bool
}

func NewGetUsersTweetsCmd(homeTimeLine, kind, path, ui string, si int, c int, mi int, tu bool, er bool, it bool, rts bool) *GetHomeTimeLineCmd {
	return &GetHomeTimeLineCmd{
		kind:         get,
		path:         "statuses/" + homeTimeLine,
		homeTimeLine: homeTimeLine, //TODO: this may not be right
		user_id:      ui,
		since_id:     si,
		count:        c,
		max_id:       mi,
		trim_user:    tu,
		er:           er,
		include_rts:  rts,
	}
}

// Create and Manage List

// developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-list
type GetListCmd struct {
	path        string
	kind        string
	name        string
	user_id     *string
	screen_name *string
	reverse     *string
}

func NewGetListCmd(n string, u *string, sn *string, r *string) *GetListCmd {
	return &GetListCmd{
		path:        "list/members",
		kind:        "GET",
		name:        n,
		user_id:     u,
		screen_name: sn,
		reverse:     r,
	}
}

type GetListMembersCmd struct {
	path              string
	kind              string
	name              string
	slug              string
	owner_screen_name *string
	owner_id          *int
	count             *int
	cursor            *int
	include_entities  *bool
	skip_status       *bool
}

func NewGetListMembersCmd(n, sl string, osn *string, oi, c, cu *int, ie, ss *bool) {
	return &GetListMemberCmd{
		path: "", 
		kind: "GET", 
		slug: sl, 
		owner_screen_name: osn, 
		owner_id: oi, 
		count: c, 
		cursor: cu, 
		include_entities: ie, 
		skip_status: ss, 
	}
}

// GetListShow Returns the specified list. Private lists will only be shown if the authenticated user owns the specified list.
type GetListShowCmd struct {
	path              string
	kind              string
	list_id           int
	name              string
	slug              int
	owner_screen_name *string
	owner_id          *string
}

func NewGetListShowCmd(n string, s int, o *string, oi *string) &GetListShowCmd {
	return &GetListShowCmd{
		path:              "lists/show",
		kind:              "GET",
		slug:              s,
		owner_screen_name: o,
		owner_id:          oi,
	}
}

type GetListCreateCmd struct {
	path        string
	kind        string
	name        string
	mode        *string
	description *string
}

// TODO: Fix this
func NewGetListCreateCmd(n string, m, d *string) *GetListCreateCmd {
	return &GetListCreateCmd{
		path:        "https://api.twitter.com/1.1/lists/create.json",
		kind:        "GET",
		name:        n,
		mode:        m,
		description: d,
	}
}
