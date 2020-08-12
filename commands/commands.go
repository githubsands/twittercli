package commands

import (
	"flag"

	"github.com/githubsands/twittercli/twitterjson/methods"
)

var commands = map[string]interface{}{

	// search commands
	"gethometimeline": methods.GetHomeTimeLineCmd{},
	"getsearchtweets": methods.GetSearchTweetsCmd{},

	// like commands
	"createlike":  methods.CreateLikeCmd{},
	"destroylike": methods.DestroyLikeCmd{},
	"getlikes":    methods.GetLikesCmd{},

	// authenticate commands
	"authenticate": methods.AuthenticateCmd{},

	// tweet commands
	"retweet":      methods.RetweetCmd{},
	"statuslookup": methods.StatusLookUpCmd{},
	"unretweetcmd": methods.UnRetweetCmd{},
}

var flags = map[string]*flag.FlagSet{
	"gethometimeline": nil,
}
