package main

import (
	"strings"
)

func cosmic() {
	version := "/v1/"
	filenamesToValues := map[string]string{
		"AppList" + version + "enable_drag_source":  "true",
		"AppList" + version + "favorites":           "[\n    com.system76.CosmicFiles\n]",
		"Comp" + version + "autotile_behavior":      "PerWorkspace",
		"Comp" + version + "cursor_follows_focus":   "false",
		"Comp" + version + "focus_follows_cursor":   "true",
		"Comp" + version + "input_default":          "(\n    state: Enabled,\n    acceleration: Some((\n        profile: Some(Flat),\n        speed: -0.10288643756187243,\n    )),\n)",
		"Comp" + version + "workspaces":             "(\n    workspace_mode: Global,\n    workspace_layout: Vertical,\n)",
		"Panel" + version + "entries":               "[\n    \"Panel\",\n]",
		"Panel.Panel" + version + "anchor_gap":      "true",
		"Panel.Panel" + version + "border_radius":   "10",
		"Panel.Panel" + version + "exclusive_zone":  "true",
		"Panel.Panel" + version + "expand_to_edges": "false",
		"Panel.Panel" + version + "opacity":         "0.75",
		"Panel.Panel" + version + "plugins_center":  "Some([])",
		"Panel.Panel" + version + "plugins_wings":   "Some(([\n    \"com.system76.CosmicAppList\",\n], [\n    \"com.system76.CosmicAppletStatusArea\",\n    \"com.system76.CosmicAppletTiling\",\n    \"com.system76.CosmicAppletAudio\",\n    \"com.system76.CosmicAppletBluetooth\",\n    \"com.system76.CosmicAppletNetwork\",\n    \"com.system76.CosmicAppletBattery\",\n    \"com.system76.CosmicAppletNotifications\",\n    \"com.system76.CosmicAppletPower\",\n    \"com.system76.CosmicAppletTime\",\n]))",
		"Panel.Panel" + version + "size":            "XS",
		"Tk" + version + "show_maximize":            "false"}

	mkdir(mountPoint + "/home/" + getUsername() + "/.config/cosmic")
	for key, value := range filenamesToValues {
		filePath := mountPoint + "/home/" + getUsername() + "/.config/cosmic/com.system76.Cosmic" + key
		mkdir(filePath[:strings.LastIndex(filePath, "/")])
		writeFile(filePath, value+"\n", getUsername(), 0755)
	}

	// Replace SDDM with cosmic-greeter
	writeFile(mountPoint+"/etc/greetd/config.toml", "[terminal]\nvt = 1\n\n[default_session]\ncommand = \"cosmic-comp /bin/cosmic-greeter\"\nuser = \"cosmic-greeter\"\n", "root", 0644)
	manageService("stop", "sddm.service")
	manageService("disable", "sddm.service")
	manageService("enable", "greetd.service")
	managePackages("-Rcns", []string{"sddm"})
}
