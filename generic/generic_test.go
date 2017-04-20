package generic

import "testing"

type person struct {
	ID   string
	Name string
}

//IsEquel 继承对别借口compare
func (this person) IsEquel(other interface{}) bool {
	if p, ok := other.(person); ok {
		return this.ID == p.ID
	}
	return false
}

func TestGeneric(t *testing.T) {
	list := List{}
	p1 := person{"1", "zhangsan"}
	p2 := person{"2", "lisi"}
	p3 := person{"3", "wangwu"}
	list.Add(p1)
	list.Add(p2)
	list.Add(p3)
	t.Log(list)
	//out
	//{[{1 zhangsan} {2 lisi} {3 wangwu}]}

	p4 := person{"1", "zhangsan11"}
	list.Remove(p4)
	t.Log(list)
	//out
	//{[{2 lisi} {3 wangwu}]}

	list.RemoveAll(func(a interface{}) bool {
		if p, ok := a.(person); ok {
			return p.ID == p2.ID
		}
		return false
	})
	t.Log(list)
	//out
	//{[{3 wangwu}]}
}
