package main

import (
    "testing"
)

func Test_Abs(t *testing.T) {
	v := Vertex{3, 4}
	result := Abs(v)
	if result != 5 {
		t.Errorf("Error")
	}
}
