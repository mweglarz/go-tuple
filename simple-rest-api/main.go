package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {

	var err error
	db, err := sql.Open("postgres", "user=psql-go dbname=psql-go-db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected ...")
	}

	e := echo.New()

	type Employee struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Salary string `json:"salary"`
		Age    string `json:"age"`
	}
	type Employees struct {
		Employees []Employee `json:"Employees"`
	}

	e.POST("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "INSERT INTO employees (name, salary, age) VALUES ($1, $2, $3)"
		res, err := db.Query(sqlStatement, u.Name, u.Salary, u.Age)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusCreated, "ok")
	})

	e.PUT("/employee", func(c echo.Context) error {
		u := new(Employee)
		if err := c.Bind(u); err != nil {
			return err
		}
		sqlStatement := "UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$4"
		res, err := db.Query(sqlStatement, u.Name, u.Salary, u.Age, u.Id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, u.Id)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
