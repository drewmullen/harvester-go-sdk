# KubevirtIoApiCoreV1MigrationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAutoConverge** | Pointer to **bool** | AllowAutoConverge allows the platform to compromise performance/availability of VMIs to guarantee successful VMI live migrations. Defaults to false | [optional] 
**AllowPostCopy** | Pointer to **bool** | AllowPostCopy enables post-copy live migrations. Such migrations allow even the busiest VMIs to successfully live-migrate. However, events like a network failure can cause a VMI crash. If set to true, migrations will still start in pre-copy, but switch to post-copy when CompletionTimeoutPerGiB triggers. Defaults to false | [optional] 
**BandwidthPerMigration** | Pointer to **string** | BandwidthPerMigration limits the amount of network bandwidth live migrations are allowed to use. The value is in quantity per second. Defaults to 0 (no limit) | [optional] 
**CompletionTimeoutPerGiB** | Pointer to **int64** | CompletionTimeoutPerGiB is the maximum number of seconds per GiB a migration is allowed to take. If a live-migration takes longer to migrate than this value multiplied by the size of the VMI, the migration will be cancelled, unless AllowPostCopy is true. Defaults to 800 | [optional] 
**DisableTLS** | Pointer to **bool** | When set to true, DisableTLS will disable the additional layer of live migration encryption provided by KubeVirt. This is usually a bad idea. Defaults to false | [optional] 
**MatchSELinuxLevelOnMigration** | Pointer to **bool** | By default, the SELinux level of target virt-launcher pods is forced to the level of the source virt-launcher. When set to true, MatchSELinuxLevelOnMigration lets the CRI auto-assign a random level to the target. That will ensure the target virt-launcher doesn&#39;t share categories with another pod on the node. However, migrations will fail when using RWX volumes that don&#39;t automatically deal with SELinux levels. | [optional] 
**Network** | Pointer to **string** | Network is the name of the CNI network to use for live migrations. By default, migrations go through the pod network. | [optional] 
**NodeDrainTaintKey** | Pointer to **string** | NodeDrainTaintKey defines the taint key that indicates a node should be drained. Note: this option relies on the deprecated node taint feature. Default: kubevirt.io/drain | [optional] 
**ParallelMigrationsPerCluster** | Pointer to **int64** | ParallelMigrationsPerCluster is the total number of concurrent live migrations allowed cluster-wide. Defaults to 5 | [optional] 
**ParallelOutboundMigrationsPerNode** | Pointer to **int64** | ParallelOutboundMigrationsPerNode is the maximum number of concurrent outgoing live migrations allowed per node. Defaults to 2 | [optional] 
**ProgressTimeout** | Pointer to **int64** | ProgressTimeout is the maximum number of seconds a live migration is allowed to make no progress. Hitting this timeout means a migration transferred 0 data for that many seconds. The migration is then considered stuck and therefore cancelled. Defaults to 150 | [optional] 
**UnsafeMigrationOverride** | Pointer to **bool** | UnsafeMigrationOverride allows live migrations to occur even if the compatibility check indicates the migration will be unsafe to the guest. Defaults to false | [optional] 

## Methods

### NewKubevirtIoApiCoreV1MigrationConfiguration

`func NewKubevirtIoApiCoreV1MigrationConfiguration() *KubevirtIoApiCoreV1MigrationConfiguration`

NewKubevirtIoApiCoreV1MigrationConfiguration instantiates a new KubevirtIoApiCoreV1MigrationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1MigrationConfigurationWithDefaults

`func NewKubevirtIoApiCoreV1MigrationConfigurationWithDefaults() *KubevirtIoApiCoreV1MigrationConfiguration`

NewKubevirtIoApiCoreV1MigrationConfigurationWithDefaults instantiates a new KubevirtIoApiCoreV1MigrationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAutoConverge

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetAllowAutoConverge() bool`

GetAllowAutoConverge returns the AllowAutoConverge field if non-nil, zero value otherwise.

### GetAllowAutoConvergeOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetAllowAutoConvergeOk() (*bool, bool)`

GetAllowAutoConvergeOk returns a tuple with the AllowAutoConverge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoConverge

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetAllowAutoConverge(v bool)`

SetAllowAutoConverge sets AllowAutoConverge field to given value.

### HasAllowAutoConverge

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasAllowAutoConverge() bool`

HasAllowAutoConverge returns a boolean if a field has been set.

### GetAllowPostCopy

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetAllowPostCopy() bool`

GetAllowPostCopy returns the AllowPostCopy field if non-nil, zero value otherwise.

### GetAllowPostCopyOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetAllowPostCopyOk() (*bool, bool)`

GetAllowPostCopyOk returns a tuple with the AllowPostCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPostCopy

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetAllowPostCopy(v bool)`

SetAllowPostCopy sets AllowPostCopy field to given value.

### HasAllowPostCopy

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasAllowPostCopy() bool`

HasAllowPostCopy returns a boolean if a field has been set.

### GetBandwidthPerMigration

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetBandwidthPerMigration() string`

GetBandwidthPerMigration returns the BandwidthPerMigration field if non-nil, zero value otherwise.

### GetBandwidthPerMigrationOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetBandwidthPerMigrationOk() (*string, bool)`

GetBandwidthPerMigrationOk returns a tuple with the BandwidthPerMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthPerMigration

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetBandwidthPerMigration(v string)`

SetBandwidthPerMigration sets BandwidthPerMigration field to given value.

### HasBandwidthPerMigration

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasBandwidthPerMigration() bool`

HasBandwidthPerMigration returns a boolean if a field has been set.

### GetCompletionTimeoutPerGiB

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetCompletionTimeoutPerGiB() int64`

GetCompletionTimeoutPerGiB returns the CompletionTimeoutPerGiB field if non-nil, zero value otherwise.

### GetCompletionTimeoutPerGiBOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetCompletionTimeoutPerGiBOk() (*int64, bool)`

GetCompletionTimeoutPerGiBOk returns a tuple with the CompletionTimeoutPerGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTimeoutPerGiB

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetCompletionTimeoutPerGiB(v int64)`

SetCompletionTimeoutPerGiB sets CompletionTimeoutPerGiB field to given value.

### HasCompletionTimeoutPerGiB

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasCompletionTimeoutPerGiB() bool`

HasCompletionTimeoutPerGiB returns a boolean if a field has been set.

### GetDisableTLS

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetDisableTLS() bool`

GetDisableTLS returns the DisableTLS field if non-nil, zero value otherwise.

### GetDisableTLSOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetDisableTLSOk() (*bool, bool)`

GetDisableTLSOk returns a tuple with the DisableTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTLS

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetDisableTLS(v bool)`

SetDisableTLS sets DisableTLS field to given value.

### HasDisableTLS

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasDisableTLS() bool`

HasDisableTLS returns a boolean if a field has been set.

### GetMatchSELinuxLevelOnMigration

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetMatchSELinuxLevelOnMigration() bool`

GetMatchSELinuxLevelOnMigration returns the MatchSELinuxLevelOnMigration field if non-nil, zero value otherwise.

### GetMatchSELinuxLevelOnMigrationOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetMatchSELinuxLevelOnMigrationOk() (*bool, bool)`

GetMatchSELinuxLevelOnMigrationOk returns a tuple with the MatchSELinuxLevelOnMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSELinuxLevelOnMigration

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetMatchSELinuxLevelOnMigration(v bool)`

SetMatchSELinuxLevelOnMigration sets MatchSELinuxLevelOnMigration field to given value.

### HasMatchSELinuxLevelOnMigration

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasMatchSELinuxLevelOnMigration() bool`

HasMatchSELinuxLevelOnMigration returns a boolean if a field has been set.

### GetNetwork

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodeDrainTaintKey

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetNodeDrainTaintKey() string`

GetNodeDrainTaintKey returns the NodeDrainTaintKey field if non-nil, zero value otherwise.

### GetNodeDrainTaintKeyOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetNodeDrainTaintKeyOk() (*string, bool)`

GetNodeDrainTaintKeyOk returns a tuple with the NodeDrainTaintKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDrainTaintKey

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetNodeDrainTaintKey(v string)`

SetNodeDrainTaintKey sets NodeDrainTaintKey field to given value.

### HasNodeDrainTaintKey

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasNodeDrainTaintKey() bool`

HasNodeDrainTaintKey returns a boolean if a field has been set.

### GetParallelMigrationsPerCluster

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetParallelMigrationsPerCluster() int64`

GetParallelMigrationsPerCluster returns the ParallelMigrationsPerCluster field if non-nil, zero value otherwise.

### GetParallelMigrationsPerClusterOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetParallelMigrationsPerClusterOk() (*int64, bool)`

GetParallelMigrationsPerClusterOk returns a tuple with the ParallelMigrationsPerCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelMigrationsPerCluster

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetParallelMigrationsPerCluster(v int64)`

SetParallelMigrationsPerCluster sets ParallelMigrationsPerCluster field to given value.

### HasParallelMigrationsPerCluster

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasParallelMigrationsPerCluster() bool`

HasParallelMigrationsPerCluster returns a boolean if a field has been set.

### GetParallelOutboundMigrationsPerNode

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetParallelOutboundMigrationsPerNode() int64`

GetParallelOutboundMigrationsPerNode returns the ParallelOutboundMigrationsPerNode field if non-nil, zero value otherwise.

### GetParallelOutboundMigrationsPerNodeOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetParallelOutboundMigrationsPerNodeOk() (*int64, bool)`

GetParallelOutboundMigrationsPerNodeOk returns a tuple with the ParallelOutboundMigrationsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelOutboundMigrationsPerNode

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetParallelOutboundMigrationsPerNode(v int64)`

SetParallelOutboundMigrationsPerNode sets ParallelOutboundMigrationsPerNode field to given value.

### HasParallelOutboundMigrationsPerNode

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasParallelOutboundMigrationsPerNode() bool`

HasParallelOutboundMigrationsPerNode returns a boolean if a field has been set.

### GetProgressTimeout

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetProgressTimeout() int64`

GetProgressTimeout returns the ProgressTimeout field if non-nil, zero value otherwise.

### GetProgressTimeoutOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetProgressTimeoutOk() (*int64, bool)`

GetProgressTimeoutOk returns a tuple with the ProgressTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTimeout

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetProgressTimeout(v int64)`

SetProgressTimeout sets ProgressTimeout field to given value.

### HasProgressTimeout

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasProgressTimeout() bool`

HasProgressTimeout returns a boolean if a field has been set.

### GetUnsafeMigrationOverride

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetUnsafeMigrationOverride() bool`

GetUnsafeMigrationOverride returns the UnsafeMigrationOverride field if non-nil, zero value otherwise.

### GetUnsafeMigrationOverrideOk

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) GetUnsafeMigrationOverrideOk() (*bool, bool)`

GetUnsafeMigrationOverrideOk returns a tuple with the UnsafeMigrationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeMigrationOverride

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) SetUnsafeMigrationOverride(v bool)`

SetUnsafeMigrationOverride sets UnsafeMigrationOverride field to given value.

### HasUnsafeMigrationOverride

`func (o *KubevirtIoApiCoreV1MigrationConfiguration) HasUnsafeMigrationOverride() bool`

HasUnsafeMigrationOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


