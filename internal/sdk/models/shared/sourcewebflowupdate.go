// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceWebflowUpdate struct {
	// The version of the Webflow API to use. See https://developers.webflow.com/#versioning
	AcceptVersion *string `json:"accept_version,omitempty"`
	// The API token for authenticating to Webflow. See https://university.webflow.com/lesson/intro-to-the-webflow-api
	APIKey string `json:"api_key"`
	// The id of the Webflow site you are requesting data from. See https://developers.webflow.com/#sites
	SiteID string `json:"site_id"`
}

func (o *SourceWebflowUpdate) GetAcceptVersion() *string {
	if o == nil {
		return nil
	}
	return o.AcceptVersion
}

func (o *SourceWebflowUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceWebflowUpdate) GetSiteID() string {
	if o == nil {
		return ""
	}
	return o.SiteID
}
