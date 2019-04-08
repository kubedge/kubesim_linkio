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
	"bufio"
	"fmt"
	"github.com/kubedge/kubesim_base/config"
	"github.com/kubedge/kubesim_linkio/pkg/linkio"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const SIM_NAME = "KUBESIM LINKIO"
const SIM_CONFIG_FILE = "/etc/kubedge/kubesim_conf.yaml"
const SIM_CONNECTED_UE_FILE = "/etc/kubedge/connected_ue.yaml"
const SIMPLE_HTTP_SERVER = false

func sim_message(msg string) {
	log.Printf("%s: %s", SIM_NAME, msg)
}

func configAPI(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = SIM_NAME + " : " + message
	w.Write([]byte(message))
}

func simulate(targeturl string, targetspeed linkio.Throughput) {
	// Create a new link at targetspeed
	link := linkio.NewLink(targetspeed)

	// Open a connection
	conn, err := net.Dial("tcp", targeturl)
	if err != nil {
		// handle error:w
	}

	// Create a link reader/writer
	linkReader := link.NewLinkReader(io.Reader(conn))
	linkWriter := link.NewLinkWriter(io.Writer(conn))

	// Use them as you would normally...
	fmt.Fprintf(linkWriter, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(linkReader).ReadString('\n')

	log.Printf("simulation status[%s] err[%s]", status, err)
}

func main() {
	sim_message("Starting")
	maintargeturl := os.Args[1]
	maintargetspeed := os.Args[2]
	log.Printf("targeturl=%s, targetspeed=%s", maintargeturl, maintargetspeed)

	var conf config.Configdata
	conf.Config(SIM_NAME, SIM_CONFIG_FILE)
	log.Printf("%s: product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s",
		SIM_NAME, conf.Product_name, conf.Product_type, conf.Product_family, conf.Product_release, conf.Feature_set1, conf.Feature_set2)

	targeturl := "google.com:80"
	targetspeed := 512 * linkio.KilobitPerSecond
	sim_message("Starting MSP Simulation")
	for {
		simulate(targeturl, targetspeed)
		time.Sleep(15 * time.Second) //every 15 seconds
	}
	// sim_message("Exiting")
}
