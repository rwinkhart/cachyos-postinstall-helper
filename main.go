package main

import (
	"os/user"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

var username string

// TODO add mode to perform EVERYTHING
func main() {
	userInfo, err := user.Current()
	if err != nil {
		other.PrintError("Failed to get current user", 1)
	}
	if userInfo.Uid != "0" {
		other.PrintError("This utility must be run as root", 1)
	}

	for {
		choice := front.InputMenuGen("Option:", []string{"universal tweaks", "system tweaks", "removals", "applications", "configuration", "cosmic"})
		switch choice {
		case 1:
			universalTweaks()
		case 2:
			system()
		case 3:
			removals()
			// TODO remove config/data
		case 4:
			applications()
			// TODO remove config/data for replacements
		case 5:
			configurations()
		case 6:
			cosmic()
			// TODO keyboard shortcuts for screenshot/sticky window/terminal
			// TODO window decoration position (not yet implemented?)
			// TODO link CachyOS wallpapers?
		}
	}
}
