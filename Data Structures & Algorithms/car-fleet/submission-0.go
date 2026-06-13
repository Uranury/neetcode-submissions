type Car struct {
	Position int
	Speed int
}

func carFleet(target int, position []int, speed []int) int {
	fleets := len(position)
	cars := make([]Car, len(position))
	for i := range position {
		cars[i] = Car{Position: position[i], Speed: speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
    	return cars[i].Position > cars[j].Position
	})
	
	stack := []float64{}
	for i := 0; i < len(cars); i++ {
		if len(stack) > 0 && computeReach(target, cars[i].Position, cars[i].Speed) <= stack[len(stack)-1] {
			fleets--
		} else {
			stack = append(stack, computeReach(target, cars[i].Position, cars[i].Speed))
		}
	}

	return fleets
}

func computeReach(target, position, speed int) float64 {
	return float64(target-position) / float64(speed)
}