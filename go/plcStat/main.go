package main

import ( 
  "fmt"
  "os"
  "strconv"
//  "time"
  "io"
	"net/http"
  "github.com/mackerelio/go-osstat/memory"
  //"github.com/mackerelio/go-osstat/cpu"
  )

func getMem(w http.ResponseWriter, r *http.Request) {
	
  memory, err := memory.Get()
  if err != nil {
 	fmt.Fprintf(os.Stderr, "%s\n", err)
 	return
 }
 //fmt.Printf("memory total: %d bytes\n", memory.Total
fmt.Printf("memory used: %d bytes\n", memory.Used)
 //fmt.Printf("memory cached: %d bytes\n", memory.Cach
 //fmt.Printf("memory free: %d bytes\n", memory.Free)
  s := strconv.FormatUint(memory.Used, 10)
  io.WriteString(w,s)
}
func getCpu(w http.ResponseWriter, r *http.Request) {
  //before, err := cpu.Get()
 // if err != nil {
	//	fmt.Fprintf(os.Stderr, "%s\n", err)
	//	return
	//}
	//time.Sleep(time.Duration(1) * time.Second)
	//after, err := cpu.Get()
	//if err != nil {
	  //fmt.Fprintf(os.Stderr, "%s\n", err)
	//total := float64(after.Total - before.Total)
	//fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	//fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	//fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/mem", getMem)
	http.HandleFunc("/cpu", getCpu)
	err := http.ListenAndServe(":3333", nil)
  if err != nil{
    fmt.Printf("ERROR")
  }
  }
