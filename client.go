package main

import (
	"fmt"
	"log"
	"studyGrpc/pkg"
)

func main() {
	var rsp string

	c, err := pkg.DialServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := c.Request(&rsp); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(rsp)
}
