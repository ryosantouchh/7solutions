package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryosantouchh/7solutions/piefiredire/internal/adapter/handlers/api"
	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/domain"
	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/ports"
)

type TestContext struct {
	*httptest.ResponseRecorder
}

func (t *TestContext) JSON(code int, body interface{}) {
	t.WriteHeader(code)
	json.NewEncoder(t.Body).Encode(body)
}

type MockBeefService struct {
	store ports.BeefService
}

var expectedBeefString string = "Pork, Beef loin. Sirloin, Beef"

func (m *MockBeefService) Get() (string, error) {
	return expectedBeefString, nil
}

func TestHandlerGetSummary(t *testing.T) {
	expectedCount := map[string]int{
		"Pork":    1,
		"Beef":    2,
		"loin":    1,
		"Sirloin": 1,
	}

	recorder := httptest.NewRecorder()
	testCtx := &TestContext{ResponseRecorder: recorder}
	mockService := &MockBeefService{}

	beefString, _ := mockService.Get()
	if beefString != expectedBeefString {
		t.Errorf("Expected '%v' but got '%v' \n", expectedBeefString, beefString)
	}

	_, err := http.NewRequest("GET", "/beef/summary", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := api.NewBeefHandler(mockService)
	handler.GetSummary(testCtx)

	if testCtx.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", testCtx.Code)
	}

	defer testCtx.Result().Body.Close()

	var response domain.BeefSummary
	err = json.NewDecoder(testCtx.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Beef == nil {
		t.Errorf("Expected 'Beef' map response \n %v", expectedCount)
	}
}
