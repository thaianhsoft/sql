package utils

import (
	"log"
	"testing"
)

func Assert(t *testing.T, have interface{}, want interface{}) {
	if have != want {
		log.Printf("have(%v)\n, want(%v)\n", have, want)
		t.FailNow()
		return
	}
}
