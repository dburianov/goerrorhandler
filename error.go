package error

import (
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func checkErrPlace(err error, errplace string) {
	if err != nil && errplace != "" {
		log.Println(errplace, err)
	}
}

func checkErrDebug(err error, debug bool) {
	if err != nil && debug == true {
		log.Println(err)
	}
}

func checkErrPlaceDebug(err error, errplace string, debug bool) {
	if err != nil && errplace != "" && debug == true {
		log.Println(errplace, err)
	}
}
