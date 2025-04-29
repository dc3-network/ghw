//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

import (
	"fmt"

	"github.com/dc3-network/ghw/v99/pkg/context"

	"github.com/dc3-network/ghw/v99/pkg/accelerator"
	"github.com/dc3-network/ghw/v99/pkg/baseboard"
	"github.com/dc3-network/ghw/v99/pkg/bios"
	"github.com/dc3-network/ghw/v99/pkg/block"
	"github.com/dc3-network/ghw/v99/pkg/chassis"
	"github.com/dc3-network/ghw/v99/pkg/cpu"
	"github.com/dc3-network/ghw/v99/pkg/gpu"
	"github.com/dc3-network/ghw/v99/pkg/marshal"
	"github.com/dc3-network/ghw/v99/pkg/memory"
	"github.com/dc3-network/ghw/v99/pkg/net"
	"github.com/dc3-network/ghw/v99/pkg/pci"
	"github.com/dc3-network/ghw/v99/pkg/product"
	"github.com/dc3-network/ghw/v99/pkg/topology"
)

// HostInfo is a wrapper struct containing information about the host system's
// memory, block storage, CPU, etc
type HostInfo struct {
	ctx         *context.Context
	Memory      *memory.Info      `json:"memory,omitempty"`
	Block       *block.Info       `json:"block,omitempty"`
	CPU         *cpu.Info         `json:"cpu,omitempty"`
	Topology    *topology.Info    `json:"topology,omitempty"`
	Network     *net.Info         `json:"network,omitempty"`
	GPU         *gpu.Info         `json:"gpu,omitempty"`
	Accelerator *accelerator.Info `json:"accelerator,omitempty"`
	Chassis     *chassis.Info     `json:"chassis,omitempty"`
	BIOS        *bios.Info        `json:"bios,omitempty"`
	Baseboard   *baseboard.Info   `json:"baseboard,omitempty"`
	Product     *product.Info     `json:"product,omitempty"`
	PCI         *pci.Info         `json:"pci,omitempty"`
}

// Host returns a pointer to a HostInfo struct that contains fields with
// information about the host system's CPU, memory, network devices, etc
func Host(opts ...*WithOption) (*HostInfo, error) {
	ctx := context.New(opts...)

	memInfo, err := memory.New(opts...)
	if err != nil {
		return nil, err
	}
	blockInfo, err := block.New(opts...)
	if err != nil {
		return nil, err
	}
	cpuInfo, err := cpu.New(opts...)
	if err != nil {
		return nil, err
	}
	topologyInfo, err := topology.New(opts...)
	if err != nil {
		return nil, err
	}
	netInfo, err := net.New(opts...)
	if err != nil {
		return nil, err
	}
	gpuInfo, err := gpu.New(opts...)
	if err != nil {
		return nil, err
	}
	acceleratorInfo, err := accelerator.New(opts...)
	if err != nil {
		return nil, err
	}
	chassisInfo, err := chassis.New(opts...)
	if err != nil {
		return nil, err
	}
	biosInfo, err := bios.New(opts...)
	if err != nil {
		return nil, err
	}
	baseboardInfo, err := baseboard.New(opts...)
	if err != nil {
		return nil, err
	}
	productInfo, err := product.New(opts...)
	if err != nil {
		return nil, err
	}
	pciInfo, err := pci.New(opts...)
	if err != nil {
		return nil, err
	}
	return &HostInfo{
		ctx:         ctx,
		CPU:         cpuInfo,
		Memory:      memInfo,
		Block:       blockInfo,
		Topology:    topologyInfo,
		Network:     netInfo,
		GPU:         gpuInfo,
		Accelerator: acceleratorInfo,
		Chassis:     chassisInfo,
		BIOS:        biosInfo,
		Baseboard:   baseboardInfo,
		Product:     productInfo,
		PCI:         pciInfo,
	}, nil
}

// String returns a newline-separated output of the HostInfo's component
// structs' String-ified output
func (info *HostInfo) String() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		info.Block.String(),
		info.CPU.String(),
		info.GPU.String(),
		info.Accelerator.String(),
		info.Memory.String(),
		info.Network.String(),
		info.Topology.String(),
		info.Chassis.String(),
		info.BIOS.String(),
		info.Baseboard.String(),
		info.Product.String(),
		info.PCI.String(),
	)
}

// YAMLString returns a string with the host information formatted as YAML
// under a top-level "host:" key
func (i *HostInfo) YAMLString() string {
	return marshal.SafeYAML(i.ctx, i)
}

// JSONString returns a string with the host information formatted as JSON
// under a top-level "host:" key
func (i *HostInfo) JSONString(indent bool) string {
	return marshal.SafeJSON(i.ctx, i, indent)
}
