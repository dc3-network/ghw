//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package memory

import (
	"fmt"
	"math"

	"github.com/dc3-network/ghw/v99/pkg/context"
	"github.com/dc3-network/ghw/v99/pkg/marshal"
	"github.com/dc3-network/ghw/v99/pkg/option"
	"github.com/dc3-network/ghw/v99/pkg/unitutil"
	"github.com/dc3-network/ghw/v99/pkg/util"
)

// Module describes a single physical memory module for a host system. Pretty
// much all modern systems contain dual in-line memory modules (DIMMs).
//
// See https://en.wikipedia.org/wiki/DIMM
type Module struct {
	Label        string `json:"label,omitempty"`
	Location     string `json:"location,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	SizeBytes    int64  `json:"size_bytes,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
}

// HugePageAmounts describes huge page info
type HugePageAmounts struct {
	Total   int64 `json:"total,omitempty"`
	Free    int64 `json:"free,omitempty"`
	Surplus int64 `json:"surplus,omitempty"`
	// Note: this field will not be populated for Topology call, since data not present in NUMA folder structure
	Reserved int64 `json:"reserved,omitempty"`
}

// Area describes a set of physical memory on a host system. Non-NUMA systems
// will almost always have a single memory area containing all memory the
// system can use. NUMA systems will have multiple memory areas, one or more
// for each NUMA node/cell in the system.
type Area struct {
	TotalPhysicalBytes int64 `json:"total_physical_bytes,omitempty"`
	TotalUsableBytes   int64 `json:"total_usable_bytes,omitempty"`
	// An array of sizes, in bytes, of memory pages supported in this area
	SupportedPageSizes []uint64 `json:"supported_page_sizes,omitempty"`
	// Default system huge page size, in bytes
	DefaultHugePageSize uint64 `json:"default_huge_page_size,omitempty"`
	// Amount of memory, in bytes, consumed by huge pages of all sizes
	TotalHugePageBytes int64 `json:"total_huge_page_bytes,omitempty"`
	// Huge page info by size
	HugePageAmountsBySize map[uint64]*HugePageAmounts `json:"huge_page_amounts_by_size,omitempty"`
	Modules               []*Module                   `json:"modules,omitempty"`
}

// String returns a short string with a summary of information for this memory
// area
func (a *Area) String() string {
	tpbs := util.UNKNOWN
	if a.TotalPhysicalBytes > 0 {
		tpb := a.TotalPhysicalBytes
		unit, unitStr := unitutil.AmountString(tpb)
		tpb = int64(math.Ceil(float64(a.TotalPhysicalBytes) / float64(unit)))
		tpbs = fmt.Sprintf("%d%s", tpb, unitStr)
	}
	tubs := util.UNKNOWN
	if a.TotalUsableBytes > 0 {
		tub := a.TotalUsableBytes
		unit, unitStr := unitutil.AmountString(tub)
		tub = int64(math.Ceil(float64(a.TotalUsableBytes) / float64(unit)))
		tubs = fmt.Sprintf("%d%s", tub, unitStr)
	}
	return fmt.Sprintf("memory (%s physical, %s usable)", tpbs, tubs)
}

// Info contains information about the memory on a host system.
type Info struct {
	ctx *context.Context
	Area
}

// New returns an Info struct that describes the memory on a host system.
func New(opts ...*option.Option) (*Info, error) {
	ctx := context.New(opts...)
	info := &Info{ctx: ctx}
	if err := ctx.Do(info.load); err != nil {
		return nil, err
	}
	return info, nil
}

// String returns a short string with a summary of memory information
func (i *Info) String() string {
	return i.Area.String()
}

// simple private struct used to encapsulate memory information in a top-level
// "memory" YAML/JSON map/object key
type memoryPrinter struct {
	Info *Info `json:"memory,omitempty"`
}

// YAMLString returns a string with the memory information formatted as YAML
// under a top-level "memory:" key
func (i *Info) YAMLString() string {
	return marshal.SafeYAML(i.ctx, memoryPrinter{i})
}

// JSONString returns a string with the memory information formatted as JSON
// under a top-level "memory:" key
func (i *Info) JSONString(indent bool) string {
	return marshal.SafeJSON(i.ctx, memoryPrinter{i}, indent)
}
