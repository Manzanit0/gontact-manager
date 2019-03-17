package main

import "./addressbook"

func main() {
	r := addressbook.DefaultRouter()
	r.Run()
}
