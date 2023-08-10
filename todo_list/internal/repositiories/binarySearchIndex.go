package repositiories

import (
	"time"

	"github.com/todo_list/internal/core/domain"
)

type indexEntry struct {
	key   string
	Id    int64
	Event string
}
type BinarySearchIndex struct {
	enteries       []indexEntry
	comparableFunc func(valueOne string, valueTwo string) bool
}

func NewBinarySearchIndex(comparableFunc func(valueOne string, valueTwo string) bool) Index {
	return &BinarySearchIndex{
		enteries:       make([]indexEntry, 0),
		comparableFunc: comparableFunc,
	}
}

func (b *BinarySearchIndex) findFirstIndex(key string) int {
	startIdx := 0
	endIdx := len(b.enteries) - 1
	index := len(b.enteries)
	mid := 0
	for startIdx <= endIdx {
		mid = startIdx + (endIdx-startIdx)/2
		if b.comparableFunc(key, b.enteries[mid].key) {
			index = mid
			endIdx = mid - 1

		} else {
			startIdx = mid + 1
		}
	}
	return index
}

func (b *BinarySearchIndex) AddEntry(key string, id int64, event string) {

	index := b.findFirstIndex(key)
	if index == len(b.enteries) {
		b.enteries = append(b.enteries, indexEntry{Id: id, key: key, Event: event})
	} else {
		tempSlice := append(b.enteries[:index], indexEntry{Id: id, key: key, Event: event})
		b.enteries = append(tempSlice, b.enteries[index:]...)
	}
}

func (b *BinarySearchIndex) FindEventInRanges(rangeOne string, rangeTwo string) []indexEntry {
	var (
		startIdx     = b.findFirstIndex(rangeOne)
		endIdx       = b.findFirstIndex(rangeTwo)
		startTime, _ = time.ParseInLocation(domain.TimeFormat, rangeOne, domain.Location)
		endTime, _   = time.ParseInLocation(domain.TimeFormat, rangeTwo, domain.Location)
		enteries     = make([]indexEntry, 0)
	)
	for curIdx := startIdx; curIdx <= endIdx && curIdx < len(b.enteries); curIdx++ {
		curTime, _ := time.ParseInLocation(domain.TimeFormat, b.enteries[curIdx].key, domain.Location)
		if startTime.Equal(curTime) || startTime.Before(curTime) {
			if endTime.Equal(curTime) || endTime.After(curTime) {
				enteries = append(enteries, b.enteries[curIdx])
			}
		}
	}
	return enteries
}
