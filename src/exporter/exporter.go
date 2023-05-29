package exporter

type IRateExporter interface {
	GetCurrentBTCPrice() (float64, error)
}

type MemoryExporter struct{}

func (exp *MemoryExporter) GetCurrentBTCPrice() (float64, error) {
	return 712.1, nil
}
