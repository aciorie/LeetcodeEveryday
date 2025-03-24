package mar

import "sort"

/*
3169. Count Days Without Meetings
Medium
Topics
Companies
Hint
You are given a positive integer days representing the total number of days an employee is available for work (starting from day 1). You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i] represents the starting and ending days of meeting i (inclusive).

Return the count of days when the employee is available for work but no meetings are scheduled.

Note: The meetings may overlap.



Example 1:

Input: days = 10, meetings = [[5,7],[1,3],[9,10]]

Output: 2

Explanation:

There is no meeting scheduled on the 4th and 8th days.

Example 2:

Input: days = 5, meetings = [[2,4],[1,3]]

Output: 1

Explanation:

There is no meeting scheduled on the 5th day.

Example 3:

Input: days = 6, meetings = [[1,6]]

Output: 0

Explanation:

Meetings are scheduled for all working days.



Constraints:

1 <= days <= 109
1 <= meetings.length <= 105
meetings[i].length == 2
1 <= meetings[i][0] <= meetings[i][1] <= days
*/
func CountDays(days int, meetings [][]int) int {
	// Merge the overlapping meetings
	mergedMeetings := mergeIntervals(meetings)

	// Count free days
	freeDays, curDay := 0, 1

	for _, meeting := range mergedMeetings {
		if curDay < meeting[0] {
			freeDays += meeting[0] - curDay
		}
		curDay = meeting[1] + 1
	}

	// Check last free days
	if curDay <= days {
		freeDays += days - curDay + 1
	}
	
	return freeDays
}

func mergeIntervals(meetings [][]int) [][]int {
	if len(meetings) == 0 {
		return [][]int{}
	}

	// Sort by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// Iterate and merge
	merged := [][]int{meetings[0]}

	for i := 1; i < len(meetings); i++ {
		curMeeting := meetings[i]
		lastMerged := merged[len(merged)-1]

		if curMeeting[0] <= lastMerged[1] {
			// Overlapped, update the end time of lastMerged
			merged[len(merged)-1][1] = max(lastMerged[1], curMeeting[1])
		} else {
			// No overlapping, add in curMeeting
			merged = append(merged, curMeeting)
		}
	}
	return merged
}
