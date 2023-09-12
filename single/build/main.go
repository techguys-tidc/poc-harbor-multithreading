package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/goombaio/namegenerator"
)

var (
	// harborURL = "harbor-s3.poc.workisboring.com"
	// project   = "knight"
	// imageName     = "helloword"
	// imageTag      = "latest"
	// username      = "robot$loadtest"
	// password      = "bpyKqEfWFCJuwFrPexqhBTK8RNhZEUyr"
	harborURL = os.Getenv("HARBOR_BASE_URL")
	project   = os.Getenv("HARBOR_PROJECT")
	imageTag  = "latest"
	username  = os.Getenv("HARBOR_USER")
	password  = os.Getenv("HARBOR_PASS")
)

func buildImage() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()
	imageName := name
	image := fmt.Sprintf("%s/%s/%s:%s", harborURL, project, imageName, imageTag)
	fmt.Printf("image name:%s", image)
	cmd := exec.Command("docker", "buildx", "build", "-t", image, ".")
	building, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to build to image: %v", err)
	}
	fmt.Printf("Building: %s\n", (building))

}

func main() {
	buildImage()
}
