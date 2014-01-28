package main

import (
	"fmt"
	"log"
	//"io/ioutil"
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
    return nil
}

func layout(g *gocui.Gui) error {
        maxX, maxY := g.Size()
        if v, err := g.SetView("side", -1, -1, 30, maxY); err != nil {
                if err != gocui.ErrorUnkView {
                        return err
                }
                v.Highlight = true
                fmt.Fprintln(v, "Item 1")
                fmt.Fprintln(v, "Item 2")
                fmt.Fprintln(v, "Item 3")
                fmt.Fprintln(v, "Item 4")
        }
        if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
                if err != gocui.ErrorUnkView {
                        return err
                }
                fmt.Fprintf(v, "%s", "Hello World")
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

func main() {
    var err error
    g := gocui.NewGui()
    if err := g.Init(); err != nil {
        log.Panicln(err)
    }
    defer g.Close()
    g.SetLayout(layout)
    if err := g.SetKeybinding("", gocui.KeyCtrlC, 0, quit); err != nil {
        log.Panicln(err)
    }

    err = g.MainLoop()
    if err != nil && err != gocui.ErrorQuit {
        log.Panicln(err)
    }

        g.SetLayout(layout)
        if err := keybindings(g); err != nil {
                log.Panicln(err)
        }
        g.SelBgColor = gocui.ColorGreen
        g.SelFgColor = gocui.ColorBlack
        g.ShowCursor = true
}