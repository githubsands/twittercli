// https://developer.twitter.com/en/docs/tweets/search/api-reference/get-search-tweets.html
// https://developer.twitter.com/en/docs/tweets/timelines/guides/working-with-timelines
package methods

// TODO: Move GetHomeTimeLindCmd somewhere else
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

const (
	MIXED = "mixed"
	RECENT = "recent"
	popular = "popular"
)

type GetSearchTweetsCmd struct {
	q string //TODO: this may need to change
	geocode *string // this is suspected to change 
	language *string // uses the standard language geocode
	locale *string 
	result_type *string
	count *int
}

func NewGetNewUsersTweetsCmd(q string, gc *string, l *string, l *string, r *string, 



	
