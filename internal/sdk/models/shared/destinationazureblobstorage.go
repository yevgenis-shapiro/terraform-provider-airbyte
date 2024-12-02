// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type AzureBlobStorage string

const (
	AzureBlobStorageAzureBlobStorage AzureBlobStorage = "azure-blob-storage"
)

func (e AzureBlobStorage) ToPointer() *AzureBlobStorage {
	return &e
}
func (e *AzureBlobStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure-blob-storage":
		*e = AzureBlobStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AzureBlobStorage: %v", v)
	}
}

type DestinationAzureBlobStorageSchemasFormatType string

const (
	DestinationAzureBlobStorageSchemasFormatTypeJsonl DestinationAzureBlobStorageSchemasFormatType = "JSONL"
)

func (e DestinationAzureBlobStorageSchemasFormatType) ToPointer() *DestinationAzureBlobStorageSchemasFormatType {
	return &e
}
func (e *DestinationAzureBlobStorageSchemasFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = DestinationAzureBlobStorageSchemasFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAzureBlobStorageSchemasFormatType: %v", v)
	}
}

type DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON struct {
	// Add file extensions to the output file.
	FileExtension *bool                                        `default:"false" json:"file_extension"`
	formatType    DestinationAzureBlobStorageSchemasFormatType `const:"JSONL" json:"format_type"`
}

func (d DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON) GetFileExtension() *bool {
	if o == nil {
		return nil
	}
	return o.FileExtension
}

func (o *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON) GetFormatType() DestinationAzureBlobStorageSchemasFormatType {
	return DestinationAzureBlobStorageSchemasFormatTypeJsonl
}

// DestinationAzureBlobStorageNormalizationFlattening - Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
type DestinationAzureBlobStorageNormalizationFlattening string

const (
	DestinationAzureBlobStorageNormalizationFlatteningNoFlattening        DestinationAzureBlobStorageNormalizationFlattening = "No flattening"
	DestinationAzureBlobStorageNormalizationFlatteningRootLevelFlattening DestinationAzureBlobStorageNormalizationFlattening = "Root level flattening"
)

func (e DestinationAzureBlobStorageNormalizationFlattening) ToPointer() *DestinationAzureBlobStorageNormalizationFlattening {
	return &e
}
func (e *DestinationAzureBlobStorageNormalizationFlattening) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No flattening":
		fallthrough
	case "Root level flattening":
		*e = DestinationAzureBlobStorageNormalizationFlattening(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAzureBlobStorageNormalizationFlattening: %v", v)
	}
}

type DestinationAzureBlobStorageFormatType string

const (
	DestinationAzureBlobStorageFormatTypeCsv DestinationAzureBlobStorageFormatType = "CSV"
)

func (e DestinationAzureBlobStorageFormatType) ToPointer() *DestinationAzureBlobStorageFormatType {
	return &e
}
func (e *DestinationAzureBlobStorageFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CSV":
		*e = DestinationAzureBlobStorageFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAzureBlobStorageFormatType: %v", v)
	}
}

type DestinationAzureBlobStorageCSVCommaSeparatedValues struct {
	// Add file extensions to the output file.
	FileExtension *bool `default:"false" json:"file_extension"`
	// Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
	Flattening *DestinationAzureBlobStorageNormalizationFlattening `default:"No flattening" json:"flattening"`
	formatType DestinationAzureBlobStorageFormatType               `const:"CSV" json:"format_type"`
}

func (d DestinationAzureBlobStorageCSVCommaSeparatedValues) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAzureBlobStorageCSVCommaSeparatedValues) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAzureBlobStorageCSVCommaSeparatedValues) GetFileExtension() *bool {
	if o == nil {
		return nil
	}
	return o.FileExtension
}

func (o *DestinationAzureBlobStorageCSVCommaSeparatedValues) GetFlattening() *DestinationAzureBlobStorageNormalizationFlattening {
	if o == nil {
		return nil
	}
	return o.Flattening
}

func (o *DestinationAzureBlobStorageCSVCommaSeparatedValues) GetFormatType() DestinationAzureBlobStorageFormatType {
	return DestinationAzureBlobStorageFormatTypeCsv
}

type DestinationAzureBlobStorageOutputFormatType string

const (
	DestinationAzureBlobStorageOutputFormatTypeDestinationAzureBlobStorageCSVCommaSeparatedValues       DestinationAzureBlobStorageOutputFormatType = "destination-azure-blob-storage_CSV: Comma-Separated Values"
	DestinationAzureBlobStorageOutputFormatTypeDestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON DestinationAzureBlobStorageOutputFormatType = "destination-azure-blob-storage_JSON Lines: newline-delimited JSON"
)

// DestinationAzureBlobStorageOutputFormat - Output data format
type DestinationAzureBlobStorageOutputFormat struct {
	DestinationAzureBlobStorageCSVCommaSeparatedValues       *DestinationAzureBlobStorageCSVCommaSeparatedValues
	DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON

	Type DestinationAzureBlobStorageOutputFormatType
}

func CreateDestinationAzureBlobStorageOutputFormatDestinationAzureBlobStorageCSVCommaSeparatedValues(destinationAzureBlobStorageCSVCommaSeparatedValues DestinationAzureBlobStorageCSVCommaSeparatedValues) DestinationAzureBlobStorageOutputFormat {
	typ := DestinationAzureBlobStorageOutputFormatTypeDestinationAzureBlobStorageCSVCommaSeparatedValues

	return DestinationAzureBlobStorageOutputFormat{
		DestinationAzureBlobStorageCSVCommaSeparatedValues: &destinationAzureBlobStorageCSVCommaSeparatedValues,
		Type: typ,
	}
}

func CreateDestinationAzureBlobStorageOutputFormatDestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON(destinationAzureBlobStorageJSONLinesNewlineDelimitedJSON DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON) DestinationAzureBlobStorageOutputFormat {
	typ := DestinationAzureBlobStorageOutputFormatTypeDestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON

	return DestinationAzureBlobStorageOutputFormat{
		DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON: &destinationAzureBlobStorageJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func (u *DestinationAzureBlobStorageOutputFormat) UnmarshalJSON(data []byte) error {

	var destinationAzureBlobStorageJSONLinesNewlineDelimitedJSON DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON = DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON{}
	if err := utils.UnmarshalJSON(data, &destinationAzureBlobStorageJSONLinesNewlineDelimitedJSON, "", true, true); err == nil {
		u.DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON = &destinationAzureBlobStorageJSONLinesNewlineDelimitedJSON
		u.Type = DestinationAzureBlobStorageOutputFormatTypeDestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON
		return nil
	}

	var destinationAzureBlobStorageCSVCommaSeparatedValues DestinationAzureBlobStorageCSVCommaSeparatedValues = DestinationAzureBlobStorageCSVCommaSeparatedValues{}
	if err := utils.UnmarshalJSON(data, &destinationAzureBlobStorageCSVCommaSeparatedValues, "", true, true); err == nil {
		u.DestinationAzureBlobStorageCSVCommaSeparatedValues = &destinationAzureBlobStorageCSVCommaSeparatedValues
		u.Type = DestinationAzureBlobStorageOutputFormatTypeDestinationAzureBlobStorageCSVCommaSeparatedValues
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationAzureBlobStorageOutputFormat", string(data))
}

func (u DestinationAzureBlobStorageOutputFormat) MarshalJSON() ([]byte, error) {
	if u.DestinationAzureBlobStorageCSVCommaSeparatedValues != nil {
		return utils.MarshalJSON(u.DestinationAzureBlobStorageCSVCommaSeparatedValues, "", true)
	}

	if u.DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON != nil {
		return utils.MarshalJSON(u.DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationAzureBlobStorageOutputFormat: all fields are null")
}

type DestinationAzureBlobStorage struct {
	// The Azure blob storage account key.
	AzureBlobStorageAccountKey string `json:"azure_blob_storage_account_key"`
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The name of the Azure blob storage container. If not exists - will be created automatically. May be empty, then will be created automatically airbytecontainer+timestamp
	AzureBlobStorageContainerName *string `json:"azure_blob_storage_container_name,omitempty"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpointDomainName *string `default:"blob.core.windows.net" json:"azure_blob_storage_endpoint_domain_name"`
	// The amount of megabytes to buffer for the output stream to Azure. This will impact memory footprint on workers, but may need adjustment for performance and appropriate block size in Azure.
	AzureBlobStorageOutputBufferSize *int64 `default:"5" json:"azure_blob_storage_output_buffer_size"`
	// The amount of megabytes after which the connector should spill the records in a new blob object. Make sure to configure size greater than individual records. Enter 0 if not applicable
	AzureBlobStorageSpillSize *int64           `default:"500" json:"azure_blob_storage_spill_size"`
	destinationType           AzureBlobStorage `const:"azure-blob-storage" json:"destinationType"`
	// Output data format
	Format DestinationAzureBlobStorageOutputFormat `json:"format"`
}

func (d DestinationAzureBlobStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAzureBlobStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAzureBlobStorage) GetAzureBlobStorageAccountKey() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountKey
}

func (o *DestinationAzureBlobStorage) GetAzureBlobStorageAccountName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountName
}

func (o *DestinationAzureBlobStorage) GetAzureBlobStorageContainerName() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageContainerName
}

func (o *DestinationAzureBlobStorage) GetAzureBlobStorageEndpointDomainName() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageEndpointDomainName
}

func (o *DestinationAzureBlobStorage) GetAzureBlobStorageOutputBufferSize() *int64 {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageOutputBufferSize
}

func (o *DestinationAzureBlobStorage) GetAzureBlobStorageSpillSize() *int64 {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageSpillSize
}

func (o *DestinationAzureBlobStorage) GetDestinationType() AzureBlobStorage {
	return AzureBlobStorageAzureBlobStorage
}

func (o *DestinationAzureBlobStorage) GetFormat() DestinationAzureBlobStorageOutputFormat {
	if o == nil {
		return DestinationAzureBlobStorageOutputFormat{}
	}
	return o.Format
}