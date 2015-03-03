package main

import (
	"github.com/hashicorp/consul/api"

	"fmt"
	"net/http"
)

var client *api.Client

func handler(w http.ResponseWriter, r *http.Request) {
	kv := client.KV()
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "The Value: %s", pair.Value)
}

func main() {
	cfg := api.DefaultConfig()
	cfg.Address = "dev-consul:8500"
	cfg.Datacenter = "sjc-dev"
	client, _ = api.NewClient(cfg)

	kv := client.KV()
	p := &api.KVPair{Key: "foo", Value: []byte("test")}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
