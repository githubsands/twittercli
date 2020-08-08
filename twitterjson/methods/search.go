// https://developer.twitter.com/en/docs/tweets/search/api-reference/get-search-tweets.html
// https://developer.twitter.com/en/docs/tweets/timelines/guides/working-with-timelines
package methods

type GetSearchTweetsCmd struct {
	q           string  //TODO: this may need to change
	geocode     *string // this is suspected to change
	language    *string // uses the standard language geocode
	locale      *string
	result_type *string
	count       *int
}
