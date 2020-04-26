package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"go-echo-vue/models"
	"net/http"
	"strconv"
)

type  H  map[string]interface{}



func GetTasks()  echo.HandlerFunc{
	return func(c echo.Context) error {
		//tasks := models.Task{}
		//db.Find(&tasks)
		return c.JSON(http.StatusOK, models.GetTasks())
	}
}



func PutTasks()  echo.HandlerFunc{
	return func(c echo.Context) error {
		//var task models.Task
		var task models.Task

		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		t := models.Task{
			Name: m["name"].(string),
			Description: m["description"].(string),
		}
		task = t
		//js, _ := json.Marshal(task)
		//fmt.Println(fmt.Sprintf("%s", js))

		err := models.PutTask(task)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": task.Name,
			})
			} else {
				return err
			}
		}



		//return c.JSON(http.StatusCreated, models.PutTask(task))
	}



func DeleteTasks()  echo.HandlerFunc{
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		err := models.DeleteTask(id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
		//return c.JSON(http.StatusOK, models.DeleteTask(id))
	}
}

func FindTaskById()  echo.HandlerFunc{
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		 err := models.FindById(id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"ID": id,
				//"Name": t.Name,

			})
		}else {
			return err
		}
		//return c.JSON(http.StatusOK, H{
		//	"task": models.FindById(id),
		//}
	}
}

func UpdateTasks(db *gorm.DB)  echo.HandlerFunc{
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "GET Tasks")
	}
}
