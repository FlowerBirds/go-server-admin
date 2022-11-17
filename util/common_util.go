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
	"strconv"
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

func CompareVersion(selfVersion string, serverVersion string) int {
	if len(selfVersion) == 0 || len(selfVersion) == 0 {
		return 0
	}
	v1 := strings.Split(selfVersion, ".")
	v2 := strings.Split(selfVersion, ".")
	if len(v1) != 3 || len(v2) != 3 {
		return 0
	}
	// 对三位版本号进行校验，逐一对比
	for i := 0; i < 3; i++ {
		n1, _ := strconv.Atoi(v1[i])
		n2, _ := strconv.Atoi(v2[i])
		if n1 != n2 {
			return n1 - n2
		}
	}

	return 0
}
