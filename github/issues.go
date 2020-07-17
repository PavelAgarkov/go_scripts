package github

import (
	"log"
	"os"
	"time"
)

func GetCategories() map[string][]*Issue {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	sort := sortByDate(result.Items)
	return sort

}

func sortByDate(items []*Issue) map[string][]*Issue {
	now := time.Now()

	timeLessMonth := now.AddDate(0, 1, 0)
	timeLessYear := now.AddDate(-1, 0, 0)

	var lessMonth []*Issue
	var lessYear []*Issue
	var moreYear []*Issue

	list := map[string][]*Issue{"lessMonth": {}, "lessYear": {}, "moreYear": {}}

	for _, v := range items {
		if v.CreatedAt.After(timeLessMonth) && v.CreatedAt.Before(now) {
			insert(&lessMonth, v)
		}
		if v.CreatedAt.Before(timeLessMonth) && v.CreatedAt.After(timeLessYear) {
			insert(&lessYear, v)
		}
		if v.CreatedAt.Before(timeLessYear) {
			insert(&moreYear, v)
		}
	}

	list["lessMonth"] = lessMonth
	list["lessYear"] = lessYear
	list["moreYear"] = moreYear

	return list
}

func insert(slice *[]*Issue, value *Issue) {
	*slice = append(*slice, value)
}
