package main

import "fmt"

//抽象工厂模式：围绕一个超级工厂，创建其他工厂，属于创建模型
//场景：一个工厂里边对应一种具体产品，有时候需要多个工厂对应多种产品
type Shape interface {
	Draw()
}

type Color interface {
	Fill()
}

type Circular struct{}

func (c Circular) Draw() {
	fmt.Println("圆形")
}

type Square struct{}

func (c Square) Draw() {
	fmt.Println("正方形")
}

type Red struct{}

func (r Red) Fill() {
	fmt.Println("红色")
}

type Green struct{}

func (r Green) Fill() {
	fmt.Println("绿色")
}

type SuperFactory interface {
	GetShape(ShapeName string) Shape
	GetColor(ColorName string) Color
}

type ShapeFactory struct{}

func (ShapeFactory) GetShape(shapeName string) Shape {
	switch shapeName {
	case "Circular":
		return &Circular{}
	case "Square":
		return &Square{}
	default:
		return nil
	}
}

func (ShapeFactory) GetColor(shapeName string) Color {
	return nil
}

type ColorFactory struct{}

func (ColorFactory) GetColor(colorName string) Color {
	switch colorName {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

func (ColorFactory) GetShape(colorName string) Shape {
	return nil
}

type SuperFactoryStruct struct {
}

func (SuperFactoryStruct) GetFactory(factoryName string) SuperFactory {
	switch factoryName {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}

func NewSuperFactoryStruct() *SuperFactoryStruct {
	return &SuperFactoryStruct{}
}

func main() {
	nsf := NewSuperFactoryStruct()
	colorFactory := nsf.GetFactory("color")
	shapeFactory := nsf.GetFactory("shape")

	red := colorFactory.GetColor("red")
	square := shapeFactory.GetShape("Square")

	red.Fill()
	square.Draw()
}
