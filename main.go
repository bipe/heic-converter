package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/HugoSmits86/nativewebp"
	"github.com/jdeng/goheif"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Syntax: %s [png|webp]\n", os.Args[0])
	}

	formatArg := os.Args[1]
	if formatArg != "png" && formatArg != "webp" {
		log.Fatalf("Invalid format: %s. Use png or webp\n", formatArg)
	}

	inputDir := "./input"
	outputDir := "./output"

	os.MkdirAll(outputDir, os.ModePerm)

	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Error reading input dir: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if ext := filepath.Ext(file.Name()); ext != ".heic" && ext != ".HEIC" {
			continue
		}

		inputPath := filepath.Join(inputDir, file.Name())
		outputName := file.Name()[0:len(file.Name())-len(filepath.Ext(file.Name()))] + "." + formatArg
		outputPath := filepath.Join(outputDir, outputName)

		inputFile, err := os.Open(inputPath)
		if err != nil {
			log.Fatal(err)
		}
		defer inputFile.Close()

		heifImage, err := goheif.Decode(inputFile)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", file.Name(), err)
			continue
		}

		outFile, err := os.Create(outputPath)
		if err != nil {
			fmt.Printf("Error creating %s: %v\n", outputName, err)
			continue
		}

		switch formatArg {
		case "png":
			if err := png.Encode(outFile, heifImage); err != nil {
				fmt.Printf("Error saving %s: %v\n", outputName, err)
			}
		case "webp":
			if err := nativewebp.Encode(outFile, heifImage, nil); err != nil {
				fmt.Printf("Error saving %s: %v\n", outputName, err)
			}
		}

		outFile.Close()
		fmt.Printf("File converted: %s -> %s\n", file.Name(), outputName)
	}
}
