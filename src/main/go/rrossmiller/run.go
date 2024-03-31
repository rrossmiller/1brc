package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

// separate
//     scanner goroutine
//     line handler goroutine
// scanner sends lines to chan
// handler reads from the chan and updates the map

// naiive
var semi = []byte(";")

func run(pth string) map[string]Station {
	f, err := os.Open(pth)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	stations := map[string]Station{}

	for scanner.Scan() {
		l := scanner.Bytes()
		spl := bytes.Split(l, semi)
		v, err := strconv.ParseFloat(string(spl[1]), 32)

		if err != nil {
			panic(err)
		}

		temp := float32(v)

		k := string(spl[0])
		if s, has := stations[k]; has {
			s.count++
			s.sum += temp
			if s.Min > temp {
				s.Min = temp
			} else if s.Max < temp {
				s.Max = temp
			}
			stations[s.Name] = s
		} else {
			stations[k] = Station{Name: k, sum: temp, Min: temp, Max: temp, count: 1}
		}
	}

	for k, v := range stations {
		v.Mean = v.sum / v.count
		stations[k] = v

	}

	return stations
}

// separate scanner and line handler into separate goroutines
// func run2(pth string) map[string]Station {
// 	f, err := os.Open(pth)
// 	defer f.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	scanner := bufio.NewScanner(f)
//
// 	// stations := sync.Map[string]Station{}
// 	var stations sync.Map
//
// 	for scanner.Scan() {
// 		l := scanner.Text()
//
// 		spl := strings.Split(l, ";")
// 		v, err := strconv.ParseFloat(spl[1], 32)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		temp := float32(v)
//
// 		if s, has := stations[spl[0]]; has {
// 			s.count++
// 			s.sum += temp
// 			if s.Min > temp {
// 				s.Min = temp
// 			} else if s.Max < temp {
// 				s.Max = temp
// 			}
// 			stations[s.Name] = s
// 		} else {
// 			stations[spl[0]] = Station{Name: spl[0], sum: temp, Min: temp, Max: temp, count: 1}
// 		}
// 	}
//
// 	for k, v := range stations {
// 		v.Mean = v.sum / v.count
// 		stations[k] = v
//
// 	}
//
// 	return stations
// }
