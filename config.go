package main

import (
	"github.com/aaronjanse/3mux/wm"
)

// Config stores all user configuration values
type Config struct {
	statusBar bool
	bindings  map[string]func(*wm.Universe)
}

var configFuncBindings = map[string]func(*wm.Universe){
	"newPane":      func(u *wm.Universe) { u.AddPane() },
	"newPaneHoriz": func(u *wm.Universe) { u.AddPaneTmux(false) },
	"newPaneVert":  func(u *wm.Universe) { u.AddPaneTmux(true) },

	"killWindow": func(u *wm.Universe) { u.KillPane() },
	"fullscreen": func(u *wm.Universe) { u.ToggleFullscreen() },
	"search":     func(u *wm.Universe) { u.ToggleSearch() },

	"resize(Up)":    func(u *wm.Universe) { u.ResizePane(wm.Up) },
	"resize(Down)":  func(u *wm.Universe) { u.ResizePane(wm.Down) },
	"resize(Left)":  func(u *wm.Universe) { u.ResizePane(wm.Left) },
	"resize(Right)": func(u *wm.Universe) { u.ResizePane(wm.Right) },

	"moveWindow(Up)":    func(u *wm.Universe) { u.MoveWindow(wm.Up) },
	"moveWindow(Down)":  func(u *wm.Universe) { u.MoveWindow(wm.Down) },
	"moveWindow(Left)":  func(u *wm.Universe) { u.MoveWindow(wm.Left) },
	"moveWindow(Right)": func(u *wm.Universe) { u.MoveWindow(wm.Right) },

	"moveSelection(Up)":    func(u *wm.Universe) { u.MoveSelection(wm.Up) },
	"moveSelection(Down)":  func(u *wm.Universe) { u.MoveSelection(wm.Down) },
	"moveSelection(Left)":  func(u *wm.Universe) { u.MoveSelection(wm.Left) },
	"moveSelection(Right)": func(u *wm.Universe) { u.MoveSelection(wm.Right) },

	"cycleSelection(Forward)":  func(u *wm.Universe) { u.CycleSelection(true) },
	"cycleSelection(Backward)": func(u *wm.Universe) { u.CycleSelection(false) },

	"splitPane(Vertical)":   func(u *wm.Universe) { u.AddPaneTmux(true) },
	"splitPane(Horizontal)": func(u *wm.Universe) { u.AddPaneTmux(false) },
}

func compileBindings(sourceBindings map[string][]string) map[string]func(*wm.Universe) {
	compiledBindings := map[string]func(*wm.Universe){}
	for funcName, keyCodes := range sourceBindings {
		fn, ok := configFuncBindings[funcName]
		if !ok {
			panic("Incorrect keybinding: " + funcName)
		}
		for _, keyCode := range keyCodes {
			compiledBindings[keyCode] = fn
		}
	}

	return compiledBindings
}

var config = Config{
	statusBar: true,
}

func init() {
	config.bindings = compileBindings(map[string][]string{
		"newPane":      []string{`Alt+N`, `Alt+Enter`},
		"newPaneHoriz": []string{`Alt+B "`},
		"newPaneVert":  []string{`Alt+B %`},

		"killWindow": []string{`Alt+Shift+Q`},
		"fullscreen": []string{`Alt+Shift+F`},
		"search":     []string{`Alt+/`},

		"resize(Up)":    []string{`Alt+R Up`},
		"resize(Down)":  []string{`Alt+R Down`},
		"resize(Left)":  []string{`Alt+R Left`},
		"resize(Right)": []string{`Alt+R Right`},

		"moveWindow(Up)":    []string{`Alt+Shift+K`, `Alt+Shift+Up`},
		"moveWindow(Down)":  []string{`Alt+Shift+J`, `Alt+Shift+Down`},
		"moveWindow(Left)":  []string{`Alt+Shift+H`, `Alt+Shift+Left`, `Ctrl+B {`},
		"moveWindow(Right)": []string{`Alt+Shift+L`, `Alt+Shift+Right`, `Ctrl+B }`},

		"moveSelection(Up)":    []string{`Alt+K`, `Alt+Up`},
		"moveSelection(Down)":  []string{`Alt+J`, `Alt+Down`},
		"moveSelection(Left)":  []string{`Alt+H`, `Alt+Left`},
		"moveSelection(Right)": []string{`Alt+L`, `Alt+Right`},
	})
}

func seiveConfigEvents(u *wm.Universe, human string) bool {
	if fn, ok := config.bindings[human]; ok {
		fn(u)
		return true
	}
	return false
}
