package main

import (
    "../vendor/gocui"
)

func reset(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("main", 5, -1, maxX, maxY - 2); err != nil {
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

func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrorQuit
}

func console(g *gocui.Gui, v *gocui.View) error {

    if err := g.SetCurrentView("console"); err != nil {
        return err
    }


    return nil
}