package main

import (
	"../vendor/gocui"
    "log"
)

func Keybindings(g *gocui.Gui) error {
    if err := g.SetKeybinding("", gocui.KeyArrowDown, 0, cursorDown); err != nil {
		return err
    }
    if err := g.SetKeybinding("", gocui.KeyArrowUp, 0, cursorUp); err != nil {
		return err
    }
    if err := g.SetKeybinding("", gocui.KeyArrowLeft, 0, cursorLeft); err != nil {
		return err
    }
    if err := g.SetKeybinding("", gocui.KeyArrowRight, 0, cursorRight); err != nil {
		return err
    }
    if err := g.SetKeybinding("", gocui.KeyTab, 0, cursorTab); err != nil {
        return err
    }

    // New
    if err := g.SetKeybinding("", gocui.KeyCtrlN, 0, New); err != nil {
        log.Panicln(err)
    }

    // Write Dialog
    if err := g.SetKeybinding("main", gocui.KeyCtrlW, 0, ShowWriteDialog); err != nil {
        log.Panicln(err)
    }

    // Quit
    if err := g.SetKeybinding("", gocui.KeyCtrlC, 0, Quit); err != nil {
        log.Panicln(err)
    }

    // TestJS
    if err := g.SetKeybinding("", gocui.KeyCtrlP, 0, testjs); err != nil {
        log.Panicln(err)
    }

    // Console related bindings
    if err := g.SetKeybinding("", gocui.KeyCtrlJ, 0, console); err != nil {
        log.Panicln(err)
    }

    // Execute JS console command
    if err := g.SetKeybinding("console", gocui.KeyEnter, 0, Execute); err != nil {
        log.Panicln(err)
    }
    return nil
}
