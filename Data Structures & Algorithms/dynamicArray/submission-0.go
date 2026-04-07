type DynamicArray struct {
    elements []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    da := &DynamicArray{}
    da.elements = make([]int, 0, capacity)

    return da
}

func (da *DynamicArray) Get(i int) int {
    return da.elements[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.elements[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if len(da.elements) == cap(da.elements) {
        da.resize()
    }
    da.elements = append(da.elements, n)
}

func (da *DynamicArray) Popback() int {
    el := da.elements[len(da.elements)-1]
    da.elements = da.elements[:len(da.elements)-1]
    return el
}

func (da *DynamicArray) resize() {
    newEls := make([]int, len(da.elements), cap(da.elements)*2)
    copy(newEls, da.elements)
    da.elements = newEls
}

func (da *DynamicArray) GetSize() int {
    return len(da.elements)
}

func (da *DynamicArray) GetCapacity() int {
    return cap(da.elements)
}