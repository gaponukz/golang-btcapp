package observer

import (
	"btcapp/src/entities"
	exporterService "btcapp/src/exporter"
	storageService "btcapp/src/storage"
	"sync"
)

type IObserverService interface {
	Notify(
		exporter exporterService.IRateExporter,
		storage storageService.IStorage[entities.User],
	) (int, int)
}

type ObserverService struct {
	Strategy func(userGmail string, BTCPrice float64, errors chan error)
}

func (obs *ObserverService) Notify(
	exporter exporterService.IRateExporter,
	storage storageService.IStorage[entities.User],
) (int, int) {
	users, err := storage.GetAll()
	if err != nil {
		return 0, 0
	}

	btcPrice, err := exporter.GetCurrentBTCPrice()
	if err != nil {
		return 0, 0
	}

	errorscCannel := make(chan error, len(users))
	errorCount := 0
	var wg sync.WaitGroup

	for _, user := range users {
		if user.HasSubscription {
			wg.Add(1)
			go func(userGmail string) {
				defer wg.Done()
				obs.Strategy(userGmail, btcPrice, errorscCannel)
			}(user.Gmail)
		}
	}

	go func() {
		wg.Wait()
		close(errorscCannel)
	}()

	for err := range errorscCannel {
		if err != nil {
			errorCount += 1
		}
	}

	return len(users) - errorCount, len(users)
}
