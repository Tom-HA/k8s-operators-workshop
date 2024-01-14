package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

const (
	TableName  = "random_data"
	ColumnName = "data"
)

var (
	pgHost      = os.Getenv("POSTGRES_HOST")
	pgPort      = os.Getenv("POSTGRES_PORT")
	pgUser      = os.Getenv("POSTGRES_USER")
	pgPassword  = os.Getenv("POSTGRES_PASSWORD")
	dbName      = os.Getenv("POSTGRES_DB")
	servicePort = os.Getenv("SERVICE_PORT")

	DBConn *sql.DB
	lock   = sync.Mutex{}
)

type PostgresConnection struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslMode  string
}

type TableData struct {
	Data string `json:"data"`
}

func init() {
	fmt.Println(os.Environ())
	pgConnection := PostgresConnection{
		host:     pgHost,
		port:     pgPort,
		user:     pgUser,
		password: pgPassword,
		dbName:   dbName,
		sslMode:  "disable",
	}

	var err error
	DBConn, err = getDBConnection(pgConnection)
	if err != nil {
		panic(err)
	}

	err = pingDatabase(DBConn)
	if err != nil {
		panic(err)
	}

	err = initDatabase(DBConn)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer DBConn.Close()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/data", func(c echo.Context) error {
		return getData(c, DBConn)
	})
	e.POST("/data", func(c echo.Context) error {
		return addData(c, DBConn)
	})

	if servicePort == "" {
		servicePort = "5001"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", servicePort)))

}

func getDBConnection(connectionInfo PostgresConnection) (dbConnection *sql.DB, err error) {
	pgConnectionInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		connectionInfo.host, connectionInfo.port, connectionInfo.user,
		connectionInfo.password, connectionInfo.dbName, connectionInfo.sslMode)
	db, err := sql.Open("postgres", pgConnectionInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func pingDatabase(dbConnection *sql.DB) error {
	waitSeconds := 5
	for counter := 0; counter < 7; counter++ {
		err := dbConnection.Ping()
		if counter == 6 {
			return fmt.Errorf("exceeded max timeout %d, could not connect to the database: %w", waitSeconds*6, err)
		}
		if err != nil {
			time.Sleep(time.Duration(waitSeconds) * time.Second)
		}
	}

	fmt.Println("Successfully connected to the database!")
	return nil
}

func initDatabase(dbConnection *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS %[1]s (
		%[2]s varchar(45) NOT NULL,
		PRIMARY KEY (%[2]s)
	)
	`
	createTableQuery := fmt.Sprintf(query, TableName, ColumnName)
	_, err := dbConnection.Exec(createTableQuery)

	return err
}

func addData(c echo.Context, dbConnection *sql.DB) error {
	lock.Lock()
	defer lock.Unlock()

	var data TableData
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	if data.Data == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Data cannot be empty"})
	}

	query := `
	insert into "%s"("%s") values($1) 
	`
	insertDataQuery := fmt.Sprintf(query, TableName, ColumnName)
	_, err = dbConnection.Exec(insertDataQuery, data.Data)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, data)
}

func getData(c echo.Context, dbConnection *sql.DB) error {
	lock.Lock()
	defer lock.Unlock()

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

	return c.JSON(http.StatusOK, dataResponse)
}
