package main

import "fmt"

type section interface {
	execute(*task)
	setNext(section)
}

type choose struct {
	next section
}

func (c *choose) execute(t *task) {
	if t.Selected {
		fmt.Println("already choose")
		c.next.execute(t)
		return
	}

	fmt.Println("choose ok")
	t.Selected = true
	c.next.execute(t) //调用下一个结构体的execute如pay, 对task任务继续逻辑判断运行
}

func (c *choose) setNext(next section) {
	c.next = next
}

type pay struct {
	next section
}

func (p *pay) execute(t *task) {
	if t.Paid {
		fmt.Println("already paid")
		p.next.execute(t)
		return
	}

	fmt.Println("paid ok")
	t.Paid = true
	p.next.execute(t)
}

func (p *pay) setNext(next section) {
	p.next = next
}

type take struct {
	next section
}

func (ta *take) execute(t *task) {
	if t.HasLeft {
		fmt.Println("already take")
		ta.next.execute(t)
		return
	}

	fmt.Println("take ok")
	t.HasLeft = true
	ta.next.execute(t)
}

func (ta *take) setNext(next section) {
	ta.next = next
}

type finised struct {
	next section
}

func (f *finised) execute(t *task) {
	if t.Finised {
		fmt.Println("already finished")
		f.next.execute(t)
		return
	}

	fmt.Println("finished ok", t.Finised)
}

func (f *finised) setNext(next section) {
	f.next = next
}

type task struct {
	name     string
	HasLeft  bool
	Paid     bool
	Selected bool
	Finised  bool
}

// 责任链替代if地狱-购物为例：挑选商品-付款-带走-完成购物
func main() {

	f := &finised{}

	t := &take{}
	t.setNext(f)

	p := &pay{}
	p.setNext(t)

	c := &choose{}
	c.setNext(p)

	task := &task{name: "shopping"}
	c.execute(task)
}
