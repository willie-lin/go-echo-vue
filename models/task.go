package models

import (

	"fmt"
	"github.com/jinzhu/gorm"
	"go-echo-vue/database"
)

type Task struct {
	gorm.Model
	Name string		`json:"name"`
	Description	 string 		`json:"description"`

}

func (t Task) Error() string {
	panic("implement me")
}

type TaskCollection struct {
	gorm.Model
	Tasks	[]Task  	`json:"items"`
}


func GetTasks()  []Task {
	db := database.Connect()
	fmt.Println(111)
	//tasks := []Task{}
	var tasks []Task
	result := db.Find(&tasks)
	//println(tasks)
	defer db.Close()
	if result.RecordNotFound(){
		fmt.Println("数据库无记录")
	}

	return tasks

}



func PutTask(task Task) *Task  {
	db := database.Connect()
	//var ta Task

	//
	//task.Name = name
	//fmt.Println(name)
	//task.Description = description
	//fmt.Println(description)
	//fmt.Println(task)
	t := task
	//fmt.Println(t)
	//js, _ := json.Marshal(t)
	//fmt.Println(fmt.Sprintf("%s", js))
	//_ = json.Unmarshal([]byte(js), &ta)
	//fmt.Println(ta)

	result := db.FirstOrCreate(&t)

	//result := db.Create(&t)

	//id := result.Exec("id")
	//fmt.Println(id)
	//return result.RowsAffected()

	defer db.Close()
	if result.RecordNotFound() {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")

	}
	return nil
}

func FindById(id int) *Task {

	db := database.Connect()
	task := new(Task)
	taskId := id
	result := db.Where("id = ?", taskId).First(&task)
	defer db.Close()
	if result.RecordNotFound() {
		fmt.Println("数据库无记录")
	}
	return nil
}


func DeleteTask(id int)  *Task{
	db := database.Connect()
	task := new(Task)
	taskId := id
	result := db.Where("id = ?", taskId).Delete(&task)
	defer db.Close()
	if result.RecordNotFound() {
		fmt.Println("数据库无记录")
	}
	return nil
}