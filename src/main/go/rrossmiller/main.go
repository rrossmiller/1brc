package main

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	// get the path to the fill
	pth := "measurements.txt"

	stations := run(pth) // run one more time to see results
	// keys := make([]string, 0, len(stations))
	// for k := range stations {
	// keys = append(keys, k)
	// }
	stationList := maps.Values(stations)
	slices.SortFunc(stationList, func(a, b Station) int {
		if a.Name < b.Name {
			return -1
		}
		return 1
	})
	var sb strings.Builder
	for i, s := range stationList {
		if i < len(stationList)-1 {
			sb.WriteString(s.String() + ", ")
		} else {
			sb.WriteString(s.String())
		}
	}
	fmt.Printf("{%v}\n", sb.String())
}

// func main() {
// 	// set the function to test
// 	f := run
//
// 	// get the path to the file
// 	pth := os.Args[1]
// 	// get the number of times to run
// 	n, err := strconv.ParseInt(os.Args[2], 10, 8)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Running %s %d times\n", pth, n)
//
// 	if len(os.Args) > 3 {
// 		fmt.Println("profiling")
// 		f, err := os.Create("cpu_profile.prof")
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer f.Close()
//
// 		if err := pprof.StartCPUProfile(f); err != nil {
// 			panic(err)
// 		}
// 		defer pprof.StopCPUProfile()
// 	}
//
// 	// run n times and store the run times
// 	if n > 1 {
// 		times := []time.Duration{}
// 		for i := range n {
// 			start := time.Now()
// 			f(pth)
// 			elapsed := time.Since(start)
// 			times = append(times, elapsed)
// 			fmt.Printf("%d, ", i+1)
// 		}
// 		fmt.Println()
// 		t := meanTime(times) // mean run time
// 		fmt.Println("avg", t)
// 	}
//
// 	//results
// 	stations := f(pth) // run one more time to see results
// 	keys := make([]string, 0, len(stations))
// 	for k := range stations {
// 		keys = append(keys, k)
// 	}
//
// 	city := keys[0]
// 	fmt.Println(stations[city])
// 	j, err := json.MarshalIndent(stations, "", "  ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	os.WriteFile("res.json", j, 0644)
// }
