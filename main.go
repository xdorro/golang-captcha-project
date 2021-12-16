package main

import (
	"log"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	// if err := client.SetWhitelist("0123456789"); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	//
	// if err := client.SetBlacklist("!?@#$%&*()<>_-+=/:;'\\\"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// if err := client.SetImage("./data/vcb/3.jpg"); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// MBB

	if err := client.SetImage("./data/mbb/3.png"); err != nil {
		log.Fatal(err)
		return
	}
	if err := client.SetWhitelist("0123456789"); err != nil {
		log.Fatal(err)
		return
	}

	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
		return
	}

	text = strings.ReplaceAll(text, " ", "")

	log.Println(text)
}
