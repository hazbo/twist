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
    "../vendor/otto"
    "../vendor/gocui"
)

func OttoConfigure(g *gocui.Gui) {

    Otto := g.Otto();

    // Print
    Otto.Set("print", func(call otto.FunctionCall) otto.Value {
        fmt.Fprintf(g.View("main"), "%s", call.Argument(0).String())
        return otto.UndefinedValue()
    })

    // Write
    Otto.Set("write", func(call otto.FunctionCall) otto.Value {
        WriteFile(call.Argument(0).String(), []byte(g.View("main").Buffer()))
        return otto.UndefinedValue()
    })

    // Load in twist.js
    js_source, err := Asset("src/js/twist.js")
    if (err == nil) {
        if (len(js_source) > 0) {
            js_source_string := string(js_source[:])
            Otto.Run(js_source_string)
        }
    }

    // Load in twist.js
    js_php_source, err := Asset("src/js/lang/php.js")
    if (err == nil) {
        if (len(js_php_source) > 0) {
            js_source_string := string(js_php_source[:])
            Otto.Run(js_source_string)
        }
    }

}

func OttoExecute(g *gocui.Gui, v *gocui.View) error {
    command_buffer := g.View("console").Buffer()
    raw_command    := strings.Replace(command_buffer, "jsc >> ", "", 1)

    if (raw_command != "") {

        Otto := g.Otto()
        Otto.Run(raw_command)
        v.Clear()
        fmt.Fprintf(v, "%s", "jsc >> ")

        // TODO: This should not skip to the next line
        cx, cy := v.Cursor()
        if err := v.SetCursor(cx + 7, cy ); err != nil {
            ox, oy := v.Origin()
            if err := v.SetOrigin(ox + 7, oy ); err != nil {
                return err
            }
        }
    }
    return nil
}