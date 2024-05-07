package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/domain"
)

func TestBeefSummary(t *testing.T) {
	expectedCount := map[string]int{
		"Beef":    2,
		"Pork":    1,
		"Sirloin": 1,
		"loin":    1,
	}

	beefSummary := domain.BeefSummary{Beef: expectedCount}

	data, err := json.Marshal(beefSummary)
	if err != nil {
		t.Errorf("Error marshal 'BeefSummary' : '%v' \n", err)
	}

	jsonString := string(data)

	expectedJSON := `{"beef":{"Beef":2,"Pork":1,"Sirloin":1,"loin":1}}`
	if jsonString != expectedJSON {
		t.Errorf("Expected JSON: %s \n but got: %s", expectedJSON, jsonString)
	}
}
