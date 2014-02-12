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
    "log"
    "../vendor/gocui"
    "../vendor/otto"
)

func main() {
    var err error
    g := gocui.NewGui()

    if err := g.Init(); err != nil {
        log.Panicln(err)
    }

    defer g.Close()

    g.SetLayout(Layout)

    g.SelBgColor = gocui.ColorWhite
    g.SelFgColor = gocui.ColorBlack
    g.ShowCursor = true

    g.SetOtto(otto.New())

    OttoConfigure(g)

    if err := Keybindings(g); err != nil {
        log.Panicln(err)
    }

    err = g.MainLoop()
    if err != nil && err != gocui.ErrorQuit {
        log.Panicln(err)
    }
}
