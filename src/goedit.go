package main

import (
	"log"
	"../vendor/gocui"
)

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

    g.SelBgColor = gocui.ColorWhite
    g.SelFgColor = gocui.ColorBlack
    g.ShowCursor = true

    if err := keybindings(g); err != nil {
            log.Panicln(err)
    }

    if err := g.SetKeybinding("", gocui.KeyCtrlC, 0, quit); err != nil {
        log.Panicln(err)
    }

    err = g.MainLoop()
    if err != nil && err != gocui.ErrorQuit {
        log.Panicln(err)
    }
}