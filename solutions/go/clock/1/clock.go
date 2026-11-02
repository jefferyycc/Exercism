package clock

import "fmt"

type Clock struct {
	minutes int
}

const minutesPerDay = 24 * 60

func normalize(minutes int) int {
	minutes %= minutesPerDay
	if minutes < 0 {
		minutes += minutesPerDay
	}
	return minutes
}

func New(h, m int) Clock {
	return Clock{minutes: normalize(h*60 + m)}
}

func (c Clock) Add(m int) Clock {
	return Clock{minutes: normalize(c.minutes + m)}
}

func (c Clock) Subtract(m int) Clock {
	return Clock{minutes: normalize(c.minutes - m)}
}

func (c Clock) String() string {
	h := c.minutes / 60
	m := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
