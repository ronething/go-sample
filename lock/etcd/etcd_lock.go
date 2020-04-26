package main

import (
	"log"

	"github.com/zieckey/etcdsync"
)

// 在我本地运行不了 can not run on my macbook pro
// 2020/04/26 23:58:32 etcdsync.Lock failed%!(EXTRA *errors.errorString=client:
//	response is invalid json. The endpoint is probably not valid etcd cluster endpoint.)
// $ curl http://127.0.0.1:2379/version
// {"etcdserver":"3.4.1","etcdcluster":"3.4.0"}%
func main() {
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed", err)
		return
	}

	log.Printf("etcdsync.Lock OK")
	log.Printf("Get the lock. Do something here.")

	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}
}
