package main

import (
	"encoding/csv"
	"os"

	"github.com/dojeto/kafka-batch-go-example/producer/utils"
)

var (
	csvpath = os.Getenv("CSV_PATH")
)

func ReadCsv() ([]utils.Records, error) {

	var rcd []utils.Records

	file, err := os.Open(csvpath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	for _, r := range records {
		rcd = append(rcd, utils.Records{
			Name:     r[0],
			Email:    r[1],
			Password: r[2],
		})
	}

	if err != nil {
		return nil, err
	}

	return rcd, nil
}
