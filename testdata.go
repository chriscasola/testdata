package testdata

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"os"
)

type mockRow struct {
	ID    int
	Cells []float64
}

func getRando() *rand.Rand {
	return rand.New(rand.NewSource(99))
}

func stream(numRows int, fileName string) {
	dataFile, _ := os.Create(fileName)
	defer dataFile.Close()
	dataFileWriter := bufio.NewWriter(dataFile)
	defer dataFileWriter.Flush()

	dataFileWriter.WriteString("[")
	defer dataFileWriter.WriteString("]")

	r := getRando()

	for i := 0; i < numRows; i++ {
		row := mockRow{ID: i, Cells: []float64{r.Float64(), r.Float64(), r.Float64()}}
		rowJSON, _ := json.Marshal(row)
		dataFileWriter.Write(rowJSON)
		if i != numRows-1 {
			dataFileWriter.WriteString(",")
		}
	}
}

// WriteTestData generates a test JSON data file with the given number of rows
func WriteTestData(numRows int, fileName string) {
	stream(numRows, fileName)
}
