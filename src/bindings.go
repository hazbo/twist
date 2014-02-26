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
    "../vendor/gocui"
    "log"
)

func Keybindings(g *gocui.Gui) error {
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

    if err := g.SetKeybinding("", gocui.KeyBackspace, 0, Backspace); err != nil {
        return err
    }

    if err := g.SetKeybinding("", gocui.KeyBackspace2, 0, Backspace); err != nil {
        return err
    }

    // Editor key enter
    if err := g.SetKeybinding("main", gocui.KeyEnter, 0, EditorKeyEnter); err != nil {
        log.Panicln(err)
    }

    // Editor key backspace
    if err := g.SetKeybinding("main", gocui.KeyBackspace, 0, EditorKeyBackspace); err != nil {
        log.Panicln(err)
    }

    if err := g.SetKeybinding("main", gocui.KeyBackspace2, 0, EditorKeyBackspace); err != nil {
        log.Panicln(err)
    }

    // New
    if err := g.SetKeybinding("", gocui.KeyCtrlN, 0, New); err != nil {
        log.Panicln(err)
    }

    // Write Dialog
    if err := g.SetKeybinding("main", gocui.KeyCtrlW, 0, ShowWriteDialog); err != nil {
        log.Panicln(err)
    }

    // Quit
    if err := g.SetKeybinding("", gocui.KeyCtrlC, 0, Quit); err != nil {
        log.Panicln(err)
    }

    // Switch to the JS console
    if err := g.SetKeybinding("", gocui.KeyCtrlJ, 0, UseConsole); err != nil {
        log.Panicln(err)
    }

    // Switch to the editor
    if err := g.SetKeybinding("console", gocui.KeyCtrlK, 0, UseEditor); err != nil {
        log.Panicln(err)
    }

    // Execute JS console command
    if err := g.SetKeybinding("console", gocui.KeyEnter, 0, OttoExecute); err != nil {
        log.Panicln(err)
    }

    // Hide write dialog
    if err := g.SetKeybinding("dialog-write", gocui.KeyEnter, 0, HideWriteDialog); err != nil {
        log.Panicln(err)
    }

    // Exit write dialog
    if err := g.SetKeybinding("dialog-write", gocui.KeyCtrlQ, 0, ExitWriteDialog); err != nil {
        log.Panicln(err)
    }
    return nil
}
