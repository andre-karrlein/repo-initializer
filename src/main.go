package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Create directories...")
	for _, dir := range getDefaultDirectories() {
		fmt.Print(dir)
		err := createFolder(dir)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("...Done")

	}

	fmt.Println("Create files...")
	for _, file := range getDefaultFiles() {
		fmt.Print(file)
		if checkIfPhpFile(file) {
			err := createFile(file, "<?php \n\ndeclare(strict_types=1);")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := createFile(file, "")
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("...Done")
	}

	fmt.Println("You are ready to code, everything is setup")
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

func createFolder(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		return errors.New(fmt.Sprintf("Can not create directory with error: %v", err))
	}

	return nil
}
