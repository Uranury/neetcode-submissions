type DynamicArray struct {
    arr []any
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{arr: make([]any, 0, capacity)}
}

func (da *DynamicArray) Get(i int) any {
    return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if len(da.arr) == cap(da.arr) {
        newCap := 1
        if cap(da.arr) > 0 {
            newCap = cap(da.arr)*2
        }
        newArr := make([]any, len(da.arr), newCap)
        copy(newArr, da.arr)
        da.arr = newArr
    }
    da.arr = da.arr[:len(da.arr)+1]
    da.arr[len(da.arr)-1] = n
}

func (da *DynamicArray) Popback() any {
    last := da.arr[len(da.arr)-1]
    da.arr = da.arr[:len(da.arr)-1]
    return last
}

func (da *DynamicArray) resize() {
    newArr := make([]any, len(da.arr), 2*cap(da.arr))
    copy(newArr, da.arr)
    da.arr = newArr
}

func (da *DynamicArray) GetSize() int {
    return len(da.arr)
}

func (da *DynamicArray) GetCapacity() int {
    return cap(da.arr)
}
