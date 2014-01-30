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
	
	if (!FilenameIsSet()) {
		if v, err := g.SetView("intro", 5, -1, maxX, maxY); err != nil {
	        if err != gocui.ErrorUnkView {
	            return err
	        }
	        fmt.Fprintf(v, "%s",`    `)

	        if err := g.SetCurrentView("intro"); err != nil {
	            return err
	        }
		}
	} else {
		if v, err := g.SetView("main", 5, -1, maxX, maxY); err != nil {
		    if err != gocui.ErrorUnkView {
		        return err
		    }

		    v.Editable = true

		    fmt.Fprintf(v, "%s", GetFileContents(GetFilenameArg()))

		    if err := g.SetCurrentView("main"); err != nil {
		        return err
		    }
		}
	}
	return nil
}