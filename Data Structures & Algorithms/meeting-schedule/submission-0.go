import (
    "cmp"
    "slices"
)

/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	answer := true

	// sort by start time (ascending)
    slices.SortFunc(intervals, func(a, b Interval) int {
        return cmp.Compare(a.start, b.start)
    })

	for idx, iVal := range(intervals) {
		if idx > 0 {
			if iVal.start < intervals[idx-1].end {
				answer = false
			}
		}
	}

	return answer
}