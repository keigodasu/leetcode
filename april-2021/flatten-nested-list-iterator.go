type NestedIterator struct {
    nums []int
    p int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    var res NestedIterator
    var flatten func([]*NestedInteger)    
    
    flatten = func(nestedList []*NestedInteger) {
        for _, nest := range nestedList {
            if nest.IsInteger() {
                res.nums = append(res.nums, nest.GetInteger())
            } else {
                flatten(nest.GetList())
            }
        }
    }
    flatten(nestedList)
    res.p = 0
    return &res
}

func (this *NestedIterator) Next() int {
    this.p++
    return this.nums[this.p - 1]
}

func (this *NestedIterator) HasNext() bool {
    return this.p < len(this.nums)
}
