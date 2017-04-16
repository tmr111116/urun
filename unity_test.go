package main

import (
	"testing"
)

func TestGetProjPath(t *testing.T) {
	result, ok := getProjPath([]string{"a", projPathArg, "c"})
	if !ok || result != "c" {
		t.Fail()
	}
}

func TestGetProjPathNotFound(t *testing.T) {
	result, ok := getProjPath([]string{"a", "b", "c"})
	if ok || result != "" {
		t.Fail()
	}
}

func TestGetProjPathNoValue(t *testing.T) {
	result, ok := getProjPath([]string{"a", "b", projPathArg})
	if ok || result != "" {
		t.Fail()
	}
}
