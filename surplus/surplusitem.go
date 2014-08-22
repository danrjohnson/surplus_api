package surplus

import (
	"fmt"
	"strconv"
	"time"
)

type SurplusItem struct {
	State           string
	County          string
	NSN             string
	ItemName        string
	Quantity        int
	UI              string
	AcquisitionCost float64
	ShipDate        time.Time
}

func (s *SurplusItem) GetFullLocation() string {
	return fmt.Sprintf("%s, %s", s.State, s.County)
}

func (s *SurplusItem) Parse(in []string) {
	s.State = in[0]
	s.County = in[1]
	s.NSN = in[2]
	s.ItemName = in[3]
	quantity, err := strconv.Atoi(in[4])
	if err != nil {
		s.Quantity = 0
	} else {
		s.Quantity = quantity
	}
	s.UI = in[5]
	ac, err := strconv.ParseFloat(in[6], 0)
	if err != nil {
		panic(err)
	}
	s.AcquisitionCost = ac
	time_layout1 := "2006-01-02"
	time_layout2 := "20060102"
	time_layout3 := "2006-01-02T15:04:05"
	ship_date, err := time.Parse(time_layout1, in[7])
	if err != nil {
		ship_date2, err2 := time.Parse(time_layout2, in[7])
		if err2 != nil {
			ship_date3, err3 := time.Parse(time_layout3, in[7])
			if err3 != nil {
				fmt.Println(err3)
				panic(err3)
			} else {
				s.ShipDate = ship_date3
			}
		} else {
			s.ShipDate = ship_date2
		}
	} else {
		s.ShipDate = ship_date
	}
}
