package company

// PaperRoll is the struct for paper roll
type PaperRoll struct {
	key
	color  int
	length int // function

	// Owner
	knifeSetting *KnifeSetting
}
