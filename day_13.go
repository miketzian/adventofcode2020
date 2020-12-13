package adventofcode2020

import (
	"fmt"
	"strconv"
	"strings"
)

// bus has an ID number that indicates how often the bus leaves

// at t=0, all busses departed

func day13(input []string) (int, error) {

	departTime, err := strconv.ParseInt(input[0], 10, 64)
	if err != nil {
		return 0, err
	}
	busRoutes := make([]int, 0)
	for _, v := range strings.Split(input[1], ",") {
		if v != "x" {
			route, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
			busRoutes = append(busRoutes, route)
		}
	}
	earliestTime := int64(-1)
	earliestRoute := -1

	for _, route := range busRoutes {
		// busTimes[ix]
		r := int64(route)
		min := (departTime / r) * r
		if min < departTime {
			min += r
		}
		if earliestRoute == -1 || min < earliestTime {
			earliestTime = min
			earliestRoute = route
		}
	}
	wait := int(earliestTime - departTime)
	return earliestRoute * wait, nil
}

func day13Part2(input string) (uint64, error) {

	busRoutes := make([]uint64, 0)
	busRouteDelay := make([]uint64, 0)

	for mins, route := range strings.Split(input, ",") {
		if route != "x" {
			route, err := strconv.Atoi(route)
			if err != nil {
				return 0, err
			}
			busRoutes = append(busRoutes, uint64(route))
			busRouteDelay = append(busRouteDelay, uint64(mins))
		}
	}
	var timeStamp uint64
	timeStamp = busRoutes[0]
	multiple := busRoutes[0]

	loops := 0

	for ix, delay := range busRouteDelay {
		if ix == 0 {
			continue
		}
		route := busRoutes[ix]
		for (timeStamp+delay)%route != 0 {
			timeStamp += multiple
			loops++
		}
		// add the route into the multiple
		// chinese remainder theorum
		multiple *= route
	}
	fmt.Printf("Complete after %d loops\n", loops)
	return timeStamp, nil
}
