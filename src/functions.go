package main

import (
	"fmt"
    "../vendor/gocui"
    "../vendor/otto"
)

func ShowWriteDialog(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("dialog-write", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		v.Editable = true
		fmt.Fprintln(v, "editor.write('');")

        cx, cy := v.Cursor()
        if err := v.SetCursor(cx + 14, cy); err != nil {
            ox, oy := v.Origin()
            if err := v.SetOrigin(ox + 14, oy); err != nil {
                return err
            }
        }

		if err := g.SetCurrentView("dialog-write"); err != nil {
			return err
		}
	}
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
