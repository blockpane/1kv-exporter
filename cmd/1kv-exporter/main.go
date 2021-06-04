package main

import (
	"flag"
	"github.com/blockpane/1kv-exporter"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)

	var validatorsIn, chain string
	var port, ticker int
	flag.StringVar(&validatorsIn, "s", "", "Validators to monitor, comma seperated list of stash addresses, alternate: 'VALIDATORS' env var")
	flag.StringVar(&chain, "c", "kusama", "Which chain to monitor, possible choices are 'kusama' or 'polkadot'")
	flag.IntVar(&port, "p", 17586, "Port to listen on")
	flag.IntVar(&ticker, "t", 5, "Update frequency in minutes")
	flag.Parse()

	if validatorsIn == "" {
		validatorsIn = os.Getenv("VALIDATORS")
		if validatorsIn == "" {
			log.Fatal("no stash keys supplied, please provide a list of validators using '-s' or the VALIDATORS environment variable")
		}
	}

	validators := strings.Split(validatorsIn, ",")

	resultsChan := make(chan *kva.ValidatorInfo)
	go kva.Fetch(chain, ticker, validators, resultsChan)
	go kva.ProcessResult(chain, resultsChan)
	kva.Serve(port)

}
