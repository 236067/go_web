package main
import(
"fmt"
"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello",r.URL.Path)
}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe("9090",nil)

}
