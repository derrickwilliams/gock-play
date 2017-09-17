package main

import (
	"fmt"
	cookie "github.com/derrickwilliams/gock-play/older"
)

func main() {
	b := NewBaker()
	b()
}

func NewBaker() func() {
	return func() {
		cookie.GimmeCookie(func(s string) {
			fmt.Println("NewBaker() cookie.GimmeCookie() says ", s)
		})
	}
}
