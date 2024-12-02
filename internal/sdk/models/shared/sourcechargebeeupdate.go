// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

// ProductCatalog - Product Catalog version of your Chargebee site. Instructions on how to find your version you may find <a href="https://apidocs.chargebee.com/docs/api?prod_cat_ver=2">here</a> under `API Version` section. If left blank, the product catalog version will be set to 2.0.
type ProductCatalog string

const (
	ProductCatalogOne0 ProductCatalog = "1.0"
	ProductCatalogTwo0 ProductCatalog = "2.0"
)

func (e ProductCatalog) ToPointer() *ProductCatalog {
	return &e
}
func (e *ProductCatalog) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1.0":
		fallthrough
	case "2.0":
		*e = ProductCatalog(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductCatalog: %v", v)
	}
}

type SourceChargebeeUpdate struct {
	// Product Catalog version of your Chargebee site. Instructions on how to find your version you may find <a href="https://apidocs.chargebee.com/docs/api?prod_cat_ver=2">here</a> under `API Version` section. If left blank, the product catalog version will be set to 2.0.
	ProductCatalog *ProductCatalog `default:"2.0" json:"product_catalog"`
	// The site prefix for your Chargebee instance.
	Site string `json:"site"`
	// Chargebee API Key. See the <a href="https://docs.airbyte.com/integrations/sources/chargebee">docs</a> for more information on how to obtain this key.
	SiteAPIKey string `json:"site_api_key"`
	// UTC date and time in the format 2017-01-25T00:00:00.000Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceChargebeeUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceChargebeeUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceChargebeeUpdate) GetProductCatalog() *ProductCatalog {
	if o == nil {
		return nil
	}
	return o.ProductCatalog
}

func (o *SourceChargebeeUpdate) GetSite() string {
	if o == nil {
		return ""
	}
	return o.Site
}

func (o *SourceChargebeeUpdate) GetSiteAPIKey() string {
	if o == nil {
		return ""
	}
	return o.SiteAPIKey
}

func (o *SourceChargebeeUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
