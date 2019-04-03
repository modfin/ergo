package fac



func Factorial(i int64) int64{
	if i < 2 {
		return 1
	}
	return Factorial(i-1)*i
}