package main

import (
	"os/exec"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func premount() {
	for {
		var choice int
		if performAllTweaks {
			choice = 2
		} else {
			choice = front.InputMenuGen("Tweak:", []string{"BACK", "enable ext4 fast_commit"})
		}
		switch choice {
		case 1:
			return
		case 2:
			cmd := exec.Command("tune2fs", "-O", "fast_commit", partitionR)
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to enable fast_commit", 1)
			}
		}
		if !performAllTweaks {
			break
		}
	}
}
