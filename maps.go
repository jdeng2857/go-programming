package main

import (
	"fmt"
	"sort"
)

var heights map[string]int

type kv struct {
	key   string
	value int
}

type kvPairs []kv

func (p kvPairs) Len() int {
	return len(p)
}

func (p kvPairs) Less(i, j int) bool {
	return p[i].value < p[j].value
}

func (p kvPairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func basicMaps() {
	heights = make(map[string]int)
	heights["Peter"] = 170
	heights["Darth Sidious"] = 170
	heights["Anakin Skywalker"] = 188
	fmt.Println(heights["Peter"])
	fmt.Println(heights["Darth Sidious"])
	fmt.Println(heights["Anakin Skywalker"])

	if v, ok := heights["Jim"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println(len(heights))
	var weights map[string]int
	fmt.Println(len(weights))

	for k, v := range heights {
		fmt.Println(k, v)
	}
	var keys []string
	for k := range heights {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	p := make(kvPairs, len(heights))
	i := 0
	for k, v := range heights {
		p[i] = kv{k, v}
		i++
	}
	fmt.Println(p)
	sort.Sort(p)
	for _, v := range p {
		fmt.Println(v)
	}
}

type dob struct {
	day   int
	month int
	year  int
}

type people struct {
	name  string
	email string
	dob   dob
}

func mapsWithStructs() {
	members := make(map[int]people)

	members[1] = people{
		name:  "Mary Smith",
		email: "marysmith@example.com",
		dob: dob{
			day:   17,
			month: 3,
			year:  1990,
		},
	}
	members[2] = people{
		name:  "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
			day:   9,
			month: 12,
			year:  1988,
		},
	}
	members[3] = people{
		name:  "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
			day:   1,
			month: 12,
			year:  1988,
		},
	}
	members[4] = people{
		name:  "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
			day:   19,
			month: 8,
			year:  2001,
		},
	}
	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	}

	var keys []int
	for k := range members {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var sliceMembers []people
	for _, k := range keys {
		sliceMembers = append(sliceMembers, members[k])
	}

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		if sliceMembers[i].dob.year != sliceMembers[j].dob.year {
			return sliceMembers[i].dob.year < sliceMembers[j].dob.year
		}
		if sliceMembers[i].dob.month != sliceMembers[j].dob.month {
			return sliceMembers[i].dob.month < sliceMembers[j].dob.month
		}
		return sliceMembers[i].dob.day < sliceMembers[j].dob.day
	})

	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}
}

func main() {
	basicMaps()
	mapsWithStructs()
}
