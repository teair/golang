package gotest

import "testing"

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("函数Test_Division_1测试不通过!")
	} else {
		t.Log("Test_Division_1测试通过!")
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		t.Error("division did not work as expected")
	} else {
		t.Log("one test passed.", e)
	}
}
