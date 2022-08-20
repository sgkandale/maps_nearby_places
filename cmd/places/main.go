package main

import (
	"context"
	"log"
	"os"
	"time"

	"places/config"
	"places/models"

	"googlemaps.github.io/maps"
)

func main() {

	// get config
	cfg := config.GetVerifiedConfig()

	// output file
	outputFile, err := os.Create(cfg.OutputFile.Name + "." + cfg.OutputFile.Type)
	if err != nil {
		log.Fatal("[ERROR] creating file : ", err)
	}

	placesRes := make([]maps.PlacesSearchResult, 0)
	ctx := context.Background()
	pageToken := ""

	// search request
	req := &maps.NearbySearchRequest{
		Radius:  cfg.RadiusInMeters,
		Keyword: cfg.Keyword,
		Name:    cfg.Name,
		OpenNow: cfg.OpenNow,
		Location: &maps.LatLng{
			Lat: cfg.Latitude,
			Lng: cfg.Longitude,
		},
		MinPrice: getPriceLevel(cfg.MinPrice),
		MaxPrice: getPriceLevel(cfg.MaxPrice),
		RankBy:   getRankBy(cfg.RankBy),
	}

	if cfg.PlaceType != "" {
		t, err := maps.ParsePlaceType(cfg.PlaceType)
		if err != nil {
			log.Fatalf("[ERROR] unknown place type : %s", cfg.PlaceType)
		}
		req.Type = t
	}

	// maps client
	client, err := maps.NewClient(
		maps.WithAPIKey(cfg.APIKey),
	)
	if err != nil {
		log.Fatal("[ERROR] creating maps client : ", err)
	}

	requestCount := 0
	for {
		if requestCount > cfg.MaxRequests {
			break
		}

		reqCtx, cancelReqCtx := context.WithTimeout(ctx, time.Second*20)
		defer cancelReqCtx()
		req.PageToken = pageToken

		resp, err := client.NearbySearch(
			reqCtx,
			req,
		)
		if err != nil {
			log.Fatal("[ERROR] searching for places : ", err)
		}
		if len(resp.Results) == 0 {
			break
		}
		placesRes = append(placesRes, resp.Results...)
		requestCount++
		pageToken = resp.NextPageToken
	}

	log.Printf("[INFO] got total %d places", len(placesRes))

	models.WriteToFile(outputFile, cfg.OutputFile.Type, placesRes)
}

func getPriceLevel(priceLevel int) maps.PriceLevel {
	switch priceLevel {
	case 0:
		return maps.PriceLevelFree
	case 1:
		return maps.PriceLevelInexpensive
	case 2:
		return maps.PriceLevelModerate
	case 3:
		return maps.PriceLevelExpensive
	case 4:
		return maps.PriceLevelVeryExpensive
	default:
		log.Fatalf("[ERROR] unknown price level: %d", priceLevel)
	}
	return maps.PriceLevelFree
}

func getRankBy(rankBy string) maps.RankBy {
	switch rankBy {
	case "prominence":
		return maps.RankByProminence
	case "distance":
		return maps.RankByDistance
	default:
		return maps.RankByDistance
	}
}
