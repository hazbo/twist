package main

import (
	"fmt"
	"../vendor/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	/*
	if v, err := g.SetView("side", -1, -1, 30, maxY); err != nil {
	    if err != gocui.ErrorUnkView {
	        return err
	    }
	    v.Highlight = true

	    // Test items
	    //fmt.Fprintln(v, "Item 1")
	    //fmt.Fprintln(v, "Item 2")
	    //fmt.Fprintln(v, "Item 3")
	    //fmt.Fprintln(v, "Item 4")
	}
*/
	if v, err := g.SetView("main", 0, 0, maxX, maxY); err != nil {
        if err != gocui.ErrorUnkView {
            return err
        }
        fmt.Fprintf(v, "%s", "")
        v.Editable = true
        if err := g.SetCurrentView("main"); err != nil {
            return err
        }
	}
	return nil
}