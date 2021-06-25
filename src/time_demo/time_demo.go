package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("t  ", time.Now())
	fmt.Println("tuc ", time.Now().UTC())

	t := time.Now()
	year, month, day := t.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	fmt.Println(today)

	fmt.Println(today.Unix())
	fmt.Println(t.Unix() - today.Unix())
	fmt.Println("--------------------------")
	thisdate := "2020-05-08T18:05:00+08:00"
	timeformatdate, err := time.Parse(time.RFC3339, thisdate)
	fmt.Println(err)
	fmt.Println(timeformatdate)
	fmt.Println(timeformatdate.Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println("--------------------------")
	const template = "2006-01-02 15:04:05"
	t2, err := time.Parse(template, "2020-05-08 18:05:00")
	fmt.Println(err)
	fmt.Println(t2.Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println("--------------------------")
	thisdate2 := "2020-05-08 17:55:00"
	timeformatdate2, err := time.Parse(template, thisdate2)
	fmt.Println(err)
	fmt.Println(timeformatdate2)
	fmt.Println(timeformatdate2.Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println("--------------------------")
	loc, err := time.LoadLocation("Local")
	fmt.Println(loc)
	dt, err := time.ParseInLocation(template, "2020-05-08 18:19:00", loc)
	fmt.Println(dt.Unix())
	fmt.Println(time.Now().Unix())
	//loc, _ := time.LoadLocation("Local") //取本地时区
	strTime := time.Unix(dt.Unix(), 0).Format(template)
	fmt.Println(strTime)
}
