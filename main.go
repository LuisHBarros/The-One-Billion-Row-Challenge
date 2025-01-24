package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Measurable struct {
	name  string
	min   float64
	max   float64
	sum   float64
	count int
}

func main() {
	start := time.Now()
	var list = []*Measurable{}
	measurement, err := os.Open("weather.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer measurement.Close()

	render, err := csv.NewReader(measurement).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range render {
		parts := strings.Split(record[0], ";")
		if len(parts) < 2 {
			continue
		}
		name := parts[0]
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		index := findInArray(name, list)
		if index != -1 {
			list[index].min = min(list[index].min, value)
			list[index].max = max(list[index].max, value)
			list[index].sum += value
			list[index].count++
		} else {
			list = append(list, &Measurable{name: name, min: value, max: value, sum: value, count: 1})
		}
	}
	for _, v := range list {
		avg := v.sum / float64(v.count)
		log.Printf("%s: min=%.2f, max=%.2f, avg=%.2f\n", v.name, v.min, v.max, avg)
	}
	duration := time.Since(start) // Calcula o tempo decorrido
	fmt.Printf("Tempo de execução: %v\n", duration)
}

func findInArray(name string, list []*Measurable) int {
	for i, v := range list {
		if v.name == name {
			return i
		}
	}
	return -1
}
