//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package product

import (
	"github.com/jaypipes/ghw/pkg/context"
	"github.com/jaypipes/ghw/pkg/marshal"
	"github.com/jaypipes/ghw/pkg/option"
	"github.com/jaypipes/ghw/pkg/util"
)

// Info defines product information
type Info struct {
	ctx          *context.Context
	Family       string `json:"family,omitempty"`
	Name         string `json:"name,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	SKU          string `json:"sku,omitempty"`
	Version      string `json:"version,omitempty"`
}

func (i *Info) String() string {
	familyStr := ""
	if i.Family != "" {
		familyStr = " family=" + i.Family
	}
	nameStr := ""
	if i.Name != "" {
		nameStr = " name=" + i.Name
	}
	vendorStr := ""
	if i.Vendor != "" {
		vendorStr = " vendor=" + i.Vendor
	}
	serialStr := ""
	if i.SerialNumber != "" && i.SerialNumber != util.UNKNOWN {
		serialStr = " serial=" + i.SerialNumber
	}
	uuidStr := ""
	if i.UUID != "" && i.UUID != util.UNKNOWN {
		uuidStr = " uuid=" + i.UUID
	}
	skuStr := ""
	if i.SKU != "" {
		skuStr = " sku=" + i.SKU
	}
	versionStr := ""
	if i.Version != "" {
		versionStr = " version=" + i.Version
	}

	return "product" + util.ConcatStrings(
		familyStr,
		nameStr,
		vendorStr,
		serialStr,
		uuidStr,
		skuStr,
		versionStr,
	)
}

// New returns a pointer to a Info struct containing information
// about the host's product
func New(opts ...*option.Option) (*Info, error) {
	ctx := context.New(opts...)
	info := &Info{ctx: ctx}
	if err := ctx.Do(info.load); err != nil {
		return nil, err
	}
	return info, nil
}

// simple private struct used to encapsulate product information in a top-level
// "product" YAML/JSON map/object key
type productPrinter struct {
	Info *Info `json:"product,omitempty"`
}

// YAMLString returns a string with the product information formatted as YAML
// under a top-level "dmi:" key
func (info *Info) YAMLString() string {
	return marshal.SafeYAML(info.ctx, productPrinter{info})
}

// JSONString returns a string with the product information formatted as JSON
// under a top-level "product:" key
func (info *Info) JSONString(indent bool) string {
	return marshal.SafeJSON(info.ctx, productPrinter{info}, indent)
}
