package main

import (
	"github.com/rwinkhart/go-boilerplate/front"
)

func configurations() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Custom config:", []string{"BACK", "ALL", "shell profile", "zshrc", "neovim", "zed", "makepkg", "pacman"})
		switch choice {
		case 1:
			return
		case 2:
			doAll = true
			fallthrough
		case 3:
			managePackage("-S", "wayland-protocols")
			writeFile(mountPoint+"/home/"+getUsername()+"/.profile", shellProfile, getUsername(), 0644)

			if !doAll {
				break
			}
			fallthrough
		case 4:
			writeFile(mountPoint+"/home/"+getUsername()+"/.zshrc", zshrc, getUsername(), 0644)

			if !doAll {
				break
			}
			fallthrough
		case 5:
			writeFile(mountPoint+"/etc/xdg/nvim/sysinit.vim", sysinitVim, "root", 0644)

			if !doAll {
				break
			}
			fallthrough
		case 6:
			mkdir(mountPoint + "/home/" + getUsername() + "/.config/zed")
			writeFile(mountPoint+"/home/"+getUsername()+"/.config/zed/settings.json", zed, getUsername(), 0600)

			if !doAll {
				break
			}
			fallthrough
		case 7:
			writeFile(mountPoint+"/etc/makepkg.conf", makepkg, "root", 0644)

			if !doAll {
				break
			}
			fallthrough
		case 8:
			managePackage("-S", "pacman-contrib")
			writeFile(mountPoint+"/etc/pacman.conf", pacman, "root", 0644)
			writeFile(mountPoint+"/etc/pacman.d/hooks/paccache-clean.hook", pacmanHookClean, "root", 0644)
			writeFile(mountPoint+"/etc/pacman.d/hooks/nvidia.hook", pacmanHookNvidia, "root", 0644)
		}
	}
}

const shellProfile = `#!/bin/sh

# WAYLAND ENV
export WAYLAND_PROTOCOLS_DATADIR="/usr/share/wayland-protocols"
export ELECTRON_OZONE_PLATFORM_HINT=auto

# XDG BASE ENV
export XDG_DATA_HOME="$HOME"/.local/share
export XDG_CONFIG_HOME="$HOME"/.config
export XDG_STATE_HOME="$HOME"/.local/state
export XDG_CACHE_HOME="$HOME"/.cache

# DOTFILE MANAGEMENT
export CUDA_CACHE_PATH="$XDG_CACHE_HOME"/nv
export GTK2_RC_FILES="$XDG_CONFIG_HOME"/gtk-2.0/gtkrc
export LESSHISTFILE="$XDG_CACHE_HOME"/less/history
export WINEPREFIX="$XDG_DATA_HOME"/wine
export ANDROID_USER_HOME="$XDG_DATA_HOME"/android
export _JAVA_OPTIONS=-Djava.util.prefs.userRoot="$XDG_CONFIG_HOME"/java
export CARGO_HOME="$XDG_DATA_HOME"/cargo
export RUSTUP_HOME="$XDG_DATA_HOME"/rustup

# GO SETTINGS
export GOPATH="$XDG_DATA_HOME"/go
export CGO_ENABLED=1
export GOFLAGS='-trimpath "-ldflags=-s -w" -buildvcs=false'
export GOAMD64=v3

# PROTON SETTINGS
export PROTON_USE_NTSYNC=1
export PROTON_ENABLE_WAYLAND=1
#export WAYLANDDRV_PRIMARY_MONITOR=DP-3
#export FSR4_UPGRADE=1

# MISC SETTINGS
export EDITOR=nvim
export BUILDDIR=/tmp/makepkg
`

const zshrc = `# Shell settings
HISTFILE="$XDG_CACHE_HOME"/zsh-histfile
HISTSIZE=2500
SAVEHIST=2000
setopt INC_APPEND_HISTORY
setopt HIST_EXPIRE_DUPS_FIRST
setopt HIST_FIND_NO_DUPS
unsetopt beep
bindkey -e
zstyle :compinstall filename "/home/$USER/.zshrc"
autoload -Uz compinit
compinit -d "$XDG_CACHE_HOME"/zcompdump
PS1='%F{yellow}%?%F{cyan}|%n%f@%F{cyan}%m:%f%1~%f%F{cyan}|%#%f '
source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh
bindkey '^[[H' beginning-of-line
bindkey '^[[F' end-of-line
bindkey '^[[3~' delete-char

# Text art on clear/new shell
art_cat='%F{#fe8019}/\_,_/\
\ 0 0 /%f'
art_dog=' %F{#d79921}/___/
| 6.6 |%f'
art_bear='%F{#d65d0e}() ()
(o.o)%f'
art_snail=' %F{#689d6a}|_/
/o o\%f'
art_bunny=' %F{#d5c4a1}/_/
|^.^|'
art_jellyfish='%F{#83a598}<===>
//|\\%f'
art_school='    %F{#fe8019}<°))><     %F{#fabd2f}><(((.>
%F{#8ec07c}><(((°>   %F{#d3869b}><((((°>%f'
art_array=("$art_cat" "$art_dog" "$art_bear" "$art_snail" "$art_bunny" "$art_jellyfish" "$art_school")
print -Pr "$art_array[$((RANDOM%7+1))]"
function term_clear() {
    clear
    print -Pr "$art_array[$((RANDOM%7+1))]"
    print -Pn $PS1
}
zle -N zsh-redraw term_clear
bindkey "^L" zsh-redraw

# Stock aliases
alias ls='ls --color=auto'
alias grep='grep --color=auto'
alias fastfetch='fastfetch -c neofetch'
alias unmv='setopt shwordsplit && lastCmdSplit=(${(@s/ /)$(fc -ln -1)}) && newCmd="mv ${lastCmdSplit[3]} ${lastCmdSplit[2]}" && $newCmd && unsetopt shwordsplit'
alias orphans='doas pacman -Rcns $(pacman -Qqdt)'
alias powersave='doas /usr/local/bin/powerset.sh powersave'
alias performance='doas /usr/local/bin/powerset.sh performance'
alias schedutil='doas /usr/local/bin/powerset.sh schedutil'
alias poweroff='doas /usr/bin/poweroff'
alias reboot='doas /usr/bin/reboot'
`

const sysinitVim = `" rwinkhart/cachyos-postinstall-helper 09/28/2025

" ensure correct basic settings
set fileencoding=utf-8

" show line numbers
set number

" enable basic spellchecking
set spell

" set keybinds (home/end, delete w/o copy)
noremap <C-a> <Home>
imap <C-a> <Home>
noremap <C-e> <End>
imap <C-e> <End>
nnoremap x "_x
nnoremap <delete> "_x

" expand all tabs and indents to 4 spaces
set tabstop=4
set shiftwidth=4
set expandtab
autocmd FileType go,html,php,css set noexpandtab

" set default clipboard register to system clipboard
set clipboard=unnamedplus

" enable visually indented text wrapping
set wrap
set breakindent
set showbreak=→

" do not auto-break lines unless they are 1000+ characters long
set textwidth=1000

" do not create automatic backups of any sort
set nobackup
set nowritebackup

" default to case insensitive search unless a capital is typed
set ignorecase
set smartcase

" tell vim to remember certain things when exited
"  '10  :  marks will be remembered for up to 10 previously edited files
"  "100 :  will save up to 100 lines for each register
"  :20  :  up to 20 lines of command-line history will be remembered
"  %    :  saves and restores the buffer list
"  n... :  where to save the viminfo files
set viminfo='10,\"100,:20,%,n~/.local/share/nviminfo

" restores cursor position in recently opened files
function! ResCur()
  if line("'\"") <= line("$")` +
	"\nnormal! g`" + `"
    return 1
  endif
endfunction

augroup resCur
  autocmd!
  autocmd BufWinEnter * call ResCur()
augroup END
`

const zed = `{
	// rwinkhart/cachyos-postinstall-helper 09/28/2025

    // AI Settings
    "telemetry": {
        "metrics": false,
        "diagnostics": false
    },
    "edit_predictions": {
        "mode": "subtle",
        "enabled_in_text_threads": false
    },
    "show_edit_predictions": true,
    "agent": {
        "default_profile": "ask",
        "default_model": {
            "provider": "copilot_chat",
            "model": "claude-sonnet-4"
        }
    },
    "features": {
        "edit_prediction_provider": "copilot"
    },
    // General Settings
    "minimap": {
        "show": "always"
    },
    "icon_theme": "JetBrains New UI Icons (Dark)",
    "middle_click_paste": false,
    "base_keymap": "VSCode",
    "autosave": "off",
    "theme": "Gruvbox Dark",
    "ui_font_size": 14,
    "buffer_font_size": 18,
    "auto_update": false,
    "terminal": {
        "shell": {
            "program": "zsh"
        }
    },
    "tab_size": 4,
    "preferred_line_length": 120,
    "git": {
        "inline_blame": {
            "enabled": false
        }
    },
    // Language-specific Settings
    "languages": {
        "Python": {
            "soft_wrap": "preferred_line_length"
        },
        "Go": {
            "hard_tabs": true
        },
        "JavaScript": {
            "format_on_save": "off"
        },
        "CSS": {
            "format_on_save": "off"
        }
    },
    // De-clutter UI
    "tab_bar": {
        "show_nav_history_buttons": false
    },
    "search": {
        "button": false
    },
    "title_bar": {
        "show_branch_name": false
    },
    "notification_panel": {
        "button": false
    },
    "collaboration_panel": {
        "button": false
    },
    "git_panel": {
        "button": false
    }
}
`

const makepkg = `#!/hint/bash
#
# /etc/makepkg.conf - rwinkhart/cachyos-postinstall-helper 09/28/2025
#

#########################################################################
# SOURCE ACQUISITION
#########################################################################
#
#-- The download utilities that makepkg should use to acquire sources
#  Format: 'protocol::agent'
DLAGENTS=('file::/usr/bin/curl -qgC - -o %o %u'
          'ftp::/usr/bin/curl -qgfC - --ftp-pasv --retry 3 --retry-delay 3 -o %o %u'
          'http::/usr/bin/curl -qgb "" -fLC - --retry 3 --retry-delay 3 -o %o %u'
          'https::/usr/bin/curl -qgb "" -fLC - --retry 3 --retry-delay 3 -o %o %u'
          'rsync::/usr/bin/rsync --no-motd -z %u %o'
          'scp::/usr/bin/scp -C %u %o')

# Other common tools:
# /usr/bin/snarf
# /usr/bin/lftpget -c
# /usr/bin/wget

#-- The package required by makepkg to download VCS sources
#  Format: 'protocol::package'
VCSCLIENTS=('bzr::breezy'
            'fossil::fossil'
            'git::git'
            'hg::mercurial'
            'svn::subversion')

#########################################################################
# ARCHITECTURE, COMPILE FLAGS
#########################################################################
#
CARCH="x86_64"
CHOST="x86_64-pc-linux-gnu"

#-- Compiler and Linker Flags
#CPPFLAGS=""
CFLAGS="-march=native -mtune=native -O3 -pipe -fno-plt -fexceptions \
        -Wp,-D_FORTIFY_SOURCE=3 -Wformat -Werror=format-security \
        -fstack-clash-protection -fcf-protection \
        -fno-omit-frame-pointer -mno-omit-leaf-frame-pointer"
CXXFLAGS="$CFLAGS -Wp,-D_GLIBCXX_ASSERTIONS"
LDFLAGS="-Wl,-O3 -Wl,--sort-common -Wl,--as-needed -Wl,-z,relro -Wl,-z,now \
         -Wl,-z,pack-relative-relocs"
LTOFLAGS="-flto=auto"
#-- Make Flags: change this for DistCC/SMP systems
MAKEFLAGS="-j$(nproc)"
NINJAFLAGS="-j$(nproc)"
#-- Debugging flags
DEBUG_CFLAGS="-g"
DEBUG_CXXFLAGS="$DEBUG_CFLAGS"

#########################################################################
# BUILD ENVIRONMENT
#########################################################################
#
# Makepkg defaults: BUILDENV=(!distcc !color !ccache check !sign)
#  A negated environment option will do the opposite of the comments below.
#
#-- distcc:   Use the Distributed C/C++/ObjC compiler
#-- color:    Colorize output messages
#-- ccache:   Use ccache to cache compilation
#-- check:    Run the check() function if present in the PKGBUILD
#-- sign:     Generate PGP signature file
#
BUILDENV=(!distcc color !ccache check !sign)
#
#-- If using DistCC, your MAKEFLAGS will also need modification. In addition,
#-- specify a space-delimited list of hosts running in the DistCC cluster.
#DISTCC_HOSTS=""
#
#-- Specify a directory for package building.
#BUILDDIR=/tmp/makepkg

#########################################################################
# GLOBAL PACKAGE OPTIONS
#   These are default values for the options=() settings
#########################################################################
#
# Makepkg defaults: OPTIONS=(!strip docs libtool staticlibs emptydirs !zipman !purge !debug !lto !autodeps)
#  A negated option will do the opposite of the comments below.
#
#-- strip:      Strip symbols from binaries/libraries
#-- docs:       Save doc directories specified by DOC_DIRS
#-- libtool:    Leave libtool (.la) files in packages
#-- staticlibs: Leave static library (.a) files in packages
#-- emptydirs:  Leave empty directories in packages
#-- zipman:     Compress manual (man and info) pages in MAN_DIRS with gzip
#-- purge:      Remove files specified by PURGE_TARGETS
#-- debug:      Add debugging flags as specified in DEBUG_* variables
#-- lto:        Add compile flags for building with link time optimization
#-- autodeps:   Automatically add depends/provides
#
OPTIONS=(strip docs !libtool !staticlibs emptydirs zipman purge !debug lto)

#-- File integrity checks to use. Valid: md5, sha1, sha224, sha256, sha384, sha512, b2
INTEGRITY_CHECK=(sha256)
#-- Options to be used when stripping binaries. See "man strip" for details.
STRIP_BINARIES="--strip-all"
#-- Options to be used when stripping shared libraries. See "man strip" for details.
STRIP_SHARED="--strip-unneeded"
#-- Options to be used when stripping static libraries. See "man strip" for details.
STRIP_STATIC="--strip-debug"
#-- Manual (man and info) directories to compress (if zipman is specified)
MAN_DIRS=({usr{,/local}{,/share},opt/*}/{man,info})
#-- Doc directories to remove (if !docs is specified)
DOC_DIRS=(usr/{,local/}{,share/}{doc,gtk-doc} opt/*/{doc,gtk-doc})
#-- Files to be removed from all packages (if purge is specified)
PURGE_TARGETS=(usr/{,share}/info/dir .packlist *.pod)
#-- Directory to store source code in for debug packages
DBGSRCDIR="/usr/src/debug"
#-- Prefix and directories for library autodeps
LIB_DIRS=('lib:usr/lib' 'lib32:usr/lib32')

#########################################################################
# PACKAGE OUTPUT
#########################################################################
#
# Default: put built package and cached source in build directory
#
#-- Destination: specify a fixed directory where all packages will be placed
#PKGDEST=/home/packages
#-- Source cache: specify a fixed directory where source files will be cached
#SRCDEST=/home/sources
#-- Source packages: specify a fixed directory where all src packages will be placed
#SRCPKGDEST=/home/srcpackages
#-- Log files: specify a fixed directory where all log files will be placed
#LOGDEST=/home/makepkglogs
#-- Packager: name/email of the person or organization building packages
#PACKAGER="John Doe <john@doe.com>"
#-- Specify a key to use for package signing
#GPGKEY=""

#########################################################################
# COMPRESSION DEFAULTS
#########################################################################
#
COMPRESSGZ=(gzip -c -f -n)
COMPRESSBZ2=(bzip2 -c -f)
COMPRESSXZ=(xz -c -z -e6 --threads=0 -)
COMPRESSZST=(zstd -c -z -q --threads=0 -)
COMPRESSLRZ=(lrzip -q)
COMPRESSLZO=(lzop -q)
COMPRESSZ=(compress -c -f)
COMPRESSLZ4=(lz4 -q)
COMPRESSLZ=(lzip -c -f)

#########################################################################
# EXTENSION DEFAULTS
#########################################################################
#
PKGEXT='.pkg.tar.xz'
SRCEXT='.src.tar.gz'

#########################################################################
# OTHER
#########################################################################
#
#-- Command used to run pacman as root, instead of trying sudo and su
PACMAN_AUTH=(doas)
`

const makepkgRS = `#!/hint/bash
#
# /etc/makepkg.conf.d/rust.conf - rwinkhart/cachyos-postinstall-helper 09/28/2025
#

#########################################################################
# RUST LANGUAGE SUPPORT
#########################################################################

# Flags used for the Rust compiler, similar in spirit to CFLAGS. Read
# linkman:rustc[1] for more details on the available flags.
RUSTFLAGS="-Cforce-frame-pointers=yes -C opt-level=3 -C target-cpu=native"

# Additional compiler flags appended to "RUSTFLAGS" for use in debugging.
# Read linkman:rustc[1] for more details on the available flags.
DEBUG_RUSTFLAGS="-C debuginfo=2"
`

const pacman = `# /etc/pacman.conf - rwinkhart/cachyos-postinstall-helper 09/28/2025

# See the pacman.conf(5) manpage for option and repository directives

# GENERAL OPTIONS
[options]
# The following paths are commented out with their default values listed.
# If you wish to use different paths, uncomment and update the paths.
#RootDir     = /
#DBPath      = /var/lib/pacman/
#CacheDir    = /var/cache/pacman/pkg/
#LogFile     = /var/log/pacman.log
#GPGDir      = /etc/pacman.d/gnupg/
#HookDir     = /etc/pacman.d/hooks/
HoldPkg     = pacman glibc
#XferCommand = /usr/bin/curl -L -C - -f -o %o %u
#XferCommand = /usr/bin/wget --passive-ftp -c -O %o %u
#CleanMethod = KeepInstalled
#UseDelta    = 0.7
Architecture = auto

# Pacman won't upgrade packages listed in IgnorePkg and members of IgnoreGroup
#IgnorePkg   =
#IgnoreGroup =

#NoUpgrade   =
#NoExtract   =

# Misc options
#UseSyslog
Color
ILoveCandy
#NoProgressBar
CheckSpace
#VerbosePkgLists
DisableDownloadTimeout
ParallelDownloads = 10
DownloadUser = alpm
#DisableSandbox

# By default, pacman accepts packages signed by keys that its local keyring
# trusts (see pacman-key and its man page), as well as unsigned packages.
SigLevel    = Required DatabaseOptional
LocalFileSigLevel = Optional
#RemoteFileSigLevel = Required

# NOTE: You must run "pacman-key --init" before first using pacman; the local
# keyring can then be populated with the keys of all official Arch Linux
# packagers with "pacman-key --populate archlinux".

#
# REPOSITORIES
#   - can be defined here or included from another file
#   - pacman will search repositories in the order defined here
#   - local/custom mirrors can be added here or in separate files
#   - repositories listed first will take precedence when packages
#     have identical names, regardless of version number
#   - URLs will have $repo replaced by the name of the current repo
#   - URLs will have $arch replaced by the name of the architecture
#
# Repository entries are of the format:
#       [repo-name]
#       Server = ServerName
#       Include = IncludePath
#
# The header [repo-name] is crucial - it must be present and
# uncommented to enable the repo.
#

# cachyos repos
[cachyos-v3]
Include = /etc/pacman.d/cachyos-v3-mirrorlist

[cachyos-core-v3]
Include = /etc/pacman.d/cachyos-v3-mirrorlist

[cachyos-extra-v3]
Include = /etc/pacman.d/cachyos-v3-mirrorlist

[cachyos]
Include = /etc/pacman.d/cachyos-mirrorlist

# The testing repositories are disabled by default. To enable, uncomment the
# repo name header and Include lines. You can add preferred servers immediately
# after the header, and they will be used before the default mirrors.

#[core-testing]
#Include = /etc/pacman.d/mirrorlist

[core]
Include = /etc/pacman.d/mirrorlist

#[extra-testing]
#Include = /etc/pacman.d/mirrorlist

[extra]
Include = /etc/pacman.d/mirrorlist

# If you want to run 32 bit applications on your x86_64 system,
# enable the multilib repositories as required here.

#[multilib-testing]
#Include = /etc/pacman.d/mirrorlist

[multilib]
Include = /etc/pacman.d/mirrorlist

# An example of a custom package repository.  See the pacman manpage for
# tips on creating your own repositories.
#[custom]
#SigLevel = Optional TrustAll
#Server = file:///home/custompkgs
`

const pacmanHookClean = `[Trigger]
Operation = Remove
Operation = Install
Operation = Upgrade
Type = Package
Target = *

[Action]
Description = Keep the last cache and the currently installed.
When = PostTransaction
Exec = /usr/bin/paccache -rvk2
`

const pacmanHookNvidia = `[Trigger]
Operation=Install
Operation=Upgrade
Operation=Remove
Type=Package
Target=nvidia-open
Target=linux-cachyos

[Action]
Description=Update Nvidia module in initcpio
Depends=mkinitcpio
When=PostTransaction
NeedsTargets
Exec=/bin/sh -c 'while read -r trg; do case $trg in linux) exit 0; esac; done; /usr/bin/mkinitcpio -P'
`
