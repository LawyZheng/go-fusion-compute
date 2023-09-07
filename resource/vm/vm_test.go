package vm

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/lawyzheng/go-fusion-compute/client"
	"github.com/lawyzheng/go-fusion-compute/resource/site"
	"github.com/lawyzheng/go-fusion-compute/resource/task"
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
		cs, err := cm.ListVm(context.Background(), s.Uri, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs)
	}
}

func TestManager_CloneVm(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect(context.Background())
	_ = "/service/sites/43BC08E8"
	m := NewManager(c)
	ts, err := m.CloneVm(context.Background(),
		"/service/sites/43BC08E8/vms/i-00000034", CloneVmRequest{
			Name:          "test-1",
			Description:   "test create vm",
			Location:      "urn:sites:43BC08E8:clusters:117",
			IsBindingHost: false,
			Config: Config{
				Cpu: Cpu{
					Quantity:    2,
					Reservation: 0,
				},
				Memory: Memory{
					QuantityMB:  2048,
					Reservation: 2048,
				},
				Disks: []Disk{
					{
						SequenceNum:  1,
						QuantityGB:   50,
						IsDataCopy:   true,
						DatastoreUrn: "urn:sites:43BC08E8:datastores:41",
						IsThin:       true,
					},
				},
				Nics: []Nic{
					{
						Name:         "vmnic1",
						PortGroupUrn: "urn:sites:43BC08E8:dvswitchs:1:portgroups:1",
					},
				},
			},
			VmCustomization: Customization{
				OsType:             "Linux",
				Hostname:           "test-1",
				IsUpdateVmPassword: false,
				NicSpecification: []NicSpecification{
					{
						SequenceNum: 1,
						Ip:          "100.199.10.88",
						Netmask:     "255.255.255.0",
						Gateway:     "100.199.10.1",
						Setdns:      "114.114.114.114",
						Adddns:      "8.8.8.8",
					},
				},
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("create vm %s", ts.Uri)
	fmt.Printf("task uri  %s", ts.TaskUri)

	tm := task.NewManager(c)
	for {
		tt, err := tm.Get(context.Background(), ts.TaskUri)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("task status %s", tt.Status)
		if tt.Status == "success" {
			break
		}
		time.Sleep(5 * time.Second)
	}

	_, err = m.DeleteVm(context.Background(), ts.Uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete vm %s", ts.Uri)

}
