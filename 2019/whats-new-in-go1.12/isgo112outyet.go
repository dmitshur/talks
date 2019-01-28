// +build ignore

package main

import (
	"context"
	"fmt"
	"log"

	"grpc.go4.org"
)

import "golang.org/x/build/maintner/maintnerd/apipb"

func main() {
	cc, err := grpc.NewClient(nil, "https://maintner.golang.org")
	if err != nil {
		log.Fatalln(err)
	}
	mc := apipb.NewMaintnerServiceClient(cc)

	err = PrintGoReleases(mc)
	if err != nil {
		log.Fatalln(err)
	}
}

// PrintGoReleases uses the provided maintner client
// to print the latest supported Go releases.
func PrintGoReleases(mc apipb.MaintnerServiceClient) error {
	resp, err := mc.ListGoReleases(context.Background(), &apipb.ListGoReleasesRequest{})
	if err != nil {
		return err
	}
	for _, r := range resp.Releases {
		fmt.Println(r.TagName)
	}
	return nil
}
