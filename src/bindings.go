package main

import (
	"../vendor/gocui"
)

func keybindings(g *gocui.Gui) error {
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
    return nil
}