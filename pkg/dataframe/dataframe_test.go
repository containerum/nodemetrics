package dataframe

import "testing"

func TestMakeDataframe(test *testing.T) {
	var dataframe = MakeDataframe("", 0, nil)
	test.Log(dataframe)
}
