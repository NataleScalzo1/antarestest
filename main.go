package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	//mux := mux.NewRouter()
	//mux.HandleFunc("/a", ReadCSVFromHttpRequest)
	//err := http.ListenAndServe(":4000", mux)
	//if err != nil {
	//	log.Println("error: "+ err.Error())
	//}

	e := echo.New()
	routes(e)
	e.Logger.Fatal(e.Start(":4040"))
}
