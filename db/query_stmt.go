package db

var (
	timeCondition = "time()"
)

const (
	// tags
	qCount = "select mean(count) from \"%s\" group by time(30s), %s;"
)
