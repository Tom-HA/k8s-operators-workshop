package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

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
	pgDBName    = os.Getenv("POSTGRES_DB")
	pgSSLMode   = os.Getenv("POSTGRES_SSLMODE")
	servicePort = os.Getenv("SERVICE_PORT")
)

type PostgresConConfig struct {
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

func main() {

	if pgSSLMode == "" {
		pgSSLMode = "disable"
	}

	dbConn, err := initDatabase(PostgresConConfig{
		host:     pgHost,
		port:     pgPort,
		user:     pgUser,
		password: pgPassword,
		dbName:   pgDBName,
		sslMode:  pgSSLMode,
	})
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/health"
		},
	}))
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		err := dbConn.Ping()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("could not connect to the database: %w", err))
		}
		return c.NoContent(http.StatusOK)
	})
	e.GET("/data", func(c echo.Context) error {
		return getData(c, dbConn)
	})
	e.POST("/data", func(c echo.Context) error {
		return addData(c, dbConn)
	})

	if servicePort == "" {
		servicePort = "5001"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", servicePort)))
}

func initDatabase(connectionInfo PostgresConConfig) (dbConnection *sql.DB, err error) {
	dbConn, err := getDBConnection(connectionInfo)
	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %w", err)
	}

	err = provisionDatabase(dbConn)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func getDBConnection(connectionInfo PostgresConConfig) (dbConnection *sql.DB, err error) {
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

func provisionDatabase(dbConnection *sql.DB) error {
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
