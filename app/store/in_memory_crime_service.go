package store

func NewInMemoryCrimeService() *InMemoryCrimeService {
	return &InMemoryCrimeService{
		map[string]int{},
		[]string{},
	}
}

type InMemoryCrimeService struct {
	stats     map[string]int
	statCalls []string
}

func (i *InMemoryCrimeService) GetCrimeStat(crime string) int {
	return i.stats[crime]
}

func (i *InMemoryCrimeService) RecordStat(crime string) {
	i.stats[crime]++
}
