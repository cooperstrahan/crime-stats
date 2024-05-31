package mocks

type StubCrimeService struct {
	Stats     map[string]int
	StatCalls []string
}

func NewMockCrimeService(stats map[string]int) *StubCrimeService {
	return &StubCrimeService{
		Stats:     stats,
		StatCalls: []string{},
	}
}

func (s *StubCrimeService) GetCrimeStat(crime string) int {
	return s.Stats[crime]
}

func (s *StubCrimeService) RecordStat(crime string) {
	s.Stats[crime]++
	s.StatCalls = append(s.StatCalls, crime)
}
