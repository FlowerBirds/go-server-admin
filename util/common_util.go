package util

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"
)

var debug = false

func GetNowUnix() int64 {
	return time.Now().UnixMilli()
}

func UpdateDebug(d bool) {
	debug = d
}

func EnableDebug() bool {
	return debug
}

func ResponseWithJson(v interface{}, w http.ResponseWriter) {
	msg, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func IsInstanceOf(o1, o2 interface{}) bool {
	return reflect.TypeOf(o1) == reflect.TypeOf(o2)
}
