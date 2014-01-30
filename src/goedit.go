package main

import (
	"log"
	"../vendor/gocui"
)

func main() {
    var err error
    g := gocui.NewGui()
    b := new(EditorBuffer)

    if err := g.Init(); err != nil {
        log.Panicln(err)
    }

    defer g.Close()

    g.SetLayout(Layout)

    g.SelBgColor = gocui.ColorWhite
    g.SelFgColor = gocui.ColorBlack
    g.ShowCursor = true

    if err := Keybindings(g, b); err != nil {
        log.Panicln(err)
    }

    err = g.MainLoop()
    if err != nil && err != gocui.ErrorQuit {
        log.Panicln(err)
    }
}