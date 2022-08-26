package osm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Create a new Postgres DB
func NewDB() (*sql.DB, error) {
	// Name the DB
	dbName := "osm"
	// Connect to the DB
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s sslmode=disable", dbName))
	if err != nil {
		return nil, err
	}
}

// Create a Postgres table to store the map
func StoreMap(db *sql.DB, data map[string]interface{}) error {
	// Use the DB from the NewDB function
	sql.ColumnType.DatabaseTypeName("osm")
	// Create a new table
	sql.CreateTable("osm")
	// Return the error
	return nil
}

// Receive the user's location or the location of the map they want to view or the starting point of their route
func Location() {
	// Get the user's location
	location := http.Get("http://localhost:8080/location")
	// Write the user's location to the DB
	sql.WriteToDB("location")
	// Flush the table if the location is not found or changes by more than 5 meters
	if location == nil {
		sql.FlushTable("location")
	}
	// Declare a 20 item array to store the user's location history
	var locationHistory [20]string
	// Store the user's location history in the array
	locationHistory = append(locationHistory, location)
	// If the user's location changes by more than 5 meters, store the new location in the array
	if location != locationHistory[0] {
		locationHistory = append(locationHistory, location)
	}
	locationOffset := 5
	if location > locationOffset {
		sql.FlushTable("location")
	}
	// Get the location of the map they want to view on openstreetmap.org
	map := http.Get("http://localhost:8080/map")
	// Write the location of the map they want to view to the DB. Drop the table and create a new one if the user wants to view a different map
	sql.WriteToDB("map")
	//flush the table if the user wants to view a different map
	if map != map {
		sql.FlushTable("map")
	}
}

	func routing() {

	// Get the coordinates of the starting point of the user's route on openstreetmap.org
	start := http.Get("http://localhost:8080/start")
	// Write the coordinates of the starting point of the user's route to the DB
	sql.WriteToDB("start")
	// Flush the table if the user wants to start a different route
	if start != start {
		sql.FlushTable("start")
	}
	// Get the coordinates of the end point of the user's route on openstreetmap.org, 
	end := http.Get("http://localhost:8080/end")
	// Write the coordinates of the end point of the user's route to the DB
	sql.WriteToDB("end")
	if end != end {
		sql.FlushTable("end")
	}
	// Apart from the standard routing algorithms, we will also implement a routing algorithm that takes into account whether there are any obstacles on the route or other users on the route and if so, fetch their speeds.
	// get the standard route from the start point to the end point on openstreetmap.org
}

// Connect to the Openstreetmap API and fetch a map
func FetchMap(lat, lon float64) (map[string]interface{}, error) {
	//First, we need to know whether to fetch map tiles for the user's current location or for the location of the map they want to view or for the starting point of their route
	var useLocation bool
	var useMap bool
	var useStart bool
	if location != nil {
		useLocation = true
		lat = location.Lat
		lon = location.Lon
	}
	if map != nil {
		useMap = true
		lat = map.lat
		lon = map.lon
	}
	if start != nil {
		useStart = true
		lat = start.Lat
		lon = start.Lon
	}

	// Connect to the API
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openstreetmap.org/api/0.6/map?bbox=%f,%f,%f,%f", lon-0.01, lat-0.01, lon+0.01, lat+0.01), nil
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "osm-fetch")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Parse the response
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	// Save the response to the osm table of the osm DB
	if err := StoreMap(data); err != nil {
		sql.InsertInto("osm")
		return nil, err
	}
	return data, nil
}

