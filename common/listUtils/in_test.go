package listUtils

import "testing"

func TestIn(t *testing.T) {
	name_list := []string{"pm", "kingname", "青南"}
	target1 := "kingname"
	target2 := "产品经理"
	result := In(target1, name_list)
	t.Log("kingname 是否在 name_list 中：", result)
	result = In(target2, name_list)
	t.Log("产品经理是否在 name_list 中：", result)
}
