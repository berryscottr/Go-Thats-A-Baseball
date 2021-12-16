package gather

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ImportData() {
	fmt.Println("Importing data")
	fmt.Println(today)
	scheduleUrl := strings.Replace(scheduleUrlTemplate, "{year}", string(rune(today.year)), 1)
	resp, err := http.Get(scheduleUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)
	fmt.Println(resp.Body)
}
