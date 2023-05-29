package observer

import (
	"../entities"
	exporterService "../exporter"
	storageService "../storage"
)

type IObserverService interface {
	Notify(
		exporter exporterService.IRateExporter,
		storage storageService.IStorage[entities.User],
	) error
}

type ObserverService struct {
	Strategy func(userGmail string, BTCPrice float64)
}

func (obs *ObserverService) Notify(
	exporter exporterService.IRateExporter,
	storage storageService.IStorage[entities.User],
) error {
	users, err := storage.GetAll()

	if err != nil {
		return err
	}

	btcPrice, err := exporter.GetCurrentBTCPrice()

	if err != nil {
		return err
	}

	for _, user := range users {
		if user.HasSubscription {
			obs.Strategy(user.Gmail, btcPrice)
		}
	}

	return nil
}
