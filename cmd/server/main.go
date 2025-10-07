package main

import (
	"Smradity/internal"
)

func main() {
	router := internal.SetupRouters()

	err := router.Run(":9080")
	if err != nil {
		return
	}
}
