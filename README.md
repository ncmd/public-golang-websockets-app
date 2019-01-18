# public-golang-websockets-app
[![Build Status](https://travis-ci.com/ncmd/nervdaq-golang-websockets-app.svg?token=yDgzEp78NY59NKQDq1hd&branch=master)](https://travis-ci.com/ncmd/nervdaq-golang-websockets-app)

- follow instructions: 
- https://docs.travis-ci.com/user/languages/go/
- https://docs.travis-ci.com/user/deployment/heroku/

# Overview
- Uses Gorilla websockets library
- Main use is to provide a real-time experience to broadcast data to all connected clients simultaneously
- This will be used for the stock trading ticker and transactions
- This is not meant for anything else; Use restapi app instead
- You can view the production app at: https://nervdaq-go-websockets-prod.herokuapp.com/

# Quickstart
1. Install Go version 1.11
2. go get -u github.com/tools/godep
3. godep save ./...
2. cd /Users/<username>/go/src/nervdaq/ && go run *.go
3. Open web browser and visit: http://localhost:8080/

# Travis deployment
- Reference: https://medium.com/@felipeluizsoares/automatically-deploy-with-travis-ci-and-heroku-ddba1361647f
1. Create a heroku app first
2. Configure .travis.yml file

