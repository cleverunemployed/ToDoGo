package test

import (
	"bytes"
	"encoding/json"
	"github/cleverunemployed/ToDoGo/internal/models"
	"io"
	"net/http"
)

// fmt.Printf("%s\n", bodyText)
func TestAddTask(client *http.Client) (response any, err error) {
	data := []byte(`{"title":"title", "completed":"false"}`)
	r := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "127.0.0.1:8080/api/v1/add_task", r)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return bodyText, nil
}

func TestGetTask(client *http.Client, id string) (response any, err error) {
	req, err := http.NewRequest("GET", "127.0.0.1:8080/api/v1/get_task/"+id, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return bodyText, nil
}

func TestGetAllTasks(client *http.Client) (response any, err error) {
	req, err := http.NewRequest("GET", "127.0.0.1:8080/api/v1/get_all_tasks/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return bodyText, nil
}

func TestUpdateTask(client *http.Client, taskData models.Task) (response any, err error) {
	updateData := map[string]interface{}{
		"id":        taskData.ID,
		"completed": taskData.Comleted,
	}

	jsonData, err := json.Marshal(updateData)
	if err != nil {
		panic(err)
	}

	// Создание PATCH запроса с телом
	req, err := http.NewRequest(
		"PATCH",
		"127.0.0.1:8080/api/v1/update_task",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return bodyText, nil
}

func TestDeleteTask(client *http.Client, id string) (response any, err error) {
	deleteData := map[string]interface{}{
		"id": id,
	}

	jsonData, err := json.Marshal(deleteData)
	if err != nil {
		panic(err)
	}

	// Создание PATCH запроса с телом
	req, err := http.NewRequest(
		"DELETE",
		"127.0.0.1:8080/api/v1/delete_task",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return bodyText, nil
}
