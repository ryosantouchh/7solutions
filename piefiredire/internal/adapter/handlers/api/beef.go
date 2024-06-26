package api

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/domain"
	"github.com/ryosantouchh/7solutions/piefiredire/internal/core/ports"
)

type BeefHandler struct {
	// TODO : injected ORM / DB here - eg.) import from storage/gorm.go
	store ports.BeefService
}

func NewBeefHandler(service ports.BeefService) *BeefHandler {
	return &BeefHandler{store: service}
}

func (b *BeefHandler) GetSummary(ctx ports.HTTPContext) {
	beefString, err := b.store.Get()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.HTTPResponse{
			"Error BeefHandler Get ::": err.Error(),
		})
		return
	}

	regX := regexp.MustCompile(`[a-zA-Z0-9_]+`)
	matchWords := regX.FindAllString(beefString, -1)

	var beefCount map[string]int32 = make(map[string]int32)
	for i := range matchWords {
		word := strings.ToLower(matchWords[i])
		if _, ok := beefCount[word]; ok {
			beefCount[word] += 1
		} else {
			beefCount[word] = 1
		}
	}

	response := domain.BeefSummary{Beef: beefCount}
	ctx.JSON(http.StatusOK, response)
}
