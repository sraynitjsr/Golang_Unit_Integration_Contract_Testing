package main

import (
    "bytes"
    "testing"
)

func TestMain_Success(t *testing.T) {
    buffer := new(bytes.Buffer)
    expected := "Adding main.go"
    mainFunction(buffer)
    result := buffer.String()
    if result != expected {
        t.Errorf("Expected: %v, but got: %v", expected, result)
    }
}
