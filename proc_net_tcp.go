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
	NetTCP []*netIPSocketLine
}

type ProcNetTCPSummary struct {
	NetTCPSummary *NetTCPSummary
}

// NetTCP returns the IPv4 kernel/networking statistics for TCP datagrams
// read from /proc/<pid>/net/tcp.
func (p Proc) NetTCP() (ProcNetTCP, error) {
	netTCP, err := newNetTCP(p.path("net/tcp"))
	if err != nil {
		return ProcNetTCP{}, err
	}
	return ProcNetTCP{NetTCP: netTCP}, err
}

// NetTCP6 returns the IPv6 kernel/networking statistics for TCP datagrams
// read from /proc/<pid>/net/tcp6.
func (p Proc) NetTCP6() (ProcNetTCP, error) {
	netTCP, err := newNetTCP(p.path("net/tcp6"))
	if err != nil {
		return ProcNetTCP{}, err
	}
	return ProcNetTCP{NetTCP: netTCP}, err
}

// NetTCPSummary returns already computed statistics like the total queue lengths
// for TCP datagrams read from /proc/<pid>/net/tcp.
func (p Proc) NetTCPSummary() (ProcNetTCPSummary, error) {
	netTCPSummary, err := newNetTCPSummary(p.path("net/tcp"))
	if err != nil {
		return ProcNetTCPSummary{}, err
	}
	return ProcNetTCPSummary{NetTCPSummary: netTCPSummary}, err
}

// NetTCP6Summary returns already computed statistics like the total queue lengths
// for TCP datagrams read from /proc/<pid>/net/tcp6.
func (p Proc) NetTCP6Summary() (ProcNetTCPSummary, error) {
	netTCPSummary, err := newNetTCPSummary(p.path("net/tcp6"))
	if err != nil {
		return ProcNetTCPSummary{}, err
	}
	return ProcNetTCPSummary{NetTCPSummary: netTCPSummary}, err
}
