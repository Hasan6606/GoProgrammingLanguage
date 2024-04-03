package mathexp

import "fmt"

//FirstExample function is discriminant calculation function.
// aVar = a || bVar = b || cVar = c
// [(b*b)-4*a*c]<0 does not have root.
// [(b*b)-4*a*c]>0 have different two roots.
// [(b*b)-4*a*c]>0 have same  two roots.
func FirstExample(aVar int16, bVar int16, cVar int16) {

	var eqResult int16 = int16(bVar)*int16(bVar) - 4*int16(aVar)*int16(cVar)

	if eqResult > 0 {
		fmt.Println("Discriminant value is biger than zero. These equation have two different roots.")
	} else if eqResult == 0 {
		fmt.Println("Discriminant value is equal to zero. These equation have two same roots.")
	} else {
		fmt.Println("Discriminant value is less than zero. These equation does not have root(s).")
	}
}
