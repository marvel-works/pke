/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// CreateAksPropertiesAks struct for CreateAksPropertiesAks
type CreateAksPropertiesAks struct {
	ResourceGroup string `json:"resourceGroup"`
	KubernetesVersion string `json:"kubernetesVersion"`
	NodePools map[string]NodePoolsAzure `json:"nodePools"`
	VnetSubnetID string `json:"vnetSubnetID,omitempty"`
}
