package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	url, err := ioutil.ReadFile("url")
	if err != nil {
		log.Fatal(err)
	}
	// bucket := "example-bucket"
	org := "example-org"
	// token := "example-token"
	// Store the URL of your InfluxDB instance
	// Create client
	userName := "zsh"
	password := "123456"

	client := influxdb2.NewClient(string(url), fmt.Sprintf("%s:%s", userName, password))
	defer client.Close()

	// Get query client
	queryAPI := client.QueryAPI(org)
	// Get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket:"my-bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// Check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
}
