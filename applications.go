package main

import (
	"os"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func applications() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Application to install:", []string{"BACK", "ALL", "pamac-aur", "yay-bin", "zsh", "htop", "neovim", "virt-manager (also installs qemu-desktop+libvirt)", "librewolf-bin", "zed", "steam"})
		switch choice {
		case 1:
			return
		case 2:
			doAll = true
			fallthrough
		case 3:
			managePackages("-S", []string{"pamac-aur"})

			if !doAll {
				break
			}
			fallthrough
		case 4:
			managePackages("-S", []string{"yay-bin"})

			if !doAll {
				break
			}
			fallthrough
		case 5:
			managePackages("-S", []string{"zsh", "zsh-autosuggestions", "zsh-syntax-highlighting"})
			writeFile(mountPoint+"/home/"+getUsername()+"/.zshrc", zshrc, getUsername(), 0644)
			err := os.RemoveAll(mountPoint + "/home/" + getUsername() + "/.bash_profile")
			if err != nil {
				other.PrintError("Failed to remove .bash_profile", 1)
			}
			err = os.RemoveAll(mountPoint + "/home/" + getUsername() + "/.bashrc")
			if err != nil {
				other.PrintError("Failed to remove .bashrc", 1)
			}

			if !doAll {
				break
			}
			fallthrough
		case 6:
			managePackages("-S", []string{"htop"})

			if !doAll {
				break
			}
			fallthrough
		case 7:
			managePackages("-S", []string{"neovim"})
			writeFile(mountPoint+"/etc/xdg/nvim/sysinit.vim", sysinitVim, "root", 0644)

			if !doAll {
				break
			}
			fallthrough
		case 8:
			managePackages("-S", []string{"virt-manager", "qemu-desktop", "libvirt"})
			manageService("enable", "libvirtd.service")
			manageService("start", "libvirtd.service")
			// TODO manage user groups!

			if !doAll {
				break
			}
			fallthrough
		case 9:
			managePackages("-S", []string{"librewolf-bin"})

			if !doAll {
				break
			}
			fallthrough
		case 10:
			managePackages("-S", []string{"zed"})
			mkdir(mountPoint + "/home/" + getUsername() + "/.config/zed")
			writeFile(mountPoint+"/home/"+getUsername()+"/.config/zed/settings.json", zed, getUsername(), 0600)

			if !doAll {
				break
			}
			fallthrough
		case 11:
			managePackages("-S", []string{"steam"})
		}
	}
}

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
