package src

import "math"

func robotSim(commands []int, obstacles [][]int) int {
	// N-0 E-1 S-2 W-3
	direction := 0
	point := [2]int{0, 0}
	for _, v := range commands {
		switch {
		case v == -1:
			direction = (direction + 1) % 4
		case v == -2:
			direction = (direction - 1) % 4
		default:
			switch direction {
			case 0:
				min := math.MaxInt64
				point[1] += v
				for _,ob := range obstacles{
					if ob[0] == point[0] && ob[1] < min{
						min = ob[1]
					}
				}
				if min < point[1]{
					point[1] = min-1
				}
			case 1:
				min := math.MaxInt64
				point[0] += v
				for _,ob := range obstacles{
					if ob[1] == point[1] && ob[0] < min{
						min = ob[0]
					}
				}
				if min < point[0]{
					point[0] = min-1
				}
			case 2:
				max := -math.MaxInt64
				point[1] -= v
				for _,ob := range obstacles{
					if ob[0] == point[0] && ob[1] > max{
						max = ob[1]
					}
				}
				if max > point[1]{
					point[1] = max+1
				}
			case 3:
				max := -math.MaxInt64
				point[0] -= v
				for _,ob := range obstacles{
					if ob[1] == point[1] && ob[0] > max{
						max = ob[0]
					}
				}
				if max > point[0]{
					point[0] = max+1
				}
			}
		}
	}
	return point[0] * point[0] +point[1] * point[1]
}
