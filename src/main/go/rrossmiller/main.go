package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	pth := os.Args[1]
	start := time.Now()
	f, err := os.Open(pth)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	stations := map[string]Station{}

	for scanner.Scan() {
		l := scanner.Text()
		spl := strings.Split(l, ";")
		v, err := strconv.ParseFloat(spl[1], 32)
		if err != nil {
			panic(err)
		}

		temp := float32(v)

		if s, has := stations[spl[0]]; has {
			s.count++
			s.sum += temp
			if s.Min > temp {
				s.Min = temp
			} else if s.Max < temp {
				s.Max = temp
			}
			stations[s.Name] = s
		} else {
			stations[spl[0]] = Station{Name: spl[0], sum: temp, Min: temp, Max: temp, count: 1}
		}
	}

	for k, v := range stations {
		v.Mean = v.sum / v.count
		stations[k] = v

	}
	//results
	elapsed := time.Since(start)
	fmt.Println(len(stations))
	n := "Stillwater"
	fmt.Println(stations[n])
	fmt.Println(elapsed)
	j, err := json.MarshalIndent(stations, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("res.json", j, 0644)
}
