package main

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileCopy(fileName string, dest string) {
	// Open a source
	srcFile, err := os.Open(fileName)
	check(err)
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		check(err)
	}(srcFile)

	// Create and open a copy
	destFile, err := os.Create(filepath.Join(dest, fileName))
	check(err)
	defer func(destFile *os.File) {
		err := destFile.Close()
		check(err)
	}(destFile)

	// Copy from source to destination file
	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	// Carefully finish
	err = destFile.Sync()
	check(err)
}

func makeProgressBar(step int, maxStep int, info string, max int) (*progressbar.ProgressBar, chan struct{}) {
	doneCh := make(chan struct{})
	infoString := fmt.Sprintf("[green][%d/%d][reset] %s", step, maxStep, info)
	bar := progressbar.NewOptions(max,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(infoString),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[cyan]=[reset]",
			SaucerHead:    "[cyan]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	return bar, doneCh
}

func addToReport(report *os.File, line string) {
	_, err := report.WriteString(line)
	check(err)
	err = report.Sync()
	check(err)
}

func main() {
	// Print info about the program
	fmt.Println("File Sorting Program v. 1.0.0")
	fmt.Println("For: KA Health and Cosmetics Acupuncture PC")
	fmt.Println("Author: Dmitriy Nelipa (https://t.me/dimankiev)")
	fmt.Println("Description:")
	fmt.Println("  Sorts patient-related files by their first and last name")
	fmt.Println("  Skips folders, creates sorted folder and places a report into it")
	fmt.Println("   - Program does copy the files into 'sorted' folder")
	fmt.Println("   - Every copied file is located in the corresponding Firstname Lastname folder")
	// Get files list
	items, err := ioutil.ReadDir(".")
	check(err)

	// Create the unsorted folder, if it does not exist
	err = os.MkdirAll(
		filepath.Join(".", "sorted", "unsorted"),
		os.ModePerm,
	)
	check(err)

	// Set report filename
	now := time.Now()
	filename := fmt.Sprintf(
		"./sort-report.%02d-%02d-%02d_%02d-%02d-%02d.txt",
		now.Month(), now.Day(), now.Year(), now.Hour(), now.Minute(), now.Second(),
	)

	// Create the report file
	f, err := os.Create(
		filepath.Join(".", "sorted", filename),
	)

	// Compile regex expression for file matching (Firstname Lastname)
	r, _ := regexp.Compile("^[\\w\\-_]+\\s[\\w\\-_]+")

	// Create array for counting matches (Filename: TotalMatches)
	matches := make(map[string]int)

	// Create a variable to count total amount of processed files
	total := 0

	// Instantiate a sorting progressbar
	sortBar, _ := makeProgressBar(1, 2, "Sorting the files", len(items))

	// Perform the copy process for every found item (except folders)
	for _, item := range items {
		// Skip the folders
		if item.IsDir() {
			continue
		}
		// Count total files
		total += 1
		// Update the progressbar
		err := sortBar.Add(1)
		check(err)
		// Clarify sorted file's folder
		match := r.FindString(item.Name())
		// Set 'unsorted' folder as a default destination path
		path := filepath.Join(".", "sorted", "unsorted")
		if len(match) != 0 {
			// Count current match
			matches[match] += 1
			// Set the corresponding folder as a destination path
			path = filepath.Join(".", "sorted", match)
			// Create the destination folder, if it does not exist
			err = os.MkdirAll(path, os.ModePerm)
			check(err)
		} else {
			// Count unsorted files
			matches["Unsorted"] += 1
		}
		// Do the copy
		fileCopy(item.Name(), path)
	}

	// Clean the memory
	items = nil
	r = nil

	// Instantiate a report progressbar
	reportBar, _ := makeProgressBar(2, 2, "Generating a report", len(matches))

	// Write the matches to the report
	for name, count := range matches {
		// Update the progressbar
		err := reportBar.Add(1)
		check(err)
		// Postpone the adding of information about total amount of unsorted files
		if name == "Unsorted" {
			continue
		}
		// Generate a report line
		summary := fmt.Sprintf(
			"%s: %d files\n",
			name, count,
		)
		// Add the line to the report
		addToReport(f, summary)
	}

	// Add information about total unsorted files
	addToReport(f, fmt.Sprintf(
		"\n%s: %d files\n",
		"Unsorted", matches["Unsorted"],
	))

	// Add information about overall amount of sorted files
	addToReport(f, fmt.Sprintf(
		"\n%s: %d files\n",
		"Total", total,
	))

	// Clean the memory
	matches = nil
	reportBar = nil
	sortBar = nil

	// Pause the program until user confirms exit
	fmt.Print("\n\nSuccess!\n\nPress 'Enter' to continue...")
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	check(err)
}
