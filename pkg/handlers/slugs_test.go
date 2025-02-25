package handlers_test

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3"
)

const (
	testingDB = "testing.db"
)

var 

var _ = Describe("slugs creation", func() {
	BeforeAll(func() {
		file, err := os.Create(testingDB) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("db created")
	})

	BeforeEach(func() {

	})

	Context("When a valid URL is submitted", func() {
		It("should generate a slug", func() {

		})
	})

	AfterAll(func() {
		os.Remove(testingDB)
	})
})
