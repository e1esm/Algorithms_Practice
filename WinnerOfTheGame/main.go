package main

import "fmt"

func winnerOfGame(colors string) bool {
	aliceCount := 0
	bobCount := 0

	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == 'A'{
			if colors[i] == colors[i-1] && colors[i] == colors[i+1] {
				aliceCount++
			}
		}else{
			if colors[i] == colors[i-1] && colors[i] == colors[i+1] {
				bobCount++
			}
		}
	}
	return aliceCount > bobCount && aliceCount != 0
}


func main(){
	fmt.Println(winnerOfGame("AAAABBBB"))
}