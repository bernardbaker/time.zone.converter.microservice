package app

import (
	"fmt"
	"time"
)

type ConverterService struct {
}

func NewConverterService() *ConverterService {
	return &ConverterService{}
}

func (s *ConverterService) Convert(timestamp string, timezone string) (string, error) {
	inputTime := timestamp
	targetZone := timezone

	utcTime, err := time.Parse(time.RFC3339, inputTime)
	if err != nil {
		fmt.Println("Error parsing input time:", err)
		return "", err
	}

	location, lerr := time.LoadLocation(targetZone)
	if lerr != nil {
		fmt.Println("Error loading target location:", err)
		return "", lerr
	}

	targetTime := utcTime.In(location)

	fmt.Println("Input Time (UTC):", utcTime)
	fmt.Println("Converted Time:", targetTime)
	return targetTime.UTC().String(), nil
}
