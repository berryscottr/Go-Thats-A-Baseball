package gather

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ImportData() {
	fmt.Println("Importing data")
	resp, err := http.Get("https://www.baseball-reference.com/leagues/majors/2021-schedule.shtml")
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
