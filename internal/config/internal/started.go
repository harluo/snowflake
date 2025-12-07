package internal

import (
	"time"
)

type Started struct {
	Year     uint32 `json:"year,omitempty"`
	Month    uint8  `json:"month,omitempty" validate:"required,min=1,max=12"`
	Day      uint8  `json:"day,omitempty" validate:"required,min=1,max=31"`
	Hour     uint8  `json:"hour,omitempty" validate:"required,min=1,max=24"`
	Minute   uint8  `json:"minute,omitempty" validate:"required,min=1,max=60"`
	Second   uint8  `json:"second,omitempty" validate:"required,min=1,max=60"`
	Location string `json:"location,omitempty"`
}

func (s *Started) Time() (timestamp time.Time, err error) {
	location := time.UTC
	if "" != s.Location {
		location, err = time.LoadLocation(s.Location)
	}
	if nil == err {
		timestamp = time.Date(
			int(s.Year), time.Month(s.Month), int(s.Day),
			int(s.Hour), int(s.Minute), int(s.Second),
			0, location,
		)
	}

	return
}
