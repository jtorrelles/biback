package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"

	dbase "biback/app/db"

	_showController "biback/app/controllers"
	_showRepo "biback/app/repository/show"
	_showService "biback/app/services/show"
)

func main() {

	/*dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())*/
	//dbConn, err := sql.Open(`mysql`, dsn)
	dbConn := dbase.ConnectDb()
	/*if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}*/

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	ar := _showRepo.NewShowRepository(dbConn)

	timeoutContext := time.Duration(2) * time.Second
	au := _showService.NewShowService(ar, timeoutContext)
	_showController.NewShowHandler(e, au)
	/*e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})*/

	/** Shows Routes */
	//e.GET("/shows", controllers.GetShows)
	//e.GET("/shows/:id", controllers.GetShowById)
	//e.POST("/shows", controllers.NewShow)

	e.Logger.Fatal(e.Start(":1323"))
}
