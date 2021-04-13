// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/dctlcty"
)

const AwsEipAssociationResourceType = "aws_eip_association"

type AwsEipAssociation struct {
	AllocationId       *string    `cty:"allocation_id" computed:"true"`
	AllowReassociation *bool      `cty:"allow_reassociation"`
	Id                 string     `cty:"id" computed:"true"`
	InstanceId         *string    `cty:"instance_id" computed:"true"`
	NetworkInterfaceId *string    `cty:"network_interface_id" computed:"true"`
	PrivateIpAddress   *string    `cty:"private_ip_address" computed:"true"`
	PublicIp           *string    `cty:"public_ip" computed:"true"`
	CtyVal             *cty.Value `diff:"-"`
}

func (r *AwsEipAssociation) TerraformId() string {
	return r.Id
}

func (r *AwsEipAssociation) TerraformType() string {
	return AwsEipAssociationResourceType
}

func (r *AwsEipAssociation) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsEipAssociationMetaData() {
	dctlcty.SetMetadata(AwsEipAssociationResourceType, AwsEipAssociationTags, AwsEipAssociationNormalizer)
}

var AwsEipAssociationTags = map[string]string{
	"allocation_id":        `computed:"true"`,
	"id":                   `computed:"true"`,
	"instance_id":          `computed:"true"`,
	"network_interface_id": `computed:"true"`,
	"private_ip_address":   `computed:"true"`,
	"public_ip":            `computed:"true"`,
}

func AwsEipAssociationNormalizer(val *dctlcty.CtyAttributes) {
}
