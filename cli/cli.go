package cli

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load("./config/.env"); err != nil {
		panic(err)
	}
	root, ok := os.LookupEnv("PIZZACOIN_ROOT")
	if !ok {
		fmt.Fprintf(os.Stderr, "Please setup a PIZZACOIN_ROOT variable\n")
		os.Exit(1)
	}
	fmt.Println(root)
}
