package network

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/resource/site"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "fit2cloud", "Huawei@1234")
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
		cm := NewManager(c, s.Uri)
		cs, err := cm.ListDVSwitch(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs[0].Uri)
		pg, err := cm.ListPortGroup(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range pg {
			ips, err := cm.ListPortGroupInUseIp(context.Background(), p.Urn)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ips)
		}
	}
}
