package storage

import (
	"fmt"
	"log"
	"testing"

	"github.com/lawyzheng/go-fusion-compute/pkg/client"
	"github.com/lawyzheng/go-fusion-compute/pkg/site"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect()

	sm := site.NewManager(c)
	ss, err := sm.ListSite()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss {
		cm := NewManager(c, s.Uri)
		cs, err := cm.ListDataStore()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs)
	}
}
