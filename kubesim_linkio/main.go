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
  "log"
  "os"
  "net/http"
  "strings"
  "time"
  "github.com/kubedge/kubesim_base/grpc/go/kubedge_client"
  "github.com/kubedge/kubesim_base/config"
  "github.com/kubedge/kubesim_base/connected"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Kubesim LTE Simulator: " + message
  w.Write([]byte(message))
}

//arguments
// arg1=demotype
// arg2=demovalue
func main() {
  log.Printf("%s", "kubesim 5G NR client is running")
  demotype := os.Args[1]
  demovalue := os.Args[2]
  log.Printf("demotype=%s, demovalue=%s", demotype, demovalue)

  var conf config.Configdata
  conf.Config()
  log.Printf("kubesim 5G NR client:  product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s",
             conf.Product_name, conf.Product_type, conf.Product_family, conf.Product_release, conf.Feature_set1, conf.Feature_set2)


  var conn connected.Connecteddata
  var message string

  log.Printf("starting for loop")
  for {
     conn.Readconnectvalues()
     log.Printf("5GC Core server:  connected=%s", conn.Connected)

     message = strings.Join(conn.Connected, ",") 
     log.Printf("message to send=" + message)
     client.Client(demotype, message)

     time.Sleep(15 * time.Second)  //every 15 seconds
  }

  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
  log.Printf("%s", "kubesim 5G NR client is exiting")
}
