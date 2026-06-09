/*
===============================================================
Exercise 1: Insert Interval
===============================================================
Question:
Insert newInterval into sorted non-overlapping intervals.

Example:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Key Idea:
Add non-overlapping before, merge overlapping, add rest.

===============================================================
Exercise 2: Merge Intervals
===============================================================
Question:
Merge all overlapping intervals.

Example:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]

Key Idea:
Sort by start time, merge if current overlaps with last merged.

===============================================================
Exercise 3: Non-overlapping Intervals
===============================================================
Question:
Find minimum number of intervals to remove to make rest non-overlapping.

Example:
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1

Key Idea:
Sort by end time, greedily keep intervals that end earliest.

===============================================================
Exercise 4: Meeting Rooms
===============================================================
Question:
Determine if person can attend all meetings.

Example:
Input: intervals = [[0,30],[5,10],[15,20]]
Output: false

Key Idea:
Sort by start time, check if any overlap.

===============================================================
Exercise 5: Meeting Rooms II
===============================================================
Question:
Find minimum number of conference rooms required.

Example:
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2

Key Idea:
Use min heap or sort start/end times separately.

===============================================================
*/
package main

import "fmt"

// ===============================================================
// Exercise 1: Insert Interval - Solution
// ===============================================================
func insert(intervals [][]int, newInterval []int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 2: Merge Intervals - Solution
// ===============================================================
func merge(intervals [][]int) [][]int {
	// TODO: Implement
	return nil
}

// Time Complexity: O(n log n)
// Space Complexity: O(n)

// ===============================================================
// Exercise 3: Non-overlapping Intervals - Solution
// ===============================================================
func eraseOverlapIntervals(intervals [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n log n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 4: Meeting Rooms - Solution
// ===============================================================
func canAttendMeetings(intervals [][]int) bool {
	// TODO: Implement
	return false
}

// Time Complexity: O(n log n)
// Space Complexity: O(1)

// ===============================================================
// Exercise 5: Meeting Rooms II - Solution
// ===============================================================
func minMeetingRooms(intervals [][]int) int {
	// TODO: Implement
	return 0
}

// Time Complexity: O(n log n)
// Space Complexity: O(n)

// ===============================================================
// Main - Test all solutions
// ===============================================================
func main() {
	fmt.Println("===== Intervals Practice =====")
	fmt.Println("Complete the TODO sections above")
}
