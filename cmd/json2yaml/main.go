package main

import (
	"fmt"
	"os"

	"github.com/morikuni/y2j2y"
)

func main() {
	c := y2j2y.JSON2YAMLConverter{}
	err := c.Convert(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
