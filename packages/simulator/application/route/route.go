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
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("Route id not provided")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)

		if err != nil {
			return nil
		}

		long, err := strconv.ParseFloat(data[1], 64)

		if err != nil {
			return nil
		}

		route.Positions = append(route.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}

func (route *Route) ExportJsonPositions() ([]string, error) {
	var partialRoutePosition PartialRoutePosition
	var result []string
	total := len(route.Positions)

	for k, v := range route.Positions {
		partialRoutePosition.ID = route.ID
		partialRoutePosition.ClientID = route.ClientID
		partialRoutePosition.Position = []float64{v.Lat, v.Long}
		partialRoutePosition.Finished = false

		if total-1 == k {
			partialRoutePosition.Finished = true
		}

		jsonRoute, err := json.Marshal(partialRoutePosition)

		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}
