package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

type Workspace struct {
	Name    string
	Focused bool
}

type KeyConfig struct {
	Key   string
	Color string
}

var (
	debug           = kingpin.Flag("debug", "Enable debug mode.").Bool()
	colorDefault    = kingpin.Flag("colorDefault", "Default color used by number-keys.").Default("009696").String()
	colorWorkspace  = kingpin.Flag("colorWorkspace", "Color used by number keys that represent a workspace.").Default("ff8800").String()
	colorFocused    = kingpin.Flag("colorFocused", "Color used by number keys representing the workspace that is focused.").Default("ff0000").String()
	keyboardModel   = kingpin.Flag("keyboardModel", "Keyboard model. Supported models are everything available in g810-led").Default("g513").String()
	keyboardCommand = "g513-led"
	wm              = kingpin.Flag("wm", "Window Manager. Supported Window Managers are i3 and sway.").HintOptions("sway", "i3").Default("sway").String()
	wmCommand       = "swaymsg"
)

// DoEvery - Call function in passed Interval
func DoEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// func helloworld(t time.Time) {
func setKeyColors() {
	// Get Sway Workspaces
	cmd := exec.Command(wmCommand, "-t", "get_workspaces")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with %s\n", err)
	}

	// Parse JSON
	var workspaces []Workspace
	json.Unmarshal([]byte(out), &workspaces)

	// Generate configs for used workspaces
	var keyConfigs []KeyConfig
	var usedWorkspaces []string
	for _, element := range workspaces {
		usedWorkspaces = append(usedWorkspaces, element.Name)
		var tmp KeyConfig
		tmp.Key = element.Name
		if element.Focused {
			tmp.Color = *colorFocused
		} else {
			tmp.Color = *colorWorkspace
		}
		keyConfigs = append(keyConfigs, tmp)
	}

	// Generate config for unused workspaces
	for i := 1; i <= 10; i++ {
		var name = strconv.Itoa(i)
		_, found := Find(usedWorkspaces, name)
		if !found {
			var tmp KeyConfig
			if i == 10 {
				tmp.Key = "0"
			} else {
				tmp.Key = name
			}
			tmp.Color = *colorDefault
			keyConfigs = append(keyConfigs, tmp)
		}
	}

	// fmt.Printf("combined out: %+v\n", keyConfigs)

	// Set key colors
	for _, element := range keyConfigs {
		cmd := exec.Command(keyboardCommand, "-k", element.Key, element.Color)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.CombinedOutput() failed with %s\n", err)
		}
	}
}

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	// Set correct keyboardCommand
	if *keyboardModel != "g513" {
		keyboardCommand = fmt.Sprintf("%s-led", *keyboardModel)
	}

	// Set correct wmCommand
	if *wm == "i3" {
		wmCommand = "i3-msg"
	}

	if *debug {
		fmt.Printf(
			"colorDefault: %s, colorWorkspace: %s, colorFocused: %s, keyboardModel: %s, keyboardCommand: %s, wm: %s, wmCommand: %s",
			*colorDefault,
			*colorWorkspace,
			*colorFocused,
			*keyboardModel,
			keyboardCommand,
			*wm,
			wmCommand)
	}

	setKeyColors()
}
