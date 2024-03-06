package main

import (
	"fmt"
	"math"
)

func main() {
	var l, x1, v1, x2, v2 int64
	_, err := fmt.Scanf("%d %d %d %d %d", &l, &x1, &v1, &x2, &v2)
	if err != nil || l < 1 || l > 1000000000 || v1 > 1000000000 || v2 > 1000000000 || x1 < 0 || x2 < 0 || x1 >= l || x2 >= l || v1 < 0 || v2 < 0 {
		fmt.Print("Read-mistake: ", err)
		return
	}

	var t float64
	X1 := float64(x1)
	X2 := float64(x2)
	V1 := float64(v1)
	V2 := float64(v2)
	L := float64(l)
	for t <= 100 {
		if ((X1+V1*t)*360/L)-float64(int64(((X1+V1*t)*360/L)/360)*360) == ((X2+V2*t)*360/L)-float64(int64(((X2+V2*t)*360/L)/360)*360) {
			fmt.Print(t)
			break
		}
		if math.Floor(((X1+V1*t)*360/L)+((X2+V2*t)*360/L)) == ((X1+V1*t)*360/L)+((X2+V2*t)*360/L) {
			if int64(((X1+V1*t)*360/L)+((X2+V2*t)*360/L))%360 == 0 {
				fmt.Print(t)
				break
			}
		}
		t += 0.0000000001
	}

	if t <= 100 {
		fmt.Print("YES\n", t)
	} else {
		fmt.Print("NO")
	}
}
