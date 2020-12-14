package main

type bus struct {
	id                 int
	nearestArrival     int
	nearestArrivalTime int
	waitTime           int
}

func NewBus(id, earliestArrival int) bus {
	b := bus{id: id}

	b.nearestArrival = (earliestArrival / id) + 1
	b.nearestArrivalTime = id * b.nearestArrival
	b.waitTime = b.nearestArrivalTime - earliestArrival

	return b
}

type buses []bus

func (s buses) Len() int { return len(s) }
func (s buses) Less(i, j int) bool {
	return s[i].waitTime < s[j].waitTime
}
func (s buses) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}