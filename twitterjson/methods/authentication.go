package twitterjson

const (
	post = "POST"
	get  = "GET"
)

// https://developer.twitter.com/en/docs/basics/authentication/api-reference/authenticate
type AuthenticateCmd struct {
	kind string
	path string
}

// TODO: forcelogin or screenName are not included in the path
func NewAuthenticateCmd(n string, forceLogin *bool, screenName *string) *AuthenticateCmd {
	return &UpdateCmd{
		kind: get,
		path: "oauth/authenticate,
	}
}
