package discovery

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"learn-go/packages/etcd/etcd"
	"log"
	"sync"
	"time"
)

type Service struct {
	Name     string
	IP       string
	Port     string
	Protocol string
}

func ServiceRegister(s *Service) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcd.GetEtcdEndpoint(),
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	var grantLease bool
	var leaseID clientv3.LeaseID
	ctx := context.Background()
	// 租约 - 声明一个租约，设置TTL为10秒
	getRes, err := client.Get(ctx, s.Name, clientv3.WithCountOnly())
	if err != nil {
		log.Println(err)
	}
	if getRes.Count == 0 { // 如果没有租约
		grantLease = true
	}
	if grantLease {
		leaseRes, err := client.Grant(ctx, 10)
		if err != nil {
			log.Println(err)
		}
		leaseID = leaseRes.ID
	}
	// 事务
	kv := clientv3.NewKV(client)
	txn := kv.Txn(ctx)
	_, err = txn.If(clientv3.Compare(clientv3.CreateRevision(s.Name), "=", 0)).
		Then(
			clientv3.OpPut(s.Name, s.Name, clientv3.WithLease(leaseID)),
			clientv3.OpPut(s.Name+".ip", s.IP, clientv3.WithLease(leaseID)),
			clientv3.OpPut(s.Name+".port", s.Port, clientv3.WithLease(leaseID)),
			clientv3.OpPut(s.Name+".protocol", s.Protocol, clientv3.WithLease(leaseID)),
		).
		Else( // WithIgnoreLease 是忽略当前租约，即将KV加到当前租约中，如果不加，会默认加入的KV是没有租约的长期KV
			clientv3.OpPut(s.Name, s.Name, clientv3.WithIgnoreLease()),
			clientv3.OpPut(s.Name+".ip", s.IP, clientv3.WithIgnoreLease()),
			clientv3.OpPut(s.Name+".port", s.Port, clientv3.WithIgnoreLease()),
			clientv3.OpPut(s.Name+".protocol", s.Protocol, clientv3.WithIgnoreLease()),
		).
		Commit()
	if err != nil {
		log.Println(err)
	}
	if grantLease {
		leaseKeepAlive, err := client.KeepAlive(ctx, leaseID)
		if err != nil {
			log.Println(err)
		}
		for lease := range leaseKeepAlive {
			fmt.Printf("lease ID: %x, ttl: %d \n", lease.ID, lease.TTL)
		}
	}
}

type Services struct {
	services map[string]*Service
	sync.RWMutex
}

var myServices = &Services{
	services: map[string]*Service{},
}

func ServiceDiscover(svcName string) *Service {
	var s *Service = nil
	myServices.RLock()
	s = myServices.services[svcName]
	myServices.RUnlock()
	return s
}

func WatchServiceName(svcName string) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcd.GetEtcdEndpoint(),
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Close()

	getRes, err := client.Get(context.Background(), svcName, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return
	}
	if getRes.Count > 0 {
		mp := SliceToMap(getRes.Kvs)
		s := &Service{}
		if kv, ok := mp[svcName]; ok {
			s.Name = string(kv.Value)
		}
		if kv, ok := mp[svcName+".ip"]; ok {
			s.IP = string(kv.Value)
		}
		if kv, ok := mp[svcName+".port"]; ok {
			s.Port = string(kv.Value)
		}
		if kv, ok := mp[svcName+".protocol"]; ok {
			s.Protocol = string(kv.Value)
		}
		myServices.Lock()
		myServices.services[svcName] = s
		myServices.Unlock()
	}

	rch := client.Watch(context.Background(), svcName, clientv3.WithPrefix())
	for wres := range rch {
		for _, ev := range wres.Events {
			if ev.Type == clientv3.EventTypeDelete {
				myServices.Lock()
				delete(myServices.services, svcName)
				myServices.Unlock()
			}
			if ev.Type == clientv3.EventTypePut {
				myServices.Lock()
				if _, ok := myServices.services[svcName]; !ok {
					myServices.services[svcName] = &Service{}
				}
				switch string(ev.Kv.Key) {
				case svcName:
					myServices.services[svcName].Name = string(ev.Kv.Value)
				case svcName + ".ip":
					myServices.services[svcName].IP = string(ev.Kv.Value)
				case svcName + ".port":
					myServices.services[svcName].Port = string(ev.Kv.Value)
				case svcName + ".protocol":
					myServices.services[svcName].Protocol = string(ev.Kv.Value)
				}
				myServices.Unlock()
			}
		}
	}
}

func SliceToMap(list []*mvccpb.KeyValue) map[string]*mvccpb.KeyValue {
	mp := make(map[string]*mvccpb.KeyValue)
	for _, item := range list {
		mp[string(item.Key)] = item
	}
	return mp
}
