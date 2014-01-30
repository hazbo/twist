package main

import (
	"fmt"
	"../vendor/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, 5, maxY); err != nil {
	    if err != gocui.ErrorUnkView {
	        return err
	    }
	    
	    fmt.Fprintln(v, "- 1")
	}

	if v, err := g.SetView("intro", 5, -1, maxX, maxY); err != nil {
        if err != gocui.ErrorUnkView {
            return err
        }
        fmt.Fprintf(v, "%s",`
        	Welcome to GoEdit v0.0.1

        	To start: Ctrl+R (reset)
		`)

        if err := g.SetCurrentView("intro"); err != nil {
            return err
        }
	}

	if v, err := g.SetView("console", -1, maxY - 2, maxX, maxY); err != nil {
	    if err != gocui.ErrorUnkView {
	        return err
	    }
	    fmt.Fprint(v, "jsc >> ")
	    v.Editable = true
	}
	return nil
}