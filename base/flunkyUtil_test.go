package base

import (
	"os"
	"testing"

	"github.com/flosch/pongo2/v6"
)

func TestReturnStringFromFile(t *testing.T) {
	// Create a file
	fileName := "testFile.txt"
	fileContents := []byte("This is a test file")
	err := os.WriteFile(fileName, fileContents, 0644)
	if err != nil {
		t.Error(err)
	}

	// Test the function
	fileContentsString := string(fileContents)
	if ReturnStringFromFile(fileName) != fileContentsString {
		t.Error("ReturnStringFromFile did not return the correct string")
	}

	// Delete the file
	err = os.Remove(fileName)
	if err != nil {
		t.Error(err)
	}
}

func TestCompileGroovy(t *testing.T) {
	// Create a context
	context := pongo2.Context{
		"test": "This is a test",
	}

	// Create a template
	template := "This is a {{ test }}"
	out := CompileGroovy(context, template)

	// Test the function
	if out != "This is a This is a test" {
		t.Error("CompileGroovy did not return the correct string")
	}
}
