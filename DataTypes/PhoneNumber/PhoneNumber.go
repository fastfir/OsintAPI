package OstinAPI

import (
	"strconv"
)

type PhoneNumber struct {
	CountryCode, AreaCode, TelephonePrefix, LineNumber int
}

func (pn PhoneNumber) FullNumber() string {
	return strconv.Itoa(pn.CountryCode) + " " + strconv.Itoa(pn.AreaCode) + "-" + strconv.Itoa(pn.TelephonePrefix) + "-" + strconv.Itoa(pn.LineNumber)
}
