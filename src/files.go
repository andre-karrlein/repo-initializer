package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func initFiles() error {
	fmt.Println("Create files...")
	for _, file := range getDefaultFiles() {
		fmt.Print(file)
		if checkIfPhpFile(file) {
			err := createFile(file, "<?php \n\ndeclare(strict_types=1);")
			if err != nil {
				return err
			}
		} else {
			err := createFile(file, "")
			if err != nil {
				return err
			}
		}
		fmt.Println("...Done")
	}
	return nil
}

func checkIfPhpFile(file string) bool {
	if file[len(file)-3:] != "php" {
		return false
	}
	return true
}

func createFile(name, content string) error {
	err := ioutil.WriteFile(name, []byte(content), 0755)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to write file: %v", err))
	}
	return nil
}
