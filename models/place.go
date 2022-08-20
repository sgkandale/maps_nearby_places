package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"googlemaps.github.io/maps"
)

func WriteToFile(file *os.File, fileType string, places []maps.PlacesSearchResult) {
	if strings.EqualFold(fileType, "csv") {
		writeToCSV(file, places)
	} else if strings.EqualFold(fileType, "json") {
		writeToJSON(file, places)
	} else {
		log.Fatalf("[ERROR] unsupported output file type : %s", fileType)
	}
}

func writeToCSV(file *os.File, places []maps.PlacesSearchResult) {
	outputArr := [][]string{
		{
			"place_id",
			"name",
			"vicinity",
			"latitude",
			"longitude",
			"rating",
			"user_ratings_total",
			"types",
			"open_now",
			"price_level",
			"permanently_closed",
			"business_status",
		},
	}

	for _, eachPlace := range places {

		openNow := "-"
		if eachPlace.OpeningHours != nil {
			openNow = fmt.Sprintf("%t", *eachPlace.OpeningHours.OpenNow)
		}

		outputArr = append(
			outputArr,
			[]string{
				eachPlace.PlaceID,
				eachPlace.Name,
				eachPlace.Vicinity,
				fmt.Sprintf("%f", eachPlace.Geometry.Location.Lat),
				fmt.Sprintf("%f", eachPlace.Geometry.Location.Lng),
				fmt.Sprintf("%f", eachPlace.Rating),
				fmt.Sprintf("%d", eachPlace.UserRatingsTotal),
				strings.Join(eachPlace.Types, ","),
				openNow,
				fmt.Sprintf("%d", eachPlace.PriceLevel),
				fmt.Sprintf("%t", eachPlace.PermanentlyClosed),
				eachPlace.BusinessStatus,
			},
		)
	}

	writer := csv.NewWriter(file)
	err := writer.WriteAll(outputArr)
	if err != nil {
		log.Fatal("[ERROR]  writing places to csv file :", err)
	}
}

func writeToJSON(file *os.File, places []maps.PlacesSearchResult) {
	// jsonPlaces, err := json.Marshal(places)
	jsonPlaces, err := json.MarshalIndent(places, "", "    ")
	if err != nil {
		log.Fatal("[ERROR]  converting places to json : ", err)
	}

	_, err = file.Write(jsonPlaces)
	if err != nil {
		log.Fatal("[ERROR] writing places to json file : ", err)
	}
}
