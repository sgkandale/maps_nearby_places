package config

import (
	"log"
	"strings"
)

func GetVerifiedConfig() config {
	parsedConfig := ParseConfig()

	if parsedConfig.APIKey == "" {
		log.Fatal("[ERROR] api key is not set")
	}
	if parsedConfig.MinPrice < 0 || parsedConfig.MinPrice > 4 {
		log.Fatal("[ERROR] invalid min price")
	}
	if parsedConfig.MaxPrice < 0 || parsedConfig.MaxPrice > 4 {
		log.Fatal("[ERROR] invalid max price")
	}
	if parsedConfig.MaxPrice == 0 {
		parsedConfig.MaxPrice = 4
	}
	if parsedConfig.RankBy != "prominence" && parsedConfig.RankBy != "distance" {
		log.Fatal("[ERROR] invalid rank by, allowed values are : 'prominence', 'distance'")
	}
	if parsedConfig.MaxRequests == 0 {
		parsedConfig.MaxRequests = 50
	}
	if parsedConfig.MaxRequests > 50 {
		log.Fatal("[ERROR] max requests should not be more than 50")
	}

	// OutputFile checks
	if parsedConfig.OutputFile.Name == "" {
		log.Fatal("[ERROR] output file name not set")
	}
	if parsedConfig.OutputFile.Type == "" {
		log.Fatal("[ERROR] output file type not set")
	}
	if !strings.EqualFold(parsedConfig.OutputFile.Type, "csv") && !strings.EqualFold(parsedConfig.OutputFile.Type, "json") {
		log.Fatal("[ERROR] invalid output file type, allowed values are : csv, json")
	}

	return parsedConfig
}
