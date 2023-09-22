package main


func findDuplicate(nums []int) int {
	numsAppearance := make([]bool, len(nums))
	for _, v := range nums{
		if numsAppearance[v]{
			return v
		}
		numsAppearance[v] = true
	}
	return -1
}



func main(){

}