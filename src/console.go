package main

import (
	"fmt"
	"strings"
	"../vendor/otto"
	"../vendor/gocui"
)

func Execute(g *gocui.Gui, v *gocui.View) error {
	
	command_buffer := g.View("console").Buffer()
	raw_command    := strings.Replace(command_buffer, "jsc >> ", "", 1)

	if (raw_command != "") {

		Otto := otto.New()

		// Testing a 'sayHello' function
		Otto.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		    fmt.Fprintf(g.View("main"), "Hello, %s.\n", call.Argument(0).String())
		    return otto.UndefinedValue()
		})

		//Otto.Run(WrapJavascript())

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