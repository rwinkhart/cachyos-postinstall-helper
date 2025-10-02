package main

func universalTweaks() {
	// Lock root account
	chrootCommandRun([]string{"passwd", "--delete", "root"})
	chrootCommandRun([]string{"usermod", "--lock", "root"})

	// Add user to uucp group for non-root serial console usage
	chrootCommandRun([]string{"usermod", "-aG", "uucp", getUsername()})

	// Force disable kernel watchdog (kernel parameters should do this as well)
	writeFile(mountPoint+"/etc/modprobe.d/blacklist.conf", "# Disable kernel watchdog\nblacklist iTCO_wdt\n", "root", 0644)

	// Install essential packages
	managePackages("-S", []string{"wl-clipboard", "qt6-wayland"})
}
