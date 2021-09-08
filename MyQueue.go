type MyQueue struct {
    len int 
    q []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{0, []int{}}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.len++
    this.q = append(this.q, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    e := this.q[0]
    this.q = this.q[1:]
    this.len--
    return e
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.q[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.len == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
