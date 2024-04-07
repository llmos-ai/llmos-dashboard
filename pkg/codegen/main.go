package main

import (
	"log"
	"os"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	os.Unsetenv("GOPATH")
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed getting pwd: %v", err)
	}

	header, err := os.ReadFile(filepath.Join(pwd, "/hack/boilerplate.go.txt"))
	if err != nil {
		log.Fatalf("failed reading header: %v", err)
	}

	config := &gen.Config{
		Header:  string(header),
		Target:  "./pkg/generated/ent",
		Package: "github.com/llmos-ai/llmos-dashboard/pkg/generated/ent",
	}
	if err = entc.Generate("./pkg/types/v1", config); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
