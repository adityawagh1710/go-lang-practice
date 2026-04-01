package main

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	// How module dependancies work
	// 1. go mod init learn
	// 2. go mod tidy
	// 3. go mod tidy -v
	// 4. 

	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
