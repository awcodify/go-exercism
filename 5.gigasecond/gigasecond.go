package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

const GIGASECOND = time.Duration(1e9) * time.Second

func AddGigasecond(today time.Time) time.Time {
	return today.Add(GIGASECOND)
}
