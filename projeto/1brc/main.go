package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	dir, _ := os.Getwd()

	measurements, err := os.Open(dir + "/projeto/1brc/" + "measurements.txt")

	if err != nil {
		panic(err)
	}

	defer measurements.Close()

	scanner := bufio.NewScanner(measurements)

	dados := make(map[string]Measurement)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := dados[location]

		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		dados[location] = measurement
	}

	locations := make([]string, 0, len(dados))

	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")
}

/*
 Podemos usar esse comando para pegar as coordenadas

 curl -L https://raw.githubusercontent.com/gunnarmorling/1brc/refs/heads/main/data/weather_stations.csv

 curl -L https://raw.githubusercontent.com/gunnarmorling/1brc/refs/heads/main/src/main/python/create_measurements.py
*/
