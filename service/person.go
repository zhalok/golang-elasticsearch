package service

import (
	"practice/helper"
	"practice/model"
)

func GetPersons(count int) []model.Person {
	



    var people []model.Person
	err:=helper.ReadData("persons",&people)
	
	if err != nil{
		panic(err.Error())
	}
	

	if count>len(people){
		count = len(people)
	}

	return people[0:count]


	
}

func GetPersonById(id string) model.Person{
  var people []model.Person
  err := helper.ReadData("persons",&people)
  if err != nil{
	panic(err.Error())
  }
  var person model.Person

  for _,p:=range people{
	if p.Id == id{
		person = p
	}
  }

  return person

}