package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/domain"
)

func TestBeefSummary(t *testing.T) {
	expectedCount := map[string]int{
		"beef":    2,
		"loin":    1,
		"pork":    1,
		"sirloin": 1,
	}

	beefSummary := domain.BeefSummary{Beef: expectedCount}

	data, err := json.Marshal(beefSummary)
	if err != nil {
		t.Errorf("Error marshal 'BeefSummary' : '%v' \n", err)
	}

	jsonString := string(data)

	expectedJSON := `{"beef":{"beef":2,"loin":1,"pork":1,"sirloin":1}}`
	if jsonString != expectedJSON {
		t.Errorf("Expected JSON: %s \n but got: %s", expectedJSON, jsonString)
	}
}
