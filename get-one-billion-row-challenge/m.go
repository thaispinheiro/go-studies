package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min float64
	Max float64
	Sum	float64
	Count int64
}

func main() {
	start := time.Now()
	measurements, err := os.Open("measurements.txt")
	panicIfErr(err)

	defer measurements.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, err := strconv.ParseFloat(rawTemp, 64)
		panicIfErr(err)

		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{
				Min: temp,
				Max:temp,
				Sum: temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++	
		}
		data[location] = measurement
	}
	panicIfErr(scanner.Err())

	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)
	}

	sort.Strings(locations)
	parts:= []string{}

	for _, name := range locations {
		parts = append(
			parts,
			formatMeasurement(data[name], name),
		)
	}

	result := strings.Join(parts, ",")
	fmt.Printf("{%s}\n", result)
	fmt.Println(time.Since(start))
}

func formatMeasurement(measurement Measurement, name string) string {
	return fmt.Sprintf(
		"%s=%.1f/%.1f/%.1f",
		name,
		measurement.Min,
		measurement.Sum/float64(measurement.Count),
		measurement.Max,
	)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}