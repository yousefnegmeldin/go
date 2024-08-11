package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestViewList(t *testing.T) {
	buffer := bytes.Buffer{}
	listToTest := []string{"hello", "test"}
	ViewList(&buffer, listToTest)
	got := buffer.String()
	want := "hello\ntest"
	if got != want {
		t.Errorf("ViewList() = %q, want %q", got, want)
	}
}

func TestAddToList(t *testing.T) {
	listToTest := []string{}
	AddToList(listToTest, "hello")
	got := listToTest[0]
	want := "hello"
	fmt.Println(listToTest[0])
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
