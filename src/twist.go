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
    "os"
    "os/user"
    "strings"
    "io/ioutil"
    "../vendor/gocui"
    "../vendor/otto"
)

func main() {
    // Concept code, abstract this out
    if (len(os.Args) > 1 && os.Args[1] == "run") {
        usr, err := user.Current()
        if err != nil {
            log.Fatal( err )
        }
        filename  := os.Args[2];
        clean_path_filename := strings.Replace(filename, ".", "/", -1)

        full_path := usr.HomeDir + "/.twist/" + clean_path_filename + ".js"

        if _, err := os.Stat(full_path); err == nil {
            g := gocui.NewGui()
            g.SetOtto(otto.New())
            OttoConfigure(g)

            file_contents, err := ioutil.ReadFile(full_path)
            if (err == nil) {
                g.Otto().Run(string(file_contents))
            }
            defer g.Close()
        }
    } else {
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
}
