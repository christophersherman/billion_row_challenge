package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("weather_stations.csv")
	check(err)

	defer file.Close()

	reader := csv.NewReader(file)

	for i := 0; i < 2; i++ {
		_, err := reader.Read()
		check(err)
	}

	records, err := reader.ReadAll()
	check(err)

	myMap := make(map[string][]float64)

	for _, eachrecord := range records {

		var city = strings.Split(eachrecord[0], ";")[0]
		var measure = strings.Split(eachrecord[0], ";")[1]
		measureInt, err := strconv.ParseFloat(measure, 64)
		check(err)

		if myMap[city] == nil {
			myMap[city] = []float64{measureInt}
		} else {
			myMap[city] = append(myMap[city], measureInt)
		}

		for key, value := range myMap {
			//calc the min, the mean, and the max for each station - then print it

			sort.Float64s(value)
			min := value[0]
			max := value[len(value)-1]

			var sum float64 = 0.0
			for _, d := range value {
				sum += d
			}

			mean := sum / float64(len(value))

			fmt.Printf("key : %s, mean %.2f, max %.2f, min %.2f \n", key, mean, max, min)
		}
	}
}
