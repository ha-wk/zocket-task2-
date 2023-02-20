package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvLine struct {
	Column1 string
	Column2 string
	Column3 string
	Column4 string //DATA STRUCTURE(STRUCT) to store csv file contents
	Column5 string
	Column6 string
	Column7 string
	Column8 string
}

func main() {

	lines, err := ReadCsv("sample.csv")
	if err != nil { //IF BLANK FILE,RETURN ERROR
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
			Column2: line[1],
			Column3: line[2],
			Column4: line[3],
			Column5: line[4],
			Column6: line[5],
			Column7: line[6],
			Column8: line[7],
		}
		fmt.Println(data.Column1 + " " + data.Column2 + " " + data.Column3 + " " + data.Column4 + " " + data.Column5 + " " + data.Column6 + " " + data.Column7 + " " + data.Column8)
	}
}

// ReadCsv accepts a file and returns its content as a multi-dimensional type
// with lines and each column(Only parses to string type).
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
