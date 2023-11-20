package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func (w *Weather) checkStatus() (resWater string, resWind string) {
	switch {
	case w.Water <= 5:
		resWater = "safe"
	case w.Water >= 6 && w.Water <= 8:
		resWater = "caution"
	case w.Water > 8:
		resWater = "danger"
	}

	switch {
	case w.Wind <= 6:
		resWind = "safe"
	case w.Wind >= 7 && w.Wind <= 15:
		resWind = "caution"
	case w.Wind >= 15:
		resWind = "danger"
	}
	return resWater, resWind
}

func generateJSON(weather *Weather) {
	res, err := json.Marshal(&weather)
	if err != nil {
		log.Fatalf("JSON marshaling error: %s", err.Error())
	}

	fmt.Printf("%s\n", res)
	resWater, resWind := weather.checkStatus()
	fmt.Printf("Water status: %s\n", resWater)
	fmt.Printf("Wind status: %s\n", resWind)
}

func main() {
	weather := Weather{}

	for {
		time.Sleep(15 * time.Second)

		// Generate a shuffled list of indices
		shuffledIndices := rand.Perm(15)

		// Use the shuffled indices to determine the order of random values
		weather.Water = shuffledIndices[0]
		weather.Wind = shuffledIndices[1]

		generateJSON(&weather)
	}
}

