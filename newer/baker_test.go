package main

import (
	g "gopkg.in/h2non/gock.v1"
	"testing"
)

func TestBaker(t *testing.T) {
	b := NewBaker()
	b()
}
