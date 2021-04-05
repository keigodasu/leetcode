type MyCircularQueue struct {
    capacity int
    index int
    head int
    queue []int
}


func Constructor(k int) MyCircularQueue {
    list := make([]int, k + 1)
    return MyCircularQueue{
        capacity: k + 1,
        index: 0,
        head: 0,
        queue: list,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }

    this.queue[this.index] = value
    this.index = (this.index+1) % this.capacity
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    
    this.head = (this.head+1) % this.capacity
    
    return true 
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    
    return this.queue[this.head]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    
    return this.queue[(this.index + this.capacity - 1)%this.capacity]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.index % this.capacity == this.head % this.capacity
}


func (this *MyCircularQueue) IsFull() bool {
    return (this.index + 1)%this.capacity == this.head%this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
