package main

// duration is a named type that represents a duration
// of time in Nanosecond.
type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

// setHours sets the specified number of hours.
func (d *duration) setHours(h float64) {
	*d = duration(h * float64(hour))
}

// hours returns the duration as a floating point number of hours.
func (d duration) hours() float64 {
	hours := d / hour
	nsec := d % hour
	return float64(hours) + (float64(nsec) * (1e-9 / 60 / 60))
}
