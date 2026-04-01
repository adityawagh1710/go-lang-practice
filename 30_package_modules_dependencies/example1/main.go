package main

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
	"github.com/fatih/color"
)

func main() {
	// How module dependancies work
	// 1. go get ******
	// 2. go mod init learn
	// 3. go mod tidy
	// 4. go mod tidy -v
	// 5.

	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
	color.Cyan(uuid)
}
