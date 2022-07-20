package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Uid         string
	Summary     string
	Description string
	StartDate   time.Time
}

type Calendar struct {
	Events []Event
}

func (e Event) Serialize() string {
	return fmt.Sprintf(`BEGIN:VEVENT
UID:%s
SUMMARY:%s
DTSTAMP:%s
DTSTART:%s
DESCRIPTION:%s
STATUS:CONFIRMED
DURATION:1H
END:VEVENT`, e.Uid, e.Summary, icsDateFormat(time.Now()), icsDateFormat(e.StartDate), e.Description)
}

func icsDateFormat(date time.Time) string {
	return fmt.Sprintf("%d%02d%02dT%02d%02d%02dZ",
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		date.Second())
}

func (c Calendar) Serialize() string {
	s := "BEGIN:VCALENDAR\n"

	for _, event := range c.Events {
		s += event.Serialize()
		s += "\n"
	}

	s += "END:VCALENDAR"
	return s
}

func main() {
	c := Calendar{
		Events: []Event{},
	}
	e := Event{
		Uid:         uuid.New().String(),
		Summary:     "Test Event",
		Description: "Test Desc",
		StartDate:   time.Date(2022, time.Month(7), 28, 0, 0, 0, 0, time.UTC),
	}
	c.Events = append(c.Events, e)

	fmt.Println(c.Serialize())
}
