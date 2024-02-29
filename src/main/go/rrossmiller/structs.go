package main

import "fmt"

type Station struct {
	count float32
	sum   float32
	Name  string  `json:"name"`
	Min   float32 `json:"min"`
	Mean  float32 `json:"mean"`
	Max   float32 `json:"max"`
}

func (s Station) String() string {
	return fmt.Sprintf("%s: %v | %v| %v", s.Name, s.Min, s.Mean, s.Max)
}
