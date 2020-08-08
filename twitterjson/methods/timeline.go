package methods

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
