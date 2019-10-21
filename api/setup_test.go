package api

import (
	"go.jinya.de/ontheroad/dummy_data"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("Setting up the tests")
	dummy_data.ClearProjects()
	dummy_data.FillProjects()
	exitCode := m.Run()
	dummy_data.ClearProjects()
	os.Exit(exitCode)
}
