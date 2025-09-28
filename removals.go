package main

import (
	"os/exec"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func removals() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Application to remove:", []string{"back", "all", "alacritty", "pavucontrol", "bash", "linux-cachyos-lts", "ufw", "modemmanager"})
		var target string
		switch choice {
		case 1:
			return
		case 2:
			cmd := exec.Command("pacman", "-Rcns", "back", "alacritty", "pavucontrol", "bash", "linux-cachyos-lts", "ufw", "modemmanager", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application removals", 1)
			}
			doAll = true
		case 3:
			target = "alacritty"
		case 4:
			target = "pavucontrol"
		case 5:
			target = "bash"
		case 6:
			target = "linux-cachyos-lts"
		case 7:
			target = "ufw"
			manageService("stop", "ufw.service")
			manageService("disable", "ufw.service")
		case 8:
			target = "modemmanager"
		}
		if !doAll {
			managePackage("-Rcns", target)
		}
	}
}
