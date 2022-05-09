package stud3

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}

func SumAllNumbers(numbers ...int) int{
	var sum = 0
	for _, value := range numbers{
		sum = sum + value
	}
	return sum
}

func SumPositiveNumbers(numbers ...int) int{
	var sum = 0
	for _, value := range numbers{
		if value < 0 {
			continue	
		}
		sum = sum + value
	}
	return sum
}