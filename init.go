package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed sample.html
var sampleHtml string

//go:embed sample.md
var sampleMd string

func Init() {
	// create outDir & srcDir
	for _, dir := range []string{outDir, srcDir} {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.Mkdir(dir, 0755)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s directory created successfully.\n", dir)
		} else {
			fmt.Printf("%s directory already exists.\n", dir)
		}
	}

    // create file base.html
	createFile("base.html", sampleHtml)

    // create file {srcDir}/base.md if srcDir is empty
	files, err := os.ReadDir(srcDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(files) == 0 {
		createFile(filepath.Join(srcDir, "sample.md"), sampleMd)
	}
}

func createFile(filePath string, content string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("created", filePath)
	} else {
		fmt.Println(filePath, "already exists.")
	}
}
