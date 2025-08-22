package test

import (
	"fmt"
	"net/http"
)

func TestMain() {
	client := &http.Client{}
	response, err := TestAddTask(client)
	if err != nil {
		fmt.Println("Error add task")
	} else {
		fmt.Println(response)
	}
	response, err = TestGetAllTasks(client)
	if err != nil {
		fmt.Println("Error add task")
	} else {
		fmt.Println(response)
	}
	// response, err = TestGetTask(client)
	// if err != nil {
	// 	fmt.Println("Error add task")
	// } else {
	// 	fmt.Println(response)
	// }
	// response, err = TestUpdateTask(client)
	// if err != nil {
	// 	fmt.Println("Error add task")
	// } else {
	// 	fmt.Println(response)
	// }
	// response, err = TestDeleteTask(client)
	// if err != nil {
	// 	fmt.Println("Error add task")
	// } else {
	// 	fmt.Println(response)
	// }

}
