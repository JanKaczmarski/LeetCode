package stack

type day struct {
	id   int
	temp int
}

func dailyTemperatures(temperatures []int) []int {
	// Going from left to right. Putting temperature on the stack and its day idx, and going forward with
	// For each new element pop every element on stack that has smaller value then current day
	stack := make([]day, 0)
	res := make([]int, len(temperatures))

	stack = append(stack, day{
		id:   0,
		temp: temperatures[0],
	})
	for i, temp := range temperatures[1:] {
		// offset as result of shifted range
		i += 1

		// going through stack removing each element that has smaller temp
		// This is guaranted to by true, because (proof) if bigger temp was found earlier
		// that earlier temp would delete element from the stack. Hence it's guaranteed that
		// current temp that is bigger that elements on the stack is the closest one
		for len(stack) > 0 && stack[len(stack)-1].temp < temp {
			prevday := stack[len(stack)-1]
			res[prevday.id] = i - prevday.id

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, day{
			id:   i,
			temp: temp,
		})
	}

	// Setting 0 for each day that bigger temp couldn't be reached
	for _, day := range stack {
		res[day.id] = 0
	}

	return res
}
