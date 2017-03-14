package main

import (
	"time"
	"log"
	"os"
	"os/signal"
	"sync"
)

import (
	"golang.org/x/net/context"
)

func main() {

	log.Println("starting CHS Client")

	ctx, cancel := context.WithCancel(context.Background())

	// handle ctrl-c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	var wg sync.WaitGroup

	// start number printer
	log.Println("start number printer")
	go printNums(&wg, ctx)
	wg.Add(1)

	// cancel all background routines
	go func() {
		log.Println("waiting for ctrl-c")
		for _ = range signalChan {
			log.Println("received ctrl-c, cancelling all background tasks")
			cancel()
			break
		}
	}()

	// wait till all routines in waitgroup are done
	wg.Wait()

	log.Println("logging out from pdws")
}

func printNums(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	t10sec := time.Tick(1 * time.Second)
	i := 1
	
	for {
		select {
		case _ = <-t10sec:
			log.Println("i :", i)
			i++
		case <-ctx.Done():
			log.Println("cancel called, exiting alertMonitor")
			return
		}
	}
}