package tasks

func TwoSum(nums []int, target int) []int {
	ids := []int{}
	l := len(nums)

	for i := 0; i < l; i++ {
		ids = append(ids, i)
	}

	array := nums

	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - i - 1; j++ {
			if array[j] > array[j + 1] {
				array[j], array[j + 1] = array[j + 1], array[j]
				ids[j], ids[j + 1] = ids[j + 1], ids[j]
			}
		}
	}

	i := 0
	j := l - 1

	for i < j {
		if nums[i] + nums[j] == target {
			return []int{ids[i], ids[j]}
		}

		if nums[i] + nums[j] > target {
			j--
		}

		if nums[i] + nums[j] < target {
			i++
		}
	}

	return []int{}
}

func TwoSumSimple(nums []int, target int) []int {
	l := len(nums)

	for i := 0; i < l - 1; i++ {
		for j := 1; j < l; j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}

	return []int{}
}