package main
import ( 
   "net/http"
   "log"
   "dynamodb-crud/router"
)


func main() {
	router := router.Initroute()
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Server Lister on :8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}