// import "fmt"
func circularArrayLoop(nums []int) bool {
	length := len(nums)
	tried := [5001][2]bool{}
	for i := 0; i < length; i++ {
		rcs := make(map[int]int)
		idir := 0
		forward := (nums[i] > 0)
		if forward {
			idir = 1
		}
		pre := -50000
		for j := i; true; {
			// j = (j % length + length) % length
			// fmt.Println("j:", j, "foward:", forward, (nums[j]>0))
			// fmt.Println("tried", tried)
			// fmt.Println("j:", j)
			if nums[j] == 0 || forward != (nums[j] > 0) || tried[j][idir] == true {
				tried[j][idir] = true
				for k, _ := range rcs {
					tried[k][idir] = true
				}
				break
			}
			if v, ok := rcs[j]; ok {
				if len(rcs) > 1 && v != j && j != pre {
					// fmt.Println("j:", j, "rcs:", rcs)
					return true
				}
				break
			}
			if pre != -50000 {
				rcs[j] = pre
			}
			pre = j
			j = (((j + nums[j]) % length) + length) % length
		}
	}
	return false
}
