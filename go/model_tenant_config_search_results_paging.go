/*
 * configuration management API  - OpenAPI 3.0
 *
 * use this to manage configurations.
 *
 * API version: 0.0.1
 * Contact: william.ohara@subscripify.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package configsapi

type TenantConfigSearchResultsPaging struct {
	// the number of pages based upon the number of items per page requested
	PageCount int32 `json:"pageCount,omitempty"`
	// the url of the previous page in a series
	Previous string `json:"previous,omitempty"`
	// the url of the next page in a series
	Next string `json:"next,omitempty"`
}
