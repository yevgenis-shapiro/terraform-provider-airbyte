// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

// SourceAmazonSqsUpdateAWSRegion - AWS Region of the SQS Queue
type SourceAmazonSqsUpdateAWSRegion string

const (
	SourceAmazonSqsUpdateAWSRegionAfSouth1     SourceAmazonSqsUpdateAWSRegion = "af-south-1"
	SourceAmazonSqsUpdateAWSRegionApEast1      SourceAmazonSqsUpdateAWSRegion = "ap-east-1"
	SourceAmazonSqsUpdateAWSRegionApNortheast1 SourceAmazonSqsUpdateAWSRegion = "ap-northeast-1"
	SourceAmazonSqsUpdateAWSRegionApNortheast2 SourceAmazonSqsUpdateAWSRegion = "ap-northeast-2"
	SourceAmazonSqsUpdateAWSRegionApNortheast3 SourceAmazonSqsUpdateAWSRegion = "ap-northeast-3"
	SourceAmazonSqsUpdateAWSRegionApSouth1     SourceAmazonSqsUpdateAWSRegion = "ap-south-1"
	SourceAmazonSqsUpdateAWSRegionApSouth2     SourceAmazonSqsUpdateAWSRegion = "ap-south-2"
	SourceAmazonSqsUpdateAWSRegionApSoutheast1 SourceAmazonSqsUpdateAWSRegion = "ap-southeast-1"
	SourceAmazonSqsUpdateAWSRegionApSoutheast2 SourceAmazonSqsUpdateAWSRegion = "ap-southeast-2"
	SourceAmazonSqsUpdateAWSRegionApSoutheast3 SourceAmazonSqsUpdateAWSRegion = "ap-southeast-3"
	SourceAmazonSqsUpdateAWSRegionApSoutheast4 SourceAmazonSqsUpdateAWSRegion = "ap-southeast-4"
	SourceAmazonSqsUpdateAWSRegionCaCentral1   SourceAmazonSqsUpdateAWSRegion = "ca-central-1"
	SourceAmazonSqsUpdateAWSRegionCaWest1      SourceAmazonSqsUpdateAWSRegion = "ca-west-1"
	SourceAmazonSqsUpdateAWSRegionCnNorth1     SourceAmazonSqsUpdateAWSRegion = "cn-north-1"
	SourceAmazonSqsUpdateAWSRegionCnNorthwest1 SourceAmazonSqsUpdateAWSRegion = "cn-northwest-1"
	SourceAmazonSqsUpdateAWSRegionEuCentral1   SourceAmazonSqsUpdateAWSRegion = "eu-central-1"
	SourceAmazonSqsUpdateAWSRegionEuCentral2   SourceAmazonSqsUpdateAWSRegion = "eu-central-2"
	SourceAmazonSqsUpdateAWSRegionEuNorth1     SourceAmazonSqsUpdateAWSRegion = "eu-north-1"
	SourceAmazonSqsUpdateAWSRegionEuSouth1     SourceAmazonSqsUpdateAWSRegion = "eu-south-1"
	SourceAmazonSqsUpdateAWSRegionEuSouth2     SourceAmazonSqsUpdateAWSRegion = "eu-south-2"
	SourceAmazonSqsUpdateAWSRegionEuWest1      SourceAmazonSqsUpdateAWSRegion = "eu-west-1"
	SourceAmazonSqsUpdateAWSRegionEuWest2      SourceAmazonSqsUpdateAWSRegion = "eu-west-2"
	SourceAmazonSqsUpdateAWSRegionEuWest3      SourceAmazonSqsUpdateAWSRegion = "eu-west-3"
	SourceAmazonSqsUpdateAWSRegionIlCentral1   SourceAmazonSqsUpdateAWSRegion = "il-central-1"
	SourceAmazonSqsUpdateAWSRegionMeCentral1   SourceAmazonSqsUpdateAWSRegion = "me-central-1"
	SourceAmazonSqsUpdateAWSRegionMeSouth1     SourceAmazonSqsUpdateAWSRegion = "me-south-1"
	SourceAmazonSqsUpdateAWSRegionSaEast1      SourceAmazonSqsUpdateAWSRegion = "sa-east-1"
	SourceAmazonSqsUpdateAWSRegionUsEast1      SourceAmazonSqsUpdateAWSRegion = "us-east-1"
	SourceAmazonSqsUpdateAWSRegionUsEast2      SourceAmazonSqsUpdateAWSRegion = "us-east-2"
	SourceAmazonSqsUpdateAWSRegionUsGovEast1   SourceAmazonSqsUpdateAWSRegion = "us-gov-east-1"
	SourceAmazonSqsUpdateAWSRegionUsGovWest1   SourceAmazonSqsUpdateAWSRegion = "us-gov-west-1"
	SourceAmazonSqsUpdateAWSRegionUsWest1      SourceAmazonSqsUpdateAWSRegion = "us-west-1"
	SourceAmazonSqsUpdateAWSRegionUsWest2      SourceAmazonSqsUpdateAWSRegion = "us-west-2"
)

func (e SourceAmazonSqsUpdateAWSRegion) ToPointer() *SourceAmazonSqsUpdateAWSRegion {
	return &e
}
func (e *SourceAmazonSqsUpdateAWSRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-south-2":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ap-southeast-3":
		fallthrough
	case "ap-southeast-4":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "ca-west-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-central-2":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-south-2":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "il-central-1":
		fallthrough
	case "me-central-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		*e = SourceAmazonSqsUpdateAWSRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonSqsUpdateAWSRegion: %v", v)
	}
}

type SourceAmazonSqsUpdate struct {
	// The Access Key ID of the AWS IAM Role to use for pulling messages
	AccessKey *string `json:"access_key,omitempty"`
	// Comma separated list of Mesage Attribute names to return
	AttributesToReturn *string `json:"attributes_to_return,omitempty"`
	// If Enabled, messages will be deleted from the SQS Queue after being read. If Disabled, messages are left in the queue and can be read more than once. WARNING: Enabling this option can result in data loss in cases of failure, use with caution, see documentation for more detail.
	DeleteMessages *bool `default:"false" json:"delete_messages"`
	// Max amount of messages to get in one batch (10 max)
	MaxBatchSize *int64 `json:"max_batch_size,omitempty"`
	// Max amount of time in seconds to wait for messages in a single poll (20 max)
	MaxWaitTime *int64 `json:"max_wait_time,omitempty"`
	// URL of the SQS Queue
	QueueURL string `json:"queue_url"`
	// AWS Region of the SQS Queue
	Region SourceAmazonSqsUpdateAWSRegion `json:"region"`
	// The Secret Key of the AWS IAM Role to use for pulling messages
	SecretKey *string `json:"secret_key,omitempty"`
	// Modify the Visibility Timeout of the individual message from the Queue's default (seconds).
	VisibilityTimeout *int64 `json:"visibility_timeout,omitempty"`
}

func (s SourceAmazonSqsUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAmazonSqsUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAmazonSqsUpdate) GetAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AccessKey
}

func (o *SourceAmazonSqsUpdate) GetAttributesToReturn() *string {
	if o == nil {
		return nil
	}
	return o.AttributesToReturn
}

func (o *SourceAmazonSqsUpdate) GetDeleteMessages() *bool {
	if o == nil {
		return nil
	}
	return o.DeleteMessages
}

func (o *SourceAmazonSqsUpdate) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *SourceAmazonSqsUpdate) GetMaxWaitTime() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxWaitTime
}

func (o *SourceAmazonSqsUpdate) GetQueueURL() string {
	if o == nil {
		return ""
	}
	return o.QueueURL
}

func (o *SourceAmazonSqsUpdate) GetRegion() SourceAmazonSqsUpdateAWSRegion {
	if o == nil {
		return SourceAmazonSqsUpdateAWSRegion("")
	}
	return o.Region
}

func (o *SourceAmazonSqsUpdate) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *SourceAmazonSqsUpdate) GetVisibilityTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.VisibilityTimeout
}
