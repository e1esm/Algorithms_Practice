package main

import (
	"sort"
)

type Spread struct {
	start int
	end   int
}

func GetGardenBeds(amount int, Spreads []Spread) []Spread {
	SortSpreadsByStart(&Spreads)
	result := make([]Spread, 0)
	
	for i := 0; i < len(Spreads); i++{
		if  i + 1 < len(Spreads) && Spreads[i].end < Spreads[i + 1].start{
			result = append(result, Spread{Spreads[i].start, Spreads[i].end})
		}else{
			start := Spreads[i].start
			end := Spreads[i].end
			for i + 1 < len(Spreads) && Spreads[i].end >= Spreads[i + 1].start{
				i++
				if Spreads[i].end > end{end = Spreads[i].end}
			}
			if i > len(Spreads){break}
			result = append(result, Spread{start, end})
		}
	}


	return result
}	

func SortSpreadsByStart(Spreads *[]Spread) {
	sort.Slice(*Spreads, func(i, j int) bool {
		return (*Spreads)[i].start < (*Spreads)[j].start
	})
}