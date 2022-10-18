package main

import (
	"fmt"
	"strconv"

	account "github.com/Lxb921006/Golang-practise/object/account"
	person "github.com/Lxb921006/Golang-practise/object/person"
)

type Student struct {
	Name string
	Age  int
	Id   int
}

type Student02 struct {
	Name  string
	Age   int
	Socre float64
}

type Visitor struct {
	Name   string
	Age    int64
	prices float64
}

type Instance struct {
	Name string
}

type Bank struct {
	Account  int
	Password string
	Balance  float64
}

func (m *Student) Say() string {
	res := fmt.Sprintf("name=%v, age=%d, id=%d", m.Name, m.Age, m.Id)
	return res
}

func (m *Visitor) Answer(b bool) {
	m.prices = 20
	for {
		if b {
			break
		}
		fmt.Printf("请输入姓名:")
		fmt.Scanln(&m.Name)
		if m.Name == "n" {
			break
		}
		fmt.Printf("请输入年龄:")
		fmt.Scanln(&m.Age)
		if (*m).Age >= 18 {
			fmt.Printf("请付款%v元\n", m.prices)
		} else {
			fmt.Println("免费")
		}
		m.Name = ""
	}
}

func (m *Bank) Query(p string) {
	if m.Password == p {
		fmt.Printf("账号:%v, 余额:%f\n", m.Account, m.Balance)
	} else {
		fmt.Println("密码错误")
	}
}

func (m *Bank) Savemoney(a float64, p string) {
	if m.Password == p {
		if a > .0 {
			m.Balance += a
			fmt.Println("存入成功")
			fmt.Printf("账号:%v, 余额:%f\n", m.Account, m.Balance)
		} else {
			fmt.Println("存入金钱必须大于1块钱")
		}

	} else {
		fmt.Println("密码错误")
	}
}

func (m *Bank) Withdraw(a float64, p string) {
	if m.Password == p {
		if a < m.Balance && a > .0 {
			m.Balance -= a
			fmt.Println("取款成功")
			fmt.Printf("账号:%v, 余额:%f\n", m.Account, m.Balance)
		} else {
			fmt.Println("取款金额不能大于余额也不能小于0")
		}

	} else {
		fmt.Println("密码错误")
	}
}

func (s1 *Student02) StuInfo() {
	fmt.Printf("name=%v, age=%d, socre=%f\n", s1.Name, s1.Age, s1.Socre)
}

func (s1 *Student02) SetSocre(soc float64) {
	s1.Socre = soc
}

func main() {
	//面向对象编程
	fmt.Println("------------------面向对象--------------------")
	o1 := Student{"lxb", 29, 1}
	o2 := (&o1).Say()
	fmt.Println("o2=", o2)
	v := Visitor{}
	(&v).Answer(true)
	fmt.Println("------------------创建结构体实例的方式--------------------")
	s := Instance{ //推荐
		Name: "lxb",
	}
	fmt.Println("s=", s)
	//创建指针类型，重要
	var s1 *Instance = &Instance{
		Name: "lqm",
	}
	s2 := &Instance{"lqm"} //这个等同于上面
	fmt.Println("s1=", (*s1).Name)
	fmt.Println("s2=", s2.Name)
	fmt.Println("s2=", *s2)
	fmt.Println("------------------面向对象编程-抽象(将一类事物的共有属性行为提取出来形成一个物理模型的研究)--------------------")
	s3 := Bank{
		Account:  88888888,
		Password: "666666",
		Balance:  100.0,
	}
	out := false
	for {
		var num int
		var money float64
		var pwd string
		if out {
			break
		}
		fmt.Printf("存钱请输入1, 取钱请输入2, 查询账户请输入3:")
		fmt.Scanln(&num)
		switch num {
		case 1:
			fmt.Printf("请输入账户密码:")
			fmt.Scanln(&pwd)
			fmt.Printf("请输入存钱金额:")
			fmt.Scanln(&money)
			s3.Savemoney(money, pwd)
			out = true
		case 2:
			fmt.Printf("请输入账户密码:")
			fmt.Scanln(&pwd)
			fmt.Printf("请输入取款金额:")
			fmt.Scanln(&money)
			s3.Withdraw(money, pwd)
			out = true
		case 3:
			fmt.Printf("请输入账户密码:")
			fmt.Scanln(&pwd)
			s3.Query(pwd)
			out = true
		}
	}
	fmt.Println("------------------面向对象编程-封装--------------------")
	p := person.Newperson("lxb")
	fmt.Println("p=", *p)
	p.Setage(20)
	p.Setsalary(6000)
	fmt.Println("p.Getage=", (*p).Getage())
	fmt.Println("p=", *p)
	ss := 12313
	sss := strconv.FormatInt(int64(ss), 10)
	fmt.Println("sss=", len(sss))
	fmt.Println("------------------面向对象编程-封装练习--------------------")
	acc, err := account.NewAccount(1515512, "lxb", "666666", 100.0)
	if acc == nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("acc=", *acc)
		fmt.Println("acc.GetAccountId()=", acc.GetAccountId())
		acc.SetAccountId(68168)
		fmt.Println("acc.GetAccountId()=", acc.GetAccountId())
	}
	fmt.Println("------------------面向对象编程-继承--------------------")
	stu02 := &Student02{
		Name:  "lxb",
		Age:   30,
		Socre: 61.0,
	}
	(*stu02).StuInfo()
	stu02.SetSocre(70.0)
	fmt.Println("stu02=", *stu02)

}
