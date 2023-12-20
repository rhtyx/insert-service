package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/rhtyx/insert-service.git/internal/config"
	"github.com/rhtyx/insert-service.git/internal/console"
	"github.com/rhtyx/insert-service.git/internal/csv"
	"github.com/rhtyx/insert-service.git/internal/db"
	"github.com/rhtyx/insert-service.git/internal/model"
	"gorm.io/gorm"
)

func main() {
	start := time.Now()
	ctx := context.Background()

	db := db.InitializePostgreSQLConn()

	CSVReader, CSVFile, err := csv.OpenCSV(config.CSVName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer CSVFile.Close()

	jobsChannel := make(chan []*model.Majestic, 0)
	wg := new(sync.WaitGroup)

	for workerIndex := 0; workerIndex < config.TotalWorker; workerIndex++ {
		wg.Add(1)
		go func(workerIdx int, db *gorm.DB, jobs <-chan []*model.Majestic, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				console.DoInsertJob(workerIdx, counter, db, job, ctx, wg)
				counter += 100
			}
			wg.Done()
		}(workerIndex, db, jobsChannel, wg)
	}
	csv.SendToWorker(CSVReader, jobsChannel)
	wg.Wait()

	duration := time.Since(start)
	log.Println("done: ", duration.Seconds(), " sec")
}
