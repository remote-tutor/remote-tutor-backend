package utils

import "github.com/speps/go-hashids"

// GenerateHash generates hash from the ID using `hashids` package
func GenerateHash(ID int) string {
	hd := hashids.NewData()
	hd.Salt = "xOBtdmJZxRcz^jkkyHfkrkT1*02bJUn+YQts0*xCeka%cGHCN1fjaC*faFtY" // adds the salt
	hd.MinLength = 8                                                         // gives the length required for the output
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{ID})
	return e
}
