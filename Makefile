build:
	GO111MODULE=on go build -o .build/gcloud main.go

update:
	go get -u all
	go mod tidy
	
.PHONEY: build update