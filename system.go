package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func system() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Tweak:", []string{"back", "all", "replace sudo with doas", "enable ipv6 privacy extensions", "enable ext4 fast_commit", "enable SysRq reboots", "disable kernel security mitigations", "unlock amdgpu tuning", "enable amd_pstate in passive mode (recommended for zenver2)"})
		switch choice {
		case 1:
			return
		case 2:
			doAll = true
			fallthrough
		case 3:
			writeFile("/etc/doas.conf", "permit persist keepenv :wheel as root", "root", 0644)
			writeFile("/tmp/PKGBUILD", `pkgname=base-devel
pkgver=1
pkgrel=2
pkgdesc='Basic tools to build Arch Linux packages'
url='https://www.archlinux.org'
arch=('any')
license=('GPL')
options=('!debug')
depends=(
  archlinux-keyring
  autoconf
  automake
  binutils
  bison
  debugedit
  fakeroot
  file
  findutils
  flex
  gawk
  gcc
  gettext
  grep
  groff
  gzip
  libtool
  m4
  make
  pacman
  pacman-contrib
  patch
  pkgconf
  sed
  opendoas
  texinfo
  which
)

# vim: ts=2 sw=2 et:
`, getUsername(), 0777)
			cmd := exec.Command("sudo", "-u", getUsername(), "makepkg", "-si", "--dir", "/tmp")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to build+install custom base-devel", 1)
			}
			cmd = exec.Command("pacman", "-S", "opendoas", "--noconfirm")
			err = cmd.Run()
			if err != nil {
				other.PrintError("Failed to install opendoas", 1)
			}
			cmd = exec.Command("pacman", "-R", "sudo", "--noconfirm")
			err = cmd.Run()
			if err != nil {
				other.PrintError("Failed to remove sudo", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 4:
			nics, err := os.ReadDir("/sys/class/net")
			if err != nil {
				other.PrintError("Failed to retrieve NICs", 1)
			}
			var sysctlConf strings.Builder
			sysctlConf.WriteString("net.ipv6.conf.all.use_tempaddr = 2\nnet.ipv6.conf.default.use_tempaddr = 2")
			for _, nic := range nics {
				sysctlConf.WriteString("\nnet.ipv6.conf." + nic.Name() + "use_tempaddr = 2")
			}
			writeFile("/etc/sysctl.d/40-ipv6-priv-ext.conf", sysctlConf.String(), "root", 0644)

			if !doAll {
				break
			}
			fallthrough
		case 5:
			cmd := exec.Command("tune2fs", "-O", "fast_commit", front.Input("/dev/driveparition:"))
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to enable fast_commit", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 6:
			writeFile("/etc/sysctl.d/35-sysrq.conf", "kernel.sysrq = 244", "root", 0644)

			if !doAll {
				break
			}
			fallthrough
		case 7:
			addKernelParams([]string{"mitigations=off"})

			if !doAll {
				break
			}
			fallthrough
		case 8:
			addKernelParams([]string{"amdgpu.ppfeaturemask=0xffffffff"})

			if !doAll {
				break
			}
			fallthrough
		case 9:
			addKernelParams([]string{"amd_pstate=passive"})
		}
	}
}

func addKernelParams(newParameters []string) {
	f, err := os.Open("/boot/loader/entries/linux-cachyos.conf")
	if err != nil {
		other.PrintError("Failed to open systemd-boot entry", 1)
	}
	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var parameters []string
	for i := 0; fScanner.Scan(); i++ {
		if i == 1 {
			parameters = strings.Split(fScanner.Text(), " ")[2:]
			break
		}
	}
	parameters = append(parameters, newParameters...)
	writeFile("/boot/loader/entries/linux-cachyos.conf", "title Linux Cachyos\noptions root=UUID=e2ff600b-0202-4620-bed3-54d7a58414f8 "+strings.Join(parameters, " ")+"\nlinux /vmlinuz-linux-cachyos\ninitrd /initramfs-linux-cachyos.img", "root", 0700)
}
