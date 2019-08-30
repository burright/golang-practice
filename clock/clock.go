package clock

import (
	"strconv"
)

// Clock structure
// Requires an hour and minute
type Clock struct {
	hour   int
	minute int
}

// New Clock with hour Hour and minute Minute
func New(Hour, Minute int) Clock {
	c := Clock{Hour, Minute}
	c = checkTime(c)
	return c
}

// Converts a Clock into a string with format 00:00
func (c Clock) String() string {
	time := ""

	// add a 0 before the hour if it is less than 10
	if c.hour < 10 {
		time += "0"
	}
	time += strconv.Itoa(c.hour) + ":"

	// add a 0 to the minute if it is less than 10
	if c.minute < 10 {
		time += "0"
	}
	time += strconv.Itoa(c.minute)

	return time
}

// Add minutes to an existing Clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes

	c = checkTime(c)
	return c
}

// Subtract minutes from an existing clock
func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes

	c = checkTime(c)
	return c
}

// Check the hour and minute of a provided Clock
func checkTime(c Clock) Clock {
	// If the minutes are greater than one hour
	if c.minute > 59 {
		// If the minutes are equal to one hour add one to hour
		// and clear minutes
		if c.minute == 60 {
			c.hour++
			c.minute = 0
		} else { // Else add an hour and put the remainder into minute
			hours := c.minute / 60
			c.hour += hours
			c.minute = c.minute % 60
		}
	}

	// If the minute is negative
	if c.minute < 0 {
		// Keep taking off minutes until you reach 0
		for c.minute < 0 {
			if c.minute < -60 {
				c.minute += 60
				c.hour--
			} else {
				c.minute = 60 + c.minute
				c.hour--
			}
		}
	}

	// If the hour is more than 23 set the remainder to hour
	if c.hour > 23 {
		if c.hour == 24 {
			c.hour = 0
		} else {
			c.hour = c.hour % 24
		}

	}

	// If the hour is negative keep taking off hours until you reach 0
	if c.hour < 0 {
		for c.hour < 0 {
			c.hour += 24
		}
	}

	return c
}
