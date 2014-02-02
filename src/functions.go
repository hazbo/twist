package main

import (
    "../vendor/gocui"
)

func Write(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func New(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("main", 5, -1, maxX, maxY); err != nil {
	    if err != gocui.ErrorUnkView {
	        return err
	    }

	    v.Editable = true

	    if err := g.SetCurrentView("main"); err != nil {
	        return err
	    }
	}
	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrorQuit
}
