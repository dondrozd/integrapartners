package main

import "fmt"

func hello(user string) string {
	if len(user) == 0 {
		return "hey you"
	}
	return fmt.Sprintf("hey %v", user)
}
