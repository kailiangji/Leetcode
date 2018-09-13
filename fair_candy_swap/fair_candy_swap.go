package fair_candy_swap

func fairCandySwap(A []int, B []int) []int {
	var sumA, sumB int
	for i := range A {
		sumA += A[i]
	}
	for j := range B {
		sumB += B[j]
	}

	tmp := (sumB - sumA) / 2

	for i := range A {
		for j := range B {
			if B[j]-A[i] == tmp {
				return []int{A[i], B[j]}
			}
		}
	}
	panic("no such pair")
}
