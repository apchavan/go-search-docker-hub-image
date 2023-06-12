package main

import (
	"fmt"

	runnerpkg "github.com/apchavan/go-search-docker-hub-image/runner"
)

func main() {
	_ = runnerpkg.GetTuiAppLayout()
	fmt.Printf("'%s' exited...\n", runnerpkg.GetApplicationName())
}
