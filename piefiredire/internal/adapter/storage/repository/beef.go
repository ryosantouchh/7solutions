package repository

import (
	"fmt"
	"io"
	"net/http"
)

type BeefRepository struct {
	// Please Read : if we use gorm , we can do like below
	// db *gorm.DB
	db interface{}
}

func NewBeefRepository(db interface{}) *BeefRepository {
	return &BeefRepository{db: db}
}

// I write this function in the way like we query the data from some sites
func (b *BeefRepository) Get() (string, error) {
	// Please Read : if we wanna do like using the gorm, we can do like this
	// b.db.<Find or Whatever methods in gorm> - and return result out of function krub

	endpoint := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	response, err := http.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("Error Get Data: %w", err)
	}

	defer response.Body.Close()

	jsonBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Error Read File: %w", err)
	}

	result := string(jsonBody)

	return result, nil
}
