package main

import (
	"os/exec"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func applications() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Application to install:", []string{"back", "all", "pamac-aur (replaces octopi)", "yay-bin (replaces paru)", "htop (replaces btop)", "neovim (replaces vi+vim+nano+micro)", "virt-manager (also installs qemu-desktop+libvirt)", "librewolf-bin", "zed", "steam"})
		switch choice {
		case 1:
			return
		case 2:
			doAll = true
			fallthrough
		case 3:
			cmd := exec.Command("pacman", "-S", "pamac-aur", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installation", 1)
			}
			cmd = exec.Command("pacman", "-Rcns", "octopi", "--noconfirm")
			err = cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application removal", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 4:
			cmd := exec.Command("pacman", "-S", "yay-bin", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installation", 1)
			}
			cmd = exec.Command("pacman", "-Rcns", "paru", "--noconfirm")
			err = cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application removal", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 5:
			cmd := exec.Command("pacman", "-S", "htop", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installation", 1)
			}
			cmd = exec.Command("pacman", "-Rcns", "btop", "--noconfirm")
			err = cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application removal", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 6:
			cmd := exec.Command("pacman", "-S", "neovim", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installation", 1)
			}
			cmd = exec.Command("pacman", "-Rcns", "vi", "vim", "nano", "micro", "--noconfirm")
			err = cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application removals", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 7:
			cmd := exec.Command("pacman", "-S", "virt-manager", "qemu-desktop", "libvirt", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installations", 1)
			}
			manageService("enable", "libvirtd.service")
			manageService("start", "libvirtd.service")

			if !doAll {
				break
			}
			fallthrough
		case 8:
			cmd := exec.Command("pacman", "-S", "librewolf-bin", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installation", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 9:
			cmd := exec.Command("pacman", "-S", "zed", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installation", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 10:
			cmd := exec.Command("pacman", "-S", "steam", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installation", 1)
			}
		}
	}
}
