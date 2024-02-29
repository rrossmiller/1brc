package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
name=min/mean/max
{Abha=-23.0/18.0/59.2
*/
func main() {
	pth := os.Args[1]
	start := time.Now()
	f, err := os.Open(pth)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	stations := []Station{}

	for scanner.Scan() {
		l := scanner.Text()
		spl := strings.Split(l, ";")
		v, err := strconv.ParseFloat(spl[1], 32)
		if err != nil {
			panic(err)
		}
		station := Station{name: spl[0], temp: float32(v)}
		stations = append(stations, station)
	}

	elapsed := time.Since(start)
	fmt.Println(len(stations))
	fmt.Println(stations[0])
	fmt.Println(elapsed)
}
