package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/daltonscharff/city-stats/internal/models"
	"github.com/daltonscharff/city-stats/internal/services"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

type FindFlags struct {
	showClimate       bool
	showNeighborhoods bool
}

var (
	flags   FindFlags
	findCmd = &cobra.Command{
		Use:   "find <city>",
		Short: "Find stats for a given city",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("climate", flags.showClimate)
			fmt.Println("neighborhoods", flags.showNeighborhoods)
			fmt.Println("args", len(args))
			find(args[0], flags)
		},
	}
)

func init() {
	findCmd.Flags().BoolVarP(&flags.showClimate, "climate", "c", false, "show climate data")
	findCmd.Flags().BoolVarP(&flags.showNeighborhoods, "neighborhoods", "n", false, "show WalkScore for individual neighborhoods")

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
}
