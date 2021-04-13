// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/dctlcty"
)

const AwsCloudfrontDistributionResourceType = "aws_cloudfront_distribution"

type AwsCloudfrontDistribution struct {
	Aliases                     *[]string         `cty:"aliases"`
	Arn                         *string           `cty:"arn" computed:"true"`
	CallerReference             *string           `cty:"caller_reference" computed:"true"`
	Comment                     *string           `cty:"comment"`
	DefaultRootObject           *string           `cty:"default_root_object"`
	DomainName                  *string           `cty:"domain_name" computed:"true"`
	Enabled                     *bool             `cty:"enabled"`
	Etag                        *string           `cty:"etag" computed:"true" diff:"-"`
	HostedZoneId                *string           `cty:"hosted_zone_id" computed:"true"`
	HttpVersion                 *string           `cty:"http_version"`
	Id                          string            `cty:"id" computed:"true"`
	InProgressValidationBatches *int              `cty:"in_progress_validation_batches" computed:"true"`
	IsIpv6Enabled               *bool             `cty:"is_ipv6_enabled"`
	LastModifiedTime            *string           `cty:"last_modified_time" computed:"true" diff:"-"`
	PriceClass                  *string           `cty:"price_class"`
	RetainOnDelete              *bool             `cty:"retain_on_delete" diff:"-"`
	Status                      *string           `cty:"status" computed:"true" diff:"-"`
	Tags                        map[string]string `cty:"tags"`
	TrustedSigners              *[]struct {
		Enabled *bool `cty:"enabled" computed:"true"`
		Items   *[]struct {
			AwsAccountNumber *string   `cty:"aws_account_number" computed:"true"`
			KeyPairIds       *[]string `cty:"key_pair_ids" computed:"true"`
		} `cty:"items" computed:"true"`
	} `cty:"trusted_signers" computed:"true"`
	WaitForDeployment   *bool   `cty:"wait_for_deployment" diff:"-"`
	WebAclId            *string `cty:"web_acl_id"`
	CustomErrorResponse *[]struct {
		ErrorCachingMinTtl *int    `cty:"error_caching_min_ttl"`
		ErrorCode          *int    `cty:"error_code"`
		ResponseCode       *int    `cty:"response_code"`
		ResponsePagePath   *string `cty:"response_page_path"`
	} `cty:"custom_error_response"`
	DefaultCacheBehavior *[]struct {
		AllowedMethods         *[]string `cty:"allowed_methods"`
		CachedMethods          *[]string `cty:"cached_methods"`
		Compress               *bool     `cty:"compress"`
		DefaultTtl             *int      `cty:"default_ttl"`
		FieldLevelEncryptionId *string   `cty:"field_level_encryption_id"`
		MaxTtl                 *int      `cty:"max_ttl"`
		MinTtl                 *int      `cty:"min_ttl"`
		SmoothStreaming        *bool     `cty:"smooth_streaming"`
		TargetOriginId         *string   `cty:"target_origin_id"`
		TrustedSigners         *[]string `cty:"trusted_signers"`
		ViewerProtocolPolicy   *string   `cty:"viewer_protocol_policy"`
		ForwardedValues        *[]struct {
			Headers              *[]string `cty:"headers"`
			QueryString          *bool     `cty:"query_string"`
			QueryStringCacheKeys *[]string `cty:"query_string_cache_keys"`
			Cookies              *[]struct {
				Forward          *string   `cty:"forward"`
				WhitelistedNames *[]string `cty:"whitelisted_names"`
			} `cty:"cookies"`
		} `cty:"forwarded_values"`
		LambdaFunctionAssociation *[]struct {
			EventType   *string `cty:"event_type"`
			IncludeBody *bool   `cty:"include_body"`
			LambdaArn   *string `cty:"lambda_arn"`
		} `cty:"lambda_function_association"`
	} `cty:"default_cache_behavior"`
	LoggingConfig *[]struct {
		Bucket         *string `cty:"bucket"`
		IncludeCookies *bool   `cty:"include_cookies"`
		Prefix         *string `cty:"prefix"`
	} `cty:"logging_config"`
	OrderedCacheBehavior *[]struct {
		AllowedMethods         *[]string `cty:"allowed_methods"`
		CachedMethods          *[]string `cty:"cached_methods"`
		Compress               *bool     `cty:"compress"`
		DefaultTtl             *int      `cty:"default_ttl"`
		FieldLevelEncryptionId *string   `cty:"field_level_encryption_id"`
		MaxTtl                 *int      `cty:"max_ttl"`
		MinTtl                 *int      `cty:"min_ttl"`
		PathPattern            *string   `cty:"path_pattern"`
		SmoothStreaming        *bool     `cty:"smooth_streaming"`
		TargetOriginId         *string   `cty:"target_origin_id"`
		TrustedSigners         *[]string `cty:"trusted_signers"`
		ViewerProtocolPolicy   *string   `cty:"viewer_protocol_policy"`
		ForwardedValues        *[]struct {
			Headers              *[]string `cty:"headers"`
			QueryString          *bool     `cty:"query_string"`
			QueryStringCacheKeys *[]string `cty:"query_string_cache_keys"`
			Cookies              *[]struct {
				Forward          *string   `cty:"forward"`
				WhitelistedNames *[]string `cty:"whitelisted_names"`
			} `cty:"cookies"`
		} `cty:"forwarded_values"`
		LambdaFunctionAssociation *[]struct {
			EventType   *string `cty:"event_type"`
			IncludeBody *bool   `cty:"include_body"`
			LambdaArn   *string `cty:"lambda_arn"`
		} `cty:"lambda_function_association"`
	} `cty:"ordered_cache_behavior"`
	Origin *[]struct {
		DomainName   *string `cty:"domain_name"`
		OriginId     *string `cty:"origin_id"`
		OriginPath   *string `cty:"origin_path"`
		CustomHeader *[]struct {
			Name  *string `cty:"name"`
			Value *string `cty:"value"`
		} `cty:"custom_header"`
		CustomOriginConfig *[]struct {
			HttpPort               *int      `cty:"http_port"`
			HttpsPort              *int      `cty:"https_port"`
			OriginKeepaliveTimeout *int      `cty:"origin_keepalive_timeout"`
			OriginProtocolPolicy   *string   `cty:"origin_protocol_policy"`
			OriginReadTimeout      *int      `cty:"origin_read_timeout"`
			OriginSslProtocols     *[]string `cty:"origin_ssl_protocols"`
		} `cty:"custom_origin_config"`
		S3OriginConfig *[]struct {
			OriginAccessIdentity *string `cty:"origin_access_identity"`
		} `cty:"s3_origin_config"`
	} `cty:"origin"`
	OriginGroup *[]struct {
		OriginId         *string `cty:"origin_id"`
		FailoverCriteria *[]struct {
			StatusCodes *[]int `cty:"status_codes"`
		} `cty:"failover_criteria"`
		Member *[]struct {
			OriginId *string `cty:"origin_id"`
		} `cty:"member"`
	} `cty:"origin_group"`
	Restrictions *[]struct {
		GeoRestriction *[]struct {
			Locations       *[]string `cty:"locations"`
			RestrictionType *string   `cty:"restriction_type"`
		} `cty:"geo_restriction"`
	} `cty:"restrictions"`
	ViewerCertificate *[]struct {
		AcmCertificateArn            *string `cty:"acm_certificate_arn"`
		CloudfrontDefaultCertificate *bool   `cty:"cloudfront_default_certificate"`
		IamCertificateId             *string `cty:"iam_certificate_id"`
		MinimumProtocolVersion       *string `cty:"minimum_protocol_version"`
		SslSupportMethod             *string `cty:"ssl_support_method"`
	} `cty:"viewer_certificate"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsCloudfrontDistribution) TerraformId() string {
	return r.Id
}

func (r *AwsCloudfrontDistribution) TerraformType() string {
	return AwsCloudfrontDistributionResourceType
}

func (r *AwsCloudfrontDistribution) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsCloudfrontDistributionMetaData() {
	dctlcty.SetMetadata(AwsCloudfrontDistributionResourceType, AwsCloudfrontDistributionTags, AwsCloudfrontDistributionNormalizer)
}

var AwsCloudfrontDistributionTags = map[string]string{
	"arn":                            "computed:\"true\"",
	"caller_reference":               "computed:\"true\"",
	"domain_name":                    "computed:\"true\"",
	"etag":                           "computed:\"true\"",
	"hosted_zone_id":                 "computed:\"true\"",
	"id":                             "computed:\"true\"",
	"in_progress_validation_batches": "computed:\"true\"",
	"last_modified_time":             "computed:\"true\"",
	"status":                         "computed:\"true\"",
	"trusted_signers":                "computed:\"true\"",
	"trusted_signers.enabled":        "computed:\"true\"",
	"trusted_signers.items":          "computed:\"true\"",
	"trusted_signers.items.aws_account_number": "computed:\"true\"",
	"trusted_signers.items.key_pair_ids":       "computed:\"true\"",
}

func AwsCloudfrontDistributionNormalizer(val *dctlcty.CtyAttributes) {
	val.SafeDelete([]string{"etag"})
	val.SafeDelete([]string{"last_modified_time"})
	val.SafeDelete([]string{"retain_on_delete"})
	val.SafeDelete([]string{"status"})
	val.SafeDelete([]string{"wait_for_deployment"})
}
