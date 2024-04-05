package recursive

func Factorial(numb int32) int32 {
	//if the enter number 0. The result have returned one.
	if numb == 0 {
		return 1
	}
	//if number biger than 0. The result have gone recycling as method.
	return numb * Factorial(numb-1)

}
