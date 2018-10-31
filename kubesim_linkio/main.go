/*
Copyright 2018 Kubedge

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/kubedge/kubesim_base/config"
	"github.com/kubedge/kubesim_linkio/linkio"
	"log"
	"os"
	"time"
	"bytes"
	"io"
)

func TestOne() {
	// a dummy buffer full of zeros to send over the link
	var y [1000]byte
	buf := bytes.NewBuffer(y[:])

	lr := linkio.NewLink(30 /* kbps */).NewLinkReader(buf)
	for {
		var x [1024]byte
		n, err := lr.Read(x[:])
		if n != 0 { 
                }
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("err %s", err)
		}
	}
}

//arguments
// arg1=demotype
// arg2=demovalue
func main() {
	log.Printf("%s", "kubesim Linkio client is running")
	demotype := os.Args[1]
	demovalue := os.Args[2]
	log.Printf("demotype=%s, demovalue=%s", demotype, demovalue)

	var conf config.Configdata
	conf.Config()
	log.Printf("kubesim 5G NR client:  product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s",
		conf.Product_name, conf.Product_type, conf.Product_family, conf.Product_release, conf.Feature_set1, conf.Feature_set2)

	log.Printf("starting for loop")
	for {
		time.Sleep(15 * time.Second) //every 15 seconds
	}
	log.Printf("%s", "kubesim 5G NR client is exiting")
}
