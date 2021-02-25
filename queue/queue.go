//Package queue 队列
package queue

import "fmt"

//Queue 队列
type Queue struct {
	data []interface{}
	size int
}

//New 初始化
func New() *Queue {
	return new(Queue)
}

//IsEmpty 是否为空
func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

//Size 队列长度
func (queue *Queue) Size() int {
	return queue.size
}

//InQueue 入队
func (queue *Queue) InQueue(data interface{}) {
	queue.data = append(queue.data, data)
	queue.size++
}

//OutQueue 出队，返回队首元素
func (queue *Queue) OutQueue() (element interface{}) {
	if queue.IsEmpty() {
		return nil
	}
	element = queue.data[0]
	queue.data = queue.data[1:]
	queue.size--
	return
}

//Front 获取队首元素
func (queue *Queue) Front() (element interface{}) {
	if queue.IsEmpty() {
		return nil
	}
	element = queue.data[0]
	return
}

//Rear 获取队尾元素
func (queue *Queue) Rear() (element interface{}) {
	if queue.IsEmpty() {
		return nil
	}
	element = queue.data[queue.size-1]
	return
}

//TestQueue 测试队列
func TestQueue() {
	queue := New()
	fmt.Println(queue)
	queue.InQueue(3)
	fmt.Println(queue)
	queue.InQueue(4)
	fmt.Println(queue)
	fmt.Println(queue.Front())
	fmt.Println(queue.Rear())
	fmt.Println(queue.Size())
	b := queue.OutQueue()
	fmt.Println(b)
	fmt.Println(queue)
}
