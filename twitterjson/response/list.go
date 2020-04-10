package main

// Create and Manage List

// developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-list
type GetListMembersResponse struct {
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

func NewGetListMembersResponse() {
	return &GetListMembersResponse{}
}

type GetListResponse struct {
}

func NewGetListResponse() *GetListResponse {
	return &GetListResponse{}
}

// GetListShow Returns the specified list. Private lists will only be shown if the authenticated user owns the specified list.
type GetListShowResponse struct {
}

func NewGetListShowResponse() *GetListShowResponse {
	return &GetListShowResponse{}
}

type GetListCreateResponse struct {
}

func NewGetListResponse() *GetListCreateResponse {
	return &GetListResponse{}
}
