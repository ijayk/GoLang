package mytest

import "testing"


func TestAdd(t *testing.T){
	ans:=Add(5,15)
	if ans!=20 {
		t.Errorf("Test case failed needs 10 got %d",ans)
	}
}

	func TestSub(t *testing.T){
		ans:=Sub(5,5)
		if ans!=0 {
			t.Errorf("Test case failed needs 10 got %d",ans)
		}
	}
