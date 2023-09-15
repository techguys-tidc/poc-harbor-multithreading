#!/bin/bash
export HARBOR_URL='harbor-s3.poc.workisboring.com'
export HARBOR_PROJECT='knight'
export HARBOR_USER='robot$loadtest'
export HARBOR_PASS='pn5vu0JHgCA6RmJ0iTfPtbfj76Mi9HcE'
export IMAGE_NAME='helloword'
export IMAGE_TAG='latest'

docker buildx build -t $IMAGE_NAME:$IMAGE_TAG . 

go run main.go