package arrays

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(elements ...[]int) []int {
	var res []int
	for _, v := range elements {
		res = append(res, Sum(v))
	}
	return res
}

func SumTails(elements ...[]int) []int {

	var sum []int
	for _, el := range elements {
		if len(el) > 1 {
			sum = append(sum, Sum(el[1:]))
		} else {
			sum = append(sum, 0)
		}
	}
	return sum
}
