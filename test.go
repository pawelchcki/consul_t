package main

import (
	"github.com/hashicorp/consul/api"

	"fmt"
	"time"
	"net/http"
)

var client *api.Client

var responseKey = "foo"

func handler(w http.ResponseWriter, r *http.Request) {
	kv := client.KV()
	pair, _, err := kv.Get(responseKey, nil)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "The Value: %s", pair.Value)
}

func healthcheck(){
    for i := 0; i < 50000; i++ {
        time.Sleep(1 * time.Second)

        fmt.Println("HC")
    }
}

func main() {
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	client, _ = api.NewClient(cfg)

	kv := client.KV()
	p := &api.KVPair{Key: responseKey, Value: []byte("test")}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}


	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":8080", nil)

        if (err != nil){
		panic(err)
	}
}
