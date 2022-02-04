package customers_test

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func init() {
	if err := os.Chdir("../"); err != nil {
		log.Fatal(err)
	}
}
