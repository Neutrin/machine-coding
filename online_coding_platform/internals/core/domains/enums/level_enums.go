package enums

const (
	Low = int64(iota + 1)
	Medium
	High
)

func ValidLevel(level int64) bool {
	return (level >= 1 && level <= 3)
}
