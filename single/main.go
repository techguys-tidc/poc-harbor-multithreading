package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	// harborURL     = "harbor-s3.poc.workisboring.com"
	// project       = "knight"
	// imageName     = "helloword"
	// imageTag      = "latest"
	// username      = "robot$loadtest"
	// password      = "bpyKqEfWFCJuwFrPexqhBTK8RNhZEUyr"
	harborURL     = os.Getenv("HARBOR_BASE_URL")
	project       = os.Getenv("HARBOR_PROJECT")
	imageName     = "helloword"
	imageTag      = "latest"
	username      = "robot$loadtest"
	password      = "bpyKqEfWFCJuwFrPexqhBTK8RNhZEUyr"
	parallelCount = 5
	loopCount     = 100 // Number of times each worker will push
)

func pushImage(workerID int, ch chan time.Duration) {
	image := fmt.Sprintf("%s/%s/%s:%s", harborURL, project, imageName, imageTag)

	var totalDuration time.Duration

	for i := 0; i < loopCount; i++ {
		startTime := time.Now()
		cmd := exec.Command("docker", "push", image)
		pushing, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Worker %d: Failed to push image on loop %d: %v", workerID, i, err)
			continue
		}
		fmt.Printf("Pushing: %s", (pushing))

		elapsedTime := time.Since(startTime)
		totalDuration += elapsedTime

		log.Printf("Worker %d: Pushed image in loop %d in %v \n\n", workerID, i, elapsedTime)
	}

	ch <- totalDuration / time.Duration(loopCount) // Send average time of this worker
}

func main() {
	cmd := exec.Command("docker", "login", harborURL, "-u", username, "-p", password)
	pushing, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to login to harbor: %v", err)
	}
	fmt.Printf("Login: %s\n", (pushing))

	ch := make(chan time.Duration, parallelCount)

	for i := 0; i < parallelCount; i++ {
		go pushImage(i, ch)
	}

	var totalDuration time.Duration
	var maxDuration time.Duration
	var minDuration = time.Hour

	for i := 0; i < parallelCount; i++ {
		avgWorkerDuration := <-ch
		if avgWorkerDuration > maxDuration {
			maxDuration = avgWorkerDuration
		}
		if avgWorkerDuration < minDuration {
			minDuration = avgWorkerDuration
		}
		totalDuration += avgWorkerDuration
	}

	avgDuration := totalDuration / time.Duration(parallelCount)
	fmt.Printf("Average push time across all workers: %v\n", avgDuration)
	fmt.Printf("Max average push time among workers: %v\n", maxDuration)
	fmt.Printf("Min average push time among workers: %v\n", minDuration)

	cmd = exec.Command("docker", "logout", harborURL)
	logout, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to logout from harbor: %v", err)
	}
	fmt.Printf("Logout: %s\n", (logout))
}