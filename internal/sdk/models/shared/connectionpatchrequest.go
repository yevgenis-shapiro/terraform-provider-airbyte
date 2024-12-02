// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type ConnectionPatchRequest struct {
	// A list of configured stream options for a connection.
	Configurations *StreamConfigurations   `json:"configurations,omitempty"`
	DataResidency  *GeographyEnumNoDefault `json:"dataResidency,omitempty"`
	// Optional name of the connection
	Name *string `json:"name,omitempty"`
	// Define the location where the data will be stored in the destination
	NamespaceDefinition *NamespaceDefinitionEnumNoDefault `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'custom_format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `default:"null" json:"namespaceFormat"`
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior *NonBreakingSchemaUpdatesBehaviorEnumNoDefault `json:"nonBreakingSchemaUpdatesBehavior,omitempty"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte_” causes “projects” => “airbyte_projects”).
	Prefix *string `json:"prefix,omitempty"`
	// schedule for when the the connection should run, per the schedule type
	Schedule *AirbyteAPIConnectionSchedule `json:"schedule,omitempty"`
	Status   *ConnectionStatusEnum         `json:"status,omitempty"`
}

func (c ConnectionPatchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionPatchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionPatchRequest) GetConfigurations() *StreamConfigurations {
	if o == nil {
		return nil
	}
	return o.Configurations
}

func (o *ConnectionPatchRequest) GetDataResidency() *GeographyEnumNoDefault {
	if o == nil {
		return nil
	}
	return o.DataResidency
}

func (o *ConnectionPatchRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionPatchRequest) GetNamespaceDefinition() *NamespaceDefinitionEnumNoDefault {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *ConnectionPatchRequest) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *ConnectionPatchRequest) GetNonBreakingSchemaUpdatesBehavior() *NonBreakingSchemaUpdatesBehaviorEnumNoDefault {
	if o == nil {
		return nil
	}
	return o.NonBreakingSchemaUpdatesBehavior
}

func (o *ConnectionPatchRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *ConnectionPatchRequest) GetSchedule() *AirbyteAPIConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *ConnectionPatchRequest) GetStatus() *ConnectionStatusEnum {
	if o == nil {
		return nil
	}
	return o.Status
}
