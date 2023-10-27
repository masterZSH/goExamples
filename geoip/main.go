package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

func main() {

	db, err := geoip2.Open("./GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("81.2.69.142")
	record, err := db.City(ip)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}
