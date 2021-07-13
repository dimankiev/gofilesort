# Go File Sorting v1.1.1
![Go Workflow](https://github.com/dimankiev/gofilesort/actions/workflows/go.yml/badge.svg) [![CodeQL](https://github.com/dimankiev/gofilesort/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dimankiev/gofilesort/actions/workflows/codeql-analysis.yml)

Program, written in golang, that sorts and group the files into categorized folders
Recursively sorts folders content, creates sorted folder and places a report (log) into it
## Features
  - Recursively sorts folders content
  - Program does copy the files into `sorted` folder
    - Every copied file is located in the corresponding `Firstname Lastname pair` folder
  - Get a report with exact numbers, how many files are sorted, unsorted and the total number of processed files
