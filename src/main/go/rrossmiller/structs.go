package main

import (
	"fmt"
	"math"
	"strconv"
)

type Station struct {
	count int
	sum   float64
	Name  string  `json:"name"`
	Min   float64 `json:"min"`
	Mean  float64 `json:"mean"`
	Max   float64 `json:"max"`
	// Min   string `json:"min"`
	// Mean  string `json:"mean"`
	// Max   string `json:"max"`
}

func (s Station) String() string {
	//					name, min mean, max
	return fmt.Sprintf("%s=%0.1f/%0.1f/%0.1f", s.Name, s.Min, s.Mean, s.Max)
	// return fmt.Sprintf("%s=%f/%f/%f", s.Name, ParseFloatFast(s.Min), ParseFloatFast(s.Mean), ParseFloatFast(s.Max))
	// return fmt.Sprintf("%s=%f/%f/%f", s.Name, s.Min, s.Mean, s.Max)
}

func ParseFloat(s []byte) float64 {
	x, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		panic(err)
	}
	return x
}

func Round(num float64) float64 {
	num *= 10
	x := int(num + math.Copysign(0.5, num))
	return float64(x) / 10
}
