package main

import (
	"calendar-start/auth"
	"calendar-start/pkg"
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := auth.GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	ranges := []pkg.TimeRange{
		{StartHour: 0, EndHour: 7, Probability: 5},
		{StartHour: 7, EndHour: 9, Probability: 20},
		{StartHour: 9, EndHour: 18, Probability: 50},
		{StartHour: 18, EndHour: 22, Probability: 15},
		{StartHour: 22, EndHour: 24, Probability: 10},
	}

	inputData, err := pkg.ReadEventInputData("./input.json")
	if err != nil {
		log.Fatalf("Error parsing input data: %v", err)
	}

	startDate, err := time.Parse("2006-01-02", inputData.StartDate)
	if err != nil {
		log.Fatalf("Error parsing start date: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", inputData.EndDate)
	if err != nil {
		log.Fatalf("Error parsing end date: %v", err)
	}

	events := pkg.GenerateEvents(startDate, endDate, inputData.EventsPerDay, ranges)

	addEventsToCalendar(srv, "primary", events)
}

func addEventsToCalendar(srv *calendar.Service, calendarID string, events []*calendar.Event) {
	for _, event := range events {
		createdEvent, err := srv.Events.Insert(calendarID, event).Do()
		if err != nil {
			fmt.Printf("Не удалось добавить событие: %v\n", err)
		} else {
			fmt.Printf("Событие добавлено: %s\n", createdEvent.HtmlLink)
		}
	}
}
