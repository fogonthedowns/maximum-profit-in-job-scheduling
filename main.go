package main

import (
	"fmt"
	"sort"
)

type job struct {
	start, end int
	profit     int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := []job{}
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, job{startTime[i], endTime[i], profit[i]})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].start < jobs[j].start
	})

	cache := map[int]*int{}
	return maxProfit(jobs, 0, cache)
}

func maxProfit(jobs []job, start int, cache map[int]*int) int {
	noProfit := -1

	fmt.Printf("len %v start: %d, cache: %v\n", len(cache), start, cache)

	// has this been calculated already?
	if cache[start] != nil {
		return *cache[start]
	}

	// handle end of job list
	if start == len(jobs) {
		fmt.Printf("start %v len(jobs) %v\n", start, len(jobs))
		return noProfit
	}

	// initialize next at -1
	next := noProfit

	// take the next job if you can.
	// take it if, it starts after the current job ends
	for i := start + 1; i < len(jobs); i++ {
		if jobs[start].end <= jobs[i].start {
			next = maxProfit(jobs, i, cache)
			break
		}
	}

	// take the job, and its profit
	with := jobs[start].profit

	// next job profit, if its not -1
	// we add up the next job
	if next != noProfit {
		with += next
	}

	// NOTE, this step does not depend on order, we do not consider next and skip.
	// do not take job, move on to the next job and consider its profit
	without := maxProfit(jobs, start+1, cache)

	// update cache with max profit for start
	maxLocal := max(with, without)
	cache[start] = &maxLocal
	return maxLocal
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
