package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID         string `json:"id"`
	Item       string `json:"title"`
	Content    string `json:"content"`
	ContentImg string `json:"contentImg"`
	Logo       string `json:"logo"`
}

var todos = []todo{
	{ID: "1", Item: "Najot ta'lim", Content: "Hammasi Najot ta'limdan boshlanadi", ContentImg: "https://firebasestorage.googleapis.com/v0/b/businessanalytics-80462.appspot.com/o/contentsPhotos%2Fhammasi.jpeg?alt=media&token=0ed62361-f32a-45df-9ccb-eadfeb0e2fbf", Logo: "https://firebasestorage.googleapis.com/v0/b/businessanalytics-80462.appspot.com/o/najot_ta'lim.png?alt=media&token=0737bf98-8ced-4047-a608-818de1565f45"},
	{ID: "2", Item: "Thompson school", Content: "7+ IELTS ni 6 oyda qo'lga kiriting", ContentImg: "https://firebasestorage.googleapis.com/v0/b/businessanalytics-80462.appspot.com/o/contentsPhotos%2Fielts.jpg?alt=media&token=bbd5ca1f-1c9a-4ef4-bfa6-6ca5c45f5afa", Logo: "https://firebasestorage.googleapis.com/v0/b/businessanalytics-80462.appspot.com/o/thompson.jpeg?alt=media&token=99948c4e-1513-4799-be98-6aab610a5341"},
	{ID: "3", Item: "PDP academy", Content: "Android kursi uchun qabul boshlandi", ContentImg: "https://firebasestorage.googleapis.com/v0/b/businessanalytics-80462.appspot.com/o/contentsPhotos%2Fandroid.jpg?alt=media&token=690b0c56-2df2-4706-b40b-eb2e99dfe4b0", Logo: "https://firebasestorage.googleapis.com/v0/b/businessanalytics-80462.appspot.com/o/pdp.png?alt=media&token=aa40ce2c-ea8a-4660-a32a-2882358ffaa8"},
	{ID: "3", Item: "PDP academy", Content: "Android kursi uchun qabul boshlandi", ContentImg: "https://www.mytrendyphone.eu/images/iPad-Air-2020-LTE-64GB-Space-Grey-0190199788763-25092020-01-p.jpg"},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
	var newTodo todo

	if err:= context.BindJSON(&newTodo);err!=nil{
		return 
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated,newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos",addTodo)
	//router.DELETE("/todos",)
	router.Run("192.168.3.124:9090")
}
