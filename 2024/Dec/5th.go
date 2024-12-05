package Dec

/*
2337. Move Pieces to Obtain a String
You are given two strings start and target, both of length n. Each string consists only of the characters 'L', 'R', and '_' where:

The characters 'L' and 'R' represent pieces, where a piece 'L' can move to the left only if there is a blank space directly to its left, and a piece 'R' can move to the right only if there is a blank space directly to its right.
The character '_' represents a blank space that can be occupied by any of the 'L' or 'R' pieces.
Return true if it is possible to obtain the string target by moving the pieces of the string start any number of times. Otherwise, return false.



Example 1:

Input: start = "_L__R__R_", target = "L______RR"
Output: true
Explanation: We can obtain the string target from start by doing the following moves:
- Move the first piece one step to the left, start becomes equal to "L___R__R_".
- Move the last piece one step to the right, start becomes equal to "L___R___R".
- Move the second piece three steps to the right, start becomes equal to "L______RR".
Since it is possible to get the string target from start, we return true.
Example 2:

Input: start = "R_L_", target = "__LR"
Output: false
Explanation: The 'R' piece in the string start can move one step to the right to obtain "_RL_".
After that, no pieces can move anymore, so it is impossible to obtain the string target from start.
Example 3:

Input: start = "_R", target = "R_"
Output: false
Explanation: The piece in the string start can move only to the right, so it is impossible to obtain the string target from start.


Constraints:

n == start.length == target.length
1 <= n <= 105
start and target consist of the characters 'L', 'R', and '_'.
*/

func CanChange(start string, target string) bool {
	n := len(start)
	lst, rst, lta, rta := 0, 0, 0, 0

	for i := 0; i < n; i++ {
		if start[i] == 'L' {
			lst++
		} else if start[i] == 'R' {
			rst++
		}

		if target[i] == 'L' {
			lta++
		} else if target[i] == 'R' {
			rta++
		}
	}

	//check counts
	if lst != lta || rst != rta {
		return false
	}

	startIndex, targetIndex := 0, 0
	for startIndex < n && targetIndex < n {
		// skip '_'
		for startIndex < n && start[startIndex] == '_' {
			startIndex++
		}
		for targetIndex < n && target[targetIndex] == '_' {
			targetIndex++
		}

		// if both of them reach the boundaries meaning success
		if startIndex == n && targetIndex == n {
			return true
		}

		// if only one of them reach the bound meaning fail
		if startIndex == n || targetIndex == n {
			return false
		}

		// check if the character matches
		if start[startIndex] != target[targetIndex] {
			return false
		}

		if start[startIndex] == 'L' && startIndex < targetIndex {
			return false // 'L' cant move to right
		}
		if start[startIndex] == 'R' && startIndex > targetIndex {
			return false // 'R' cant move to left
		}

		startIndex++
		targetIndex++
	}

	return true
}
