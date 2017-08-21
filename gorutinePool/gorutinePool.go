package gorutinePool

import "fmt"

type GoroutinePool struct {
	Queue  chan func() error
	Number int
	Total  int

	result         chan error
	finishCallback func()
}

// 初始化
func (self *GoroutinePool) Init(number int, total int) {
	self.Queue = make(chan func() error, total)
	self.Number = number
	self.Total = total
	self.result = make(chan error, total)
}

// 开门接客
func (self *GoroutinePool) Start() {
	// 开启Number个goroutine
	for i := 0; i < self.Number; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}
				err := task()
				self.result <- err
			}
		}()
	}
	// 获得每个work的执行结果
	for j := 0; j < self.Total; j++ {
		res, ok := <-self.result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}
	// 所有任务都执行完成，回调函数
	if self.finishCallback != nil {
		self.finishCallback()
	}
}

// 关门送客
func (self *GoroutinePool) Stop() {
	close(self.Queue)
	close(self.result)
}

// 添加任务
func (self *GoroutinePool) AddTask(task func() error) {
	self.Queue <- task
}

// 设置结束回调
func (self *GoroutinePool) SetFinishCallback(callback func()) {
	self.finishCallback = callback
}
