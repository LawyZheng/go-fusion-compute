package site

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lawyzheng/go-fusion-compute/client"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect()
	m := NewManager(c)
	ss, err := m.ListSite(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ss)
}
