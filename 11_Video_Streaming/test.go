/*
A sample program that just streams video
*/
package main;

import (
  "net/http"
  "fmt"
  "io/ioutil"
)

func lol(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Serving File for :",req)
  http.ServeFile(w,req,"./test.mp4")
}

func index(w http.ResponseWriter, req *http.Request) {
    content,err := ioutil.ReadFile("./sample.html")
    if (err != nil) {
      fmt.Fprint(w,"Error Reading File");
      fmt.Println(err)
      return;
    }
    fmt.Fprint(w,string(content));
}
func main() {
  http.HandleFunc("/",index)
  http.HandleFunc("/test.mp4",lol);
  http.ListenAndServe(":8080",nil);
}
