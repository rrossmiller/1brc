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
	n, err := strconv.ParseInt(os.Args[2], 10, 8)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Running %s %d times\n", pth, n)

	times := []time.Duration{}
	for i := range n {
		start := time.Now()
		run(pth)
		elapsed := time.Since(start)
		times = append(times, elapsed)
		fmt.Println(i)
	}
	t := meanTime(times)
	fmt.Println("avg", t)

	//results
	stations := run(pth)
	keys := make([]string, 0, len(stations))
	for k := range stations {
		keys = append(keys, k)
	}

	city := keys[0]
	fmt.Println(stations[city])
	j, err := json.MarshalIndent(stations, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("res.json", j, 0644)
}

func run(pth string) map[string]Station {
	f, err := os.Open(pth)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	stations := map[string]Station{}

	// separate
	//     scanner goroutine
	//     line handler goroutine
	// scanner sends lines to chan
	// handler reads from the chan and updates the map
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

	return stations
}
