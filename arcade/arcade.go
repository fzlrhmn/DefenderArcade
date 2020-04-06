package arcade

import (
	"math"
	"sort"
)

// Schedule is a struct for schedule
type Schedule struct {
	Schedules []Interval
}

// Interval is a vector for store coordinates in a vector of pair mapped with character x and y
type Interval struct {
	time       float64
	coordinate string
}

// NewSchedule is a function for initiating new Interval instance
func NewSchedule(schedule [][]float64) *Schedule {
	// initiate empty data for holds interval schedule
	var data []Interval

	// store the x and y coordinates in data vector
	for _, v := range schedule {
		// create struct of coordinates x
		x := Interval{
			time:       v[0],
			coordinate: "x",
		}

		// create struct of coordinates y
		y := Interval{
			time:       v[1],
			coordinate: "y",
		}

		// pushing the coordinate to data
		data = append(data, x, y)
	}

	return &Schedule{
		Schedules: data,
	}
}

// CountMinimum is a function that will get the minu
func (a *Schedule) CountMinimum() float64 {

	// initiate new variable for total, count and data
	var total float64
	var count float64

	// set total and count to 0
	total = 0
	count = 0

	// sort the data
	sort.Slice(a.Schedules, func(i, j int) bool {
		if a.Schedules[i].time != a.Schedules[j].time {
			return a.Schedules[i].time < a.Schedules[j].time
		}
		return a.Schedules[i].coordinate < a.Schedules[j].coordinate
	})

	// Traverse the a.Schedules to count number of overlaps
	for _, v := range a.Schedules {

		// if x occur, it means a new range detected (user start play)
		// then increase the count
		if v.coordinate == "x" {
			count++
		}

		// if y occur, it means a range is ended (user stop play)
		// then decrease the count
		if v.coordinate == "y" {
			count--
		}

		// updating the value of total for every traversal
		// always take the maximum
		total = math.Max(total, count)
	}

	// return the total
	return total
}
