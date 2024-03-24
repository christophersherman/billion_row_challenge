package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	counter := 0
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
		counter++
		var city = strings.Split(eachrecord[0], ";")[0]
		var measure = strings.Split(eachrecord[0], ";")[1]
		measureInt, err := strconv.ParseFloat(measure, 64)
		check(err)

		if myMap[city] == nil {
			myMap[city] = []float64{measureInt, 1, measureInt, measureInt} // Initialize sum, total_entries, min, max
		} else {
			// Update sum and total_entries
			myMap[city][0] += measureInt // Update sum
			myMap[city][1] += 1          // Update total_entries

			if measureInt > myMap[city][3] {
				myMap[city][3] = measureInt // Update max
			} else if measureInt < myMap[city][2] {
				myMap[city][2] = measureInt // Update min
			}
		}
	}
	for key, value := range myMap {
		mean := value[0] / value[1]
		max := value[2]
		min := value[3]

		fmt.Printf("key : %s, mean %.2f, max %.2f, min %.2f \n", key, mean, max, min)
	}

	fmt.Println(counter)

}
