package disjointset

type DisjoinSet struct {
	nums []int
}

func (d *DisjoinSet) find(i int) int {
	if d.nums[i] == i {
		return i
	} else {
		d.nums[i] = d.find(d.nums[i])
	}

	return d.nums[i]
}

func (d *DisjoinSet) union(i int, j int) {
	x := d.find(i)
	y := d.find(j)

	d.nums[x] = y
}
