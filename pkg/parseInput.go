package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

type EventInputData struct {
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	EventsPerDay int    `json:"eventsPerDay"`
}

func ReadEventInputData(filePath string) (*EventInputData, error) {
	var inputData EventInputData

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %w", err)
	}
	if err := json.Unmarshal(fileData, &inputData); err != nil {
		return nil, err
	}
	return &inputData, nil
}
