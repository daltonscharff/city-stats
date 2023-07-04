package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/daltonscharff/city-stats/internal/models"
	"github.com/daltonscharff/city-stats/internal/utils"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slices"
)

type NumbeoCostOfLivingRecord struct {
	Location                  string
	CostOfLivingIndex         float32
	RentIndex                 float32
	CostOfLivingPlusRentIndex float32
	GroceriesIndex            float32
	RestaurantPriceIndex      float32
	LocalPurchasingPowerIndex float32
}

type NumbeoService struct {
	Client *resty.Client
}

func (n NumbeoService) parseLocationTable(body string) ([]NumbeoCostOfLivingRecord, error) {
	var records []NumbeoCostOfLivingRecord
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return []NumbeoCostOfLivingRecord{}, err
	}

	doc.Find("table#t2 tbody tr").Each(func(i int, s *goquery.Selection) {
		var l string
		var f []float32
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				break
			case 1:
				l = s.Text()
			default:
				stat, err := strconv.ParseFloat(s.Text(), 32)
				if err != nil {
					stat = -1
				}
				f = append(f, float32(stat))
			}
		})

		records = append(records, NumbeoCostOfLivingRecord{l, f[0], f[1], f[2], f[3], f[4], f[5]})
	})

	return records, nil
}

func (n NumbeoService) LocationSearch(location models.Location) (NumbeoCostOfLivingRecord, error) {
	resp, err := n.Client.R().
		Get(utils.NumbeoUrl)
	if err != nil {
		return NumbeoCostOfLivingRecord{}, err
	}

	records, err := n.parseLocationTable(string(resp.Body()))
	if err != nil {
		return NumbeoCostOfLivingRecord{}, err
	}

	index := slices.IndexFunc(records, func(row NumbeoCostOfLivingRecord) bool {
		l := strings.ToLower((row.Location))
		loc := strings.ToLower(fmt.Sprintf("%s, %s, %s", location.City, location.StateAbbrev, location.Country))
		return strings.Contains(l, loc)
	})

	if index == -1 {
		return NumbeoCostOfLivingRecord{}, errors.New("location not found")
	}

	return records[index], nil
}
