#!/bin/bash
export HARBOR_URL='harbor.papermint.io'
export HARBOR_PROJECT='knight'
export HARBOR_USER='robot$loadtest'
export HARBOR_PASS='BAYo9ANN6z3eJrXrcLwj222capkz9s5d'
export IMAGE_NAME='helloword'
export IMAGE_TAG='latest'

docker buildx build -t $IMAGE_NAME:$IMAGE_TAG . 

go run main.go