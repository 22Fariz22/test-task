package pkg

import "errors"

var (
	ErrNoAgeData = errors.New("no data age")
	ErrNoGenderData = errors.New("no gender age")
	ErrNoNationalityData = errors.New("no nationality age")
)