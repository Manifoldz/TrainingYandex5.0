package main

import (
	"fmt"
	"math"
)

func main() {
	var l, x1, v1, x2, v2 int64
	_, err := fmt.Scanf("%d %d %d %d %d", &l, &x1, &v1, &x2, &v2)
	var t float64
	X1 := float64(x1)
	X2 := float64(x2)
	V1 := float64(v1)
	V2 := float64(v2)
	L := float64(l)

	if err != nil || l < 1 || l > 1000000000 || int64(math.Abs(V1)) > 1000000000 || int64(math.Abs(V2)) > 1000000000 || x1 < 0 || x2 < 0 || x1 >= l || x2 >= l {
		fmt.Print("Read-mistake: ", err)
		return
	}
	if x1 == x2 {
		t = 0
	} else if V1 == 0 && V2 == 0 { //точки неподвижны в разных местах и исхода нет
		fmt.Print("NO")
		return
	} else if (V1 == 0 && X1 < X2) || (V2 == 0 && X1 > X2) || V1 == V2 || (V1 < 0 && V2 < 0) { //если точки с одинаковой скоростью движутся или одна из них неподвижна
		//L-X1-V1t = X2+V2t
		if V1 < 0 && V2 < 0 { // особый случай - в обратку
			t = (-X1 - X2) / (V2 + V1)
		} else { //из замеченного условия что всегда в сумме суперпозиций получается сумма круга
			t = (L - (X1 + X2)) / (V2 + V1)
		}
	} else {
		if X2 > X1 && V1 <= 0 && V2 > 0 { // в обратку через длинную сторону
			t = (L - X2 + X1) / (V1 - V2)
		} else if X1 > X2 && V2 <= 0 && V1 > 0 {
			t = (L - X1 + X2) / (V1 - V2)
		} else { //стандартный случай
			if ((L-X1 < X1-X2) || L-X2 < X2-X1) && (V1 >= 0 && V2 >= 0 || V1 <= 0 && V2 <= 0) {
				t = (L - X1 - X2) / (V1 - V2)
				temp := float64((L - (X1 + X2 - float64((x1+x2)/l*l))) / (V2 + V1))
				if temp < math.Abs(t) {
					t = temp
				}
				//fmt.Print("\n", temp < t, "\n")
			} else {
				t = (X2 - X1) / (V1 - V2)
			}
		}
	}

	fmt.Print("YES\n", math.Abs(t))

}
