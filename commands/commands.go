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

// example gethometimeline cmd:
//
//    gethometimeline hometimeline=ww userid=12 sinceid=12 count=12 max_id=15 trim_user=16 er=a include_rts=12
//
//

// var genFlags = BuildFlagSetFromCommands()

/*
// TODO: place in init for autogeneration of flag list
func BuildFlagSetFromCommands() map[string]*flag.FlagSet {
	var flags map[string]*flag.FlagSet
	for k, _ := range commands {
		// TODO: add actual flag.Set
		flags[k] = nil
	}

	return flags
}
*/

func GetFlagSets() map[string]*flag.FlagSet {
	flagSets := make(map[string]*flag.FlagSet)
	for i, _ := range commands {
		flagSets[i] = flag.NewFlagSet(i, flag.ExitOnError)
	}

	return flagSets
}

// The Command is unknown at compile time; it must be built using reflection purely off the name of given command.

/*
func FlagSetToCommand(f *flag.FlagSet) interface{} {
	// get command
	name := f.Name()
	cmd := commands[name]

	// get each field of cmd
	// range through the struct to get each option, we need the type and name of each methods' or structures fields to
	// automate the creation of their respective flags.

	ty := reflect.TypeOf(cmd).Elem()
	val := reflect.ValueOf(cmd).Elem()
	for i := 0; i < ty.NumField(); i++ {
		tyField := ty.Field(i)
		cmd.tyField.Name.String() = getValue(tyField)
		// opt = option{method: k, name: tyField.Name, itsType: tyField.Type.String()}
		// options = append(options, opt)
	}

	return cmd
}

func getValue(v reflect.Value) string {
	var value interface{}
	i := v.Interface()

	switch z := i.(type) {
	case string:
		return v.String()
	default:
		return ""
	}
}

// func FlagSetToCommand() interface{} {
*/
