package mar

import "math"

/*
2594. Minimum Time to Repair Cars
Medium
Topics
Companies
Hint
You are given an integer array ranks representing the ranks of some mechanics. ranksi is the rank of the ith mechanic. A mechanic with a rank r can repair n cars in r * n2 minutes.

You are also given an integer cars representing the total number of cars waiting in the garage to be repaired.

Return the minimum time taken to repair all the cars.

Note: All the mechanics can repair the cars simultaneously.



Example 1:

Input: ranks = [4,2,3,1], cars = 10
Output: 16
Explanation:
- The first mechanic will repair two cars. The time required is 4 * 2 * 2 = 16 minutes.
- The second mechanic will repair two cars. The time required is 2 * 2 * 2 = 8 minutes.
- The third mechanic will repair two cars. The time required is 3 * 2 * 2 = 12 minutes.
- The fourth mechanic will repair four cars. The time required is 1 * 4 * 4 = 16 minutes.
It can be proved that the cars cannot be repaired in less than 16 minutes.​​​​​
Example 2:

Input: ranks = [5,1,8], cars = 6
Output: 16
Explanation:
- The first mechanic will repair one car. The time required is 5 * 1 * 1 = 5 minutes.
- The second mechanic will repair four cars. The time required is 1 * 4 * 4 = 16 minutes.
- The third mechanic will repair one car. The time required is 8 * 1 * 1 = 8 minutes.
It can be proved that the cars cannot be repaired in less than 16 minutes.​​​​​


Constraints:

1 <= ranks.length <= 105
1 <= ranks[i] <= 100
1 <= cars <= 106
*/
func RepairCars(ranks []int, cars int) int64 {
	l, h := int64(1), min(ranks)*cars*cars

	for int(l) < h {
		m := (int(l) + h) / 2
		if canRepair(ranks, cars, m) {
			h = m
		} else {
			l = int64(m + 1)
		}
	}
	return l
}

func canRepair(ranks []int, cars, m int) bool {
	totalCars := 0
	for _, r := range ranks {
		totalCars += int(math.Sqrt(float64(m) / float64(r)))
	}
	return totalCars >= cars
}

func min(arr []int) int {
	minVal := arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func RepairCars2(ranks []int, cars int) int64 {
	// Find the maximum rank
	maxRank := 0
	for _, rank := range ranks {
		if rank > maxRank {
			maxRank = rank
		}
	}

	l, r := int64(1), int64(maxRank)*int64(cars)*int64(cars)
	for int64(l) <= r {
		// middle time
		m := l + (r-l)/2
		if check(ranks, cars, m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return int64(l)
}

func check(ranks []int, cars int, time int64) bool {
	totalCarsFixed := int64(0)
	for _, rank := range ranks {
		carsFixed := int64(math.Floor(math.Sqrt(float64(time) / float64(rank))))
		totalCarsFixed += carsFixed
	}
	return totalCarsFixed >= int64(cars)
}
