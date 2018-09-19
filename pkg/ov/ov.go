package ov

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiEndpoint = "http://v0.ovapi.nl/tpc/"
	timeLayout  = "2006-01-02T15:04:05"
)

var Location *time.Location

func init() {
	var err error
	Location, err = time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		panic(err)
	}
}

type TimingPoint struct {
	Passes map[string]Pass //`json:"Passes"`
}

type Pass struct {
	ExpectedArrivalTime time.Time //time.Time //`json:"ExpectedArrivalTime"`
}

func (pass *Pass) UnmarshalJSON(data []byte) error {
	var stringyPass struct {
		ExpectedArrivalTime string
	}

	err := json.Unmarshal(data, &stringyPass)
	if err != nil {
		return err
	}

	expectedArrivalTime, err := time.ParseInLocation(
		timeLayout,
		stringyPass.ExpectedArrivalTime,
		Location,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	pass.ExpectedArrivalTime = expectedArrivalTime
	return nil
}

func (timingPoint *TimingPoint) NextArrivalAfter(now time.Time) time.Time {
	var nextArrival time.Time
	for _, pass := range timingPoint.Passes {
		if pass.ExpectedArrivalTime.After(now) && pass.ExpectedArrivalTime.Before(nextArrival) {
			nextArrival = pass.ExpectedArrivalTime
		}
	}

	return nextArrival
}

func NextArrivalAt(timingPointCode string) (*time.Time, error) {
	response, err := http.Get(apiEndpoint + timingPointCode)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var timingPoints map[string]TimingPoint
	err = json.Unmarshal(body, &timingPoints)
	if err != nil {
		return nil, err
	}

	timingPoint := timingPoints[timingPointCode]

	now := time.Now()
	nextArrival := now.Add(24 * time.Hour)
	noArrivals := true
	for _, pass := range timingPoint.Passes {
		if pass.ExpectedArrivalTime.After(now) && pass.ExpectedArrivalTime.Before(nextArrival) {
			nextArrival = pass.ExpectedArrivalTime
			noArrivals = false
		}
	}

	if noArrivals {
		return nil, &NoArrivalsError{timingPointCode}
	}

	return &nextArrival, nil
}
