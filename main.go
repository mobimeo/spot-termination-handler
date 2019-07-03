package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ECSMetadata struct {
	Cluster              string `json:"Cluster"`
	ContainerInstanceArn string `json:"ContainerInstanceArn"`
}

func main() {
	client := &http.Client{Timeout: time.Second * 10}

	for {
		time.Sleep(5 * time.Second)
		response, err := client.Get("http://169.254.169.254/latest/meta-data/spot/termination-time")
		if err != nil {
			panic(err)
		}

		if response.StatusCode == 200 {
			log.Print("Spot instance termination notice detected")
			break
		}
	}

	response, err := client.Get("http://localhost:51678/v1/metadata")
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var ecsMetadata ECSMetadata
	err = json.Unmarshal(buf, &ecsMetadata)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(session.New(&aws.Config{}))
	status := "DRAINING"
	state := &ecs.UpdateContainerInstancesStateInput{
		Cluster:            &ecsMetadata.Cluster,
		ContainerInstances: []*string{&ecsMetadata.ContainerInstanceArn},
		Status:             &status,
	}
	log.Print("Putting instance in state DRAINING")
	_, err = svc.UpdateContainerInstancesState(state)
	if err != nil {
		panic(err)
	}

	log.Print("Sleeping for 120s until termination")
	time.Sleep(120 * time.Second)
}
