package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Away0x/7yue_api_server/router"
	"log"
	"net/http"
	"github.com/Away0x/7yue_api_server/model"
	"github.com/Away0x/7yue_api_server/config"
)

func main() {
	g := gin.New()
	g.Static("/static", "static")
	g.LoadHTMLGlob("templates/*")

	config.InitConfig()

	db := model.InitDB()
	db.AutoMigrate(
		&model.Classic{},
		&model.Book{},
		&model.BookComment{},
		&model.Favor{},
		&model.User{},
		&model.HotKeyword{})
	defer db.Close()

	// mock
	mockInit()

	router.Register(g)

	log.Printf("Start to listening the incoming requests on http address: %s", ":8886")
	log.Fatal(http.ListenAndServe(":8886", g).Error())
}

// mock 数据
func mockInit() {
	// mock user
	//mock.PushDataIntoUserTable()

	// mock classic
	//mock.PushDataIntoClassicTable()

	// mock book
	//mock.PushDataIntoBookTable()

	// mock book short comment

	// mock hot keyword
	//mock.PushDataIntoHotKeyWordTable()
}