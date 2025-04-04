package module1

import (
	"fmt"
	"slices"
	"sort"
)

func countDays(days int, meetings [][2]int) int {
	dates := []int{}
	for _, value := range meetings {
		dates = append(dates, value[0], value[1])
	}
	count := 0
	sort.Ints(dates)
	fmt.Println("dates", dates)
	for j := 1; j <= days; j++ {
		if !slices.Contains(dates, j) {
			count++
		}
	}

	return count
}

func countDays_v2(days int, meetings [][2]int) int {
	occupied := make([]int, days+2) // Use a difference array (size days+2 to avoid out-of-bounds)

	// Mark start and end+1 of occupied ranges
	for _, meeting := range meetings {
		occupied[meeting[0]]++   // Mark start of a meeting
		occupied[meeting[1]+1]-- // Mark end of a meeting
	}
	fmt.Println("", occupied)
	count := 0
	activeMeetings := 0

	// Compute prefix sum and count unoccupied days
	for i := 1; i <= days; i++ {
		activeMeetings += occupied[i]
		fmt.Print(" - ", i, activeMeetings)
		if activeMeetings == 0 {
			count++ // Free day
		}
	}

	return count
}
func countDays_v3(days int, meetings [][2]int) int {
	events := make(map[int]int)

	// Mark start (+1) and end (+1 to indicate the end of the meeting range)
	for _, meeting := range meetings {
		events[meeting[0]]++   // Meeting starts
		events[meeting[1]+1]-- // Meeting ends
	}
	fmt.Println("", events)
	// Sort event keys (days where changes happen)
	uniqueDays := make([]int, 0, len(events))
	for day := range events {
		if day <= days {
			uniqueDays = append(uniqueDays, day)
		}
	}
	sort.Ints(uniqueDays) // Process days in order

	freeDays := 0
	activeMeetings := 0
	prevDay := 1
	fmt.Println("unique", uniqueDays)
	for _, day := range uniqueDays {
		// Count free days before the next event day
		if activeMeetings == 0 {
			freeDays += day - prevDay
		}
		fmt.Print(" +", freeDays, activeMeetings)
		// Update the number of active meetings
		activeMeetings += events[day]
		prevDay = day
	}

	// Count remaining free days after the last meeting event
	if prevDay <= days && activeMeetings == 0 {
		freeDays += days - prevDay + 1
	}

	return freeDays
}

func countDays_v4(days int, meetings [][2]int) int {
	if len(meetings) == 0 { // Handle case when no meetings exist
		return days
	}

	// Sort meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	fmt.Println("me", meetings)
	merged := [][2]int{}
	start, end := meetings[0][0], meetings[0][1]

	for _, meeting := range meetings[1:] {
		if meeting[0] <= end { // Overlapping meeting
			if meeting[1] > end {
				end = meeting[1]
			}
		} else { // No overlap, push previous and reset
			merged = append(merged, [2]int{start, end})
			start, end = meeting[0], meeting[1]
		}
		fmt.Println("merged", merged, start, end)
	}
	merged = append(merged, [2]int{start, end}) // Add last interval

	// Calculate occupied days
	occupied := 0
	for _, m := range merged {
		occupied += (m[1] - m[0] + 1)
		fmt.Println("occupied", occupied)
	}

	// Total free days
	return days - occupied
}
