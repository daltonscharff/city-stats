package sources

import (
	"errors"
	"strconv"
	"strings"

	"github.com/daltonscharff/city-stats/internal/utils"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

type NumbeoDataRow struct {
	Location                  string
	CostOfLivingIndex         float32
	RentIndex                 float32
	CostOfLivingPlusRentIndex float32
	GroceriesIndex            float32
	RestaurantPriceIndex      float32
	LocalPurchasingPowerIndex float32
}

type Numbeo struct {
	Rows []NumbeoDataRow
}

func (n *Numbeo) parse(body string) {
	tkn := html.NewTokenizer(strings.NewReader(body))
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return

		case tt == html.StartTagToken:
			t := tkn.Token()
			if t.Data == "tr" {
				row := []string{}
				for {
					tt = tkn.Next()
					t = tkn.Token()
					d := strings.TrimSpace(t.Data)
					if tt == html.ErrorToken {
						return
					}
					if tt == html.EndTagToken && d == "tr" {
						break
					}
					if tt == html.TextToken && len(d) > 0 {
						row = append(row, d)
					}
				}
				if len(row) == 7 {
					stats := []float32{}
					for i := 1; i < len(row); i++ {
						f, _ := strconv.ParseFloat(row[i], 8)
						stats = append(stats, float32(f))
					}

					r := NumbeoDataRow{row[0], stats[0], stats[1], stats[2], stats[3], stats[4], stats[5]}

					n.Rows = append(n.Rows, r)
				}
			}
		}
	}
}

func (n *Numbeo) Find(location string) (NumbeoDataRow, error) {
	body, err := utils.Scrape(utils.NumbeoUrl)
	if err != nil {
		return NumbeoDataRow{}, err
	}
	n.parse(body)

	index := slices.IndexFunc(n.Rows, func(row NumbeoDataRow) bool {
		l := strings.ToLower((row.Location))
		loc := strings.ToLower(location)
		return strings.Contains(l, loc)
	})

	if index == -1 {
		return NumbeoDataRow{}, errors.New("location not found")
	}

	return n.Rows[index], nil
}
