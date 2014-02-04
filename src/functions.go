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

        cx, cy := v.Cursor()
        if err := v.SetCursor(cx + 21, cy); err != nil {
            ox, oy := v.Origin()
            if err := v.SetOrigin(ox + 21, oy); err != nil {
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
    Execute(g, g.View("console"))
    g.DeleteView("dialog-write")
    g.SetCurrentView("main")
    return nil
}
