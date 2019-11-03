package main

import (
	"errors"
	"fmt"
)

func initComposer() error {
	err := getComposer()
	if err != nil {
		return errors.New(fmt.Sprintf("Can not download composer phar: %v", err))
	}
	return nil
}

func getComposer() error {
	fileUrl := "https://getcomposer.org/download/1.9.1/composer.phar"
	return DownloadFile("composer.phar", fileUrl)
}
