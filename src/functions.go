/**
 * Twist v0.1-dev
 *
 * (c) Harry Lawrence
 *
 * License: MIT
 * 
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package main

import (
    "fmt"
    "strings"
    "strconv"
    "../vendor/gocui"
)

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

func UseConsole(g *gocui.Gui, v *gocui.View) error {
    if err := g.SetCurrentView("console"); err != nil {
        return err
    }
    return nil
}

func Backspace(g *gocui.Gui, v *gocui.View) error {
    cx, cy := v.Cursor()
    spaceCount := 0

    var s []byte

    for i := 0; i < 5; i++ {
        rune, err := g.Rune(cx + (6 - i), cy)
        if (err == nil) {
            res := strconv.AppendQuoteRune(s, rune)
            if (string(res) == "' '") {
                spaceCount++
            }
        }

        if (spaceCount == 4) {
            for c := 1; c < 4; c++ {
                v.DeleteRune(cx - c , cy)
                v.SetCursor(cx - c, cy)
            }
        }
    }
    return nil
}

func UseEditor(g *gocui.Gui, v *gocui.View) error {
    if err := g.SetCurrentView("main"); err != nil {
        return err
    }
    return nil
}

func ShowWriteDialog(g *gocui.Gui, v *gocui.View) error {
    maxX, maxY := g.Size()
    if v, err := g.SetView("dialog-write", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
        if err != gocui.ErrorUnkView {
            return err
        }
        v.Editable = true
        fmt.Fprintf(v, "jsc >> editor.write('%s');", g.Filename())

        cursor_position := 21 + strings.Index(g.Filename(), ".")
        if (cursor_position == 20) {
            cursor_position++
        }

        cx, cy := v.Cursor()
        if err := v.SetCursor(cx + cursor_position, cy); err != nil {
            ox, oy := v.Origin()
            if err := v.SetOrigin(ox + cursor_position, oy); err != nil {
                return err
            }
        }

        if err := g.SetCurrentView("dialog-write"); err != nil {
            return err
        }
    }
    return nil
}

func HideWriteDialog(g *gocui.Gui, v *gocui.View) error {
    g.View("console").Clear()
    fmt.Fprint(g.View("console"), v.Buffer())
    OttoExecute(g, g.View("console"))
    g.DeleteView("dialog-write")
    g.SetCurrentView("main")
    return nil
}

func ExitWriteDialog(g *gocui.Gui, v *gocui.View) error {
    g.DeleteView("dialog-write")
    g.SetCurrentView("main")
    return nil
}
