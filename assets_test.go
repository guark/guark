package guark

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/guark/guark/internal/generator"
)

var (
	dir  string
	file string
)

func TestAssets(t *testing.T) {

	setup()
	defer teardown()

	if err := generator.Assets(dir+"/assets", dir+"/guark_assets", file, "main", dir); err != nil {
		t.Error(err)
	}

	t.Run("Create Main", createTestMain)
	t.Run("Verify Assets", verify)
}

func createTestMain(t *testing.T) {

	if err := createMain(); err != nil {
		t.Error(err)
	}
}

func createMain() error {

	f, err := os.Create(dir + "/main.go")

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(`package main

import (
	"log"

	"github.com/guark/guark/app/utils"
)

func main() {

	files := []string{"/assets/file1", "/assets/file2"}

	for i := range files {
		if Assets.Has(files[i]) == false {
			log.Fatalf("could not find: %s", files[i])
		} else if utils.IsFile("guark_assets/"+ Assets.Index[files[i]]) == false {
			log.Fatalf("could not find static asset %s of %s", Assets.Index[files[i]], files[i])
		}
	}
}
`)

	return err
}

func verify(t *testing.T) {

	cmd := exec.Command("go", "run", "guark_assets.go", "main.go")
	cmd.Dir = dir

	b, err := cmd.CombinedOutput()

	if err != nil {
		t.Log(string(b))
		t.Error(err)
	}
}

func setup() {

	dir = getTestDataDir()
	file = dir + "/guark_assets.go"
	os.Mkdir(dir+"/guark_assets", 0760)
}

func teardown() {
	os.Remove(file)
	os.Remove(dir + "/main.go")
	os.RemoveAll(dir + "/guark_assets")
}

func getTestDataDir() string {

	_, path, _, ok := runtime.Caller(0)

	if !ok {
		panic("could not get caller!")
	}

	return filepath.Dir(path) + "/testdata"
}
