package main

import (
	"fmt"
	"os"

	"github.com/morikuni/y2j2y"
)

func main() {
	c := y2j2y.YAML2JSONConverter{}
	err := c.Convert(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
