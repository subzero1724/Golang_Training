package service

import (
	"GoBus/data"
	"encoding/json"
	"errors"
)

func LoadDestinations() (map[string]int, error) {
	var result map[string]int
	err := json.Unmarshal([]byte(data.Destination), &result)
	if err != nil {
		return nil, errors.New("failed to load destinations")
	}
	return result, nil
}
