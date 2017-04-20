package generic

import (
	"fmt"
)

//List 用golang构建一个list.
type List struct {
	//golang的interface相当于object类型
	list []interface{}
}

//Add 列表添加元素方法.
func (l *List) Add(e interface{}) {
	l.list = append(l.list, e)
}

//Remove 列表删除元素方法，如果实现compare接口，那么就调用该接口的IsEquel方来来判断
func (l *List) Remove(e interface{}) {
	for k, v := range l.list {
		if isEquel(v, e) {
			l.list = append(l.list[:k], l.list[k+1:]...)
			return
		}
	}
}

//RemoveAll 删除所有匹配条件的元素，入参时函数，交给使用者自己定义
func (l *List) RemoveAll(f func(interface{}) bool) {
	for k, v := range l.list {
		if ok := f(v); ok {
			//remove v here
			l.list = append(l.list[:k], l.list[k+1:]...)
		}
	}
}

//FindFirst 查找泛型匹配的第一个元素
func (l *List) FindFirst(e interface{}) (interface{}, int) {
	for k, v := range l.list {
		if isEquel(k, e) {
			return v, k
		}
	}

	return nil, -1
}

//Compare 定义一个接口由泛型元素继承，用于比较2个元素是否相等.
type Compare interface {
	IsEquel(other interface{}) bool
}

//isEquel 函数独立出来用于比较2个元素
func isEquel(a, b interface{}) bool {
	if c, ok := a.(Compare); ok {
		fmt.Println("realization compare interface")
		return c.IsEquel(b)
	}
	fmt.Println("not realization compare interface")
	return a == b
}
