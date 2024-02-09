package pkg

import (
	"fmt"
	"google.golang.org/api/calendar/v3"
	"math/rand"
	"time"
)

type TimeRange struct {
	StartHour, EndHour int
	Probability        int
}

func GenerateEvents(startDate, endDate time.Time, eventsPerDay int, ranges []TimeRange) []*calendar.Event {
	var events []*calendar.Event

	currentDate := startDate
	for !currentDate.After(endDate) {
		for i := 0; i < eventsPerDay; i++ {
			eventTime := selectRandomTime(ranges, currentDate)
			event := &calendar.Event{
				Summary: fmt.Sprintf("Event on %s #%d", currentDate.Format("2006-01-02"), i+1),
				Start: &calendar.EventDateTime{
					DateTime: eventTime.Format(time.RFC3339),
					TimeZone: "UTC",
				},
				End: &calendar.EventDateTime{
					DateTime: eventTime.Add(1 * time.Hour).Format(time.RFC3339),
					TimeZone: "UTC",
				},
			}
			events = append(events, event)
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return events
}

func selectRandomTime(ranges []TimeRange, currentDate time.Time) time.Time {
	rand.Seed(time.Now().UnixNano())
	totalProbability := 0
	for _, r := range ranges {
		totalProbability += r.Probability
	}

	randomProb := rand.Intn(totalProbability)
	sum := 0

	for _, r := range ranges {
		sum += r.Probability
		if randomProb < sum {
			randomHour := rand.Intn(r.EndHour-r.StartHour) + r.StartHour
			return time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), randomHour, 0, 0, 0, currentDate.Location())
		}
	}
	return time.Time{}
}
