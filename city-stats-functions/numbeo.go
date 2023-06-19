package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

var numbeoUrl string = "https://www.numbeo.com/cost-of-living/rankings_current.jsp"

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
	Rows      []NumbeoDataRow
	scrapedAt time.Time
}

func (n *Numbeo) Scrape() {
	resp, err := http.Get(numbeoUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	n.scrapedAt = time.Now().UTC()

	tkn := html.NewTokenizer(strings.NewReader((string(body))))
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
						fmt.Println("Error token")
						return
					}
					if tt == html.EndTagToken && d == "tr" {
						break
					}
					if tt == html.TextToken && len(d) > 0 {
						if d == "Rank" {
							continue
						}
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
	if n.scrapedAt.Before(time.Now().AddDate(0, 0, -1)) {
		n.Scrape()
	}

	index := slices.IndexFunc(n.Rows, func(row NumbeoDataRow) bool {
		l := strings.ToLower((row.Location))
		return strings.Contains(l, location)
	})

	if index == -1 {
		return NumbeoDataRow{}, errors.New("location not found")
	}

	return n.Rows[index], nil
}
