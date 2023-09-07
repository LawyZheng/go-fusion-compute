package storage

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/resource/site"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect(context.Background())

	sm := site.NewManager(c)
	ss, err := sm.ListSite(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss {
		cm := NewManager(c)
		cs, err := cm.ListDataStore(context.Background(), s.Uri)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs)
	}
}
