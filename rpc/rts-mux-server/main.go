package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	iot_realtime "./iot_realtime"
)

const (
	defaultWriteTimeout = 15 * time.Second
	defaultReadTimeout  = 15 * time.Second
)

func getRealTimeStruct(t time.Time) *iot_realtime.RealTime {
	return &iot_realtime.RealTime{
		Timestamp: t.Unix(),
		Parsed: &iot_realtime.RealTime_ParsedTime{
			Year:      int32(t.Year()),
			Month:     int32(t.Month()),
			Day:       int32(t.Day()),
			DayOfWeek: t.Weekday().String(),
			Hour:      int32(t.Hour()),
			Minute:    int32(t.Minute()),
			Second:    int32(t.Second()),
		},
	}
}

func getLocalTime(w http.ResponseWriter, r *http.Request) {
	realTime := getRealTimeStruct(time.Now())
	byteData, err := proto.Marshal(realTime)

	if err != nil {
		log.Println(err)
		return
	}

	w.Write(byteData)
}

func main() {
	var port = flag.String("port", "50001", "Port number to run the server on")

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/localTime", getLocalTime).Methods("GET")

	handler := handlers.LoggingHandler(os.Stdout, r)

	rtsServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:" + *port,
		WriteTimeout: defaultWriteTimeout,
		ReadTimeout:  defaultReadTimeout,
	}

	log.Println("Starting server on port", *port)
	log.Fatal(rtsServer.ListenAndServe())
}
