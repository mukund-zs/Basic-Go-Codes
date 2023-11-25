package main

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type data struct {
	EnrollmentNumber int    `json:"enrollmentNumber"`
	Name             string `json:"name"`
}

func main() {
	// Creating GoFr's instance
	app := gofr.New()

	// POST endpoint using GoFr
	app.POST("/post", func(ctx *gofr.Context) (interface{}, error) {
		var d data

		// Reading and converting request data to go struct type
		err := ctx.Bind(&d)
		if err != nil {
			return nil, err
		}

		// Insert operation for database
		_, dbErr := ctx.DB().Exec("INSERT INTO students (enrollment_number,name) values(?,?)", d.EnrollmentNumber, d.Name)
		if dbErr != nil {
			return nil, &errors.DB{Err: err}
		}

		return "Student added successfully!", nil
	})

	// Starting the server
	app.Start()
}
