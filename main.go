package main

import (
	"otus_project/internal/repository"
	"otus_project/internal/service"
)

func main() {
	repository.StartSliceLogger()
	service.GenerateData()
	
	select {}
}
