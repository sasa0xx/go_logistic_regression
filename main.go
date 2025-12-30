package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	iris_path := "./Iris.csv"

	file, err := os.Open(iris_path)
	if err != nil {
		log.Fatalf("File %s does not exist.", iris_path)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	encode := func(label string) int {
		switch label {
		case "Iris-setosa":
			return 0
		case "Iris-versicolor":
			return 1
		default:
			return 2
		}
	}

	X := NewMatrix(150, 4)
	y := make([]int, 150)

	csvReader.Read()
	for n := range 150 {
		row, err := csvReader.Read()
		if err != nil {
			log.Fatal(err)
		}

		for i := range 4 {
			f, err := strconv.ParseFloat(row[i+1], 64)
			if err != nil {
				log.Fatal(err)
			}
			X.Set(n, i, f)
		}
		y[n] = encode(row[5])
	}

	model := NewLogisticRegressor(150, 4, 3)
	model.fit(X, y, 0.1, 16, 1000)
}
