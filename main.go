package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"github.com/bitrise-io/go-utils/fileutil"
)

func main() {
	var (
		inputFile                     = os.Getenv("file")
		inputMatchPattern             = os.Getenv("match_pattern")
		inputSubstitution             = os.Getenv("substitution_value")
		inputIsShowFileContent        = os.Getenv("show_file") == "true"
	)

	if inputFile == "" {
		log.Fatal("No file input specified")
	}
	if inputMatchPattern == "" {
		log.Fatal("No match_pattern input specified")
	}
	if inputSubstitution == "" {
		log.Fatal("No substitution_value input specified")
	}

	origContent, err := fileutil.ReadStringFromFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read from specified file, error: %s", err)
	}

	if inputIsShowFileContent {
		fmt.Println()
		fmt.Println("------------------------------------------")
		fmt.Println("-------------OLD  FILE--------------------")
		fmt.Println("------------------------------------------")
		fmt.Print(origContent)
		fmt.Println()
		fmt.Println("------------------------------------------")
	}

	// replace
	fmt.Println(" (i) Replacing...")
	r, _ := regexp.Compile(inputMatchPattern)
	replacedContent := r.ReplaceAllString(origContent, inputSubstitution)

	if inputIsShowFileContent {
		fmt.Println()
		fmt.Println("------------------------------------------")
		fmt.Println("-------------NEW  FILE--------------------")
		fmt.Println("------------------------------------------")
		fmt.Print(replacedContent)
		fmt.Println()
		fmt.Println("------------------------------------------")
	}

	// write back to file
	if err := fileutil.WriteStringToFile(inputFile, replacedContent); err != nil {
		log.Printf("Failed to write replaced content back to file, error: %s", err)
	}
	fmt.Println(" (i) Done")
}