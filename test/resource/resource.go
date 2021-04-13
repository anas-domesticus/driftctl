package resource

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/dctlcty"
)

type FakeResource struct {
	Id        string `cty:"id"`
	FooBar    string `cty:"foo_bar"`
	BarFoo    string `cty:"bar_foo" computed:"true"`
	Json      string `cty:"json" jsonstring:"true"`
	Type      string
	Tags      map[string]string `cty:"tags"`
	CustomMap map[string]struct {
		Tag string `cty:"tag"`
	} `cty:"custom_map"`
	Slice  []string `cty:"slice"`
	Struct struct {
		Baz string `cty:"baz" computed:"true"`
		Bar string `cty:"bar"`
	} `cty:"struct"`
	StructSlice []struct {
		String string   `cty:"string" computed:"true"`
		Array  []string `cty:"array" computed:"true"`
	} `cty:"struct_slice"`
	CtyVal *cty.Value `diff:"-"`
}

func (d FakeResource) TerraformId() string {
	return d.Id
}

func (d FakeResource) TerraformType() string {
	if d.Type != "" {
		dctlcty.SetMetadata(d.Type, FakeResourceTags, FakeResourceNormalizer)
		return d.Type
	}
	return "FakeResource"
}

func (r FakeResource) CtyValue() *cty.Value {
	return r.CtyVal
}

type FakeResourceStringer struct {
	Id     string     `cty:"id"`
	Name   string     `cty:"name"`
	CtyVal *cty.Value `diff:"-"`
}

func (d *FakeResourceStringer) TerraformId() string {
	return d.Id
}

func (d *FakeResourceStringer) TerraformType() string {
	return "FakeResourceStringer"
}

func (r *FakeResourceStringer) CtyValue() *cty.Value {
	return r.CtyVal
}

func (d *FakeResourceStringer) String() string {
	return fmt.Sprintf("Name: '%s'", d.Name)
}

func InitFakeResourceMetadata() {
	dctlcty.SetMetadata("FakeResource", FakeResourceTags, FakeResourceNormalizer)
}

var FakeResourceTags = map[string]string{
	"bar_foo":             `computed:"true"`,
	"json":                `jsonstring:"true"`,
	"struct.baz":          `computed:"true"`,
	"struct_slice.string": `computed:"true"`,
	"struct_slice.array":  `computed:"true"`,
}

func FakeResourceNormalizer(val *dctlcty.CtyAttributes) {
}
