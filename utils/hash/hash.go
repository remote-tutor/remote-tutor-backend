package hash

import (
	"github.com/speps/go-hashids"
)

// GenerateHash generates hash from the ID using `hashids` package
func GenerateHash(IDs []uint, salt string) string {
	hd := hashids.NewData()
	hd.Salt = salt    // adds the salt
	hd.MinLength = 15 // gives the length required for the output
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode(convertFromUIntToIntArray(IDs))
	return e
}

func convertFromUIntToIntArray(IDs []uint) []int {
	intIDs := make([]int, len(IDs))
	for i := range IDs {
		intIDs[i] = int(IDs[i])
	}
	return intIDs
}
