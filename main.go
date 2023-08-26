package main

import (
	"fmt"

	yabai "github.com/irth/goyabai"
)

type WorkspaceManager struct {
}

func main() {
	y := yabai.Yabai{}
	displays, err := y.Displays()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hiiiiiiiii ;3")
	fmt.Println("You've got", len(displays), "displays")
}
