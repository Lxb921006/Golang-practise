package main

import "fmt"

type User struct {
	Name string
}

func main() {
	m := make(map[string]User)
	u := User{}
	vd, ok := m["user"]
	if !ok {
		u.Name = "lxb"
		m["user"] = u
		vd = m["user"]
	}

	fmt.Println("vd = ", vd.Name)

}
