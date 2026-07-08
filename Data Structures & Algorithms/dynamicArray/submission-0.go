type DynamicArray struct {
    arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    dyArr := DynamicArray{
        arr: make([]int, 0, capacity),
    }
    return &dyArr
}

func (da *DynamicArray) Get(i int) int {
    return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.GetCapacity() == da.GetSize() {
        da.resize()
    }
    
    da.arr = append(da.arr, n)
}

func (da *DynamicArray) Popback() int {
    last := da.arr[da.GetSize()-1]
    da.arr = da.arr[:len(da.arr)-1]
    return last
}

func (da *DynamicArray) resize() {
    newCap := da.GetCapacity()*2
    if newCap == 0 {
        newCap = 1
    }

    newArr := make([]int, da.GetSize(), newCap)
    copy(newArr, da.arr)
    da.arr = newArr
}

func (da *DynamicArray) GetSize() int {
    return len(da.arr)
}

func (da *DynamicArray) GetCapacity() int {
    return cap(da.arr)
}
