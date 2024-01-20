package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func addData(ctx echo.Context, dbConnection *sql.DB) error {
	var data TableData
	err := ctx.Bind(&data)
	if err != nil {
		return err
	}

	if data.Data == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Data cannot be empty"})
	}

	query := `
	insert into "%s"("%s") values($1) 
	`
	insertDataQuery := fmt.Sprintf(query, TableName, ColumnName)
	_, err = dbConnection.Exec(insertDataQuery, data.Data)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

func getData(ctx echo.Context, dbConnection *sql.DB) error {
	query := `
	select * from "%s"
	`
	getDataQuery := fmt.Sprintf(query, TableName)
	rows, err := dbConnection.Query(getDataQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	var dataResponse []TableData
	for rows.Next() {
		var data TableData
		rows.Scan(&data.Data)
		dataResponse = append(dataResponse, data)
	}

	return ctx.JSON(http.StatusOK, dataResponse)
}
