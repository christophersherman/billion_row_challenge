package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	counter := 0
	file, err := os.Open("measurements.csv")
	check(err)

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = ';'
	// for i := 0; i < 2; i++ {
	// 	_, err := reader.Read()
	// 	check(err)
	// }

	myMap := make(map[string][]float64)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		counter++
		var city = record[0]
		var measure = record[1]
		measureInt, err := strconv.ParseFloat(measure, 64)

		if err != nil {
			fmt.Println("this shit right here boy")

		}
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
		max := value[3]
		min := value[2]

		fmt.Printf("key : %s, mean %.2f, max %.2f, min %.2f \n", key, mean, max, min)
	}

	fmt.Println(counter)

}
