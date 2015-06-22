package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			//fmt.Println(string(k))
			if k != "Title" {
				v[k] = 0
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
