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

type ProcNetTCP struct {
	// The process ID.
	PID    int
	NetTCP []*netIPSocketLine
}

func (p Proc) NetTCP() (ProcNetTCP, error) {
	NetTCP, err := newNetTCP(p.path("net/tcp"))
	if err != nil {
		return ProcNetTCP{PID: p.PID}, err
	}
	return ProcNetTCP{NetTCP: NetTCP, PID: p.PID}, err
}
func (p Proc) NetTCP6() (ProcNetTCP, error) {
	NetTCP, err := newNetTCP(p.path("net/tcp6"))
	if err != nil {
		return ProcNetTCP{PID: p.PID}, err
	}
	return ProcNetTCP{NetTCP: NetTCP, PID: p.PID}, err
}
