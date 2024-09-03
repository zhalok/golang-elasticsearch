package service

import (
	"fmt"
	"os"
	"practice/model"
	"testing"
)


type TestCase[T1 string|int,T2 model.Person|int] struct{
	name string
	input T1
	expected T2
}

func TestGetPersons(t *testing.T){
    tests := []TestCase[int,int]{
     {"test case 1",1, 1},
	 {"test case 2",2, 2},
	 {"test case 3",3, 3},

	}

currentWorkingDirectory,err := os.Getwd()
if err!= nil{
	panic("project root not found")
}
fmt.Println(currentWorkingDirectory)

 for _,testCase:=range tests{
	t.Run(testCase.name,func(t *testing.T) {
		result := GetPersons(int(testCase.input))
		if len(result) != testCase.expected{
		   t.Errorf("expected 2 but got %d",len(result))
		}
   
	 })
 }


	
}

func TestGetPersonById(t *testing.T){
	testCases := []TestCase[string,model.Person]{
		{"test case 1","1", model.Person{Id:"1",Name:"Alice",Age:22}},
		{"test case 2","2", model.Person{Id:"2",Name:"Bob",Age:34}},
		{"test case 3","3", model.Person{Id:"3",Name:"Charlie",Age:22}},
		{"test case 4","4", model.Person{Id:"4",Name:"Diana",Age:30}},
	}
	for _,testCase:=range testCases{
		t.Run(testCase.name,func(t *testing.T) {
			result := GetPersonById(testCase.input)
			if result.Id != testCase.expected.Id || result.Name != testCase.expected.Name || result.Age != testCase.expected.Age{
				t.Errorf("expected %v but got %v",testCase.expected,result)
			}
		})
	}
}