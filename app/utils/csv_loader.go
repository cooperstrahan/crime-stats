package utils

import (
	"crime-stats/app/domain"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/jszwec/csvutil"
)

func LoadCrimeData() error {
	file, err := os.Open("Crime_Data_from_2020_to_Present.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	dec, err := csvutil.NewDecoder(reader)
	if err != nil {
		return err
	}

	header := dec.Header()
	crimes := []domain.CrimeDecoder{}

	for {
		crime := domain.CrimeDecoder{}

		if err := dec.Decode(&crime); err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		for _, i := range dec.Unused() {
			crime.OtherData[header[i]] = dec.Record()[i]
		}

		crimes = append(crimes, crime)
	}

	for i, record := range crimes {
		if i > 100 {
			break
		}
		fmt.Printf("%+v\n", record)
	}

	return nil
}
