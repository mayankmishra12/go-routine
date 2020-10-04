package utility

func IsPrimeNumber(number int) bool{

	isPrime := true
	if number == 0 || number == 1 {
		isPrime = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}
	return isPrime
}
