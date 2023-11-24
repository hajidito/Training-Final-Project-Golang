package controller

import (
	"fmt"
	"html/template"
	"net/http"

	// "webserver/config"
	"webserver/config"
	"webserver/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func HelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello")
}

func JsonMap(ctx echo.Context) error {
	data := model.M{"message": "hello from/json", "counter": 2, "statusKode": http.StatusOK}
	return ctx.JSON(http.StatusOK, data)
}

func Page(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	data := fmt.Sprintf("helo %s", name)
	return ctx.String(http.StatusOK, data)
}

func User(ctx echo.Context) error {
	user := model.Item{}
	if err := ctx.Bind(&user); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

// func CreateEmployee(ctx echo.Context) error{
// 	db, err := config.Connect()
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	employee := model.Employee{}
// 	if err:= ctx.Bind(&employee); err != nil{
// 		return err
// 	}

// 	sqlStatement := `INSERT INTO employees (name, email, age, division) values ($1, $2, $3, $4)`

// 	_, err = db.Exec(sqlStatement, employee.Name, employee.Email, employee.Age, employee.Division)
// 	if err!= nil{
// 		panic(err)
// 	}
// 	return ctx.JSON(http.StatusOK, employee)
// }

// CreateEmployee godoc
// @Summary Create Employee
// @Description Create Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param model.Employee body model.Employee true "create employee"
// @Success 200 {object} model.Employee
// @Router /employee [post]
func CreateEmployee(ctx echo.Context) error {
	db := config.GetDB()
	employee := model.Employee{}
	if err := ctx.Bind(&employee); err != nil {
		return err
	}
	err := db.Create(&employee)
	if err.Error != nil {
		return ctx.JSON(http.StatusBadRequest, "tidak berhasil input")
	}
	fmt.Println("created employee")
	return ctx.JSON(http.StatusOK, employee)
}

func CreateItem(ctx echo.Context) error {
	db := config.GetDB()
	item := model.Item{}

	userData, ok := ctx.Get("userData").(jwt.MapClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   userData,
			"message": "failed to get user data",
		})
	}
	userID := uint(userData["id"].(float64))
	if err := ctx.Bind(&item); err != nil {
		return err
	}
	item.EmployeeId = int(userID)
	err := db.Create(&item)
	if err.Error != nil {
		return ctx.JSON(http.StatusBadRequest, "tidak berhasil input")
	}
	fmt.Println("created item")
	return ctx.JSON(http.StatusOK, item)
}

func UpdateEmployee(ctx echo.Context) error {
	db := config.GetDB()
	employee := model.Employee{}
	if err := ctx.Bind(&employee); err != nil {
		return err
	}
	db.Model(&employee).Where("id = ?", employee.ID).Updates(model.Employee{
		Name:     employee.Name,
		Email:    employee.Email,
		Age:      employee.Age,
		Division: employee.Division,
	})
	fmt.Println("updated employee")
	return ctx.JSON(http.StatusOK, employee)
}

func DeleteEmployee(ctx echo.Context) error {
	db := config.GetDB()
	employee := model.Employee{}
	delRes := model.DeleteResponse{
		Status:  http.StatusOK,
		Message: "delete success",
	}
	if err := ctx.Bind(&employee); err != nil {
		return err
	}
	paramId := ctx.Param("id")
	db.Model(&employee).Where("id = ?", paramId).Delete(&employee)
	fmt.Println("updated employee")
	return ctx.JSON(http.StatusOK, delRes)
}

func Index(ctx echo.Context) error {
	tmpl := template.Must(template.ParseGlob("template/*.html"))
	type M map[string]interface{}
	data := make(M)
	data[config.CSRFKey] = ctx.Get(config.CSRFKey)
	return tmpl.Execute(ctx.Response(), data)
}

func SayHello(ctx echo.Context) error {
	data := make(map[string]interface{})
	if err := ctx.Bind(&data); err != nil {
		return err
	}
	message :=
		fmt.Sprintf("hello %s, My gender %s", data["name"], data["gender"])
	return ctx.JSON(http.StatusOK, message)
}
