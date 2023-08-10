package repositiories

type Index interface {
	AddEntry(key string, id int64, event string)
	FindEventInRanges(rangeOne string, rangeTwo string) []indexEntry
}
