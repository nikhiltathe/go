package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type DashboardViewRes struct {
	NumAlerts  int
	NumChassis int
	NumFans    int
}
type ResultSetNumAlerts interface {
	SetNumAlerts(int)
}
type ResultSetNumChassis interface {
	SetNumChassis(int)
}
type ResultSetNumFans interface {
	SetNumFans(int)
}

func (r *DashboardViewRes) SetNumAlerts(n int) {
	r.NumAlerts = n
}
func (r *DashboardViewRes) SetNumChassis(n int) {
	r.NumChassis = n
}
func (r *DashboardViewRes) SetNumFans(n int) {
	r.NumFans = n
}

type ChassisViewRes struct {
	NumAlerts int
}

func (r *ChassisViewRes) SetNumAlerts(n int) {
	r.NumAlerts = n
}

func GetNumAlerts() (int, error) {
	pause := time.Duration(rand.Intn(10)) * time.Second
	log.Println("na will take", pause, "secs")
	time.Sleep(pause)
	return 10, nil
}

func GetNumChassis() (int, error) {
	pause := time.Duration(rand.Intn(10)) * time.Second
	log.Println("nc will take", pause, "secs")
	time.Sleep(pause)
	return 1, nil
}

func GetNumFans() (int, error) {
	pause := time.Duration(rand.Intn(10)) * time.Second
	log.Println("nf will take", pause, "secs")
	time.Sleep(pause)
	return 2, nil
}

func getNumFansPara(wg *sync.WaitGroup, r ResultSetNumFans) {
	defer wg.Done()
	nf, err := GetNumFans()
	if err != nil {
		return
	}
	log.Println("got num fans:", nf)
	r.SetNumFans(nf)

}
func getNumAlertsPara(wg *sync.WaitGroup, r ResultSetNumAlerts) {
	defer wg.Done()
	na, err := GetNumAlerts()
	if err != nil {
		return
	}
	log.Println("got num alerts:", na)
	r.SetNumAlerts(na)

}
func getNumChassisPara(wg *sync.WaitGroup, r ResultSetNumChassis) {
	defer wg.Done()
	nc, err := GetNumChassis()
	if err != nil {
		return
	}

	log.Println("got num chassis:", nc)
	r.SetNumChassis(nc)
}

// start all the queries to run in parallel
// wait for N seconds for the results. whatever becomes available
// in those N seconds, take that and display the results
func main() {
	log.Println("==== dashboard")
	rand.Seed(time.Now().Unix())
	var dWg sync.WaitGroup
	r := new(DashboardViewRes)

	log.Println("getting num alerts")
	go getNumAlertsPara(&dWg, r)
	dWg.Add(1)

	log.Println("getting num chassis")
	go getNumChassisPara(&dWg, r)
	dWg.Add(1)

	log.Println("getting num fans")
	go getNumFansPara(&dWg, r)
	dWg.Add(1)

	waitChan := make(chan struct{})
	go func() {
		dWg.Wait()
		close(waitChan)
	}()

	log.Println("dashboard waiting for 5 seconds")
	slaTime := time.NewTimer(5 * time.Second)
	select {
	case <-waitChan:
		log.Println("dashboard finished before time!")
	case <-slaTime.C:
		log.Println("dashboard sla reached. displaying.")
	}
	log.Printf("%+v\n", r)

	log.Println("==== chassis")
	var cWg sync.WaitGroup
	rand.Seed(time.Now().Unix())
	cRes := new(ChassisViewRes)

	log.Println("getting num alerts")
	go getNumAlertsPara(&cWg, cRes)
	cWg.Add(1)

	waitChan = make(chan struct{})
	go func() {
		cWg.Wait()
		close(waitChan)
	}()

	log.Println("chassis waiting for 5 seconds")
	slaTime = time.NewTimer(5 * time.Second)
	select {
	case <-waitChan:
		log.Println("chassis finished before time!!")
	case <-slaTime.C:
		log.Println("chassis sla reached. displaying.")
	}
	log.Printf("%+v\n", cRes)

}
