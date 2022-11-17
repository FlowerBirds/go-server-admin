package util

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
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

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
