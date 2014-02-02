package main

import (
    "../vendor/gocui"
)

func cursorDown(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        cx, cy := v.Cursor()
        if err := v.SetCursor(cx, cy+1); err != nil {
            ox, oy := v.Origin()
            if err := v.SetOrigin(ox, oy+1); err != nil {
                return err
            }
        }
    }
    return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        ox, oy := v.Origin()
        cx, cy := v.Cursor()
        if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
            if err := v.SetOrigin(ox, oy-1); err != nil {
                    return err
            }
        }
    }
    return nil
}

func cursorLeft(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        ox, oy := v.Origin()
        cx, cy := v.Cursor()
        if err := v.SetCursor(cx-1, cy); err != nil && ox > 0 {
            if err := v.SetOrigin(ox-1, oy); err != nil {
                return err
            }
        }
    }
    return nil
}

func cursorRight(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        cx, cy := v.Cursor()
        if err := v.SetCursor(cx+1, cy); err != nil {
            ox, oy := v.Origin()
            if err := v.SetOrigin(ox+1, oy); err != nil {
                return err
            }
        }
    }
    return nil
}

func cursorTab(g *gocui.Gui, v *gocui.View) error {
    // Insert spaces
    return nil
}
