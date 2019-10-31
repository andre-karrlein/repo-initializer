package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	for _, dir := range getDefaultDirectories() {
		_ = createFolder(dir)
	}

}

func createFiles(name string) error {
	return nil
}

func createFolder(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		return errors.New(fmt.Sprintf("Can not create directory with error: %s", err.Error()))
	}

	return nil
}

func getDefaultDirectories() []string {
	return []string{
		"config",
		"docker",
		"docker/nginx",
		"docker/php-fpm",
		"public",
		"src",
		"test",
	}
}

func getDefaultFiles() []string {
	return []string{
		"",
	}
}
