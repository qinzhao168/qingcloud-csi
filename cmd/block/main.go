// +-------------------------------------------------------------------------
// | Copyright (C) 2018 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package main

import (
	"flag"
	"github.com/yunify/qingcloud-csi/pkg/block"
	"os"
)

func init() {
	flag.Set("logtostderr", "true")
}

var (
	endpoint   = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	driverName = flag.String("drivername", "csi-qingcloud", "name of the driver")
	nodeID     = flag.String("nodeid", "", "node id")
	configPath = flag.String("config", "/etc/config/config.yaml", "server config file path")
)

func main() {
	flag.Parse()
	handle()
	os.Exit(0)
}

func handle() {
	block.ConfigFilePath = *configPath
	driver := block.GetBlockDriver()
	driver.Run(*driverName, *nodeID, *endpoint)
}
