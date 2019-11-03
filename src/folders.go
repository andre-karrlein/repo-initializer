package main

import (
	"errors"
	"fmt"
	"os"
)

func initFolders() error {
	fmt.Println("Create directories...")
	for _, dir := range getDefaultDirectories() {
		fmt.Print(dir)
		err := createFolder(dir)
		if err != nil {
			return err
		}
		fmt.Println("...Done")
	}
	return nil
}

func createFolder(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		return errors.New(fmt.Sprintf("Can not create directory with error: %v", err))
	}

	return nil
}
