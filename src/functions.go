package main

import (
    "../vendor/gocui"
    "../vendor/otto"
)

func Write(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func New(g *gocui.Gui, v *gocui.View) error {
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

func Quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrorQuit
}

func console(g *gocui.Gui, v *gocui.View) error {
    if err := g.SetCurrentView("console"); err != nil {
        return err
    }
    return nil
}

func testjs(g *gocui.Gui, v *gocui.View) error {
	Otto := otto.New()

	Otto.Run(`
		console.log('Hello World');
	`)
	return nil
}
