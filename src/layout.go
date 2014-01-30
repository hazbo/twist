package main

import (
	"fmt"
	"../vendor/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, 5, maxY); err != nil {
	    if err != gocui.ErrorUnkView {
	        return err
	    }
	    
	    fmt.Fprintln(v, "- 1")
	}

	if v, err := g.SetView("intro", 5, -1, maxX, maxY); err != nil {
        if err != gocui.ErrorUnkView {
            return err
        }
        fmt.Fprintf(v, "%s",`


			        GGGGGGGGGGGGG                 EEEEEEEEEEEEEEEEEEEEEE
			     GGG::::::::::::G                 E::::::::::::::::::::E
			   GG:::::::::::::::G                 E::::::::::::::::::::E
			  G:::::GGGGGGGG::::G                 EE::::::EEEEEEEEE::::E
			 G:::::G       GGGGGG   ooooooooooo     E:::::E       EEEEEE
			G:::::G               oo:::::::::::oo   E:::::E             
			G:::::G              o:::::::::::::::o  E::::::EEEEEEEEEE   
			G:::::G    GGGGGGGGGGo:::::ooooo:::::o  E:::::::::::::::E   
			G:::::G    G::::::::Go::::o     o::::o  E:::::::::::::::E   
			G:::::G    GGGGG::::Go::::o     o::::o  E::::::EEEEEEEEEE   
			G:::::G        G::::Go::::o     o::::o  E:::::E             
			 G:::::G       G::::Go::::o     o::::o  E:::::E       EEEEEE
			  G:::::GGGGGGGG::::Go:::::ooooo:::::oEE::::::EEEEEEEE:::::E
			   GG:::::::::::::::Go:::::::::::::::oE::::::::::::::::::::E
			     GGG::::::GGG:::G oo:::::::::::oo E::::::::::::::::::::E
			        GGGGGG   GGGG   ooooooooooo   EEEEEEEEEEEEEEEEEEEEEE

			Version 0.0.1

		`)

        if err := g.SetCurrentView("intro"); err != nil {
            return err
        }
	}
	return nil
}