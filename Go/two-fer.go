package main

import "fmt"


func ShareWith(name string) string {
	msg := `One for %s, one for me.`
	if len(name) > 0 {
		return fmt.Sprintf(msg, name)
	}
	return fmt.Sprintf(msg, "you")
}
