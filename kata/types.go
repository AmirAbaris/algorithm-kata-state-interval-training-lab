package kata

type Interval struct {
	Start int
	End   int
}

type Event struct {
	Start int
	End   int
	Delta int
}

type StateSegment struct {
	Interval Interval
	State    int
}
