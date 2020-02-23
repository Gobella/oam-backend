package common

const (
	OK                   = 0
	PARAM_ERROR CodeEnum = 2000 + iota
	UNKNOWN_EXCEPTION
	DB_EXCEPTION
)

type CodeEnum int
