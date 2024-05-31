package control

type CrimeStore interface {
	GetCrimeStat(crime string) int
	RecordStat(crime string)
}
