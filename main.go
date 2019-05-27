package main

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("on.ico"))
	if updateTitle {
		systray.SetTitle("Hosts are blocked")
	}
	systray.SetTooltip("Hosts are blocked")

	mToggle := systray.AddMenuItem("Toggle", "Toggles hosts block")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			select {
			case <-systray.ClickedCh:
				toggle()

			case <-mToggle.ClickedCh:
				toggle()

			case <-mQuit.ClickedCh:
				os.Exit(0)
			}
		}
	}()

	updBlocked(isBlocked())
}

func onExit() {
}

func getIcon(s string) []byte {
	f, _ := assets.Open(s)
	b, _ := ioutil.ReadAll(f)
	return b
}

func toggle() {
	h, err := ioutil.ReadFile(hostsPath)
	if err != nil {
		systray.SetTooltip(err.Error())
		return
	}

	blocked := false
	if bytes.Contains(h, []byte("# etchoststoggle")) {
		blocked = true
		h = bytes.Replace(h, []byte("\n# etchoststoggle 127."), []byte("\n127."), -1)
	} else {
		h = bytes.Replace(h, []byte("\n127."), []byte("\n# etchoststoggle 127."), -1)
	}

	err = ioutil.WriteFile(hostsPath, h, 0640)
	if err != nil {
		systray.SetTooltip(err.Error())
		return
	}
	updBlocked(blocked)
}

func updBlocked(blocked bool) {
	if blocked {
		systray.SetIcon(getIcon("on.ico"))
		if updateTitle {
			systray.SetTitle("Hosts are blocked")
		}
		systray.SetTooltip("Hosts are blocked")
	} else {
		systray.SetIcon(getIcon("off.ico"))
		if updateTitle {
			systray.SetTitle("Hosts are not blocked")
		}
		systray.SetTooltip("Hosts are not blocked")
	}
}

func isBlocked() bool {
	h, err := ioutil.ReadFile(hostsPath)
	if err != nil {
		systray.SetTooltip(err.Error())
		return true
	}

	return !bytes.Contains(h, []byte("# etchoststoggle"))
}
