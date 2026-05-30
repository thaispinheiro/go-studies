# Get One Billion Row Challenge

This project tackles the challenge of processing one billion rows of data.

## Structure
- `create.py` - Script for data generation
- `data.csv` - Generated data file
- `m.go` - Go code for data processing
- `go.mod` - Dependency management

## How to run

### Generate data
1. Make sure you have Python installed.
2. Run:
   ```sh
   python3 create.py 1_000_000
   ```

   > **Warning:** The current script generates all data in memory before saving to the file. For one billion rows, this can consume almost 15GB of RAM and may freeze your machine. In the future, I plan to implement a solution that generates and writes data in chunks to avoid this issue.

### Process data with Go
1. Install Go (https://golang.org/doc/install)
2. Run:
   ```sh
   go run m.go
   ```