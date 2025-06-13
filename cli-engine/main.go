package main

import (
	"fmt"
	"os"

	"cli-engine/parser"
	"cli-engine/utils"
)

func main() {
	os.MkdirAll("../parsed-output", os.ModePerm)

	rootPath := "../jcl-files"
	files := utils.FindJCLFiles(rootPath)

	for _, file := range files {
		fmt.Printf("Parsing %s\n", file)
		jobMap := parser.ParseJCLFile(file)
		utils.SaveJSON(jobMap, file)
	}
}
