// Package twofer implements a function to return a string
// if the message is correct
package twofer

import (
	"fmt"
)

// ShareWith returns the string correct
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
