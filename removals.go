package main

import (
	"os/exec"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func removals() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Application to remove:", []string{"back", "all", "bash", "linux-cachyos-lts"})
		var target string
		switch choice {
		case 1:
			return
		case 2:
			cmd := exec.Command("pacman", "-Rcns", "bash", "linux-cachyos-lts", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application removals", 1)
			}
			doAll = true
		case 3:
			target = "bash"
		case 4:
			target = "linux-cachyos-lts"
		}
		if !doAll {
			managePackage("-Rcns", target)
		}
	}
}
