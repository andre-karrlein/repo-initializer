package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func initComposer() error {
	err := getComposer()
	if err != nil {
		return errors.New(fmt.Sprintf("Can not download composer phar: %v", err))
	}
	php := checkIfPhpBinAvailable()
	if !php {
		return nil
	}
	err = initializeComposer()
	if err != nil {
		return errors.New(fmt.Sprintf("Can not initialize composer: %v", err))
	}
	return nil
}

func getComposer() error {
	fileUrl := "https://getcomposer.org/download/1.9.1/composer.phar"
	return DownloadFile("composer.phar", fileUrl)
}

func checkIfPhpBinAvailable() bool {
	_, err := exec.LookPath("php")
	return err == nil
}

func initializeComposer() error {
	return errors.New("NO INIT METHOD")
}
