package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/daltonscharff/city-stats/internal/models"
	"github.com/daltonscharff/city-stats/internal/services"
	"github.com/go-resty/resty/v2"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

type FindFlags struct {
	showTables bool
}

var (
	flags   FindFlags
	findCmd = &cobra.Command{
		Use:   "find <city>",
		Short: "Find stats for a given city",
		Run: func(cmd *cobra.Command, args []string) {
			find(args[0], flags)
		},
	}
)

func init() {
	findCmd.Flags().BoolVarP(&flags.showTables, "tables", "t", false, "show tables for climate and neighborhood walkscore data")

	rootCmd.AddCommand(findCmd)
}

func find(location string, flags FindFlags) {
	client := resty.New()

	wikipediaService := services.WikipediaService{Client: client}
	wikipediaResponse, err := wikipediaService.LocationSearch(location)
	if err != nil {
		fmt.Println("wikipediaService.LocationSearch error:", err)
	}

	loc := models.Location{}
	loc.New(wikipediaResponse.City, wikipediaResponse.State, "United States")

	walkscoreService := services.WalkscoreService{Client: client}
	walkscoreResponse, err := walkscoreService.LocationSearch(loc)
	if err != nil {
		fmt.Println("walkscoreService.LocationSearch error:", err)
	}

	numbeoService := services.NumbeoService{Client: client}
	numbeoResponse, err := numbeoService.LocationSearch(loc)
	if err != nil {
		fmt.Println("numbeoService.LocationSearch error:", err)
	}

	response := map[string]any{
		"location":  loc,
		"wikipedia": wikipediaResponse,
		"walkscore": walkscoreResponse,
		"numbeo":    numbeoResponse,
	}

	json, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error formatting json:", err)
	}

	fmt.Println(string(json))

	if flags.showTables {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December", "Year"})

		for _, record := range wikipediaResponse.ClimateTable {
			t.AppendRow([]interface{}{
				record.Name,
				record.Month[0],
				record.Month[1],
				record.Month[2],
				record.Month[3],
				record.Month[4],
				record.Month[5],
				record.Month[6],
				record.Month[7],
				record.Month[8],
				record.Month[9],
				record.Month[10],
				record.Month[11],
				record.Year,
			})
		}
		t.Render()

		tt := table.NewWriter()
		tt.SetOutputMirror(os.Stdout)
		tt.AppendHeader(table.Row{"Neighborhood", "Population", "Walk", "Transit", "Bike"})

		for _, record := range walkscoreResponse.NeighborhoodScores {
			tt.AppendRow([]interface{}{
				record.Name,
				record.Population,
				record.Score.Walk,
				record.Score.Transit,
				record.Score.Bike,
			})
		}
		tt.Render()
	}
}
