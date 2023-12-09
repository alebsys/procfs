// Copyright 2022 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package procfs

import "fmt"

type ProcNetTCP struct {
	PID    int
	NetTCP []*netIPSocketLine
}

func (p Proc) NetTCP(pid int) (ProcNetTCP, error) {
	path := fmt.Sprintf("/proc/%d/net/tcp", pid)
	NetTCP, err := newNetTCP(path)
	return ProcNetTCP{NetTCP: NetTCP, PID: pid}, err
}
