// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2013 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	//"flag"
	"fmt"
	"github.com/bluele/rpc/v2"
	"github.com/bluele/rpc/v2/json2"
	"log"
	"strings"
)

type Counter struct {
	Count int
}

type IncrReq struct {
	Delta int
}

// Notification.
func (c *Counter) Incr(req *IncrReq, res *json2.EmptyResponse) error {
	log.Printf("<- Incr %+v", *req)
	c.Count += req.Delta
	return nil
}

type GetReq struct {
}

func (c *Counter) Get(req *GetReq, res *Counter) error {
	log.Printf("<- Get %+v", *req)
	*res = *c
	log.Printf("-> %v", *res)
	return nil
}

func main() {
	incrJson := `
{"jsonrpc":"2.0", "method":"Counter.Incr", "params": {"delta": 10}, "id": "1"}
`
getJson := `
{"jsonrpc":"2.0", "method":"Counter.Get", "params": {}, "id": "1"}
`

	//address := flag.String("address", ":65534", "")
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json")
	fmt.Println(s.RegisterService(new(Counter), ""))
	for i := 0; i < 5; i++ {
			s.ServeRequest("application/json", strings.NewReader(incrJson))
	}
	s.ServeRequest("application/json", strings.NewReader(getJson))	
	/*
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
		http.Handle("/jsonrpc/", s)
		log.Fatal(http.ListenAndServe(*address, nil))
	*/
}
