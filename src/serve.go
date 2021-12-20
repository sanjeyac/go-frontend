package main
 
import "net/http"
import "os"

func main(){

	dir:= "."

	if len(os.Args) > 1 { 
        dir = os.Args[1]
    }

	http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))

}