package main

import (
	"os"
	"regexp"
	"upay/api/v2/src/connections"
	"upay/api/v2/src/models"
)

func main() {
	args := os.Args
	runserver, _ := regexp.Compile(`runserver`)
	mirate, _ := regexp.Compile(`migrate`)

	if runserver.MatchString(args[1]) {
		connections.RunServer()
	} else if mirate.MatchString(args[1]) {
		models.Migrations()
	}
}
