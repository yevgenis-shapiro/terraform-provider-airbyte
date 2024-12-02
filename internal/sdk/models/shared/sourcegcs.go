// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceGcsGcs string

const (
	SourceGcsGcsGcs SourceGcsGcs = "gcs"
)

func (e SourceGcsGcs) ToPointer() *SourceGcsGcs {
	return &e
}
func (e *SourceGcsGcs) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gcs":
		*e = SourceGcsGcs(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsGcs: %v", v)
	}
}

type SourceGcsSchemasStreamsFormatFormat6Filetype string

const (
	SourceGcsSchemasStreamsFormatFormat6FiletypeExcel SourceGcsSchemasStreamsFormatFormat6Filetype = "excel"
)

func (e SourceGcsSchemasStreamsFormatFormat6Filetype) ToPointer() *SourceGcsSchemasStreamsFormatFormat6Filetype {
	return &e
}
func (e *SourceGcsSchemasStreamsFormatFormat6Filetype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "excel":
		*e = SourceGcsSchemasStreamsFormatFormat6Filetype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasStreamsFormatFormat6Filetype: %v", v)
	}
}

type SourceGcsExcelFormat struct {
	filetype *SourceGcsSchemasStreamsFormatFormat6Filetype `const:"excel" json:"filetype"`
}

func (s SourceGcsExcelFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsExcelFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsExcelFormat) GetFiletype() *SourceGcsSchemasStreamsFormatFormat6Filetype {
	return SourceGcsSchemasStreamsFormatFormat6FiletypeExcel.ToPointer()
}

type SourceGcsSchemasStreamsFormatFormatFiletype string

const (
	SourceGcsSchemasStreamsFormatFormatFiletypeUnstructured SourceGcsSchemasStreamsFormatFormatFiletype = "unstructured"
)

func (e SourceGcsSchemasStreamsFormatFormatFiletype) ToPointer() *SourceGcsSchemasStreamsFormatFormatFiletype {
	return &e
}
func (e *SourceGcsSchemasStreamsFormatFormatFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unstructured":
		*e = SourceGcsSchemasStreamsFormatFormatFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasStreamsFormatFormatFiletype: %v", v)
	}
}

type SourceGcsSchemasMode string

const (
	SourceGcsSchemasModeAPI SourceGcsSchemasMode = "api"
)

func (e SourceGcsSchemasMode) ToPointer() *SourceGcsSchemasMode {
	return &e
}
func (e *SourceGcsSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api":
		*e = SourceGcsSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasMode: %v", v)
	}
}

type SourceGcsAPIParameterConfigModel struct {
	// The name of the unstructured API parameter to use
	Name string `json:"name"`
	// The value of the parameter
	Value string `json:"value"`
}

func (o *SourceGcsAPIParameterConfigModel) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGcsAPIParameterConfigModel) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// SourceGcsViaAPI - Process files via an API, using the `hi_res` mode. This option is useful for increased performance and accuracy, but requires an API key and a hosted instance of unstructured.
type SourceGcsViaAPI struct {
	// The API key to use matching the environment
	APIKey *string `default:"" json:"api_key"`
	// The URL of the unstructured API to use
	APIURL *string               `default:"https://api.unstructured.io" json:"api_url"`
	mode   *SourceGcsSchemasMode `const:"api" json:"mode"`
	// List of parameters send to the API
	Parameters []SourceGcsAPIParameterConfigModel `json:"parameters,omitempty"`
}

func (s SourceGcsViaAPI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsViaAPI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsViaAPI) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *SourceGcsViaAPI) GetAPIURL() *string {
	if o == nil {
		return nil
	}
	return o.APIURL
}

func (o *SourceGcsViaAPI) GetMode() *SourceGcsSchemasMode {
	return SourceGcsSchemasModeAPI.ToPointer()
}

func (o *SourceGcsViaAPI) GetParameters() []SourceGcsAPIParameterConfigModel {
	if o == nil {
		return nil
	}
	return o.Parameters
}

type SourceGcsMode string

const (
	SourceGcsModeLocal SourceGcsMode = "local"
)

func (e SourceGcsMode) ToPointer() *SourceGcsMode {
	return &e
}
func (e *SourceGcsMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "local":
		*e = SourceGcsMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsMode: %v", v)
	}
}

// SourceGcsLocal - Process files locally, supporting `fast` and `ocr` modes. This is the default option.
type SourceGcsLocal struct {
	mode *SourceGcsMode `const:"local" json:"mode"`
}

func (s SourceGcsLocal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsLocal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsLocal) GetMode() *SourceGcsMode {
	return SourceGcsModeLocal.ToPointer()
}

type SourceGcsProcessingType string

const (
	SourceGcsProcessingTypeSourceGcsLocal  SourceGcsProcessingType = "source-gcs_Local"
	SourceGcsProcessingTypeSourceGcsViaAPI SourceGcsProcessingType = "source-gcs_via API"
)

// SourceGcsProcessing - Processing configuration
type SourceGcsProcessing struct {
	SourceGcsLocal  *SourceGcsLocal
	SourceGcsViaAPI *SourceGcsViaAPI

	Type SourceGcsProcessingType
}

func CreateSourceGcsProcessingSourceGcsLocal(sourceGcsLocal SourceGcsLocal) SourceGcsProcessing {
	typ := SourceGcsProcessingTypeSourceGcsLocal

	return SourceGcsProcessing{
		SourceGcsLocal: &sourceGcsLocal,
		Type:           typ,
	}
}

func CreateSourceGcsProcessingSourceGcsViaAPI(sourceGcsViaAPI SourceGcsViaAPI) SourceGcsProcessing {
	typ := SourceGcsProcessingTypeSourceGcsViaAPI

	return SourceGcsProcessing{
		SourceGcsViaAPI: &sourceGcsViaAPI,
		Type:            typ,
	}
}

func (u *SourceGcsProcessing) UnmarshalJSON(data []byte) error {

	var sourceGcsLocal SourceGcsLocal = SourceGcsLocal{}
	if err := utils.UnmarshalJSON(data, &sourceGcsLocal, "", true, true); err == nil {
		u.SourceGcsLocal = &sourceGcsLocal
		u.Type = SourceGcsProcessingTypeSourceGcsLocal
		return nil
	}

	var sourceGcsViaAPI SourceGcsViaAPI = SourceGcsViaAPI{}
	if err := utils.UnmarshalJSON(data, &sourceGcsViaAPI, "", true, true); err == nil {
		u.SourceGcsViaAPI = &sourceGcsViaAPI
		u.Type = SourceGcsProcessingTypeSourceGcsViaAPI
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceGcsProcessing", string(data))
}

func (u SourceGcsProcessing) MarshalJSON() ([]byte, error) {
	if u.SourceGcsLocal != nil {
		return utils.MarshalJSON(u.SourceGcsLocal, "", true)
	}

	if u.SourceGcsViaAPI != nil {
		return utils.MarshalJSON(u.SourceGcsViaAPI, "", true)
	}

	return nil, errors.New("could not marshal union type SourceGcsProcessing: all fields are null")
}

// SourceGcsParsingStrategy - The strategy used to parse documents. `fast` extracts text directly from the document which doesn't work for all files. `ocr_only` is more reliable, but slower. `hi_res` is the most reliable, but requires an API key and a hosted instance of unstructured and can't be used with local mode. See the unstructured.io documentation for more details: https://unstructured-io.github.io/unstructured/core/partition.html#partition-pdf
type SourceGcsParsingStrategy string

const (
	SourceGcsParsingStrategyAuto    SourceGcsParsingStrategy = "auto"
	SourceGcsParsingStrategyFast    SourceGcsParsingStrategy = "fast"
	SourceGcsParsingStrategyOcrOnly SourceGcsParsingStrategy = "ocr_only"
	SourceGcsParsingStrategyHiRes   SourceGcsParsingStrategy = "hi_res"
)

func (e SourceGcsParsingStrategy) ToPointer() *SourceGcsParsingStrategy {
	return &e
}
func (e *SourceGcsParsingStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "fast":
		fallthrough
	case "ocr_only":
		fallthrough
	case "hi_res":
		*e = SourceGcsParsingStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsParsingStrategy: %v", v)
	}
}

// SourceGcsUnstructuredDocumentFormat - Extract text from document formats (.pdf, .docx, .md, .pptx) and emit as one record per file.
type SourceGcsUnstructuredDocumentFormat struct {
	filetype *SourceGcsSchemasStreamsFormatFormatFiletype `const:"unstructured" json:"filetype"`
	// Processing configuration
	Processing *SourceGcsProcessing `json:"processing,omitempty"`
	// If true, skip files that cannot be parsed and pass the error message along as the _ab_source_file_parse_error field. If false, fail the sync.
	SkipUnprocessableFiles *bool `default:"true" json:"skip_unprocessable_files"`
	// The strategy used to parse documents. `fast` extracts text directly from the document which doesn't work for all files. `ocr_only` is more reliable, but slower. `hi_res` is the most reliable, but requires an API key and a hosted instance of unstructured and can't be used with local mode. See the unstructured.io documentation for more details: https://unstructured-io.github.io/unstructured/core/partition.html#partition-pdf
	Strategy *SourceGcsParsingStrategy `default:"auto" json:"strategy"`
}

func (s SourceGcsUnstructuredDocumentFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsUnstructuredDocumentFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsUnstructuredDocumentFormat) GetFiletype() *SourceGcsSchemasStreamsFormatFormatFiletype {
	return SourceGcsSchemasStreamsFormatFormatFiletypeUnstructured.ToPointer()
}

func (o *SourceGcsUnstructuredDocumentFormat) GetProcessing() *SourceGcsProcessing {
	if o == nil {
		return nil
	}
	return o.Processing
}

func (o *SourceGcsUnstructuredDocumentFormat) GetSkipUnprocessableFiles() *bool {
	if o == nil {
		return nil
	}
	return o.SkipUnprocessableFiles
}

func (o *SourceGcsUnstructuredDocumentFormat) GetStrategy() *SourceGcsParsingStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

type SourceGcsSchemasStreamsFormatFiletype string

const (
	SourceGcsSchemasStreamsFormatFiletypeParquet SourceGcsSchemasStreamsFormatFiletype = "parquet"
)

func (e SourceGcsSchemasStreamsFormatFiletype) ToPointer() *SourceGcsSchemasStreamsFormatFiletype {
	return &e
}
func (e *SourceGcsSchemasStreamsFormatFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "parquet":
		*e = SourceGcsSchemasStreamsFormatFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasStreamsFormatFiletype: %v", v)
	}
}

type SourceGcsParquetFormat struct {
	// Whether to convert decimal fields to floats. There is a loss of precision when converting decimals to floats, so this is not recommended.
	DecimalAsFloat *bool                                  `default:"false" json:"decimal_as_float"`
	filetype       *SourceGcsSchemasStreamsFormatFiletype `const:"parquet" json:"filetype"`
}

func (s SourceGcsParquetFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsParquetFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsParquetFormat) GetDecimalAsFloat() *bool {
	if o == nil {
		return nil
	}
	return o.DecimalAsFloat
}

func (o *SourceGcsParquetFormat) GetFiletype() *SourceGcsSchemasStreamsFormatFiletype {
	return SourceGcsSchemasStreamsFormatFiletypeParquet.ToPointer()
}

type SourceGcsSchemasStreamsFiletype string

const (
	SourceGcsSchemasStreamsFiletypeJsonl SourceGcsSchemasStreamsFiletype = "jsonl"
)

func (e SourceGcsSchemasStreamsFiletype) ToPointer() *SourceGcsSchemasStreamsFiletype {
	return &e
}
func (e *SourceGcsSchemasStreamsFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "jsonl":
		*e = SourceGcsSchemasStreamsFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasStreamsFiletype: %v", v)
	}
}

type SourceGcsJsonlFormat struct {
	filetype *SourceGcsSchemasStreamsFiletype `const:"jsonl" json:"filetype"`
}

func (s SourceGcsJsonlFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsJsonlFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsJsonlFormat) GetFiletype() *SourceGcsSchemasStreamsFiletype {
	return SourceGcsSchemasStreamsFiletypeJsonl.ToPointer()
}

type SourceGcsSchemasFiletype string

const (
	SourceGcsSchemasFiletypeCsv SourceGcsSchemasFiletype = "csv"
)

func (e SourceGcsSchemasFiletype) ToPointer() *SourceGcsSchemasFiletype {
	return &e
}
func (e *SourceGcsSchemasFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		*e = SourceGcsSchemasFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasFiletype: %v", v)
	}
}

type SourceGcsSchemasStreamsHeaderDefinitionType string

const (
	SourceGcsSchemasStreamsHeaderDefinitionTypeUserProvided SourceGcsSchemasStreamsHeaderDefinitionType = "User Provided"
)

func (e SourceGcsSchemasStreamsHeaderDefinitionType) ToPointer() *SourceGcsSchemasStreamsHeaderDefinitionType {
	return &e
}
func (e *SourceGcsSchemasStreamsHeaderDefinitionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "User Provided":
		*e = SourceGcsSchemasStreamsHeaderDefinitionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasStreamsHeaderDefinitionType: %v", v)
	}
}

type SourceGcsUserProvided struct {
	// The column names that will be used while emitting the CSV records
	ColumnNames          []string                                     `json:"column_names"`
	headerDefinitionType *SourceGcsSchemasStreamsHeaderDefinitionType `const:"User Provided" json:"header_definition_type"`
}

func (s SourceGcsUserProvided) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsUserProvided) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsUserProvided) GetColumnNames() []string {
	if o == nil {
		return []string{}
	}
	return o.ColumnNames
}

func (o *SourceGcsUserProvided) GetHeaderDefinitionType() *SourceGcsSchemasStreamsHeaderDefinitionType {
	return SourceGcsSchemasStreamsHeaderDefinitionTypeUserProvided.ToPointer()
}

type SourceGcsSchemasHeaderDefinitionType string

const (
	SourceGcsSchemasHeaderDefinitionTypeAutogenerated SourceGcsSchemasHeaderDefinitionType = "Autogenerated"
)

func (e SourceGcsSchemasHeaderDefinitionType) ToPointer() *SourceGcsSchemasHeaderDefinitionType {
	return &e
}
func (e *SourceGcsSchemasHeaderDefinitionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Autogenerated":
		*e = SourceGcsSchemasHeaderDefinitionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsSchemasHeaderDefinitionType: %v", v)
	}
}

type SourceGcsAutogenerated struct {
	headerDefinitionType *SourceGcsSchemasHeaderDefinitionType `const:"Autogenerated" json:"header_definition_type"`
}

func (s SourceGcsAutogenerated) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsAutogenerated) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsAutogenerated) GetHeaderDefinitionType() *SourceGcsSchemasHeaderDefinitionType {
	return SourceGcsSchemasHeaderDefinitionTypeAutogenerated.ToPointer()
}

type SourceGcsHeaderDefinitionType string

const (
	SourceGcsHeaderDefinitionTypeFromCsv SourceGcsHeaderDefinitionType = "From CSV"
)

func (e SourceGcsHeaderDefinitionType) ToPointer() *SourceGcsHeaderDefinitionType {
	return &e
}
func (e *SourceGcsHeaderDefinitionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "From CSV":
		*e = SourceGcsHeaderDefinitionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsHeaderDefinitionType: %v", v)
	}
}

type SourceGcsFromCSV struct {
	headerDefinitionType *SourceGcsHeaderDefinitionType `const:"From CSV" json:"header_definition_type"`
}

func (s SourceGcsFromCSV) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsFromCSV) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsFromCSV) GetHeaderDefinitionType() *SourceGcsHeaderDefinitionType {
	return SourceGcsHeaderDefinitionTypeFromCsv.ToPointer()
}

type SourceGcsCSVHeaderDefinitionType string

const (
	SourceGcsCSVHeaderDefinitionTypeSourceGcsFromCSV       SourceGcsCSVHeaderDefinitionType = "source-gcs_From CSV"
	SourceGcsCSVHeaderDefinitionTypeSourceGcsAutogenerated SourceGcsCSVHeaderDefinitionType = "source-gcs_Autogenerated"
	SourceGcsCSVHeaderDefinitionTypeSourceGcsUserProvided  SourceGcsCSVHeaderDefinitionType = "source-gcs_User Provided"
)

// SourceGcsCSVHeaderDefinition - How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.
type SourceGcsCSVHeaderDefinition struct {
	SourceGcsFromCSV       *SourceGcsFromCSV
	SourceGcsAutogenerated *SourceGcsAutogenerated
	SourceGcsUserProvided  *SourceGcsUserProvided

	Type SourceGcsCSVHeaderDefinitionType
}

func CreateSourceGcsCSVHeaderDefinitionSourceGcsFromCSV(sourceGcsFromCSV SourceGcsFromCSV) SourceGcsCSVHeaderDefinition {
	typ := SourceGcsCSVHeaderDefinitionTypeSourceGcsFromCSV

	return SourceGcsCSVHeaderDefinition{
		SourceGcsFromCSV: &sourceGcsFromCSV,
		Type:             typ,
	}
}

func CreateSourceGcsCSVHeaderDefinitionSourceGcsAutogenerated(sourceGcsAutogenerated SourceGcsAutogenerated) SourceGcsCSVHeaderDefinition {
	typ := SourceGcsCSVHeaderDefinitionTypeSourceGcsAutogenerated

	return SourceGcsCSVHeaderDefinition{
		SourceGcsAutogenerated: &sourceGcsAutogenerated,
		Type:                   typ,
	}
}

func CreateSourceGcsCSVHeaderDefinitionSourceGcsUserProvided(sourceGcsUserProvided SourceGcsUserProvided) SourceGcsCSVHeaderDefinition {
	typ := SourceGcsCSVHeaderDefinitionTypeSourceGcsUserProvided

	return SourceGcsCSVHeaderDefinition{
		SourceGcsUserProvided: &sourceGcsUserProvided,
		Type:                  typ,
	}
}

func (u *SourceGcsCSVHeaderDefinition) UnmarshalJSON(data []byte) error {

	var sourceGcsFromCSV SourceGcsFromCSV = SourceGcsFromCSV{}
	if err := utils.UnmarshalJSON(data, &sourceGcsFromCSV, "", true, true); err == nil {
		u.SourceGcsFromCSV = &sourceGcsFromCSV
		u.Type = SourceGcsCSVHeaderDefinitionTypeSourceGcsFromCSV
		return nil
	}

	var sourceGcsAutogenerated SourceGcsAutogenerated = SourceGcsAutogenerated{}
	if err := utils.UnmarshalJSON(data, &sourceGcsAutogenerated, "", true, true); err == nil {
		u.SourceGcsAutogenerated = &sourceGcsAutogenerated
		u.Type = SourceGcsCSVHeaderDefinitionTypeSourceGcsAutogenerated
		return nil
	}

	var sourceGcsUserProvided SourceGcsUserProvided = SourceGcsUserProvided{}
	if err := utils.UnmarshalJSON(data, &sourceGcsUserProvided, "", true, true); err == nil {
		u.SourceGcsUserProvided = &sourceGcsUserProvided
		u.Type = SourceGcsCSVHeaderDefinitionTypeSourceGcsUserProvided
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceGcsCSVHeaderDefinition", string(data))
}

func (u SourceGcsCSVHeaderDefinition) MarshalJSON() ([]byte, error) {
	if u.SourceGcsFromCSV != nil {
		return utils.MarshalJSON(u.SourceGcsFromCSV, "", true)
	}

	if u.SourceGcsAutogenerated != nil {
		return utils.MarshalJSON(u.SourceGcsAutogenerated, "", true)
	}

	if u.SourceGcsUserProvided != nil {
		return utils.MarshalJSON(u.SourceGcsUserProvided, "", true)
	}

	return nil, errors.New("could not marshal union type SourceGcsCSVHeaderDefinition: all fields are null")
}

type SourceGcsCSVFormat struct {
	// The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.
	Delimiter *string `default:"," json:"delimiter"`
	// Whether two quotes in a quoted CSV value denote a single quote in the data.
	DoubleQuote *bool `default:"true" json:"double_quote"`
	// The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.
	Encoding *string `default:"utf8" json:"encoding"`
	// The character used for escaping special characters. To disallow escaping, leave this field blank.
	EscapeChar *string `json:"escape_char,omitempty"`
	// A set of case-sensitive strings that should be interpreted as false values.
	FalseValues []string                  `json:"false_values,omitempty"`
	filetype    *SourceGcsSchemasFiletype `const:"csv" json:"filetype"`
	// How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.
	HeaderDefinition *SourceGcsCSVHeaderDefinition `json:"header_definition,omitempty"`
	// Whether to ignore errors that occur when the number of fields in the CSV does not match the number of columns in the schema.
	IgnoreErrorsOnFieldsMismatch *bool `default:"false" json:"ignore_errors_on_fields_mismatch"`
	// A set of case-sensitive strings that should be interpreted as null values. For example, if the value 'NA' should be interpreted as null, enter 'NA' in this field.
	NullValues []string `json:"null_values,omitempty"`
	// The character used for quoting CSV values. To disallow quoting, make this field blank.
	QuoteChar *string `default:"\"" json:"quote_char"`
	// The number of rows to skip after the header row.
	SkipRowsAfterHeader *int64 `default:"0" json:"skip_rows_after_header"`
	// The number of rows to skip before the header row. For example, if the header row is on the 3rd row, enter 2 in this field.
	SkipRowsBeforeHeader *int64 `default:"0" json:"skip_rows_before_header"`
	// Whether strings can be interpreted as null values. If true, strings that match the null_values set will be interpreted as null. If false, strings that match the null_values set will be interpreted as the string itself.
	StringsCanBeNull *bool `default:"true" json:"strings_can_be_null"`
	// A set of case-sensitive strings that should be interpreted as true values.
	TrueValues []string `json:"true_values,omitempty"`
}

func (s SourceGcsCSVFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsCSVFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsCSVFormat) GetDelimiter() *string {
	if o == nil {
		return nil
	}
	return o.Delimiter
}

func (o *SourceGcsCSVFormat) GetDoubleQuote() *bool {
	if o == nil {
		return nil
	}
	return o.DoubleQuote
}

func (o *SourceGcsCSVFormat) GetEncoding() *string {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *SourceGcsCSVFormat) GetEscapeChar() *string {
	if o == nil {
		return nil
	}
	return o.EscapeChar
}

func (o *SourceGcsCSVFormat) GetFalseValues() []string {
	if o == nil {
		return nil
	}
	return o.FalseValues
}

func (o *SourceGcsCSVFormat) GetFiletype() *SourceGcsSchemasFiletype {
	return SourceGcsSchemasFiletypeCsv.ToPointer()
}

func (o *SourceGcsCSVFormat) GetHeaderDefinition() *SourceGcsCSVHeaderDefinition {
	if o == nil {
		return nil
	}
	return o.HeaderDefinition
}

func (o *SourceGcsCSVFormat) GetIgnoreErrorsOnFieldsMismatch() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreErrorsOnFieldsMismatch
}

func (o *SourceGcsCSVFormat) GetNullValues() []string {
	if o == nil {
		return nil
	}
	return o.NullValues
}

func (o *SourceGcsCSVFormat) GetQuoteChar() *string {
	if o == nil {
		return nil
	}
	return o.QuoteChar
}

func (o *SourceGcsCSVFormat) GetSkipRowsAfterHeader() *int64 {
	if o == nil {
		return nil
	}
	return o.SkipRowsAfterHeader
}

func (o *SourceGcsCSVFormat) GetSkipRowsBeforeHeader() *int64 {
	if o == nil {
		return nil
	}
	return o.SkipRowsBeforeHeader
}

func (o *SourceGcsCSVFormat) GetStringsCanBeNull() *bool {
	if o == nil {
		return nil
	}
	return o.StringsCanBeNull
}

func (o *SourceGcsCSVFormat) GetTrueValues() []string {
	if o == nil {
		return nil
	}
	return o.TrueValues
}

type SourceGcsFiletype string

const (
	SourceGcsFiletypeAvro SourceGcsFiletype = "avro"
)

func (e SourceGcsFiletype) ToPointer() *SourceGcsFiletype {
	return &e
}
func (e *SourceGcsFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "avro":
		*e = SourceGcsFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsFiletype: %v", v)
	}
}

type SourceGcsAvroFormat struct {
	// Whether to convert double fields to strings. This is recommended if you have decimal numbers with a high degree of precision because there can be a loss precision when handling floating point numbers.
	DoubleAsString *bool              `default:"false" json:"double_as_string"`
	filetype       *SourceGcsFiletype `const:"avro" json:"filetype"`
}

func (s SourceGcsAvroFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsAvroFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsAvroFormat) GetDoubleAsString() *bool {
	if o == nil {
		return nil
	}
	return o.DoubleAsString
}

func (o *SourceGcsAvroFormat) GetFiletype() *SourceGcsFiletype {
	return SourceGcsFiletypeAvro.ToPointer()
}

type SourceGcsFormatType string

const (
	SourceGcsFormatTypeSourceGcsAvroFormat                 SourceGcsFormatType = "source-gcs_Avro Format"
	SourceGcsFormatTypeSourceGcsCSVFormat                  SourceGcsFormatType = "source-gcs_CSV Format"
	SourceGcsFormatTypeSourceGcsJsonlFormat                SourceGcsFormatType = "source-gcs_Jsonl Format"
	SourceGcsFormatTypeSourceGcsParquetFormat              SourceGcsFormatType = "source-gcs_Parquet Format"
	SourceGcsFormatTypeSourceGcsUnstructuredDocumentFormat SourceGcsFormatType = "source-gcs_Unstructured Document Format"
	SourceGcsFormatTypeSourceGcsExcelFormat                SourceGcsFormatType = "source-gcs_Excel Format"
)

// SourceGcsFormat - The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.
type SourceGcsFormat struct {
	SourceGcsAvroFormat                 *SourceGcsAvroFormat
	SourceGcsCSVFormat                  *SourceGcsCSVFormat
	SourceGcsJsonlFormat                *SourceGcsJsonlFormat
	SourceGcsParquetFormat              *SourceGcsParquetFormat
	SourceGcsUnstructuredDocumentFormat *SourceGcsUnstructuredDocumentFormat
	SourceGcsExcelFormat                *SourceGcsExcelFormat

	Type SourceGcsFormatType
}

func CreateSourceGcsFormatSourceGcsAvroFormat(sourceGcsAvroFormat SourceGcsAvroFormat) SourceGcsFormat {
	typ := SourceGcsFormatTypeSourceGcsAvroFormat

	return SourceGcsFormat{
		SourceGcsAvroFormat: &sourceGcsAvroFormat,
		Type:                typ,
	}
}

func CreateSourceGcsFormatSourceGcsCSVFormat(sourceGcsCSVFormat SourceGcsCSVFormat) SourceGcsFormat {
	typ := SourceGcsFormatTypeSourceGcsCSVFormat

	return SourceGcsFormat{
		SourceGcsCSVFormat: &sourceGcsCSVFormat,
		Type:               typ,
	}
}

func CreateSourceGcsFormatSourceGcsJsonlFormat(sourceGcsJsonlFormat SourceGcsJsonlFormat) SourceGcsFormat {
	typ := SourceGcsFormatTypeSourceGcsJsonlFormat

	return SourceGcsFormat{
		SourceGcsJsonlFormat: &sourceGcsJsonlFormat,
		Type:                 typ,
	}
}

func CreateSourceGcsFormatSourceGcsParquetFormat(sourceGcsParquetFormat SourceGcsParquetFormat) SourceGcsFormat {
	typ := SourceGcsFormatTypeSourceGcsParquetFormat

	return SourceGcsFormat{
		SourceGcsParquetFormat: &sourceGcsParquetFormat,
		Type:                   typ,
	}
}

func CreateSourceGcsFormatSourceGcsUnstructuredDocumentFormat(sourceGcsUnstructuredDocumentFormat SourceGcsUnstructuredDocumentFormat) SourceGcsFormat {
	typ := SourceGcsFormatTypeSourceGcsUnstructuredDocumentFormat

	return SourceGcsFormat{
		SourceGcsUnstructuredDocumentFormat: &sourceGcsUnstructuredDocumentFormat,
		Type:                                typ,
	}
}

func CreateSourceGcsFormatSourceGcsExcelFormat(sourceGcsExcelFormat SourceGcsExcelFormat) SourceGcsFormat {
	typ := SourceGcsFormatTypeSourceGcsExcelFormat

	return SourceGcsFormat{
		SourceGcsExcelFormat: &sourceGcsExcelFormat,
		Type:                 typ,
	}
}

func (u *SourceGcsFormat) UnmarshalJSON(data []byte) error {

	var sourceGcsJsonlFormat SourceGcsJsonlFormat = SourceGcsJsonlFormat{}
	if err := utils.UnmarshalJSON(data, &sourceGcsJsonlFormat, "", true, true); err == nil {
		u.SourceGcsJsonlFormat = &sourceGcsJsonlFormat
		u.Type = SourceGcsFormatTypeSourceGcsJsonlFormat
		return nil
	}

	var sourceGcsExcelFormat SourceGcsExcelFormat = SourceGcsExcelFormat{}
	if err := utils.UnmarshalJSON(data, &sourceGcsExcelFormat, "", true, true); err == nil {
		u.SourceGcsExcelFormat = &sourceGcsExcelFormat
		u.Type = SourceGcsFormatTypeSourceGcsExcelFormat
		return nil
	}

	var sourceGcsAvroFormat SourceGcsAvroFormat = SourceGcsAvroFormat{}
	if err := utils.UnmarshalJSON(data, &sourceGcsAvroFormat, "", true, true); err == nil {
		u.SourceGcsAvroFormat = &sourceGcsAvroFormat
		u.Type = SourceGcsFormatTypeSourceGcsAvroFormat
		return nil
	}

	var sourceGcsParquetFormat SourceGcsParquetFormat = SourceGcsParquetFormat{}
	if err := utils.UnmarshalJSON(data, &sourceGcsParquetFormat, "", true, true); err == nil {
		u.SourceGcsParquetFormat = &sourceGcsParquetFormat
		u.Type = SourceGcsFormatTypeSourceGcsParquetFormat
		return nil
	}

	var sourceGcsUnstructuredDocumentFormat SourceGcsUnstructuredDocumentFormat = SourceGcsUnstructuredDocumentFormat{}
	if err := utils.UnmarshalJSON(data, &sourceGcsUnstructuredDocumentFormat, "", true, true); err == nil {
		u.SourceGcsUnstructuredDocumentFormat = &sourceGcsUnstructuredDocumentFormat
		u.Type = SourceGcsFormatTypeSourceGcsUnstructuredDocumentFormat
		return nil
	}

	var sourceGcsCSVFormat SourceGcsCSVFormat = SourceGcsCSVFormat{}
	if err := utils.UnmarshalJSON(data, &sourceGcsCSVFormat, "", true, true); err == nil {
		u.SourceGcsCSVFormat = &sourceGcsCSVFormat
		u.Type = SourceGcsFormatTypeSourceGcsCSVFormat
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceGcsFormat", string(data))
}

func (u SourceGcsFormat) MarshalJSON() ([]byte, error) {
	if u.SourceGcsAvroFormat != nil {
		return utils.MarshalJSON(u.SourceGcsAvroFormat, "", true)
	}

	if u.SourceGcsCSVFormat != nil {
		return utils.MarshalJSON(u.SourceGcsCSVFormat, "", true)
	}

	if u.SourceGcsJsonlFormat != nil {
		return utils.MarshalJSON(u.SourceGcsJsonlFormat, "", true)
	}

	if u.SourceGcsParquetFormat != nil {
		return utils.MarshalJSON(u.SourceGcsParquetFormat, "", true)
	}

	if u.SourceGcsUnstructuredDocumentFormat != nil {
		return utils.MarshalJSON(u.SourceGcsUnstructuredDocumentFormat, "", true)
	}

	if u.SourceGcsExcelFormat != nil {
		return utils.MarshalJSON(u.SourceGcsExcelFormat, "", true)
	}

	return nil, errors.New("could not marshal union type SourceGcsFormat: all fields are null")
}

// SourceGcsValidationPolicy - The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema.
type SourceGcsValidationPolicy string

const (
	SourceGcsValidationPolicyEmitRecord      SourceGcsValidationPolicy = "Emit Record"
	SourceGcsValidationPolicySkipRecord      SourceGcsValidationPolicy = "Skip Record"
	SourceGcsValidationPolicyWaitForDiscover SourceGcsValidationPolicy = "Wait for Discover"
)

func (e SourceGcsValidationPolicy) ToPointer() *SourceGcsValidationPolicy {
	return &e
}
func (e *SourceGcsValidationPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Emit Record":
		fallthrough
	case "Skip Record":
		fallthrough
	case "Wait for Discover":
		*e = SourceGcsValidationPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsValidationPolicy: %v", v)
	}
}

type SourceGcsFileBasedStreamConfig struct {
	// When the state history of the file store is full, syncs will only read files that were last modified in the provided day range.
	DaysToSyncIfHistoryIsFull *int64 `default:"3" json:"days_to_sync_if_history_is_full"`
	// The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.
	Format SourceGcsFormat `json:"format"`
	// The pattern used to specify which files should be selected from the file system. For more information on glob pattern matching look <a href="https://en.wikipedia.org/wiki/Glob_(programming)">here</a>.
	Globs []string `json:"globs,omitempty"`
	// The schema that will be used to validate records extracted from the file. This will override the stream schema that is auto-detected from incoming files.
	InputSchema *string `json:"input_schema,omitempty"`
	// The name of the stream.
	Name string `json:"name"`
	// The number of resent files which will be used to discover the schema for this stream.
	RecentNFilesToReadForSchemaDiscovery *int64 `json:"recent_n_files_to_read_for_schema_discovery,omitempty"`
	// When enabled, syncs will not validate or structure records against the stream's schema.
	Schemaless *bool `default:"false" json:"schemaless"`
	// The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema.
	ValidationPolicy *SourceGcsValidationPolicy `default:"Emit Record" json:"validation_policy"`
}

func (s SourceGcsFileBasedStreamConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcsFileBasedStreamConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcsFileBasedStreamConfig) GetDaysToSyncIfHistoryIsFull() *int64 {
	if o == nil {
		return nil
	}
	return o.DaysToSyncIfHistoryIsFull
}

func (o *SourceGcsFileBasedStreamConfig) GetFormat() SourceGcsFormat {
	if o == nil {
		return SourceGcsFormat{}
	}
	return o.Format
}

func (o *SourceGcsFileBasedStreamConfig) GetGlobs() []string {
	if o == nil {
		return nil
	}
	return o.Globs
}

func (o *SourceGcsFileBasedStreamConfig) GetInputSchema() *string {
	if o == nil {
		return nil
	}
	return o.InputSchema
}

func (o *SourceGcsFileBasedStreamConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGcsFileBasedStreamConfig) GetRecentNFilesToReadForSchemaDiscovery() *int64 {
	if o == nil {
		return nil
	}
	return o.RecentNFilesToReadForSchemaDiscovery
}

func (o *SourceGcsFileBasedStreamConfig) GetSchemaless() *bool {
	if o == nil {
		return nil
	}
	return o.Schemaless
}

func (o *SourceGcsFileBasedStreamConfig) GetValidationPolicy() *SourceGcsValidationPolicy {
	if o == nil {
		return nil
	}
	return o.ValidationPolicy
}

// SourceGcs - NOTE: When this Spec is changed, legacy_config_transformer.py must also be
// modified to uptake the changes because it is responsible for converting
// legacy GCS configs into file based configs using the File-Based CDK.
type SourceGcs struct {
	// Name of the GCS bucket where the file(s) exist.
	Bucket string `json:"bucket"`
	// Enter your Google Cloud <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">service account key</a> in JSON format
	ServiceAccount string       `json:"service_account"`
	sourceType     SourceGcsGcs `const:"gcs" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00.000000Z. Any file modified before this date will not be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Each instance of this configuration defines a <a href="https://docs.airbyte.com/cloud/core-concepts#stream">stream</a>. Use this to define which files belong in the stream, their format, and how they should be parsed and validated. When sending data to warehouse destination such as Snowflake or BigQuery, each stream is a separate table.
	Streams []SourceGcsFileBasedStreamConfig `json:"streams"`
}

func (s SourceGcs) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGcs) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGcs) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *SourceGcs) GetServiceAccount() string {
	if o == nil {
		return ""
	}
	return o.ServiceAccount
}

func (o *SourceGcs) GetSourceType() SourceGcsGcs {
	return SourceGcsGcsGcs
}

func (o *SourceGcs) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceGcs) GetStreams() []SourceGcsFileBasedStreamConfig {
	if o == nil {
		return []SourceGcsFileBasedStreamConfig{}
	}
	return o.Streams
}
