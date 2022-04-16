package config

import "errors"

type DateScope string

const (
	Hour  DateScope = "hour"
	Day   DateScope = "day"
	Month DateScope = "month"
	Year  DateScope = "year"
)

var (
	DataScopes = []DateScope{Hour, Day, Month, Year}
)

// String is used both by fmt.Print and by Cobra in help text
func (e *DateScope) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *DateScope) Set(v string) error {
	switch v {
	case "hour", "day", "month", "year":
		*e = DateScope(v)
		return nil
	default:
		return errors.New(`must be one of "hour", "day", "month", "year"`)
	}
}

// Type is only used in help text
func (e *DateScope) Type() string {
	return "DateScope"
}
