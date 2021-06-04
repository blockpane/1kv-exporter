package kva

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	k = "https://kusama.w3f.community/candidate/"
	d = "https://polkadot.w3f.community/candidate/"
)

func Fetch(chain string, ticker int, stashes []string, results chan *ValidatorInfo) {
	var baseUrl string
	switch chain {
	case "kusama":
		baseUrl = k
	case "polkadot":
		baseUrl = d
	default:
		log.Fatalf("Invalid chain '%s' choices are 'kusama' or 'polkadot'\n", chain)
	}

	fetch := func(stash string) (*ValidatorInfo, error) {
		resp, err := http.Get(baseUrl + stash)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
		info := &ValidatorInfo{}
		err = json.Unmarshal(body, info)
		if err != nil {
			return nil, err
		}
		return info, err
	}

	fetchAll := func() {
		for i := range stashes {
			go func(j int) {
				r, e := fetch(stashes[j])
				if e != nil {
					log.Println(e)
					return
				}
				results <- r
			}(i)
		}
	}
	fetchAll()

	tick := time.NewTicker(time.Duration(ticker) * time.Minute)
	for {
		select {
		case <- tick.C:
			fetchAll()
		}
	}
}
