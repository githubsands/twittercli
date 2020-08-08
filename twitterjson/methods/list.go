package methods

// Create and Manage List

// developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-list
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

func NewGetListMembersCmd(n, sl string, osn *string, oi, c, cu *int, ie, ss *bool) *GetListMembersCmd {
	return &GetListMembersCmd{
		path:              "",
		kind:              "GET",
		slug:              sl,
		owner_screen_name: osn,
		owner_id:          oi,
		count:             c,
		cursor:            cu,
		include_entities:  ie,
		skip_status:       ss,
	}
}

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
