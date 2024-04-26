package main

import (
	"testing"
)

func Test(t *testing.T) {
	f := Round(1.22)
	if f != 1.2 {
		t.Fatalf("1.2 != %v", f)
	}
	f = Round(-1.2)
	if f != -1.2 {
		t.Fatalf("-1.2 != %v", f)
	}

	i := 25.449999999999996
	ans := 25.4
	f = Round(i)
	if f != ans {
		t.Fatalf("25.5 != %v", f)
	}

	i = 17.94999999999999
	ans = 17.9
	f = Round(i)
	if f != ans {
		t.Fatalf("18.0 != %v", f)
	}

}

// func Test1(t *testing.T) {
// 	ans := 1.3
// 	stations := run("test_data/test_bosaso.txt")
// 	if x := stations["Bosaso"].Mean; ans != x {
// 		fmt.Printf("%f\n", x)
// 		t.Fatalf("%.1f", x)
// 	}
// }
//
// func Test_rounding(t *testing.T) {
// 	ans := map[string]float64{
// 		"ham": 25.5,
// 		"jel": 18.0,
// 	}
// 	stations := run("test_data/test_rounding.txt")
//
// 	for k, v := range stations {
// 		if ans[k] != v.Mean {
// 			fmt.Printf("%q: %f | %f\n", k, v.Mean, ans[k])
// 			t.Fatalf("%q: %.1f | %f", k, v.Mean, ans[k])
// 		}
//
// 	}
// }
