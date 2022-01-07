/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PatchScheduleObservation struct {
}

type PatchScheduleParameters struct {

	// +kubebuilder:validation:Required
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	StartHourUtc *int64 `json:"startHourUtc,omitempty" tf:"start_hour_utc,omitempty"`
}

type RedisCacheObservation struct {
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	SSLPort *int64 `json:"sslPort,omitempty" tf:"ssl_port,omitempty"`
}

type RedisCacheParameters struct {

	// +kubebuilder:validation:Required
	Capacity *int64 `json:"capacity" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Optional
	EnableNonSSLPort *bool `json:"enableNonSslPort,omitempty" tf:"enable_non_ssl_port,omitempty"`

	// +kubebuilder:validation:Required
	Family *string `json:"family" tf:"family,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	// +kubebuilder:validation:Optional
	PatchSchedule []PatchScheduleParameters `json:"patchSchedule,omitempty" tf:"patch_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateStaticIPAddress *string `json:"privateStaticIpAddress,omitempty" tf:"private_static_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RedisConfiguration []RedisConfigurationParameters `json:"redisConfiguration,omitempty" tf:"redis_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	RedisVersion *string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicasPerMaster *int64 `json:"replicasPerMaster,omitempty" tf:"replicas_per_master,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicasPerPrimary *int64 `json:"replicasPerPrimary,omitempty" tf:"replicas_per_primary,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ShardCount *int64 `json:"shardCount,omitempty" tf:"shard_count,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TenantSettings map[string]*string `json:"tenantSettings,omitempty" tf:"tenant_settings,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type RedisConfigurationObservation struct {
	Maxclients *int64 `json:"maxclients,omitempty" tf:"maxclients,omitempty"`
}

type RedisConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AofBackupEnabled *bool `json:"aofBackupEnabled,omitempty" tf:"aof_backup_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AofStorageConnectionString0SecretRef *v1.SecretKeySelector `json:"aofStorageConnectionString0SecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AofStorageConnectionString1SecretRef *v1.SecretKeySelector `json:"aofStorageConnectionString1SecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	EnableAuthentication *bool `json:"enableAuthentication,omitempty" tf:"enable_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	MaxfragmentationmemoryReserved *int64 `json:"maxfragmentationmemoryReserved,omitempty" tf:"maxfragmentationmemory_reserved,omitempty"`

	// +kubebuilder:validation:Optional
	MaxmemoryDelta *int64 `json:"maxmemoryDelta,omitempty" tf:"maxmemory_delta,omitempty"`

	// +kubebuilder:validation:Optional
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// +kubebuilder:validation:Optional
	MaxmemoryReserved *int64 `json:"maxmemoryReserved,omitempty" tf:"maxmemory_reserved,omitempty"`

	// +kubebuilder:validation:Optional
	NotifyKeySpaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// +kubebuilder:validation:Optional
	RdbBackupEnabled *bool `json:"rdbBackupEnabled,omitempty" tf:"rdb_backup_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RdbBackupFrequency *int64 `json:"rdbBackupFrequency,omitempty" tf:"rdb_backup_frequency,omitempty"`

	// +kubebuilder:validation:Optional
	RdbBackupMaxSnapshotCount *int64 `json:"rdbBackupMaxSnapshotCount,omitempty" tf:"rdb_backup_max_snapshot_count,omitempty"`

	// +kubebuilder:validation:Optional
	RdbStorageConnectionStringSecretRef *v1.SecretKeySelector `json:"rdbStorageConnectionStringSecretRef,omitempty" tf:"-"`
}

// RedisCacheSpec defines the desired state of RedisCache
type RedisCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisCacheParameters `json:"forProvider"`
}

// RedisCacheStatus defines the observed state of RedisCache.
type RedisCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCache is the Schema for the RedisCaches API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type RedisCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisCacheSpec   `json:"spec"`
	Status            RedisCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCacheList contains a list of RedisCaches
type RedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCache `json:"items"`
}

// Repository type metadata.
var (
	RedisCache_Kind             = "RedisCache"
	RedisCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisCache_Kind}.String()
	RedisCache_KindAPIVersion   = RedisCache_Kind + "." + CRDGroupVersion.String()
	RedisCache_GroupVersionKind = CRDGroupVersion.WithKind(RedisCache_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisCache{}, &RedisCacheList{})
}