package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var app = kingpin.New("App", "Simple app")

var (
	commandAdd             = app.Command("add", "add new user")
	commandAddFlagOverride = commandAdd.Flag("override", "override existing user").Short('o').Bool()
	commandAddArgUser      = commandAdd.Arg("user", "username").Required().String()
)

var (
	commandUpdate           = app.Command("update", "update user")
	commandUpdateArgOldUser = commandUpdate.Arg("old", "old username").Required().String()
	commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().String()
)

var (
	commandDelete          = app.Command("delete", "delete user")
	commandDeleteFlagForce = commandDelete.Flag("force", "force deletion").Short('f').Bool()
	commandDeleteArgUser   = commandDelete.Arg("user", "username").Required().String()
)

func main() {
	// commandAdd.Action(func(ctx *kingpin.ParseContext) error {
	// 	user := *commandAddArgUser
	// 	override := *commandAddFlagOverride
	// 	fmt.Printf("adding user %s with override %t\n", user, override)
	// 	return nil
	// })

	// commandUpdate.Action(func(ctx *kingpin.ParseContext) error {
	// 	oldUser := *commandUpdateArgOldUser
	// 	newUser := *commandUpdateArgNewUser
	// 	fmt.Printf("updating user from %s %s\n", oldUser, newUser)
	// 	return nil
	// })

	// commandDelete.Action(func(ctx *kingpin.ParseContext) error {
	// 	user := *commandDeleteArgUser
	// 	force := *commandDeleteFlagForce
	// 	fmt.Printf("deleting user %s, force %t\n", user, force)
	// 	return nil
	// })

	// kingpin.MustParse(app.Parse(os.Args[1:]))

	commandInString := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch commandInString {
	case commandAdd.FullCommand():
		user := *commandAddArgUser
		override := *commandAddFlagOverride
		fmt.Printf("adding user %s with override %t\n", user, override)

	case commandUpdate.FullCommand():
		oldUser := *commandUpdateArgOldUser
		newUser := *commandUpdateArgNewUser
		fmt.Printf("updating user from %s %s\n", oldUser, newUser)
	case commandDelete.FullCommand():
		user := *commandDeleteArgUser
		force := *commandDeleteFlagForce
		fmt.Printf("deleting user %s, force %t\n", user, force)
	}
}

// go run command.go add --help
// go run command.go --help-long
// go run command.go add --override "Setia"
// go run command.go update "Setia" "tia"
// go run command.go delete --force "Setia"
