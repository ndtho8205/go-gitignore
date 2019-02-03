package main

import "flag"

var (
	createFlag string
	listFlag   string
	searchFlag string
)

func init() {
	flag.StringVar(&createFlag, "create", "", "Create .gitignore")
	flag.StringVar(&listFlag, "list", "", "List all supported templates")
	flag.StringVar(&searchFlag, "search", "", "Show all templates match string")
}

func main() {
	flag.Parse()

}
