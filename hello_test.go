package main

import "testing"

func TestHello(t *testing.T) {
	result := HelloWorld()
	expect := "Hello, world"

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}
}
