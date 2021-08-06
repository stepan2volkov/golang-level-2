package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"searchdupl/scan"
	"strings"
)

var (
	removeFlag *bool
	path       *string
)

func init() {
	removeFlag = flag.Bool("remove", false, "remove found duplicates")
	path = flag.String("path", ".", "path to scan")
	flag.Parse()
}

func main() {
	duplicates := scan.ScanDir(*path)
	for duplicateFiles := range duplicates {
		fmt.Printf("[%s] (%d):\n", filepath.Base(duplicateFiles[0]), len(duplicateFiles))
		for _, duplicateFile := range duplicateFiles {
			fmt.Println("\t", duplicateFile)
		}
		if *removeFlag {
			fmt.Print("Are you sure, that you want to remove all files besides first? (y/N) ")
			var answer string
			// fmt.Scan(answer)
			_, err := fmt.Scanln(&answer)
			if err != nil {
				fmt.Println(err)
			}

			if strings.ToLower(answer) == "y" {
				for _, f := range duplicateFiles[1:] {
					err := os.Remove(f)
					if err != nil {
						fmt.Printf("Couldn't remove file: %v", err)
						os.Exit(1)
					}
				}
			}
		}
	}
}
