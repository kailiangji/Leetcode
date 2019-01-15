package degree_of_an_array

func findShortestSubArray(nums []int) int {
	var l int = len(nums)
	var nonVisited int = l
	var currentDegree, degree int
	var end int
	var length int = l
	var tmp int
	for i := range nums {
		if nonVisited == 0 {
			return length
		}

		if nums[i] != -1 {
			tmp = nums[i]
			currentDegree = 0
			for j := i; j < l; j++ {
				if nums[j] == tmp {
					currentDegree++
					nonVisited--
					nums[j] = -1
					end = j
				}
			}
			currentLength := end - i + 1
			if currentDegree == degree {
				if currentLength < length {
					length = currentLength
				}
			} else if currentDegree > degree {
				degree = currentDegree
				length = currentLength
			}
		}
	}
	return length
}

type Tmp struct {
	degree int
	start  int
	end    int
	length int
}

func (this *Tmp) SetLength() {
	this.length = this.end - this.start + 1
}

func findShortestSubArray1(nums []int) int {
	m := make(map[int]*Tmp, 0)
	var degree, length int
	length = len(nums)

	for i := range nums {
		if m[nums[i]] == nil {
			m[nums[i]] = &Tmp{degree: 1, start: i, end: i, length: 1}
		} else {
			m[nums[i]].degree++
			m[nums[i]].end = i
			m[nums[i]].SetLength()
		}
	}

	for _, v := range m {
		if v.degree == degree {
			if v.length < length {
				length = v.length
			}
		}
		if v.degree > degree {
			degree = v.degree
			length = v.length
		}
	}

	return length

}
