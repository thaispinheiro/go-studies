# Get One Billion Row Challenge

A Go project created to process large datasets containing weather station measurements.

The application reads a file with temperature records and calculates the minimum, maximum, and average temperature for each station.

This challenge focuses on processing files with millions or billions of rows efficiently, avoiding excessive memory usage that could crash or freeze the machine.

In the future, I plan to implement a solution that generates and writes data in chunks to avoid this issue.

---

## Project Overview

Each line in the dataset follows this format:

```txt
<station_name>;<temperature>
```

Example:

```txt
Hamburg;12.0
São Paulo;31.4
Tokyo;18.2
```

The project processes all rows and groups the results by station name.

---

## Features

- Large file processing
- Temperature aggregation by station
- Minimum, maximum, and average calculation
- Concurrent processing with goroutines
- Optimized memory usage
- Buffered file reading

---

## Concepts Practiced

- Goroutines
- File reading
- Maps and data aggregation
- Structs
- Memory optimization
- Parallel processing
- Performance optimization

---

## Project Structure

```bash
get-one-billion-row-challenge/
├── go.mod
├── main.go
├── create.py
└── measurements.txt
```

---

## Generate Data

1. Make sure Python is installed.
2. Run:

```sh
python3 create.py 1_000_000
```

> **Warning:** The current script generates all data in memory before saving the file. Generating one billion rows may consume almost 15GB of RAM and freeze the machine. A future improvement is to generate and write the data in chunks to reduce memory usage.

---

## How to Run

### Prerequisites

- Go installed on the machine

Installation guide:

https://golang.org/doc/install

---

### Running the Application

```bash
go run main.go
```

---

## Challenge Reference

- https://github.com/1brc
- https://1brc.dev
