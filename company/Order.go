package company

// Order struct
type Order struct {
	key
	ID                int
	color             int
	quantity          int
	dueDate           int
	fulfilledQuantity int // function

	// Owner
	company *Company

	// Relation
	knifeSetting *KnifeSetting
}
