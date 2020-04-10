// https://developer.twitter.com/en/docs/tweets/search/api-reference/get-search-tweets.html
// https://developer.twitter.com/en/docs/tweets/timelines/guides/working-with-timelines
package methods

// TODO: Move GetHomeTimeLindCmd somewhere else
type GetHomeTimeLineResponse struct {
}

func NewGetHomeTimeLineResponse() *GetHomeTimeLineResponse {
	return &GetHomeTimeLineResponse{}
}

type GetSearchTweetsResponse struct {
}

func NewGetSearchTweetsResponse() *GetSearchTweetsResponse {
	return &GetSearchTweetsResponse{}
}
