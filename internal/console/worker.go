package console

import (
	"context"
	"fmt"
	"sync"

	"github.com/rhtyx/insert-service.git/internal/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func DoInsertJob(workerIdx, counter int, db *gorm.DB, majestic []*model.Majestic, ctx context.Context, wg *sync.WaitGroup) {
	for {
		var outterError error

		func(outterError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outterError = fmt.Errorf("%s", err)
				}
			}()

			err := db.WithContext(ctx).Create(majestic).Error
			if err != nil {
				log.Error(err.Error())
			}
		}(&outterError)

		if outterError == nil {
			break
		}
	}

	if counter%100 == 0 {
		log.Println("=> worker", workerIdx, "inserted", counter, "data")
	}
}
