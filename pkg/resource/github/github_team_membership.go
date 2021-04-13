// GENERATED, DO NOT EDIT THIS FILE
package github

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/dctlcty"
)

const GithubTeamMembershipResourceType = "github_team_membership"

type GithubTeamMembership struct {
	Etag     *string    `cty:"etag" computed:"true" diff:"-"`
	Id       string     `cty:"id" computed:"true"`
	Role     *string    `cty:"role"`
	TeamId   *string    `cty:"team_id"`
	Username *string    `cty:"username"`
	CtyVal   *cty.Value `diff:"-"`
}

func (r *GithubTeamMembership) TerraformId() string {
	return r.Id
}

func (r *GithubTeamMembership) TerraformType() string {
	return GithubTeamMembershipResourceType
}

func (r *GithubTeamMembership) CtyValue() *cty.Value {
	return r.CtyVal
}

func initGithubTeamMembershipMetadata() {
	dctlcty.SetMetadata(GithubTeamMembershipResourceType, GithubTeamMembershipTags, GithubTeamMembershipNormalizer)
}

var GithubTeamMembershipTags = map[string]string{}

func GithubTeamMembershipNormalizer(val *dctlcty.CtyAttributes) {
	val.SafeDelete([]string{"etag"})
}
