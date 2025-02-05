package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSampleFile(t *testing.T) {
    os.Setenv("TEST_MODE", "true")
    if err := testSampleFile(); err != nil {
        t.Errorf("Sample file test failed: %v", err)
    }
}

func testSampleFile() error {
    // Open the file
    file, err := openFile("sample.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    // Parse and filter the data
    data := parseFile(file)
    matches := filterData(data)

    // Calculate the sum of products
    output, err := getOutput(matches)
    if err != nil {
        return err
    }

    // Check if the output matches the expected output for sample.txt
    if output != 48 {
        return fmt.Errorf("expected output 48, but got %d", output)
    }

    return nil
}
