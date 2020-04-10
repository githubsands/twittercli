package commands

import (
	"fmt"
	"testing"
)

/*
func TestGetFlagSets(t *testing.T) {
	flags := GetFlagSets()
	spew.Dump(flags)
}
*/

/*
func TestFlagSetToCommand(t *testing.T) {
	var (
		cmdName = "gethometimeline"
		f       = flag.NewFlagSet(cmdName, flag.ContinueOnError)
		_       = f.Int("user_id", 4, "the user to be looked up")
		_       = f.Int("since_id", 3, "the date sense")
	)

	spew.Dump(f)
	actual := FlagSetToCommand(f)
	expected := GetHomeTimeLineCmd{user_id: 4, since_id: 3}

	reflect.DeepEqual(actual, expected)
}
*/

func Test(t *testing.T) {
	actual := BuildFlagSetFromCommands()
	fmt.Print(actual)
}
