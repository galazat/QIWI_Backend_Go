package commands

import (
	"fmt"
	"log"
	"time"
)

func GetUrl(date string) string {
	tm, err := time.Parse("2006-01-02", date[0:10])

	if err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?date_req=%v", tm.Format("02/01/2006"))
}
