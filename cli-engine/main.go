package main

import (
	"fmt"
	"os"

	"cli-engine/parser"
	"cli-engine/utils"
)

func main() {
	ensureOutputFolders()

	rootPath := "../jcl-files"
	files := utils.FindJCLFiles(rootPath)

	for _, file := range files {
		processJobFile(file)
	}
}

// ensureOutputFolders creates necessary directories
func ensureOutputFolders() {
	os.MkdirAll("../parsed-output", os.ModePerm)
	os.MkdirAll("../parsed-output/errors", os.ModePerm)
}

// processJobFile parses, saves, and logs one JCL file
func processJobFile(file string) {
	fmt.Printf("Parsing %s\n", file)

	jobMap := parser.ParseJCLFile(file)
	utils.SaveJSON(jobMap, file)

	if len(jobMap.Errors) > 0 {
		utils.LogErrors(jobMap, file)
	}
}
