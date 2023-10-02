package enums

const (
	Created = int64(iota + 1)
	Active
	Ended
)

func ValidStatus(level int64) bool {
	return (level >= 1 && level <= 3)
}
