package main
	
import (
	"fmt"
)

type person struct {
    name string
    age  int
}

func main()  {
	var me = person{"Gustavo", 26}
	var friend1 = person{"Eduado", 27}
	var friend2 = person{"Adrian", 27}
	var friend3 = person{"Andres", 38}
	//people_information(me, friend1, friend2, friend3)
	//fmt.Println(me)
	//fmt.Println(born_on(me))
	//fmt.Println(return_text())
	arrays(me, friend1, friend1, friend2, friend3)
}

func arrays(people ...person)  {

	var friends[5] person
	for i, current_person := range people {
		//fmt.Println(i)
		//fmt.Println(current_person)
		friends[i] = current_person
	}
	//fmt.Println(friends)
	fmt.Println(len(people))
	fmt.Println(len(friends))
	fmt.Println(friends[0:2])
}

func people_information(people ...person)  {
	for _, current_person := range people {
		fmt.Println(born_on(current_person))
	}
}
func born_on(me person) (string, int, string, int) {
	born := func () (int) {
		return 2019 - me.age
	}
	return "La edad de " + me.name + " es", me.age, " y nacion en ", born()
}

func return_text() (value1 string, value2 int)  {
	value1 = "hello"
	value2 = 24
	return
}