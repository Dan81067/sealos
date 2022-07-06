// Copyright © 2022 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package care

import (
	"github.com/labring/lvscare/pkg/glog"
)

//createVsAndRs is
func (care *LvsCare) createVsAndRs() {
	//ip, _ := utils.SplitServer(vs)
	if care.lvs == nil {
		care.lvs = BuildLvscare()
	}
	var errs []string
	isAvailable := care.lvs.IsVirtualServerAvailable(care.VirtualServer)
	if !isAvailable {
		err := care.lvs.CreateVirtualServer(care.VirtualServer, true)
		//virtual server is exists
		if err != nil {
			//can't return
			errs = append(errs, err.Error())
		}
	}
	for _, r := range care.RealServer {
		err := care.lvs.CreateRealServer(r, true)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) != 0 {
		glog.Errorf("createVsAndRs error: %v", errs)
	}
}