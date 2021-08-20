package main
import (
	"fmt"
	"gee"
	"net/http"
)


func main() {
	//engine := new(Engine)
	g :=gee.New()
	g.GET("/",func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	g.GET("/hello",func(w http.ResponseWriter, req *http.Request){
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}})

	g.Run(":9999")
	//log.Fatal(http.ListenAndServe(":9999", engine))
}
