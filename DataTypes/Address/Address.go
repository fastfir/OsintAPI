package OsintAPI

import (
	"strconv"
)

type Address struct {
	HouseNumber, PostalCode                 int
	StreetName, StreetType, City, StateCode string
}

func (a Address) FullAddress() string {
	return strconv.Itoa(a.HouseNumber) + " " + a.StreetName + " " + a.StreetType + " " + a.City + " " + a.StateCode + " " + strconv.Itoa(a.PostalCode)
}
