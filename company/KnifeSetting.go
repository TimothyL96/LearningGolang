package company

import (
	keyConfiguration "github.com/ttimt/GolangWebSocket/key"
)

// KnifeSetting struct holds the field of a knife setting and is owned by company
type KnifeSetting struct {
	key
	numberOfCut int
	color       int

	// Owner
	company *Company

	// Owning objects
	paperRoll *PaperRoll

	// Relation
	orders []*Order
}

// CreatePaperRoll creates a paper roll owned by this knife setting
func (knifeSetting *KnifeSetting) CreatePaperRoll(color int, length int) *PaperRoll {
	paperRoll := &PaperRoll{
		key:          keyConfiguration.NewKey(),
		color:        color,
		knifeSetting: knifeSetting,
		length:       length,
	}

	return paperRoll
}
