package main

import (
	"github.com/rAlexander89/channel-meter-challenge/data"
	"github.com/rAlexander89/channel-meter-challenge/router"
)

func main() {
	go data.LoadData()
	router.Router()
}
