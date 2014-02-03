package main

import (
	"fmt"
	"../vendor/gocui"
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	file 	   := File{name:""}

	if v, err := g.SetView("side", -1, -1, 5, maxY); err != nil {
	    if err != gocui.ErrorUnkView {
	        return err
	    }
	    
	    // Hide line numbers for now
	    //fmt.Fprintln(v, "- 1")
	    fmt.Fprintf(v, "%s", "")
	}

	if (!FilenameIsSet()) {
		if v, err := g.SetView("intro", 5, -1, maxX, maxY); err != nil {
	        if err != gocui.ErrorUnkView {
	            return err
	        }
	        fmt.Fprintf(v, "%s", `
	        	Welcome to GoEdit v0.0.1

	        	To start : Ctrl+N (new)
	        	To quit  : Ctrl+C
			`)

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

		    filename := GetFilenameArg()
		    file.SetName(filename)

		    fmt.Fprintf(v, "%s", GetFileContents(file.Name()))

		    if err := g.SetCurrentView("main"); err != nil {
		        return err
		    }
		}
	}
	return nil
}
