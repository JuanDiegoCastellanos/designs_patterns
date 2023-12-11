package models

import "fmt"

type Button interface{
	render()
	onClick()
}
type WindowsButton struct{}
type HTMLButton struct{}

func (WindowsButton) render(){
	fmt.Println("Rendering WindowsButton")
}
func (HTMLButton) render(){
	fmt.Println("Rendering HTMLButton")
}

func (WindowsButton) onClick(){
	fmt.Println("WindowsButton has click")
}
func (HTMLButton) onClick(){
	fmt.Println("HTMLButton has click")
}
