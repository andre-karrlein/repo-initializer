package main

import (
	"fmt"
)

func main() {
	err := initFolders()
	if err != nil {
		fmt.Println(err)
	}
	err = initFiles()
	if err != nil {
		fmt.Println(err)
	}
	err = initComposer()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You are ready to code, everything is setup")
}
