package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/rhtyx/insert-service.git/internal/model"
	log "github.com/sirupsen/logrus"
)

func OpenCSV(CSVName string) (*csv.Reader, *os.File, error) {
	log.Println("=== Open CSV file ===")

	file, err := os.Open(CSVName)
	if err != nil {
		log.Error(err.Error())
		return nil, nil, err
	}

	reader := csv.NewReader(file)
	return reader, file, nil
}

func SendToWorker(CSVReader *csv.Reader, jobs chan<- []*model.Majestic) {
	jobData := []*model.Majestic{}
	for {
		row, err := CSVReader.Read()
		if err == io.EOF {
			err = nil
			break
		}

		if err != nil {
			log.Error(err.Error())
			break
		}

		firstRow, err := strconv.Atoi(row[0])
		if err != nil {
			log.Error(err.Error())
		}

		if firstRow == 0 {
			continue
		}

		jobData = append(jobData, model.NewMajestic(row))

		if len(jobData) == 100 {
			jobs <- jobData
			jobData = []*model.Majestic{}
		}
	}
	close(jobs)
}
