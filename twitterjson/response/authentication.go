package twitterjson

const (
	post = "POST"
	get  = "GET"
)

// https://developer.twitter.com/en/docs/basics/authentication/api-reference/authenticate
type AuthenticateResponse struct {
	kind string
	path string
}

// TODO: forcelogin or screenName are not included in the path
func NewAuthenticateResponse() *AuthenticateResponse {
	return &UpdateCmd{}
}
