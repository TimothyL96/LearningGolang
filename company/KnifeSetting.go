package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/LearningGolang/key"
)

// KnifeSetting struct holds the field of a knife setting and is owned by company
type KnifeSetting struct {
	key
	numberOfCut int
	color       int
	repetition  int // function

	// Owner
	company *Company

	// Owning objects
	paperRoll *PaperRoll

	// Relation
	orders []*Order
}

// CreatePaperRoll creates a paper roll owned by this knife setting
func (ks *KnifeSetting) CreatePaperRoll(color int, length int) *PaperRoll {
	paperRoll := &PaperRoll{
		key:          keyConfiguration.NewKey(),
		color:        color,
		length:       length,
		knifeSetting: ks,
	}

	// Create operations for paper roll
	paperRoll.createOperations()

	// Append paper roll
	ks.paperRoll = paperRoll

	return paperRoll
}

// Company returns the company of knife setting
func (ks *KnifeSetting) Company() *Company {
	if ks == nil {
		return nil
	}

	return ks.company
}

// PaperRoll returns the paper roll of the knife setting
func (ks *KnifeSetting) PaperRoll() *PaperRoll {
	if ks == nil {
		return nil
	}

	return ks.paperRoll
}

// Orders return the orders of the knife setting
func (ks *KnifeSetting) Orders() []*Order {
	if ks == nil {
		return nil
	}

	return ks.orders
}

// Color returns the color of the knife setting
func (ks *KnifeSetting) Color() int {
	if ks == nil {
		panic(errors.New("knife setting is nil").Error())
	}

	return ks.color
}

// NumberOfCut returns the number of cut of the knife setting
func (ks *KnifeSetting) NumberOfCut() int {
	if ks == nil {
		panic(errors.New("knife setting is nil").Error())
	}

	return ks.numberOfCut
}

// Repetition returns the repetition of the knife setting
func (ks *KnifeSetting) Repetition() int {
	if ks == nil {
		panic(errors.New("knife setting is nil").Error())
	}

	return ks.repetition
}

// AssignOrder will assign order to this knife setting
func (ks *KnifeSetting) AssignOrder(order *Order) {
	if len(ks.orders) > 3 {
		panic(errors.New("knife setting max order reached").Error()) // constraint
	}

	order.SetKnifeSetting(ks)
	ks.SetOrders(append(ks.Orders(), order))
}

// SetOrders will set the orders to this knife setting
func (ks *KnifeSetting) SetOrders(orders []*Order) {
	ks.orders = orders
}
