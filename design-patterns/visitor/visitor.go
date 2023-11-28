package visitor

import "fmt"

type IVisitor interface {
	Visit()
}

type WeiboVisitor struct {
}

func (w WeiboVisitor) Visit() {
	fmt.Println("visit weibo")
}

type IqiyiVisitor struct {
}

func (i IqiyiVisitor) Visit() {
	fmt.Println("visit iqiyi")
}

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct {
}

func (e Element) Accept(v IVisitor) {
	v.Visit()
}
