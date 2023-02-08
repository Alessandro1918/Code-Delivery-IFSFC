package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID 				string 			`json:"routeId"`
	ClientID 	string 			`json:"clientId"`
	Positions []Position 	`json:"position"`
}

type Position struct {
	Lat 			float64 		`json:"lat"`
	Long 			float64 		`json:"long"`
}

type PartialRoutePosition struct {
	ID 				string 			`json:"routeId"`
	ClientID 	string 			`json:"clientId"`
	Position 	[]float64 	`json:"position"`
	Finished 	bool 				`json:"finished"`
}

//Creates a *Route struct
func NewRoute() *Route {
	return &Route{}
}

//object method: load data from txt file to my "route" object's Positions
//returns: error
func(r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route ID not provided")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(r.Positions, Position{
			Lat: lat,
			Long: long,
		})
	}
	return nil
}

//object method: send positions to Kafka as a list of strings 
//returns: array of strings, error
func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string

	//k: key, v: value
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if len(r.Positions) - 1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)				//assemble json
		if err != nil {
			return nil, err														//empty list element, the error that cause it
		}
		result = append(result, string(jsonRoute))	//stringify json
	}
	return result, nil														//list of lat/longs, error (nil, because it was successfull) 
}