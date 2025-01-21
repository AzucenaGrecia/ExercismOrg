package main

import "fmt"

func share(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("one for " + name + ", one for me.")
}
