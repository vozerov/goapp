package main

import (
    "log"
    "net/http"
	"strconv"
	"encoding/json"
)

type Reply struct {
  Status string
  Code int
  Result int
}

func handler(w http.ResponseWriter, r *http.Request) {
  var a,b int
  var err error

  response := Reply{Status: "OK", Code: 0, Result: 0}

  keys, ok := r.URL.Query()["a"]
  if !ok || len(keys[0]) < 1 {
    response.Status = "param A is missing"
	response.Code = 1
  } else {
    a, err = strconv.Atoi(keys[0])
    if err != nil {
      response.Status = "param A is not a number"
	  response.Code = 2
    }
  }

  keys, ok = r.URL.Query()["b"]
  if !ok || len(keys[0]) < 1 {
    response.Status = "param B is missing"
	response.Code = 3
  } else {
    b, err = strconv.Atoi(keys[0])
    if err != nil {
      response.Status = "param B is not a number"
	  response.Code = 4
    }
  }

  if response.Code == 0 {
    response.Result = sum(a,b)
  }

  w.Header().Set("Content-Type", "application/json")

  data, err := json.Marshal(response)
  if err != nil {
    log.Println("ERROR: Can't marshal data")
  }
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(data)
}

func main() {
    http.HandleFunc("/sum/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func sum (a, b int) int {
  t := a + b
  return t
}
