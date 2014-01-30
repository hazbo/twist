package main

import (
	"../vendor/gocui"
    "log"
)

func keybindings(g *gocui.Gui, b *EditorBuffer) error {
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

    // Line count related functions
    if err := g.SetKeybinding("main", gocui.KeyEnter, 0, setLineCount); err != nil {
        log.Panicln(err)
    }

    // Reset
    if err := g.SetKeybinding("", gocui.KeyCtrlR, 0, reset); err != nil {
        log.Panicln(err)
    }

    return nil
}