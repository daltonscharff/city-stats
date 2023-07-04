package models

import (
	"fmt"
	"strings"
)

type Location struct {
	City        string
	State       string
	StateAbbrev string
	Country     string
}

var states = [](struct {
	State        string
	Abbreviation string
}){
	{"Arizona", "AZ"},
	{"Alabama", "AL"},
	{"Alaska", "AK"},
	{"Arkansas", "AR"},
	{"California", "CA"},
	{"Colorado", "CO"},
	{"Connecticut", "CT"},
	{"Delaware", "DE"},
	{"Florida", "FL"},
	{"Georgia", "GA"},
	{"Hawaii", "HI"},
	{"Idaho", "ID"},
	{"Illinois", "IL"},
	{"Indiana", "IN"},
	{"Iowa", "IA"},
	{"Kansas", "KS"},
	{"Kentucky", "KY"},
	{"Louisiana", "LA"},
	{"Maine", "ME"},
	{"Maryland", "MD"},
	{"Massachusetts", "MA"},
	{"Michigan", "MI"},
	{"Minnesota", "MN"},
	{"Mississippi", "MS"},
	{"Missouri", "MO"},
	{"Montana", "MT"},
	{"Nebraska", "NE"},
	{"Nevada", "NV"},
	{"New Hampshire", "NH"},
	{"New Jersey", "NJ"},
	{"New Mexico", "NM"},
	{"New York", "NY"},
	{"North Carolina", "NC"},
	{"North Dakota", "ND"},
	{"Ohio", "OH"},
	{"Oklahoma", "OK"},
	{"Oregon", "OR"},
	{"Pennsylvania", "PA"},
	{"Rhode Island", "RI"},
	{"South Carolina", "SC"},
	{"South Dakota", "SD"},
	{"Tennessee", "TN"},
	{"Texas", "TX"},
	{"Utah", "UT"},
	{"Vermont", "VT"},
	{"Virginia", "VA"},
	{"Washington", "WA"},
	{"West Virginia", "WV"},
	{"Wisconsin", "WI"},
	{"Wyoming", "WY"}}

func (l *Location) New(city string, state string, country string) {
	l.City = city
	l.State = state
	l.Country = country

	for _, s := range states {
		if strings.EqualFold(state, s.State) {
			l.StateAbbrev = s.Abbreviation
			break
		}
	}
}

func (l Location) String() string {
	return fmt.Sprintf("%s, %s, %s", l.City, l.StateAbbrev, l.Country)
}
