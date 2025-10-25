package files

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSVFiles(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, re := range records[1:] {
		fmt.Println(re)
	}
}