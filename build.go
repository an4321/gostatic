package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Build() {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			filename := filepath.Join(srcDir, file.Name())
			data, err := os.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}

			metaData, content := separateMetadataData(string(data))
			metaDataMap, err := ParseMetadata(metaData)
			if err != nil {
				fmt.Println("Error parsing metadata:", err)
				return
			}
			contentHtml := Convert(content)

			baseHtml, err := os.ReadFile("./base.html")
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			updatedContent := strings.ReplaceAll(string(baseHtml), "{{content}}", contentHtml)
			updatedContent = strings.ReplaceAll(string(updatedContent), "{{title}}", metaDataMap["title"])
			updatedContent = strings.ReplaceAll(string(updatedContent), "{{description}}", metaDataMap["description"])
			updatedContent = strings.ReplaceAll(string(updatedContent), "{{date}}", metaDataMap["date"])
			updatedContent = strings.ReplaceAll(string(updatedContent), "{{author}}", metaDataMap["author"])

			outputFile := strings.TrimSuffix(file.Name(), ".md") + ".html"
			outputPath := filepath.Join(outDir, outputFile)

			// create output file
			outFile, err := os.Create(outputPath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer outFile.Close()

			// write to file
			_, err = outFile.WriteString(updatedContent)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}

			fmt.Printf("%s -> %s\n", filepath.Join(srcDir, file.Name()), outputPath)
		}
	}
}

func separateMetadataData(input string) (metadata string, data string) {
	parts := strings.Split(input, "---")

	// if there are less than 3 parts, return empty strings
	if len(parts) < 3 {
		return "", input
	}

	// the metadata is the second part
	metadata = strings.TrimSpace(parts[1])

	// the data is the concatenation of the remaining parts
	data = strings.TrimSpace(strings.Join(parts[2:], "---"))

	return metadata, data
}
