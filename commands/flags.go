package commands

import (
	"flag"
	"fmt"
	"os"
)

// example gethometimeline cmd:
//
// gethometimeline hometimeline=ww userid=12 sinceid=12 count=12 max_id=15 trim_user=16 er=a include_rts=12
var (
	local_flag           = flag.NewFlagSet("local", flag.ExitOnError)
	gethometimeline_flag = flag.NewFlagSet("gethometimeline", flag.ExitOnError)
)

func GetLocalFlags() *bool {
	return local_flag.Bool("local", false, "if twittercli is running local or not")
}

func GetHomeTimeLineFlags() {
	gethometimeline_flag.String("hometimeline", "", "hometimeline of user")
	gethometimeline_flag.Int("user_id", 0, "user id of the user")
	gethometimeline_flag.Int("since_id", 0, "sense this id")
	gethometimeline_flag.Int("count", 0, "number of tweets")
	gethometimeline_flag.Int("max_id", 0, "maximum number of tweets")
	gethometimeline_flag.Int("trim_user", 0, "user to be trimmed")
	gethometimeline_flag.Bool("include_rts", true, "to include retweets or not")
}

func RunFlags() {
	if len(os.Args) == 1 {
		fmt.Printf("No command given\n")
		os.Exit(1)
	}

	GetHomeTimeLineFlags()
}
