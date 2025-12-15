package db_service

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellServiceDTO *model.TessellServiceDTO) error {

	if err := d.Set("id", tessellServiceDTO.Id); err != nil {
		return err
	}

	if err := d.Set("availability_machine_id", tessellServiceDTO.AvailabilityMachineId); err != nil {
		return err
	}

	if err := d.Set("name", tessellServiceDTO.Name); err != nil {
		return err
	}

	if err := d.Set("description", tessellServiceDTO.Description); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellServiceDTO.EngineType); err != nil {
		return err
	}

	if err := d.Set("topology", tessellServiceDTO.Topology); err != nil {
		return err
	}

	if err := d.Set("num_of_instances", tessellServiceDTO.NumOfInstances); err != nil {
		return err
	}

	if err := d.Set("status", tessellServiceDTO.Status); err != nil {
		return err
	}

	if err := d.Set("context_info", parseTessellServiceContextInfoWithResData(tessellServiceDTO.ContextInfo, d)); err != nil {
		return err
	}

	if err := d.Set("license_type", tessellServiceDTO.LicenseType); err != nil {
		return err
	}

	if err := d.Set("auto_minor_version_update", tessellServiceDTO.AutoMinorVersionUpdate); err != nil {
		return err
	}

	if err := d.Set("enable_deletion_protection", tessellServiceDTO.EnableDeletionProtection); err != nil {
		return err
	}

	if err := d.Set("enable_stop_protection", tessellServiceDTO.EnableStopProtection); err != nil {
		return err
	}

	if err := d.Set("edition", tessellServiceDTO.Edition); err != nil {
		return err
	}

	if err := d.Set("software_image", tessellServiceDTO.SoftwareImage); err != nil {
		return err
	}

	if err := d.Set("software_image_version", tessellServiceDTO.SoftwareImageVersion); err != nil {
		return err
	}

	if err := d.Set("software_image_version_family", tessellServiceDTO.SoftwareImageVersionFamily); err != nil {
		return err
	}

	if err := d.Set("tenant_id", tessellServiceDTO.TenantId); err != nil {
		return err
	}

	if err := d.Set("subscription", tessellServiceDTO.Subscription); err != nil {
		return err
	}

	if err := d.Set("owner", tessellServiceDTO.Owner); err != nil {
		return err
	}

	if err := d.Set("logged_in_user_role", tessellServiceDTO.LoggedInUserRole); err != nil {
		return err
	}

	if err := d.Set("date_created", tessellServiceDTO.DateCreated); err != nil {
		return err
	}

	if err := d.Set("started_at", tessellServiceDTO.StartedAt); err != nil {
		return err
	}

	if err := d.Set("stopped_at", tessellServiceDTO.StoppedAt); err != nil {
		return err
	}

	if err := d.Set("cloned_from_info", parseTessellServiceClonedFromInfoWithResData(tessellServiceDTO.ClonedFromInfo, d)); err != nil {
		return err
	}

	if err := d.Set("refresh_info", parseRefreshServiceInfoWithResData(tessellServiceDTO.RefreshInfo, d)); err != nil {
		return err
	}

	if err := d.Set("service_connectivity", parseTessellServiceConnectivityInfoWithResData(tessellServiceDTO.ServiceConnectivity, d)); err != nil {
		return err
	}

	if err := d.Set("tessell_genie_status", tessellServiceDTO.TessellGenieStatus); err != nil {
		return err
	}

	if err := d.Set("infrastructure", parseTessellServiceInfrastructureInfoWithResData(tessellServiceDTO.Infrastructure, d)); err != nil {
		return err
	}

	if err := d.Set("maintenance_window", parseTessellServiceMaintenanceWindowWithResData(tessellServiceDTO.MaintenanceWindow, d)); err != nil {
		return err
	}

	if err := d.Set("engine_configuration", parseTessellServiceEngineInfoWithResData(tessellServiceDTO.EngineConfiguration, d)); err != nil {
		return err
	}

	if err := d.Set("integrations_config", parseTessellServiceIntegrationsInfoWithResData(tessellServiceDTO.IntegrationsConfig, d)); err != nil {
		return err
	}

	if err := d.Set("deletion_config", parseTessellServiceDeletionConfigWithResData(tessellServiceDTO.DeletionConfig, d)); err != nil {
		return err
	}

	if err := d.Set("tags", parseTessellTagListWithResData(tessellServiceDTO.Tags, d)); err != nil {
		return err
	}

	if err := d.Set("updates_in_progress", parseTessellResourceUpdateInfoListWithResData(tessellServiceDTO.UpdatesInProgress, d)); err != nil {
		return err
	}

	if err := d.Set("instances", parseTessellServiceInstanceDTOListWithResData(tessellServiceDTO.Instances, d)); err != nil {
		return err
	}

	if err := d.Set("databases", parseTessellDatabaseDTOListWithResData(tessellServiceDTO.Databases, d)); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingInfoWithResData(tessellServiceDTO.SharedWith, d)); err != nil {
		return err
	}

	if err := d.Set("deletion_schedule", parseDeletionScheduleDTOWithResData(tessellServiceDTO.DeletionSchedule, d)); err != nil {
		return err
	}

	if err := d.Set("upcoming_scheduled_actions", parseServiceUpcomingScheduledActionsWithResData(tessellServiceDTO.UpcomingScheduledActions, d)); err != nil {
		return err
	}

	return nil
}

func parseTessellServiceContextInfoWithResData(contextInfo *model.TessellServiceContextInfo, d *schema.ResourceData) []interface{} {
	if contextInfo == nil {
		return nil
	}
	parsedContextInfo := make(map[string]interface{})
	if d.Get("context_info") != nil {
		contextInfoResourceData := d.Get("context_info").([]interface{})
		if len(contextInfoResourceData) > 0 {
			parsedContextInfo = (contextInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedContextInfo["sub_status"] = contextInfo.SubStatus
	parsedContextInfo["description"] = contextInfo.Description

	return []interface{}{parsedContextInfo}
}

func parseTessellServiceContextInfo(contextInfo *model.TessellServiceContextInfo) interface{} {
	if contextInfo == nil {
		return nil
	}
	parsedContextInfo := make(map[string]interface{})
	parsedContextInfo["sub_status"] = contextInfo.SubStatus
	parsedContextInfo["description"] = contextInfo.Description

	return parsedContextInfo
}

func parseTessellServiceClonedFromInfoWithResData(clonedFromInfo *model.TessellServiceClonedFromInfo, d *schema.ResourceData) []interface{} {
	if clonedFromInfo == nil {
		return nil
	}
	parsedClonedFromInfo := make(map[string]interface{})
	if d.Get("cloned_from_info") != nil {
		clonedFromInfoResourceData := d.Get("cloned_from_info").([]interface{})
		if len(clonedFromInfoResourceData) > 0 {
			parsedClonedFromInfo = (clonedFromInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedClonedFromInfo["clone_type"] = clonedFromInfo.CloneType
	parsedClonedFromInfo["content_type"] = clonedFromInfo.ContentType
	parsedClonedFromInfo["tessell_service_id"] = clonedFromInfo.TessellServiceId
	parsedClonedFromInfo["availability_machine_id"] = clonedFromInfo.AvailabilityMachineId
	parsedClonedFromInfo["tessell_service"] = clonedFromInfo.TessellService
	parsedClonedFromInfo["availability_machine"] = clonedFromInfo.AvailabilityMachine
	parsedClonedFromInfo["snapshot_name"] = clonedFromInfo.SnapshotName
	parsedClonedFromInfo["snapshot_id"] = clonedFromInfo.SnapshotId
	parsedClonedFromInfo["snapshot_time"] = clonedFromInfo.SnapshotTime
	parsedClonedFromInfo["pitr_time"] = clonedFromInfo.PITRTime
	parsedClonedFromInfo["maximum_recoverability"] = clonedFromInfo.MaximumRecoverability
	parsedClonedFromInfo["storage_provider"] = clonedFromInfo.StorageProvider

	return []interface{}{parsedClonedFromInfo}
}

func parseTessellServiceClonedFromInfo(clonedFromInfo *model.TessellServiceClonedFromInfo) interface{} {
	if clonedFromInfo == nil {
		return nil
	}
	parsedClonedFromInfo := make(map[string]interface{})
	parsedClonedFromInfo["clone_type"] = clonedFromInfo.CloneType
	parsedClonedFromInfo["content_type"] = clonedFromInfo.ContentType
	parsedClonedFromInfo["tessell_service_id"] = clonedFromInfo.TessellServiceId
	parsedClonedFromInfo["availability_machine_id"] = clonedFromInfo.AvailabilityMachineId
	parsedClonedFromInfo["tessell_service"] = clonedFromInfo.TessellService
	parsedClonedFromInfo["availability_machine"] = clonedFromInfo.AvailabilityMachine
	parsedClonedFromInfo["snapshot_name"] = clonedFromInfo.SnapshotName
	parsedClonedFromInfo["snapshot_id"] = clonedFromInfo.SnapshotId
	parsedClonedFromInfo["snapshot_time"] = clonedFromInfo.SnapshotTime
	parsedClonedFromInfo["pitr_time"] = clonedFromInfo.PITRTime
	parsedClonedFromInfo["maximum_recoverability"] = clonedFromInfo.MaximumRecoverability
	parsedClonedFromInfo["storage_provider"] = clonedFromInfo.StorageProvider

	return parsedClonedFromInfo
}

func parseRefreshServiceInfoWithResData(refreshInfo *model.RefreshServiceInfo, d *schema.ResourceData) []interface{} {
	if refreshInfo == nil {
		return nil
	}
	parsedRefreshInfo := make(map[string]interface{})
	if d.Get("refresh_info") != nil {
		refreshInfoResourceData := d.Get("refresh_info").([]interface{})
		if len(refreshInfoResourceData) > 0 {
			parsedRefreshInfo = (refreshInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedRefreshInfo["content_type"] = refreshInfo.ContentType
	parsedRefreshInfo["snapshot_name"] = refreshInfo.SnapshotName
	parsedRefreshInfo["snapshot_time"] = refreshInfo.SnapshotTime
	parsedRefreshInfo["pitr"] = refreshInfo.PITR

	parsedRefreshInfo["schedule_id"] = refreshInfo.ScheduleId
	parsedRefreshInfo["last_successful_refresh_time"] = refreshInfo.LastSuccessfulRefreshTime

	var scriptInfo *model.PrePostScriptInfo
	if refreshInfo.ScriptInfo != scriptInfo {
		parsedRefreshInfo["script_info"] = []interface{}{parsePrePostScriptInfo(refreshInfo.ScriptInfo)}
	}

	return []interface{}{parsedRefreshInfo}
}

func parseRefreshServiceInfo(refreshInfo *model.RefreshServiceInfo) interface{} {
	if refreshInfo == nil {
		return nil
	}
	parsedRefreshInfo := make(map[string]interface{})
	parsedRefreshInfo["content_type"] = refreshInfo.ContentType
	parsedRefreshInfo["snapshot_name"] = refreshInfo.SnapshotName
	parsedRefreshInfo["snapshot_time"] = refreshInfo.SnapshotTime
	parsedRefreshInfo["pitr"] = refreshInfo.PITR

	parsedRefreshInfo["schedule_id"] = refreshInfo.ScheduleId
	parsedRefreshInfo["last_successful_refresh_time"] = refreshInfo.LastSuccessfulRefreshTime

	var scriptInfo *model.PrePostScriptInfo
	if refreshInfo.ScriptInfo != scriptInfo {
		parsedRefreshInfo["script_info"] = []interface{}{parsePrePostScriptInfo(refreshInfo.ScriptInfo)}
	}

	return parsedRefreshInfo
}

func parsePrePostScriptInfo(prePostScriptInfo *model.PrePostScriptInfo) interface{} {
	if prePostScriptInfo == nil {
		return nil
	}
	parsedPrePostScriptInfo := make(map[string]interface{})

	var preScriptInfo *[]model.ScriptInfo
	if prePostScriptInfo.PreScriptInfo != preScriptInfo {
		parsedPrePostScriptInfo["pre_script_info"] = parseScriptInfoList(prePostScriptInfo.PreScriptInfo)
	}

	var postScriptInfo *[]model.ScriptInfo
	if prePostScriptInfo.PostScriptInfo != postScriptInfo {
		parsedPrePostScriptInfo["post_script_info"] = parseScriptInfoList(prePostScriptInfo.PostScriptInfo)
	}

	return parsedPrePostScriptInfo
}

func parseScriptInfoList(scriptInfo *[]model.ScriptInfo) []interface{} {
	if scriptInfo == nil {
		return nil
	}
	scriptInfoList := make([]interface{}, 0)

	if scriptInfo != nil {
		scriptInfoList = make([]interface{}, len(*scriptInfo))
		for i, scriptInfoItem := range *scriptInfo {
			scriptInfoList[i] = parseScriptInfo(&scriptInfoItem)
		}
	}

	return scriptInfoList
}

func parseScriptInfo(scriptInfo *model.ScriptInfo) interface{} {
	if scriptInfo == nil {
		return nil
	}
	parsedScriptInfo := make(map[string]interface{})
	parsedScriptInfo["script_id"] = scriptInfo.ScriptId
	parsedScriptInfo["script_version"] = scriptInfo.ScriptVersion

	return parsedScriptInfo
}

func parseTessellServiceConnectivityInfoWithResData(serviceConnectivity *model.TessellServiceConnectivityInfo, d *schema.ResourceData) []interface{} {
	if serviceConnectivity == nil {
		return nil
	}
	parsedServiceConnectivity := make(map[string]interface{})
	if d.Get("service_connectivity") != nil {
		serviceConnectivityResourceData := d.Get("service_connectivity").([]interface{})
		if len(serviceConnectivityResourceData) > 0 {
			parsedServiceConnectivity = (serviceConnectivityResourceData[0]).(map[string]interface{})
		}
	}
	parsedServiceConnectivity["enable_ssl"] = serviceConnectivity.EnableSSL
	parsedServiceConnectivity["ca_cert_id"] = serviceConnectivity.CaCertId
	parsedServiceConnectivity["dns_prefix"] = serviceConnectivity.DNSPrefix
	parsedServiceConnectivity["service_port"] = serviceConnectivity.ServicePort
	parsedServiceConnectivity["enable_public_access"] = serviceConnectivity.EnablePublicAccess
	parsedServiceConnectivity["allowed_ip_addresses"] = serviceConnectivity.AllowedIpAddresses

	var connectStrings *[]model.TessellServiceConnectString
	if serviceConnectivity.ConnectStrings != connectStrings {
		parsedServiceConnectivity["connect_strings"] = parseTessellServiceConnectStringList(serviceConnectivity.ConnectStrings)
	}

	var privateLink *model.ServiceConnectivityPrivateLink
	if serviceConnectivity.PrivateLink != privateLink {
		parsedServiceConnectivity["private_link"] = []interface{}{parseServiceConnectivityPrivateLink(serviceConnectivity.PrivateLink)}
	}

	var computesConnectivity *[]model.ComputeConnectivityConfig
	if serviceConnectivity.ComputesConnectivity != computesConnectivity {
		parsedServiceConnectivity["computes_connectivity"] = parseComputeConnectivityConfigList(serviceConnectivity.ComputesConnectivity)
	}

	var updateInProgressInfo *model.TessellServiceConnectivityUpdateInProgressInfo
	if serviceConnectivity.UpdateInProgressInfo != updateInProgressInfo {
		parsedServiceConnectivity["update_in_progress_info"] = []interface{}{parseTessellServiceConnectivityUpdateInProgressInfo(serviceConnectivity.UpdateInProgressInfo)}
	}

	return []interface{}{parsedServiceConnectivity}
}

func parseTessellServiceConnectivityInfo(serviceConnectivity *model.TessellServiceConnectivityInfo) interface{} {
	if serviceConnectivity == nil {
		return nil
	}
	parsedServiceConnectivity := make(map[string]interface{})
	parsedServiceConnectivity["enable_ssl"] = serviceConnectivity.EnableSSL
	parsedServiceConnectivity["ca_cert_id"] = serviceConnectivity.CaCertId
	parsedServiceConnectivity["dns_prefix"] = serviceConnectivity.DNSPrefix
	parsedServiceConnectivity["service_port"] = serviceConnectivity.ServicePort
	parsedServiceConnectivity["enable_public_access"] = serviceConnectivity.EnablePublicAccess
	parsedServiceConnectivity["allowed_ip_addresses"] = serviceConnectivity.AllowedIpAddresses

	var connectStrings *[]model.TessellServiceConnectString
	if serviceConnectivity.ConnectStrings != connectStrings {
		parsedServiceConnectivity["connect_strings"] = parseTessellServiceConnectStringList(serviceConnectivity.ConnectStrings)
	}

	var privateLink *model.ServiceConnectivityPrivateLink
	if serviceConnectivity.PrivateLink != privateLink {
		parsedServiceConnectivity["private_link"] = []interface{}{parseServiceConnectivityPrivateLink(serviceConnectivity.PrivateLink)}
	}

	var computesConnectivity *[]model.ComputeConnectivityConfig
	if serviceConnectivity.ComputesConnectivity != computesConnectivity {
		parsedServiceConnectivity["computes_connectivity"] = parseComputeConnectivityConfigList(serviceConnectivity.ComputesConnectivity)
	}

	var updateInProgressInfo *model.TessellServiceConnectivityUpdateInProgressInfo
	if serviceConnectivity.UpdateInProgressInfo != updateInProgressInfo {
		parsedServiceConnectivity["update_in_progress_info"] = []interface{}{parseTessellServiceConnectivityUpdateInProgressInfo(serviceConnectivity.UpdateInProgressInfo)}
	}

	return parsedServiceConnectivity
}

func parseTessellServiceConnectStringList(tessellServiceConnectString *[]model.TessellServiceConnectString) []interface{} {
	if tessellServiceConnectString == nil {
		return nil
	}
	tessellServiceConnectStringList := make([]interface{}, 0)

	if tessellServiceConnectString != nil {
		tessellServiceConnectStringList = make([]interface{}, len(*tessellServiceConnectString))
		for i, tessellServiceConnectStringItem := range *tessellServiceConnectString {
			tessellServiceConnectStringList[i] = parseTessellServiceConnectString(&tessellServiceConnectStringItem)
		}
	}

	return tessellServiceConnectStringList
}

func parseTessellServiceConnectString(tessellServiceConnectString *model.TessellServiceConnectString) interface{} {
	if tessellServiceConnectString == nil {
		return nil
	}
	parsedTessellServiceConnectString := make(map[string]interface{})
	parsedTessellServiceConnectString["type"] = tessellServiceConnectString.Type
	parsedTessellServiceConnectString["usage_type"] = tessellServiceConnectString.UsageType
	parsedTessellServiceConnectString["connect_descriptor"] = tessellServiceConnectString.ConnectDescriptor
	parsedTessellServiceConnectString["endpoint"] = tessellServiceConnectString.Endpoint
	parsedTessellServiceConnectString["master_user"] = tessellServiceConnectString.MasterUser
	parsedTessellServiceConnectString["service_port"] = tessellServiceConnectString.ServicePort

	return parsedTessellServiceConnectString
}

func parseServiceConnectivityPrivateLink(serviceConnectivityPrivateLink *model.ServiceConnectivityPrivateLink) interface{} {
	if serviceConnectivityPrivateLink == nil {
		return nil
	}
	parsedServiceConnectivityPrivateLink := make(map[string]interface{})
	parsedServiceConnectivityPrivateLink["status"] = serviceConnectivityPrivateLink.Status
	parsedServiceConnectivityPrivateLink["service_principals"] = serviceConnectivityPrivateLink.ServicePrincipals
	parsedServiceConnectivityPrivateLink["endpoint_service_name"] = serviceConnectivityPrivateLink.EndpointServiceName
	parsedServiceConnectivityPrivateLink["client_azure_subscription_ids"] = serviceConnectivityPrivateLink.ClientAzureSubscriptionIds
	parsedServiceConnectivityPrivateLink["private_link_service_alias"] = serviceConnectivityPrivateLink.PrivateLinkServiceAlias

	return parsedServiceConnectivityPrivateLink
}

func parseComputeConnectivityConfigList(computeConnectivityConfig *[]model.ComputeConnectivityConfig) []interface{} {
	if computeConnectivityConfig == nil {
		return nil
	}
	computeConnectivityConfigList := make([]interface{}, 0)

	if computeConnectivityConfig != nil {
		computeConnectivityConfigList = make([]interface{}, len(*computeConnectivityConfig))
		for i, computeConnectivityConfigItem := range *computeConnectivityConfig {
			computeConnectivityConfigList[i] = parseComputeConnectivityConfig(&computeConnectivityConfigItem)
		}
	}

	return computeConnectivityConfigList
}

func parseComputeConnectivityConfig(computeConnectivityConfig *model.ComputeConnectivityConfig) interface{} {
	if computeConnectivityConfig == nil {
		return nil
	}
	parsedComputeConnectivityConfig := make(map[string]interface{})
	parsedComputeConnectivityConfig["compute_resource_id"] = computeConnectivityConfig.ComputeResourceId

	var portAccessConfig *[]model.PortAccessConfig
	if computeConnectivityConfig.PortAccessConfig != portAccessConfig {
		parsedComputeConnectivityConfig["port_access_config"] = parsePortAccessConfigList(computeConnectivityConfig.PortAccessConfig)
	}

	return parsedComputeConnectivityConfig
}

func parsePortAccessConfigList(portAccessConfig *[]model.PortAccessConfig) []interface{} {
	if portAccessConfig == nil {
		return nil
	}
	portAccessConfigList := make([]interface{}, 0)

	if portAccessConfig != nil {
		portAccessConfigList = make([]interface{}, len(*portAccessConfig))
		for i, portAccessConfigItem := range *portAccessConfig {
			portAccessConfigList[i] = parsePortAccessConfig(&portAccessConfigItem)
		}
	}

	return portAccessConfigList
}

func parsePortAccessConfig(portAccessConfig *model.PortAccessConfig) interface{} {
	if portAccessConfig == nil {
		return nil
	}
	parsedPortAccessConfig := make(map[string]interface{})
	parsedPortAccessConfig["port"] = portAccessConfig.Port
	parsedPortAccessConfig["enable_public_access"] = portAccessConfig.EnablePublicAccess
	parsedPortAccessConfig["allowed_ip_addresses"] = portAccessConfig.AllowedIpAddresses

	return parsedPortAccessConfig
}

func parseTessellServiceConnectivityUpdateInProgressInfo(tessellServiceConnectivityUpdateInProgressInfo *model.TessellServiceConnectivityUpdateInProgressInfo) interface{} {
	if tessellServiceConnectivityUpdateInProgressInfo == nil {
		return nil
	}
	parsedTessellServiceConnectivityUpdateInProgressInfo := make(map[string]interface{})
	parsedTessellServiceConnectivityUpdateInProgressInfo["dns_prefix"] = tessellServiceConnectivityUpdateInProgressInfo.DNSPrefix
	parsedTessellServiceConnectivityUpdateInProgressInfo["enable_public_access"] = tessellServiceConnectivityUpdateInProgressInfo.EnablePublicAccess
	parsedTessellServiceConnectivityUpdateInProgressInfo["allowed_ip_addresses"] = tessellServiceConnectivityUpdateInProgressInfo.AllowedIpAddresses

	var privateLink *model.ServiceConnectivityUpdateInProgressInfo
	if tessellServiceConnectivityUpdateInProgressInfo.PrivateLink != privateLink {
		parsedTessellServiceConnectivityUpdateInProgressInfo["private_link"] = []interface{}{parseServiceConnectivityUpdateInProgressInfo(tessellServiceConnectivityUpdateInProgressInfo.PrivateLink)}
	}

	var computesConnectivity *[]model.ComputeConnectivityConfig
	if tessellServiceConnectivityUpdateInProgressInfo.ComputesConnectivity != computesConnectivity {
		parsedTessellServiceConnectivityUpdateInProgressInfo["computes_connectivity"] = parseComputeConnectivityConfigList(tessellServiceConnectivityUpdateInProgressInfo.ComputesConnectivity)
	}

	return parsedTessellServiceConnectivityUpdateInProgressInfo
}

func parseServiceConnectivityUpdateInProgressInfo(serviceConnectivityUpdateInProgressInfo *model.ServiceConnectivityUpdateInProgressInfo) interface{} {
	if serviceConnectivityUpdateInProgressInfo == nil {
		return nil
	}
	parsedServiceConnectivityUpdateInProgressInfo := make(map[string]interface{})
	parsedServiceConnectivityUpdateInProgressInfo["service_principals"] = serviceConnectivityUpdateInProgressInfo.ServicePrincipals
	parsedServiceConnectivityUpdateInProgressInfo["client_azure_subscription_ids"] = serviceConnectivityUpdateInProgressInfo.ClientAzureSubscriptionIds

	return parsedServiceConnectivityUpdateInProgressInfo
}

func parseTessellServiceInfrastructureInfoWithResData(infrastructure *model.TessellServiceInfrastructureInfo, d *schema.ResourceData) []interface{} {
	if infrastructure == nil {
		return nil
	}
	parsedInfrastructure := make(map[string]interface{})
	if d.Get("infrastructure") != nil {
		infrastructureResourceData := d.Get("infrastructure").([]interface{})
		if len(infrastructureResourceData) > 0 {
			parsedInfrastructure = (infrastructureResourceData[0]).(map[string]interface{})
		}
	}
	parsedInfrastructure["cloud"] = infrastructure.Cloud
	parsedInfrastructure["region"] = infrastructure.Region
	parsedInfrastructure["availability_zone"] = infrastructure.AvailabilityZone

	parsedInfrastructure["vpc"] = infrastructure.VPC
	parsedInfrastructure["enable_encryption"] = infrastructure.EnableEncryption
	parsedInfrastructure["encryption_key"] = infrastructure.EncryptionKey
	parsedInfrastructure["compute_type"] = infrastructure.ComputeType

	parsedInfrastructure["enable_compute_sharing"] = infrastructure.EnableComputeSharing
	parsedInfrastructure["iops"] = infrastructure.Iops
	parsedInfrastructure["throughput"] = infrastructure.Throughput
	parsedInfrastructure["storage"] = infrastructure.Storage
	parsedInfrastructure["additional_storage"] = infrastructure.AdditionalStorage
	parsedInfrastructure["timezone"] = infrastructure.Timezone
	parsedInfrastructure["multi_disk"] = infrastructure.MultiDisk
	parsedInfrastructure["storage_provider"] = infrastructure.StorageProvider

	parsedInfrastructure["compute_provider"] = infrastructure.ComputeProvider

	var cloudAvailability *[]model.CloudRegionInfo
	if infrastructure.CloudAvailability != cloudAvailability {
		parsedInfrastructure["cloud_availability"] = parseCloudRegionInfoList(infrastructure.CloudAvailability)
	}

	var awsInfraConfig *model.AwsInfraConfig
	if infrastructure.AwsInfraConfig != awsInfraConfig {
		parsedInfrastructure["aws_infra_config"] = []interface{}{parseAwsInfraConfig(infrastructure.AwsInfraConfig)}
	}

	var storageConfig *model.ServiceStorageConfig
	if infrastructure.StorageConfig != storageConfig {
		parsedInfrastructure["storage_config"] = []interface{}{parseServiceStorageConfig(infrastructure.StorageConfig)}
	}

	var archiveStorageConfig *model.ServiceStorageConfig
	if infrastructure.ArchiveStorageConfig != archiveStorageConfig {
		parsedInfrastructure["archive_storage_config"] = []interface{}{parseServiceStorageConfig(infrastructure.ArchiveStorageConfig)}
	}

	return []interface{}{parsedInfrastructure}
}

func parseTessellServiceInfrastructureInfo(infrastructure *model.TessellServiceInfrastructureInfo) interface{} {
	if infrastructure == nil {
		return nil
	}
	parsedInfrastructure := make(map[string]interface{})
	parsedInfrastructure["cloud"] = infrastructure.Cloud
	parsedInfrastructure["region"] = infrastructure.Region
	parsedInfrastructure["availability_zone"] = infrastructure.AvailabilityZone

	parsedInfrastructure["vpc"] = infrastructure.VPC
	parsedInfrastructure["enable_encryption"] = infrastructure.EnableEncryption
	parsedInfrastructure["encryption_key"] = infrastructure.EncryptionKey
	parsedInfrastructure["compute_type"] = infrastructure.ComputeType

	parsedInfrastructure["enable_compute_sharing"] = infrastructure.EnableComputeSharing
	parsedInfrastructure["iops"] = infrastructure.Iops
	parsedInfrastructure["throughput"] = infrastructure.Throughput
	parsedInfrastructure["storage"] = infrastructure.Storage
	parsedInfrastructure["additional_storage"] = infrastructure.AdditionalStorage
	parsedInfrastructure["timezone"] = infrastructure.Timezone
	parsedInfrastructure["multi_disk"] = infrastructure.MultiDisk
	parsedInfrastructure["storage_provider"] = infrastructure.StorageProvider

	parsedInfrastructure["compute_provider"] = infrastructure.ComputeProvider

	var cloudAvailability *[]model.CloudRegionInfo
	if infrastructure.CloudAvailability != cloudAvailability {
		parsedInfrastructure["cloud_availability"] = parseCloudRegionInfoList(infrastructure.CloudAvailability)
	}

	var awsInfraConfig *model.AwsInfraConfig
	if infrastructure.AwsInfraConfig != awsInfraConfig {
		parsedInfrastructure["aws_infra_config"] = []interface{}{parseAwsInfraConfig(infrastructure.AwsInfraConfig)}
	}

	var storageConfig *model.ServiceStorageConfig
	if infrastructure.StorageConfig != storageConfig {
		parsedInfrastructure["storage_config"] = []interface{}{parseServiceStorageConfig(infrastructure.StorageConfig)}
	}

	var archiveStorageConfig *model.ServiceStorageConfig
	if infrastructure.ArchiveStorageConfig != archiveStorageConfig {
		parsedInfrastructure["archive_storage_config"] = []interface{}{parseServiceStorageConfig(infrastructure.ArchiveStorageConfig)}
	}

	return parsedInfrastructure
}

func parseCloudRegionInfoList(cloudRegionInfo *[]model.CloudRegionInfo) []interface{} {
	if cloudRegionInfo == nil {
		return nil
	}
	cloudRegionInfoList := make([]interface{}, 0)

	if cloudRegionInfo != nil {
		cloudRegionInfoList = make([]interface{}, len(*cloudRegionInfo))
		for i, cloudRegionInfoItem := range *cloudRegionInfo {
			cloudRegionInfoList[i] = parseCloudRegionInfo(&cloudRegionInfoItem)
		}
	}

	return cloudRegionInfoList
}

func parseCloudRegionInfo(cloudRegionInfo *model.CloudRegionInfo) interface{} {
	if cloudRegionInfo == nil {
		return nil
	}
	parsedCloudRegionInfo := make(map[string]interface{})
	parsedCloudRegionInfo["cloud"] = cloudRegionInfo.Cloud

	var regions *[]model.RegionInfo
	if cloudRegionInfo.Regions != regions {
		parsedCloudRegionInfo["regions"] = parseRegionInfoList(cloudRegionInfo.Regions)
	}

	return parsedCloudRegionInfo
}

func parseRegionInfoList(regionInfo *[]model.RegionInfo) []interface{} {
	if regionInfo == nil {
		return nil
	}
	regionInfoList := make([]interface{}, 0)

	if regionInfo != nil {
		regionInfoList = make([]interface{}, len(*regionInfo))
		for i, regionInfoItem := range *regionInfo {
			regionInfoList[i] = parseRegionInfo(&regionInfoItem)
		}
	}

	return regionInfoList
}

func parseRegionInfo(regionInfo *model.RegionInfo) interface{} {
	if regionInfo == nil {
		return nil
	}
	parsedRegionInfo := make(map[string]interface{})
	parsedRegionInfo["region"] = regionInfo.Region
	parsedRegionInfo["availability_zones"] = regionInfo.AvailabilityZones

	return parsedRegionInfo
}

func parseAwsInfraConfig(awsInfraConfig *model.AwsInfraConfig) interface{} {
	if awsInfraConfig == nil {
		return nil
	}
	parsedAwsInfraConfig := make(map[string]interface{})

	var awsCpuOptions *model.AwsCpuOptions
	if awsInfraConfig.AwsCpuOptions != awsCpuOptions {
		parsedAwsInfraConfig["aws_cpu_options"] = []interface{}{parseAwsCpuOptions(awsInfraConfig.AwsCpuOptions)}
	}

	return parsedAwsInfraConfig
}

func parseAwsCpuOptions(awsCpuOptions *model.AwsCpuOptions) interface{} {
	if awsCpuOptions == nil {
		return nil
	}
	parsedAwsCpuOptions := make(map[string]interface{})
	parsedAwsCpuOptions["vcpus"] = awsCpuOptions.Vcpus

	return parsedAwsCpuOptions
}

func parseServiceStorageConfig(serviceStorageConfig *model.ServiceStorageConfig) interface{} {
	if serviceStorageConfig == nil {
		return nil
	}
	parsedServiceStorageConfig := make(map[string]interface{})
	parsedServiceStorageConfig["provider"] = serviceStorageConfig.Provider

	var azureNetAppConfig *model.ServiceAzureNetAppConfig
	if serviceStorageConfig.AzureNetAppConfig != azureNetAppConfig {
		parsedServiceStorageConfig["azure_net_app_config"] = []interface{}{parseServiceAzureNetAppConfig(serviceStorageConfig.AzureNetAppConfig)}
	}

	return parsedServiceStorageConfig
}

func parseServiceAzureNetAppConfig(serviceAzureNetAppConfig *model.ServiceAzureNetAppConfig) interface{} {
	if serviceAzureNetAppConfig == nil {
		return nil
	}
	parsedServiceAzureNetAppConfig := make(map[string]interface{})
	parsedServiceAzureNetAppConfig["service_level"] = serviceAzureNetAppConfig.ServiceLevel

	return parsedServiceAzureNetAppConfig
}

func parseTessellServiceMaintenanceWindowWithResData(maintenanceWindow *model.TessellServiceMaintenanceWindow, d *schema.ResourceData) []interface{} {
	if maintenanceWindow == nil {
		return nil
	}
	parsedMaintenanceWindow := make(map[string]interface{})
	if d.Get("maintenance_window") != nil {
		maintenanceWindowResourceData := d.Get("maintenance_window").([]interface{})
		if len(maintenanceWindowResourceData) > 0 {
			parsedMaintenanceWindow = (maintenanceWindowResourceData[0]).(map[string]interface{})
		}
	}
	parsedMaintenanceWindow["day"] = maintenanceWindow.Day
	parsedMaintenanceWindow["time"] = maintenanceWindow.Time
	parsedMaintenanceWindow["duration"] = maintenanceWindow.Duration

	return []interface{}{parsedMaintenanceWindow}
}

func parseTessellServiceMaintenanceWindow(maintenanceWindow *model.TessellServiceMaintenanceWindow) interface{} {
	if maintenanceWindow == nil {
		return nil
	}
	parsedMaintenanceWindow := make(map[string]interface{})
	parsedMaintenanceWindow["day"] = maintenanceWindow.Day
	parsedMaintenanceWindow["time"] = maintenanceWindow.Time
	parsedMaintenanceWindow["duration"] = maintenanceWindow.Duration

	return parsedMaintenanceWindow
}

func parseTessellServiceEngineInfoWithResData(engineConfiguration *model.TessellServiceEngineInfo, d *schema.ResourceData) []interface{} {
	if engineConfiguration == nil {
		return nil
	}
	parsedEngineConfiguration := make(map[string]interface{})
	if d.Get("engine_configuration") != nil {
		engineConfigurationResourceData := d.Get("engine_configuration").([]interface{})
		if len(engineConfigurationResourceData) > 0 {
			parsedEngineConfiguration = (engineConfigurationResourceData[0]).(map[string]interface{})
		}
	}

	parsedEngineConfiguration["backup_url"] = engineConfiguration.BackupUrl
	parsedEngineConfiguration["ignore_post_script_failure"] = engineConfiguration.IgnorePostScriptFailure
	parsedEngineConfiguration["ignore_pre_script_failure"] = engineConfiguration.IgnorePreScriptFailure

	var oracleConfig *model.TessellServiceOracleEngineConfig
	if engineConfiguration.OracleConfig != oracleConfig {
		parsedEngineConfiguration["oracle_config"] = []interface{}{parseTessellServiceOracleEngineConfig(engineConfiguration.OracleConfig)}
	}

	var postgresqlConfig *model.TessellServicePostgresqlEngineConfig
	if engineConfiguration.PostgresqlConfig != postgresqlConfig {
		parsedEngineConfiguration["postgresql_config"] = []interface{}{parseTessellServicePostgresqlEngineConfig(engineConfiguration.PostgresqlConfig)}
	}

	var mysqlConfig *model.TessellServiceMysqlEngineConfig
	if engineConfiguration.MysqlConfig != mysqlConfig {
		parsedEngineConfiguration["mysql_config"] = []interface{}{parseTessellServiceMysqlEngineConfig(engineConfiguration.MysqlConfig)}
	}

	var sqlServerConfig *model.TessellServiceSqlServerEngineConfig
	if engineConfiguration.SqlServerConfig != sqlServerConfig {
		parsedEngineConfiguration["sql_server_config"] = []interface{}{parseTessellServiceSqlServerEngineConfig(engineConfiguration.SqlServerConfig)}
	}

	var apacheKafkaConfig *model.TessellServiceApacheKafkaEngineConfig
	if engineConfiguration.ApacheKafkaConfig != apacheKafkaConfig {
		parsedEngineConfiguration["apache_kafka_config"] = []interface{}{parseTessellServiceApacheKafkaEngineConfig(engineConfiguration.ApacheKafkaConfig)}
	}

	var mongoDBConfig *model.TessellServiceMongoDBEngineConfig
	if engineConfiguration.MongoDBConfig != mongoDBConfig {
		parsedEngineConfiguration["mongodb_config"] = []interface{}{parseTessellServiceMongoDBEngineConfig(engineConfiguration.MongoDBConfig)}
	}

	var milvusConfig *model.TessellServiceMilvusEngineConfig
	if engineConfiguration.MilvusConfig != milvusConfig {
		parsedEngineConfiguration["milvus_config"] = []interface{}{parseTessellServiceMilvusEngineConfig(engineConfiguration.MilvusConfig)}
	}

	var preScriptInfo *model.ScriptInfo
	if engineConfiguration.PreScriptInfo != preScriptInfo {
		parsedEngineConfiguration["pre_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PreScriptInfo)}
	}

	var postScriptInfo *model.ScriptInfo
	if engineConfiguration.PostScriptInfo != postScriptInfo {
		parsedEngineConfiguration["post_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PostScriptInfo)}
	}

	var collationConfig *model.DBEngineCollationConfig
	if engineConfiguration.CollationConfig != collationConfig {
		parsedEngineConfiguration["collation_config"] = []interface{}{parseDBEngineCollationConfig(engineConfiguration.CollationConfig)}
	}

	return []interface{}{parsedEngineConfiguration}
}

func parseTessellServiceEngineInfo(engineConfiguration *model.TessellServiceEngineInfo) interface{} {
	if engineConfiguration == nil {
		return nil
	}
	parsedEngineConfiguration := make(map[string]interface{})

	parsedEngineConfiguration["backup_url"] = engineConfiguration.BackupUrl
	parsedEngineConfiguration["ignore_post_script_failure"] = engineConfiguration.IgnorePostScriptFailure
	parsedEngineConfiguration["ignore_pre_script_failure"] = engineConfiguration.IgnorePreScriptFailure

	var oracleConfig *model.TessellServiceOracleEngineConfig
	if engineConfiguration.OracleConfig != oracleConfig {
		parsedEngineConfiguration["oracle_config"] = []interface{}{parseTessellServiceOracleEngineConfig(engineConfiguration.OracleConfig)}
	}

	var postgresqlConfig *model.TessellServicePostgresqlEngineConfig
	if engineConfiguration.PostgresqlConfig != postgresqlConfig {
		parsedEngineConfiguration["postgresql_config"] = []interface{}{parseTessellServicePostgresqlEngineConfig(engineConfiguration.PostgresqlConfig)}
	}

	var mysqlConfig *model.TessellServiceMysqlEngineConfig
	if engineConfiguration.MysqlConfig != mysqlConfig {
		parsedEngineConfiguration["mysql_config"] = []interface{}{parseTessellServiceMysqlEngineConfig(engineConfiguration.MysqlConfig)}
	}

	var sqlServerConfig *model.TessellServiceSqlServerEngineConfig
	if engineConfiguration.SqlServerConfig != sqlServerConfig {
		parsedEngineConfiguration["sql_server_config"] = []interface{}{parseTessellServiceSqlServerEngineConfig(engineConfiguration.SqlServerConfig)}
	}

	var apacheKafkaConfig *model.TessellServiceApacheKafkaEngineConfig
	if engineConfiguration.ApacheKafkaConfig != apacheKafkaConfig {
		parsedEngineConfiguration["apache_kafka_config"] = []interface{}{parseTessellServiceApacheKafkaEngineConfig(engineConfiguration.ApacheKafkaConfig)}
	}

	var mongoDBConfig *model.TessellServiceMongoDBEngineConfig
	if engineConfiguration.MongoDBConfig != mongoDBConfig {
		parsedEngineConfiguration["mongodb_config"] = []interface{}{parseTessellServiceMongoDBEngineConfig(engineConfiguration.MongoDBConfig)}
	}

	var milvusConfig *model.TessellServiceMilvusEngineConfig
	if engineConfiguration.MilvusConfig != milvusConfig {
		parsedEngineConfiguration["milvus_config"] = []interface{}{parseTessellServiceMilvusEngineConfig(engineConfiguration.MilvusConfig)}
	}

	var preScriptInfo *model.ScriptInfo
	if engineConfiguration.PreScriptInfo != preScriptInfo {
		parsedEngineConfiguration["pre_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PreScriptInfo)}
	}

	var postScriptInfo *model.ScriptInfo
	if engineConfiguration.PostScriptInfo != postScriptInfo {
		parsedEngineConfiguration["post_script_info"] = []interface{}{parseScriptInfo(engineConfiguration.PostScriptInfo)}
	}

	var collationConfig *model.DBEngineCollationConfig
	if engineConfiguration.CollationConfig != collationConfig {
		parsedEngineConfiguration["collation_config"] = []interface{}{parseDBEngineCollationConfig(engineConfiguration.CollationConfig)}
	}

	return parsedEngineConfiguration
}

func parseTessellServiceOracleEngineConfig(tessellServiceOracleEngineConfig *model.TessellServiceOracleEngineConfig) interface{} {
	if tessellServiceOracleEngineConfig == nil {
		return nil
	}
	parsedTessellServiceOracleEngineConfig := make(map[string]interface{})
	parsedTessellServiceOracleEngineConfig["multi_tenant"] = tessellServiceOracleEngineConfig.MultiTenant
	parsedTessellServiceOracleEngineConfig["parameter_profile_id"] = tessellServiceOracleEngineConfig.ParameterProfileId
	parsedTessellServiceOracleEngineConfig["options_profile"] = tessellServiceOracleEngineConfig.OptionsProfile
	parsedTessellServiceOracleEngineConfig["option_profile_id"] = tessellServiceOracleEngineConfig.OptionProfileId
	parsedTessellServiceOracleEngineConfig["sid"] = tessellServiceOracleEngineConfig.Sid
	parsedTessellServiceOracleEngineConfig["character_set"] = tessellServiceOracleEngineConfig.CharacterSet
	parsedTessellServiceOracleEngineConfig["national_character_set"] = tessellServiceOracleEngineConfig.NationalCharacterSet
	parsedTessellServiceOracleEngineConfig["enable_archive_mode"] = tessellServiceOracleEngineConfig.EnableArchiveMode

	return parsedTessellServiceOracleEngineConfig
}

func parseTessellServicePostgresqlEngineConfig(tessellServicePostgresqlEngineConfig *model.TessellServicePostgresqlEngineConfig) interface{} {
	if tessellServicePostgresqlEngineConfig == nil {
		return nil
	}
	parsedTessellServicePostgresqlEngineConfig := make(map[string]interface{})
	parsedTessellServicePostgresqlEngineConfig["parameter_profile_id"] = tessellServicePostgresqlEngineConfig.ParameterProfileId
	parsedTessellServicePostgresqlEngineConfig["ad_domain_id"] = tessellServicePostgresqlEngineConfig.AdDomainId
	parsedTessellServicePostgresqlEngineConfig["proxy_port"] = tessellServicePostgresqlEngineConfig.ProxyPort
	parsedTessellServicePostgresqlEngineConfig["option_profile_name"] = tessellServicePostgresqlEngineConfig.OptionProfileName
	parsedTessellServicePostgresqlEngineConfig["option_profile_id"] = tessellServicePostgresqlEngineConfig.OptionProfileId

	return parsedTessellServicePostgresqlEngineConfig
}

func parseTessellServiceMysqlEngineConfig(tessellServiceMySqlEngineConfig *model.TessellServiceMysqlEngineConfig) interface{} {
	if tessellServiceMySqlEngineConfig == nil {
		return nil
	}
	parsedTessellServiceMySqlEngineConfig := make(map[string]interface{})
	parsedTessellServiceMySqlEngineConfig["parameter_profile_id"] = tessellServiceMySqlEngineConfig.ParameterProfileId
	parsedTessellServiceMySqlEngineConfig["ad_domain_id"] = tessellServiceMySqlEngineConfig.AdDomainId
	parsedTessellServiceMySqlEngineConfig["option_profile_id"] = tessellServiceMySqlEngineConfig.OptionProfileId

	return parsedTessellServiceMySqlEngineConfig
}

func parseTessellServiceSqlServerEngineConfig(tessellServiceSqlServerEngineConfig *model.TessellServiceSqlServerEngineConfig) interface{} {
	if tessellServiceSqlServerEngineConfig == nil {
		return nil
	}
	parsedTessellServiceSqlServerEngineConfig := make(map[string]interface{})
	parsedTessellServiceSqlServerEngineConfig["parameter_profile_id"] = tessellServiceSqlServerEngineConfig.ParameterProfileId
	parsedTessellServiceSqlServerEngineConfig["ad_domain_id"] = tessellServiceSqlServerEngineConfig.AdDomainId
	parsedTessellServiceSqlServerEngineConfig["service_account_user"] = tessellServiceSqlServerEngineConfig.ServiceAccountUser
	parsedTessellServiceSqlServerEngineConfig["agent_service_account_user"] = tessellServiceSqlServerEngineConfig.AgentServiceAccountUser
	parsedTessellServiceSqlServerEngineConfig["instance_name"] = tessellServiceSqlServerEngineConfig.InstanceName

	return parsedTessellServiceSqlServerEngineConfig
}

func parseTessellServiceApacheKafkaEngineConfig(tessellServiceApacheKafkaEngineConfig *model.TessellServiceApacheKafkaEngineConfig) interface{} {
	if tessellServiceApacheKafkaEngineConfig == nil {
		return nil
	}
	parsedTessellServiceApacheKafkaEngineConfig := make(map[string]interface{})
	parsedTessellServiceApacheKafkaEngineConfig["parameter_profile_id"] = tessellServiceApacheKafkaEngineConfig.ParameterProfileId

	return parsedTessellServiceApacheKafkaEngineConfig
}

func parseTessellServiceMongoDBEngineConfig(tessellServiceMongodbEngineConfig *model.TessellServiceMongoDBEngineConfig) interface{} {
	if tessellServiceMongodbEngineConfig == nil {
		return nil
	}
	parsedTessellServiceMongodbEngineConfig := make(map[string]interface{})
	parsedTessellServiceMongodbEngineConfig["cluster_name"] = tessellServiceMongodbEngineConfig.ClusterName
	parsedTessellServiceMongodbEngineConfig["parameter_profile_id"] = tessellServiceMongodbEngineConfig.ParameterProfileId

	return parsedTessellServiceMongodbEngineConfig
}

func parseTessellServiceMilvusEngineConfig(tessellServiceMilvusEngineConfig *model.TessellServiceMilvusEngineConfig) interface{} {
	if tessellServiceMilvusEngineConfig == nil {
		return nil
	}
	parsedTessellServiceMilvusEngineConfig := make(map[string]interface{})
	parsedTessellServiceMilvusEngineConfig["parameter_profile_id"] = tessellServiceMilvusEngineConfig.ParameterProfileId

	return parsedTessellServiceMilvusEngineConfig
}

func parseDBEngineCollationConfig(dbEngineCollationConfig *model.DBEngineCollationConfig) interface{} {
	if dbEngineCollationConfig == nil {
		return nil
	}
	parsedDbEngineCollationConfig := make(map[string]interface{})
	parsedDbEngineCollationConfig["collation_name"] = dbEngineCollationConfig.CollationName

	return parsedDbEngineCollationConfig
}

func parseTessellServiceIntegrationsInfoWithResData(integrationsConfig *model.TessellServiceIntegrationsInfo, d *schema.ResourceData) []interface{} {
	if integrationsConfig == nil {
		return nil
	}
	parsedIntegrationsConfig := make(map[string]interface{})
	if d.Get("integrations_config") != nil {
		integrationsConfigResourceData := d.Get("integrations_config").([]interface{})
		if len(integrationsConfigResourceData) > 0 {
			parsedIntegrationsConfig = (integrationsConfigResourceData[0]).(map[string]interface{})
		}
	}
	parsedIntegrationsConfig["integrations"] = integrationsConfig.Integrations

	return []interface{}{parsedIntegrationsConfig}
}

func parseTessellServiceIntegrationsInfo(integrationsConfig *model.TessellServiceIntegrationsInfo) interface{} {
	if integrationsConfig == nil {
		return nil
	}
	parsedIntegrationsConfig := make(map[string]interface{})
	parsedIntegrationsConfig["integrations"] = integrationsConfig.Integrations

	return parsedIntegrationsConfig
}

func parseTessellServiceDeletionConfigWithResData(deletionConfig *model.TessellServiceDeletionConfig, d *schema.ResourceData) []interface{} {
	if deletionConfig == nil {
		return nil
	}
	parsedDeletionConfig := make(map[string]interface{})
	if d.Get("deletion_config") != nil {
		deletionConfigResourceData := d.Get("deletion_config").([]interface{})
		if len(deletionConfigResourceData) > 0 {
			parsedDeletionConfig = (deletionConfigResourceData[0]).(map[string]interface{})
		}
	}
	parsedDeletionConfig["retain_availability_machine"] = deletionConfig.RetainAvailabilityMachine

	return []interface{}{parsedDeletionConfig}
}

func parseTessellServiceDeletionConfig(deletionConfig *model.TessellServiceDeletionConfig) interface{} {
	if deletionConfig == nil {
		return nil
	}
	parsedDeletionConfig := make(map[string]interface{})
	parsedDeletionConfig["retain_availability_machine"] = deletionConfig.RetainAvailabilityMachine

	return parsedDeletionConfig
}

func parseTessellTagListWithResData(tags *[]model.TessellTag, d *schema.ResourceData) []interface{} {
	if tags == nil {
		return nil
	}
	tessellTagList := make([]interface{}, 0)

	if tags != nil {
		tessellTagList = make([]interface{}, len(*tags))
		for i, tessellTagItem := range *tags {
			tessellTagList[i] = parseTessellTag(&tessellTagItem)
		}
	}

	return tessellTagList
}

func parseTessellTagList(tags *[]model.TessellTag) []interface{} {
	if tags == nil {
		return nil
	}
	tessellTagList := make([]interface{}, 0)

	if tags != nil {
		tessellTagList = make([]interface{}, len(*tags))
		for i, tessellTagItem := range *tags {
			tessellTagList[i] = parseTessellTag(&tessellTagItem)
		}
	}

	return tessellTagList
}

func parseTessellTag(tags *model.TessellTag) interface{} {
	if tags == nil {
		return nil
	}
	parsedTags := make(map[string]interface{})
	parsedTags["name"] = tags.Name
	parsedTags["value"] = tags.Value

	return parsedTags
}

func parseTessellResourceUpdateInfoListWithResData(updatesInProgress *[]model.TessellResourceUpdateInfo, d *schema.ResourceData) []interface{} {
	if updatesInProgress == nil {
		return nil
	}
	tessellResourceUpdateInfoList := make([]interface{}, 0)

	if updatesInProgress != nil {
		tessellResourceUpdateInfoList = make([]interface{}, len(*updatesInProgress))
		for i, tessellResourceUpdateInfoItem := range *updatesInProgress {
			tessellResourceUpdateInfoList[i] = parseTessellResourceUpdateInfo(&tessellResourceUpdateInfoItem)
		}
	}

	return tessellResourceUpdateInfoList
}

func parseTessellResourceUpdateInfoList(updatesInProgress *[]model.TessellResourceUpdateInfo) []interface{} {
	if updatesInProgress == nil {
		return nil
	}
	tessellResourceUpdateInfoList := make([]interface{}, 0)

	if updatesInProgress != nil {
		tessellResourceUpdateInfoList = make([]interface{}, len(*updatesInProgress))
		for i, tessellResourceUpdateInfoItem := range *updatesInProgress {
			tessellResourceUpdateInfoList[i] = parseTessellResourceUpdateInfo(&tessellResourceUpdateInfoItem)
		}
	}

	return tessellResourceUpdateInfoList
}

func parseTessellResourceUpdateInfo(updatesInProgress *model.TessellResourceUpdateInfo) interface{} {
	if updatesInProgress == nil {
		return nil
	}
	parsedUpdatesInProgress := make(map[string]interface{})
	parsedUpdatesInProgress["update_type"] = updatesInProgress.UpdateType
	parsedUpdatesInProgress["reference_id"] = updatesInProgress.ReferenceId
	parsedUpdatesInProgress["submitted_at"] = updatesInProgress.SubmittedAt
	parsedUpdatesInProgress["update_info"] = updatesInProgress.UpdateInfo

	return parsedUpdatesInProgress
}

func parseTessellServiceInstanceDTOListWithResData(instances *[]model.TessellServiceInstanceDTO, d *schema.ResourceData) []interface{} {
	if instances == nil {
		return nil
	}
	tessellServiceInstanceDTOList := make([]interface{}, 0)

	if instances != nil {
		tessellServiceInstanceDTOList = make([]interface{}, len(*instances))
		for i, tessellServiceInstanceDTOItem := range *instances {
			tessellServiceInstanceDTOList[i] = parseTessellServiceInstanceDTO(&tessellServiceInstanceDTOItem)
		}
	}

	return tessellServiceInstanceDTOList
}

func parseTessellServiceInstanceDTOList(instances *[]model.TessellServiceInstanceDTO) []interface{} {
	if instances == nil {
		return nil
	}
	tessellServiceInstanceDTOList := make([]interface{}, 0)

	if instances != nil {
		tessellServiceInstanceDTOList = make([]interface{}, len(*instances))
		for i, tessellServiceInstanceDTOItem := range *instances {
			tessellServiceInstanceDTOList[i] = parseTessellServiceInstanceDTO(&tessellServiceInstanceDTOItem)
		}
	}

	return tessellServiceInstanceDTOList
}

func parseTessellServiceInstanceDTO(instances *model.TessellServiceInstanceDTO) interface{} {
	if instances == nil {
		return nil
	}
	parsedInstances := make(map[string]interface{})
	parsedInstances["id"] = instances.Id
	parsedInstances["name"] = instances.Name
	parsedInstances["instance_group_name"] = instances.InstanceGroupName
	parsedInstances["type"] = instances.Type
	parsedInstances["role"] = instances.Role
	parsedInstances["status"] = instances.Status
	parsedInstances["tessell_service_id"] = instances.TessellServiceId
	parsedInstances["cloud"] = instances.Cloud
	parsedInstances["region"] = instances.Region
	parsedInstances["availability_zone"] = instances.AvailabilityZone
	parsedInstances["instance_group_id"] = instances.InstanceGroupId
	parsedInstances["compute_type"] = instances.ComputeType

	parsedInstances["compute_id"] = instances.ComputeId
	parsedInstances["compute_name"] = instances.ComputeName
	parsedInstances["storage"] = instances.Storage
	parsedInstances["data_volume_iops"] = instances.DataVolumeIops
	parsedInstances["throughput"] = instances.Throughput
	parsedInstances["enable_perf_insights"] = instances.EnablePerfInsights

	parsedInstances["vpc"] = instances.VPC
	parsedInstances["public_subnet"] = instances.PublicSubnet
	parsedInstances["private_subnet"] = instances.PrivateSubnet
	parsedInstances["encryption_key"] = instances.EncryptionKey
	parsedInstances["software_image"] = instances.SoftwareImage
	parsedInstances["software_image_version"] = instances.SoftwareImageVersion
	parsedInstances["date_created"] = instances.DateCreated

	parsedInstances["last_started_at"] = instances.LastStartedAt
	parsedInstances["last_stopped_at"] = instances.LastStoppedAt
	parsedInstances["sync_mode"] = instances.SyncMode

	var awsInfraConfig *model.AwsInfraConfig
	if instances.AwsInfraConfig != awsInfraConfig {
		parsedInstances["aws_infra_config"] = []interface{}{parseAwsInfraConfig(instances.AwsInfraConfig)}
	}

	var parameterProfile *model.ParameterProfile
	if instances.ParameterProfile != parameterProfile {
		parsedInstances["parameter_profile"] = []interface{}{parseParameterProfile(instances.ParameterProfile)}
	}

	var optionProfile *model.OptionProfile
	if instances.OptionProfile != optionProfile {
		parsedInstances["option_profile"] = []interface{}{parseOptionProfile(instances.OptionProfile)}
	}

	var monitoringConfig *model.MonitoringConfig
	if instances.MonitoringConfig != monitoringConfig {
		parsedInstances["monitoring_config"] = []interface{}{parseMonitoringConfig(instances.MonitoringConfig)}
	}

	var connectString *model.TessellServiceInstanceConnectString
	if instances.ConnectString != connectString {
		parsedInstances["connect_string"] = []interface{}{parseTessellServiceInstanceConnectString(instances.ConnectString)}
	}

	var updatesInProgress *[]model.TessellResourceUpdateInfo
	if instances.UpdatesInProgress != updatesInProgress {
		parsedInstances["updates_in_progress"] = parseTessellResourceUpdateInfoList(instances.UpdatesInProgress)
	}

	var engineConfiguration *model.ServiceInstanceEngineInfo
	if instances.EngineConfiguration != engineConfiguration {
		parsedInstances["engine_configuration"] = []interface{}{parseServiceInstanceEngineInfo(instances.EngineConfiguration)}
	}

	var computeConfig *model.InstanceComputeConfig
	if instances.ComputeConfig != computeConfig {
		parsedInstances["compute_config"] = []interface{}{parseInstanceComputeConfig(instances.ComputeConfig)}
	}

	var storageConfig *model.InstanceStorageConfig
	if instances.StorageConfig != storageConfig {
		parsedInstances["storage_config"] = []interface{}{parseInstanceStorageConfig(instances.StorageConfig)}
	}

	var archiveStorageConfig *model.InstanceStorageConfig
	if instances.ArchiveStorageConfig != archiveStorageConfig {
		parsedInstances["archive_storage_config"] = []interface{}{parseInstanceStorageConfig(instances.ArchiveStorageConfig)}
	}

	var privateLinkInfo *model.PrivateLinkInfo
	if instances.PrivateLinkInfo != privateLinkInfo {
		parsedInstances["private_link_info"] = []interface{}{parsePrivateLinkInfo(instances.PrivateLinkInfo)}
	}

	return parsedInstances
}

func parseParameterProfile(parameterProfile *model.ParameterProfile) interface{} {
	if parameterProfile == nil {
		return nil
	}
	parsedParameterProfile := make(map[string]interface{})
	parsedParameterProfile["id"] = parameterProfile.Id
	parsedParameterProfile["name"] = parameterProfile.Name
	parsedParameterProfile["version"] = parameterProfile.Version
	parsedParameterProfile["status"] = parameterProfile.Status

	return parsedParameterProfile
}

func parseOptionProfile(optionProfile *model.OptionProfile) interface{} {
	if optionProfile == nil {
		return nil
	}
	parsedOptionProfile := make(map[string]interface{})
	parsedOptionProfile["id"] = optionProfile.Id
	parsedOptionProfile["name"] = optionProfile.Name
	parsedOptionProfile["version"] = optionProfile.Version
	parsedOptionProfile["status"] = optionProfile.Status

	return parsedOptionProfile
}

func parseMonitoringConfig(monitoringConfig *model.MonitoringConfig) interface{} {
	if monitoringConfig == nil {
		return nil
	}
	parsedMonitoringConfig := make(map[string]interface{})

	var perfInsights *model.PerfInsightsConfig
	if monitoringConfig.PerfInsights != perfInsights {
		parsedMonitoringConfig["perf_insights"] = []interface{}{parsePerfInsightsConfig(monitoringConfig.PerfInsights)}
	}

	return parsedMonitoringConfig
}

func parsePerfInsightsConfig(perfInsightsConfig *model.PerfInsightsConfig) interface{} {
	if perfInsightsConfig == nil {
		return nil
	}
	parsedPerfInsightsConfig := make(map[string]interface{})
	parsedPerfInsightsConfig["perf_insights_enabled"] = perfInsightsConfig.PerfInsightsEnabled
	parsedPerfInsightsConfig["monitoring_deployment_id"] = perfInsightsConfig.MonitoringDeploymentId
	parsedPerfInsightsConfig["status"] = perfInsightsConfig.Status

	return parsedPerfInsightsConfig
}

func parseTessellServiceInstanceConnectString(tessellServiceInstanceConnectString *model.TessellServiceInstanceConnectString) interface{} {
	if tessellServiceInstanceConnectString == nil {
		return nil
	}
	parsedTessellServiceInstanceConnectString := make(map[string]interface{})
	parsedTessellServiceInstanceConnectString["connect_descriptor"] = tessellServiceInstanceConnectString.ConnectDescriptor
	parsedTessellServiceInstanceConnectString["master_user"] = tessellServiceInstanceConnectString.MasterUser
	parsedTessellServiceInstanceConnectString["endpoint"] = tessellServiceInstanceConnectString.Endpoint
	parsedTessellServiceInstanceConnectString["service_port"] = tessellServiceInstanceConnectString.ServicePort

	return parsedTessellServiceInstanceConnectString
}

func parseServiceInstanceEngineInfo(engineConfiguration *model.ServiceInstanceEngineInfo) interface{} {
	if engineConfiguration == nil {
		return nil
	}
	parsedEngineConfiguration := make(map[string]interface{})

	var oracleConfig *model.ServiceInstanceOracleEngineConfig
	if engineConfiguration.OracleConfig != oracleConfig {
		parsedEngineConfiguration["oracle_config"] = []interface{}{parseServiceInstanceOracleEngineConfig(engineConfiguration.OracleConfig)}
	}

	return parsedEngineConfiguration
}

func parseServiceInstanceOracleEngineConfig(serviceInstanceOracleEngineConfig *model.ServiceInstanceOracleEngineConfig) interface{} {
	if serviceInstanceOracleEngineConfig == nil {
		return nil
	}
	parsedServiceInstanceOracleEngineConfig := make(map[string]interface{})
	parsedServiceInstanceOracleEngineConfig["access_mode"] = serviceInstanceOracleEngineConfig.AccessMode

	return parsedServiceInstanceOracleEngineConfig
}

func parseInstanceComputeConfig(instanceComputeConfig *model.InstanceComputeConfig) interface{} {
	if instanceComputeConfig == nil {
		return nil
	}
	parsedInstanceComputeConfig := make(map[string]interface{})
	parsedInstanceComputeConfig["provider"] = instanceComputeConfig.Provider

	var exadataConfig *model.InstanceExadataComputeConfig
	if instanceComputeConfig.ExadataConfig != exadataConfig {
		parsedInstanceComputeConfig["exadata_config"] = []interface{}{parseInstanceExadataComputeConfig(instanceComputeConfig.ExadataConfig)}
	}

	return parsedInstanceComputeConfig
}

func parseInstanceExadataComputeConfig(instanceExadataComputeConfig *model.InstanceExadataComputeConfig) interface{} {
	if instanceExadataComputeConfig == nil {
		return nil
	}
	parsedInstanceExadataComputeConfig := make(map[string]interface{})
	parsedInstanceExadataComputeConfig["infrastructure_id"] = instanceExadataComputeConfig.InfrastructureId
	parsedInstanceExadataComputeConfig["infrastructure_name"] = instanceExadataComputeConfig.InfrastructureName
	parsedInstanceExadataComputeConfig["vm_cluster_id"] = instanceExadataComputeConfig.VmClusterId
	parsedInstanceExadataComputeConfig["vm_cluster_name"] = instanceExadataComputeConfig.VmClusterName
	parsedInstanceExadataComputeConfig["vcpus"] = instanceExadataComputeConfig.Vcpus
	parsedInstanceExadataComputeConfig["memory"] = instanceExadataComputeConfig.Memory

	return parsedInstanceExadataComputeConfig
}

func parseInstanceStorageConfig(instanceStorageConfig *model.InstanceStorageConfig) interface{} {
	if instanceStorageConfig == nil {
		return nil
	}
	parsedInstanceStorageConfig := make(map[string]interface{})
	parsedInstanceStorageConfig["provider"] = instanceStorageConfig.Provider
	parsedInstanceStorageConfig["volume_type"] = instanceStorageConfig.VolumeType

	var fsxNetAppConfig *model.InstanceFsxNetAppConfig
	if instanceStorageConfig.FsxNetAppConfig != fsxNetAppConfig {
		parsedInstanceStorageConfig["fsx_net_app_config"] = []interface{}{parseInstanceFsxNetAppConfig(instanceStorageConfig.FsxNetAppConfig)}
	}

	var azureNetAppConfig *model.InstanceAzureNetAppConfig
	if instanceStorageConfig.AzureNetAppConfig != azureNetAppConfig {
		parsedInstanceStorageConfig["azure_net_app_config"] = []interface{}{parseInstanceAzureNetAppConfig(instanceStorageConfig.AzureNetAppConfig)}
	}

	return parsedInstanceStorageConfig
}

func parseInstanceFsxNetAppConfig(instanceFsxNetAppConfig *model.InstanceFsxNetAppConfig) interface{} {
	if instanceFsxNetAppConfig == nil {
		return nil
	}
	parsedInstanceFsxNetAppConfig := make(map[string]interface{})
	parsedInstanceFsxNetAppConfig["file_system_name"] = instanceFsxNetAppConfig.FileSystemName
	parsedInstanceFsxNetAppConfig["svm_name"] = instanceFsxNetAppConfig.SvmName
	parsedInstanceFsxNetAppConfig["volume_name"] = instanceFsxNetAppConfig.VolumeName
	parsedInstanceFsxNetAppConfig["file_system_id"] = instanceFsxNetAppConfig.FileSystemId
	parsedInstanceFsxNetAppConfig["svm_id"] = instanceFsxNetAppConfig.SvmId

	return parsedInstanceFsxNetAppConfig
}

func parseInstanceAzureNetAppConfig(instanceAzureNetAppConfig *model.InstanceAzureNetAppConfig) interface{} {
	if instanceAzureNetAppConfig == nil {
		return nil
	}
	parsedInstanceAzureNetAppConfig := make(map[string]interface{})
	parsedInstanceAzureNetAppConfig["azure_net_app_name"] = instanceAzureNetAppConfig.AzureNetAppName
	parsedInstanceAzureNetAppConfig["capacity_pool_name"] = instanceAzureNetAppConfig.CapacityPoolName
	parsedInstanceAzureNetAppConfig["volume_name"] = instanceAzureNetAppConfig.VolumeName
	parsedInstanceAzureNetAppConfig["azure_net_app_id"] = instanceAzureNetAppConfig.AzureNetAppId
	parsedInstanceAzureNetAppConfig["capacity_pool_id"] = instanceAzureNetAppConfig.CapacityPoolId
	parsedInstanceAzureNetAppConfig["delegated_subnet_id"] = instanceAzureNetAppConfig.DelegatedSubnetId
	parsedInstanceAzureNetAppConfig["delegated_subnet_name"] = instanceAzureNetAppConfig.DelegatedSubnetName

	parsedInstanceAzureNetAppConfig["network_features"] = instanceAzureNetAppConfig.NetworkFeatures
	parsedInstanceAzureNetAppConfig["service_level"] = instanceAzureNetAppConfig.ServiceLevel

	var encryptionKeyInfo *model.AzureNetAppEncryptionKeyInfo
	if instanceAzureNetAppConfig.EncryptionKeyInfo != encryptionKeyInfo {
		parsedInstanceAzureNetAppConfig["encryption_key_info"] = []interface{}{parseAzureNetAppEncryptionKeyInfo(instanceAzureNetAppConfig.EncryptionKeyInfo)}
	}

	return parsedInstanceAzureNetAppConfig
}

func parseAzureNetAppEncryptionKeyInfo(azureNetAppEncryptionKeyInfo *model.AzureNetAppEncryptionKeyInfo) interface{} {
	if azureNetAppEncryptionKeyInfo == nil {
		return nil
	}
	parsedAzureNetAppEncryptionKeyInfo := make(map[string]interface{})
	parsedAzureNetAppEncryptionKeyInfo["id"] = azureNetAppEncryptionKeyInfo.Id
	parsedAzureNetAppEncryptionKeyInfo["name"] = azureNetAppEncryptionKeyInfo.Name
	parsedAzureNetAppEncryptionKeyInfo["key_vault_cloud_resource_id"] = azureNetAppEncryptionKeyInfo.KeyVaultCloudResourceId
	parsedAzureNetAppEncryptionKeyInfo["key_source"] = azureNetAppEncryptionKeyInfo.KeySource

	return parsedAzureNetAppEncryptionKeyInfo
}

func parsePrivateLinkInfo(privateLinkInfo *model.PrivateLinkInfo) interface{} {
	if privateLinkInfo == nil {
		return nil
	}
	parsedPrivateLinkInfo := make(map[string]interface{})
	parsedPrivateLinkInfo["id"] = privateLinkInfo.Id
	parsedPrivateLinkInfo["status"] = privateLinkInfo.Status
	parsedPrivateLinkInfo["endpoint_service_name"] = privateLinkInfo.EndpointServiceName
	parsedPrivateLinkInfo["private_link_service_alias"] = privateLinkInfo.PrivateLinkServiceAlias
	parsedPrivateLinkInfo["service_principals"] = privateLinkInfo.ServicePrincipals
	parsedPrivateLinkInfo["client_azure_subscription_ids"] = privateLinkInfo.ClientAzureSubscriptionIds

	return parsedPrivateLinkInfo
}

func parseTessellDatabaseDTOListWithResData(databases *[]model.TessellDatabaseDTO, d *schema.ResourceData) []interface{} {
	if databases == nil {
		return nil
	}
	tessellDatabaseDTOList := make([]interface{}, 0)

	if databases != nil {
		tessellDatabaseDTOList = make([]interface{}, len(*databases))
		for i, tessellDatabaseDTOItem := range *databases {
			tessellDatabaseDTOList[i] = parseTessellDatabaseDTO(&tessellDatabaseDTOItem)
		}
	}

	return tessellDatabaseDTOList
}

func parseTessellDatabaseDTOList(databases *[]model.TessellDatabaseDTO) []interface{} {
	if databases == nil {
		return nil
	}
	tessellDatabaseDTOList := make([]interface{}, 0)

	if databases != nil {
		tessellDatabaseDTOList = make([]interface{}, len(*databases))
		for i, tessellDatabaseDTOItem := range *databases {
			tessellDatabaseDTOList[i] = parseTessellDatabaseDTO(&tessellDatabaseDTOItem)
		}
	}

	return tessellDatabaseDTOList
}

func parseTessellDatabaseDTO(databases *model.TessellDatabaseDTO) interface{} {
	if databases == nil {
		return nil
	}
	parsedDatabases := make(map[string]interface{})
	parsedDatabases["id"] = databases.Id
	parsedDatabases["database_name"] = databases.DatabaseName
	parsedDatabases["description"] = databases.Description
	parsedDatabases["tessell_service_id"] = databases.TessellServiceId
	parsedDatabases["engine_type"] = databases.EngineType
	parsedDatabases["status"] = databases.Status
	parsedDatabases["date_created"] = databases.DateCreated
	parsedDatabases["tessell_created"] = databases.TessellCreated

	var clonedFromInfo *model.TessellDatabaseClonedFromInfo
	if databases.ClonedFromInfo != clonedFromInfo {
		parsedDatabases["cloned_from_info"] = []interface{}{parseTessellDatabaseClonedFromInfo(databases.ClonedFromInfo)}
	}

	var databaseConfiguration *model.DatabaseConfiguration
	if databases.DatabaseConfiguration != databaseConfiguration {
		parsedDatabases["database_configuration"] = []interface{}{parseDatabaseConfiguration(databases.DatabaseConfiguration)}
	}

	var connectString *model.TessellServiceDatabaseConnectString
	if databases.ConnectString != connectString {
		parsedDatabases["connect_string"] = []interface{}{parseTessellServiceDatabaseConnectString(databases.ConnectString)}
	}

	return parsedDatabases
}

func parseTessellDatabaseClonedFromInfo(clonedFromInfo *model.TessellDatabaseClonedFromInfo) interface{} {
	if clonedFromInfo == nil {
		return nil
	}
	parsedClonedFromInfo := make(map[string]interface{})
	parsedClonedFromInfo["database_id"] = clonedFromInfo.DatabaseId

	return parsedClonedFromInfo
}

func parseDatabaseConfiguration(databaseConfiguration *model.DatabaseConfiguration) interface{} {
	if databaseConfiguration == nil {
		return nil
	}
	parsedDatabaseConfiguration := make(map[string]interface{})

	var oracleConfig *model.OracleDatabaseConfig
	if databaseConfiguration.OracleConfig != oracleConfig {
		parsedDatabaseConfiguration["oracle_config"] = []interface{}{parseOracleDatabaseConfig(databaseConfiguration.OracleConfig)}
	}

	var postgresqlConfig *model.PostgresqlDatabaseConfig
	if databaseConfiguration.PostgresqlConfig != postgresqlConfig {
		parsedDatabaseConfiguration["postgresql_config"] = []interface{}{parsePostgresqlDatabaseConfig(databaseConfiguration.PostgresqlConfig)}
	}

	var mysqlConfig *model.MysqlDatabaseConfig
	if databaseConfiguration.MysqlConfig != mysqlConfig {
		parsedDatabaseConfiguration["mysql_config"] = []interface{}{parseMysqlDatabaseConfig(databaseConfiguration.MysqlConfig)}
	}

	var sqlServerConfig *model.SqlServerDatabaseConfig
	if databaseConfiguration.SqlServerConfig != sqlServerConfig {
		parsedDatabaseConfiguration["sql_server_config"] = []interface{}{parseSqlServerDatabaseConfig(databaseConfiguration.SqlServerConfig)}
	}

	var mongoDBConfig *model.MongoDBDatabaseConfig
	if databaseConfiguration.MongoDBConfig != mongoDBConfig {
		parsedDatabaseConfiguration["mongodb_config"] = []interface{}{parseMongoDBDatabaseConfig(databaseConfiguration.MongoDBConfig)}
	}

	var milvusConfig *model.MilvusDatabaseConfig
	if databaseConfiguration.MilvusConfig != milvusConfig {
		parsedDatabaseConfiguration["milvus_config"] = []interface{}{parseMilvusDatabaseConfig(databaseConfiguration.MilvusConfig)}
	}

	return parsedDatabaseConfiguration
}

func parseOracleDatabaseConfig(oracleDatabaseConfig *model.OracleDatabaseConfig) interface{} {
	if oracleDatabaseConfig == nil {
		return nil
	}
	parsedOracleDatabaseConfig := make(map[string]interface{})
	parsedOracleDatabaseConfig["parameter_profile_id"] = oracleDatabaseConfig.ParameterProfileId
	parsedOracleDatabaseConfig["options_profile"] = oracleDatabaseConfig.OptionsProfile
	parsedOracleDatabaseConfig["username"] = oracleDatabaseConfig.Username
	parsedOracleDatabaseConfig["option_profile_id"] = oracleDatabaseConfig.OptionProfileId

	return parsedOracleDatabaseConfig
}

func parsePostgresqlDatabaseConfig(postgresqlDatabaseConfig *model.PostgresqlDatabaseConfig) interface{} {
	if postgresqlDatabaseConfig == nil {
		return nil
	}
	parsedPostgresqlDatabaseConfig := make(map[string]interface{})
	parsedPostgresqlDatabaseConfig["parameter_profile_id"] = postgresqlDatabaseConfig.ParameterProfileId
	parsedPostgresqlDatabaseConfig["option_profile_id"] = postgresqlDatabaseConfig.OptionProfileId

	return parsedPostgresqlDatabaseConfig
}

func parseMysqlDatabaseConfig(mySqlDatabaseConfig *model.MysqlDatabaseConfig) interface{} {
	if mySqlDatabaseConfig == nil {
		return nil
	}
	parsedMySqlDatabaseConfig := make(map[string]interface{})
	parsedMySqlDatabaseConfig["parameter_profile_id"] = mySqlDatabaseConfig.ParameterProfileId
	parsedMySqlDatabaseConfig["option_profile_id"] = mySqlDatabaseConfig.OptionProfileId

	return parsedMySqlDatabaseConfig
}

func parseSqlServerDatabaseConfig(sqlServerDatabaseConfig *model.SqlServerDatabaseConfig) interface{} {
	if sqlServerDatabaseConfig == nil {
		return nil
	}
	parsedSqlServerDatabaseConfig := make(map[string]interface{})
	parsedSqlServerDatabaseConfig["parameter_profile_id"] = sqlServerDatabaseConfig.ParameterProfileId

	return parsedSqlServerDatabaseConfig
}

func parseMongoDBDatabaseConfig(mongodbDatabaseConfig *model.MongoDBDatabaseConfig) interface{} {
	if mongodbDatabaseConfig == nil {
		return nil
	}
	parsedMongodbDatabaseConfig := make(map[string]interface{})
	parsedMongodbDatabaseConfig["parameter_profile_id"] = mongodbDatabaseConfig.ParameterProfileId

	return parsedMongodbDatabaseConfig
}

func parseMilvusDatabaseConfig(milvusDatabaseConfig *model.MilvusDatabaseConfig) interface{} {
	if milvusDatabaseConfig == nil {
		return nil
	}
	parsedMilvusDatabaseConfig := make(map[string]interface{})
	parsedMilvusDatabaseConfig["parameter_profile_id"] = milvusDatabaseConfig.ParameterProfileId

	return parsedMilvusDatabaseConfig
}

func parseTessellServiceDatabaseConnectString(tessellServiceDatabaseConnectString *model.TessellServiceDatabaseConnectString) interface{} {
	if tessellServiceDatabaseConnectString == nil {
		return nil
	}
	parsedTessellServiceDatabaseConnectString := make(map[string]interface{})
	parsedTessellServiceDatabaseConnectString["connect_descriptor"] = tessellServiceDatabaseConnectString.ConnectDescriptor
	parsedTessellServiceDatabaseConnectString["master_user"] = tessellServiceDatabaseConnectString.MasterUser
	parsedTessellServiceDatabaseConnectString["endpoint"] = tessellServiceDatabaseConnectString.Endpoint
	parsedTessellServiceDatabaseConnectString["service_port"] = tessellServiceDatabaseConnectString.ServicePort

	return parsedTessellServiceDatabaseConnectString
}

func parseEntityAclSharingInfoWithResData(sharedWith *model.EntityAclSharingInfo, d *schema.ResourceData) []interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})
	if d.Get("shared_with") != nil {
		sharedWithResourceData := d.Get("shared_with").([]interface{})
		if len(sharedWithResourceData) > 0 {
			parsedSharedWith = (sharedWithResourceData[0]).(map[string]interface{})
		}
	}

	var users *[]model.EntityUserAclSharingInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseEntityUserAclSharingInfoList(sharedWith.Users)
	}

	return []interface{}{parsedSharedWith}
}

func parseEntityAclSharingInfo(sharedWith *model.EntityAclSharingInfo) interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})

	var users *[]model.EntityUserAclSharingInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseEntityUserAclSharingInfoList(sharedWith.Users)
	}

	return parsedSharedWith
}

func parseEntityUserAclSharingInfoList(entityUserAclSharingInfo *[]model.EntityUserAclSharingInfo) []interface{} {
	if entityUserAclSharingInfo == nil {
		return nil
	}
	entityUserAclSharingInfoList := make([]interface{}, 0)

	if entityUserAclSharingInfo != nil {
		entityUserAclSharingInfoList = make([]interface{}, len(*entityUserAclSharingInfo))
		for i, entityUserAclSharingInfoItem := range *entityUserAclSharingInfo {
			entityUserAclSharingInfoList[i] = parseEntityUserAclSharingInfo(&entityUserAclSharingInfoItem)
		}
	}

	return entityUserAclSharingInfoList
}

func parseEntityUserAclSharingInfo(entityUserAclSharingInfo *model.EntityUserAclSharingInfo) interface{} {
	if entityUserAclSharingInfo == nil {
		return nil
	}
	parsedEntityUserAclSharingInfo := make(map[string]interface{})
	parsedEntityUserAclSharingInfo["email_id"] = entityUserAclSharingInfo.EmailId
	parsedEntityUserAclSharingInfo["role"] = entityUserAclSharingInfo.Role

	return parsedEntityUserAclSharingInfo
}

func parseDeletionScheduleDTOWithResData(deletionSchedule *model.DeletionScheduleDTO, d *schema.ResourceData) []interface{} {
	if deletionSchedule == nil {
		return nil
	}
	parsedDeletionSchedule := make(map[string]interface{})
	if d.Get("deletion_schedule") != nil {
		deletionScheduleResourceData := d.Get("deletion_schedule").([]interface{})
		if len(deletionScheduleResourceData) > 0 {
			parsedDeletionSchedule = (deletionScheduleResourceData[0]).(map[string]interface{})
		}
	}
	parsedDeletionSchedule["id"] = deletionSchedule.Id
	parsedDeletionSchedule["delete_at"] = deletionSchedule.DeleteAt

	var deletionConfig *model.TessellServiceDeletionConfig
	if deletionSchedule.DeletionConfig != deletionConfig {
		parsedDeletionSchedule["deletion_config"] = []interface{}{parseTessellServiceDeletionConfig(deletionSchedule.DeletionConfig)}
	}

	return []interface{}{parsedDeletionSchedule}
}

func parseDeletionScheduleDTO(deletionSchedule *model.DeletionScheduleDTO) interface{} {
	if deletionSchedule == nil {
		return nil
	}
	parsedDeletionSchedule := make(map[string]interface{})
	parsedDeletionSchedule["id"] = deletionSchedule.Id
	parsedDeletionSchedule["delete_at"] = deletionSchedule.DeleteAt

	var deletionConfig *model.TessellServiceDeletionConfig
	if deletionSchedule.DeletionConfig != deletionConfig {
		parsedDeletionSchedule["deletion_config"] = []interface{}{parseTessellServiceDeletionConfig(deletionSchedule.DeletionConfig)}
	}

	return parsedDeletionSchedule
}

func parseServiceUpcomingScheduledActionsWithResData(upcomingScheduledActions *model.ServiceUpcomingScheduledActions, d *schema.ResourceData) []interface{} {
	if upcomingScheduledActions == nil {
		return nil
	}
	parsedUpcomingScheduledActions := make(map[string]interface{})
	if d.Get("upcoming_scheduled_actions") != nil {
		upcomingScheduledActionsResourceData := d.Get("upcoming_scheduled_actions").([]interface{})
		if len(upcomingScheduledActionsResourceData) > 0 {
			parsedUpcomingScheduledActions = (upcomingScheduledActionsResourceData[0]).(map[string]interface{})
		}
	}

	var startStop *model.ServiceUpcomingScheduledActionsStartStop
	if upcomingScheduledActions.StartStop != startStop {
		parsedUpcomingScheduledActions["start_stop"] = []interface{}{parseServiceUpcomingScheduledActionsStartStop(upcomingScheduledActions.StartStop)}
	}

	var patch *model.ServiceUpcomingScheduledActionsPatch
	if upcomingScheduledActions.Patch != patch {
		parsedUpcomingScheduledActions["patch"] = []interface{}{parseServiceUpcomingScheduledActionsPatch(upcomingScheduledActions.Patch)}
	}

	var delete *model.ServiceUpcomingScheduledActionsDelete
	if upcomingScheduledActions.Delete != delete {
		parsedUpcomingScheduledActions["delete"] = []interface{}{parseServiceUpcomingScheduledActionsDelete(upcomingScheduledActions.Delete)}
	}

	return []interface{}{parsedUpcomingScheduledActions}
}

func parseServiceUpcomingScheduledActions(upcomingScheduledActions *model.ServiceUpcomingScheduledActions) interface{} {
	if upcomingScheduledActions == nil {
		return nil
	}
	parsedUpcomingScheduledActions := make(map[string]interface{})

	var startStop *model.ServiceUpcomingScheduledActionsStartStop
	if upcomingScheduledActions.StartStop != startStop {
		parsedUpcomingScheduledActions["start_stop"] = []interface{}{parseServiceUpcomingScheduledActionsStartStop(upcomingScheduledActions.StartStop)}
	}

	var patch *model.ServiceUpcomingScheduledActionsPatch
	if upcomingScheduledActions.Patch != patch {
		parsedUpcomingScheduledActions["patch"] = []interface{}{parseServiceUpcomingScheduledActionsPatch(upcomingScheduledActions.Patch)}
	}

	var delete *model.ServiceUpcomingScheduledActionsDelete
	if upcomingScheduledActions.Delete != delete {
		parsedUpcomingScheduledActions["delete"] = []interface{}{parseServiceUpcomingScheduledActionsDelete(upcomingScheduledActions.Delete)}
	}

	return parsedUpcomingScheduledActions
}

func parseServiceUpcomingScheduledActionsStartStop(serviceUpcomingScheduledActions_startStop *model.ServiceUpcomingScheduledActionsStartStop) interface{} {
	if serviceUpcomingScheduledActions_startStop == nil {
		return nil
	}
	parsedServiceUpcomingScheduledActions_startStop := make(map[string]interface{})
	parsedServiceUpcomingScheduledActions_startStop["action"] = serviceUpcomingScheduledActions_startStop.Action
	parsedServiceUpcomingScheduledActions_startStop["at"] = serviceUpcomingScheduledActions_startStop.At

	return parsedServiceUpcomingScheduledActions_startStop
}

func parseServiceUpcomingScheduledActionsPatch(serviceUpcomingScheduledActions_patch *model.ServiceUpcomingScheduledActionsPatch) interface{} {
	if serviceUpcomingScheduledActions_patch == nil {
		return nil
	}
	parsedServiceUpcomingScheduledActions_patch := make(map[string]interface{})
	parsedServiceUpcomingScheduledActions_patch["at"] = serviceUpcomingScheduledActions_patch.At
	parsedServiceUpcomingScheduledActions_patch["message"] = serviceUpcomingScheduledActions_patch.Message

	return parsedServiceUpcomingScheduledActions_patch
}

func parseServiceUpcomingScheduledActionsDelete(serviceUpcomingScheduledActions_delete *model.ServiceUpcomingScheduledActionsDelete) interface{} {
	if serviceUpcomingScheduledActions_delete == nil {
		return nil
	}
	parsedServiceUpcomingScheduledActions_delete := make(map[string]interface{})
	parsedServiceUpcomingScheduledActions_delete["at"] = serviceUpcomingScheduledActions_delete.At

	return parsedServiceUpcomingScheduledActions_delete
}

func formatTfInputInstances(d *schema.ResourceData) *[]model.AddDBServiceInstancePayloadV2 {
	tfInstances := d.Get("instances")
	instanceMaps, _ := tfInstances.([]interface{})

	// parse tf instance input
	instances := make([]model.AddDBServiceInstancePayloadV2, 0)
	for _, instanceMap := range instanceMaps {
		inputInstance := instanceMap.(map[string]interface{})
		instance := model.AddDBServiceInstancePayloadV2{
			InstanceGroupName:  helper.GetStringPointer(inputInstance["instance_group_name"]),
			Name:               helper.GetStringPointer(inputInstance["name"]),
			Region:             helper.GetStringPointer(inputInstance["region"]),
			VPC:                helper.GetStringPointer(inputInstance["vpc"]),
			ComputeType:        helper.GetStringPointer(inputInstance["compute_type"]),
			ComputeId:          helper.GetStringPointer(inputInstance["compute_id"]),
			EnablePerfInsights: helper.GetBoolPointer(inputInstance["enable_perf_insights"]),
			AwsInfraConfig:     formAwsInfraConfig(inputInstance["aws_infra_config"]),
			Role:               helper.GetStringPointer(inputInstance["role"]),
			AvailabilityZone:   helper.GetStringPointer(inputInstance["availability_zone"]),
		}
		if inputInstance["private_subnet"] != nil && inputInstance["private_subnet"] != "" {
			instance.PrivateSubnet = helper.GetStringPointer(inputInstance["private_subnet"])
		}
		if inputInstance["data_volume_iops"] != nil && inputInstance["data_volume_iops"] != 0 {
			instance.Iops = helper.GetIntPointer(inputInstance["data_volume_iops"])
		}
		if inputInstance["throughput"] != nil && inputInstance["throughput"] != 0 {
			instance.Throughput = helper.GetIntPointer(inputInstance["throughput"])
		}
		if inputInstance["compute_name"] != nil && inputInstance["compute_name"] != "" {
			instance.ComputeName = helper.GetStringPointer(inputInstance["compute_name"])
		}
		if inputInstance["storage_config"] != nil {
			instance.StorageConfig = formStorageConfigPayload(inputInstance["storage_config"])
		}
		if inputInstance["archive_storage_config"] != nil {
			instance.ArchiveStorageConfig = formStorageConfigPayload(inputInstance["archive_storage_config"])
		}

		instances = append(instances, instance)
	}
	return &instances
}

func getNewTFInstances(d *schema.ResourceData, remoteInstances *[]model.TessellServiceInstanceDTO) *[]model.AddDBServiceInstancePayloadV2 {
	// parse tf instance input
	tfInstances := formatTfInputInstances(d)

	// create map for name of remote instances
	remoteNames := make(map[string]struct{})
	for _, ri := range *remoteInstances {
		remoteNames[*ri.Name] = struct{}{}
	}

	// filter out new tf instance
	var newInstances []model.AddDBServiceInstancePayloadV2
	for _, inst := range *tfInstances {
		if _, found := remoteNames[*inst.Name]; !found {
			newInstances = append(newInstances, inst)
		}
	}
	return &newInstances
}

func formPayloadForAddTessellServiceInstances(d *schema.ResourceData, tfInstancePayload *model.AddDBServiceInstancePayloadV2) *model.AddDBServiceInstancesPayload {
	addDBServiceInstancesPayloadFormed := model.AddDBServiceInstancesPayload{
		InstanceNamePrefix: tfInstancePayload.InstanceGroupName,
		Cloud:              helper.GetStringPointer(d.Get("infrastructure.0.cloud")),
		Region:             tfInstancePayload.Region,
		VPC:                tfInstancePayload.VPC,
		ComputeType:        tfInstancePayload.ComputeType,
		EnablePerfInsights: tfInstancePayload.EnablePerfInsights,
		AwsInfraConfig:     tfInstancePayload.AwsInfraConfig,
		Instances:          formAddDBServiceInstancePayloadList(tfInstancePayload),
	}

	return &addDBServiceInstancesPayloadFormed
}

func formPayloadForCloneTessellService(d *schema.ResourceData) model.CloneTessellServicePayload {
	cloneTessellServicePayloadFormed := model.CloneTessellServicePayload{
		SnapshotId:               helper.GetStringPointer(d.Get("snapshot_id")),
		PITR:                     helper.GetStringPointer(d.Get("pitr")),
		RefreshSchedule:          formCreateUpdateRefreshSchedulePayload(d.Get("refresh_schedule")),
		Name:                     helper.GetStringPointer(d.Get("name")),
		Description:              helper.GetStringPointer(d.Get("description")),
		Subscription:             helper.GetStringPointer(d.Get("subscription")),
		Edition:                  helper.GetStringPointer(d.Get("edition")),
		EngineType:               helper.GetStringPointer(d.Get("engine_type")),
		Topology:                 helper.GetStringPointer(d.Get("topology")),
		NumOfInstances:           helper.GetIntPointer(d.Get("num_of_instances")),
		SoftwareImage:            helper.GetStringPointer(d.Get("software_image")),
		SoftwareImageVersion:     helper.GetStringPointer(d.Get("software_image_version")),
		AutoMinorVersionUpdate:   helper.GetBoolPointer(d.Get("auto_minor_version_update")),
		EnableDeletionProtection: helper.GetBoolPointer(d.Get("enable_deletion_protection")),
		EnableStopProtection:     helper.GetBoolPointer(d.Get("enable_stop_protection")),
		EnablePerfInsights:       helper.GetBoolPointer(d.Get("enable_perf_insights")),
		Infrastructure:           formProvisionInfraPayload(d.Get("infrastructure")),
		Instances:                formAddDBServiceInstancePayloadV2List(d.Get("instances")),
		ServiceConnectivity:      formTessellServiceConnectivityInfoPayload(d.Get("service_connectivity")),
		Creds:                    formTessellServiceCredsPayload(d.Get("creds")),
		MaintenanceWindow:        formTessellServiceMaintenanceWindow(d.Get("maintenance_window")),
		DeletionConfig:           formTessellServiceDeletionConfig(d.Get("deletion_config")),
		SnapshotConfiguration:    formSnapshotConfigurationPayload(d.Get("snapshot_configuration")),
		RPOPolicyConfig:          formRPOPolicyConfig(d.Get("rpo_policy_config")),
		EngineConfiguration:      formTessellServiceEngineConfigurationPayload(d.Get("engine_configuration")),
		Databases:                formCreateDatabasePayloadList(d.Get("databases")),
		IntegrationsConfig:       formTessellServiceIntegrationsPayload(d.Get("integrations_config")),
		Tags:                     formTessellTagList(d.Get("tags")),
	}

	return cloneTessellServicePayloadFormed
}

func formPayloadForDeleteTessellService(d *schema.ResourceData) model.DeleteTessellServicePayload {
	deleteTessellServicePayloadFormed := model.DeleteTessellServicePayload{
		DeletionConfig:  formTessellServiceDeletionConfig(d.Get("deletion_config")),
		Comment:         helper.GetStringPointer(d.Get("comment")),
		PublishEventLog: helper.GetBoolPointer(d.Get("publish_event_log")),
	}

	return deleteTessellServicePayloadFormed
}

// func formPayloadForDeleteTessellServiceInstances(d *schema.ResourceData) model.DeleteTessellServiceInstancePayload {
// 	deleteTessellServiceInstancePayloadFormed := model.DeleteTessellServiceInstancePayload{
// 		InstanceIds: form(d.Get("instance_ids")),
// 	}

// 	return deleteTessellServiceInstancePayloadFormed
// }

func formPayloadForProvisionTessellService(d *schema.ResourceData) model.ProvisionServicePayload {
	provisionServicePayloadFormed := model.ProvisionServicePayload{
		Name:                     helper.GetStringPointer(d.Get("name")),
		Description:              helper.GetStringPointer(d.Get("description")),
		Subscription:             helper.GetStringPointer(d.Get("subscription")),
		Edition:                  helper.GetStringPointer(d.Get("edition")),
		EngineType:               helper.GetStringPointer(d.Get("engine_type")),
		Topology:                 helper.GetStringPointer(d.Get("topology")),
		NumOfInstances:           helper.GetIntPointer(d.Get("num_of_instances")),
		SoftwareImage:            helper.GetStringPointer(d.Get("software_image")),
		SoftwareImageVersion:     helper.GetStringPointer(d.Get("software_image_version")),
		AutoMinorVersionUpdate:   helper.GetBoolPointer(d.Get("auto_minor_version_update")),
		EnableDeletionProtection: helper.GetBoolPointer(d.Get("enable_deletion_protection")),
		EnableStopProtection:     helper.GetBoolPointer(d.Get("enable_stop_protection")),
		EnablePerfInsights:       helper.GetBoolPointer(d.Get("enable_perf_insights")),
		Infrastructure:           formProvisionInfraPayload(d.Get("infrastructure")),
		Instances:                formAddDBServiceInstancePayloadV2List(d.Get("instances")),
		ServiceConnectivity:      formTessellServiceConnectivityInfoPayload(d.Get("service_connectivity")),
		Creds:                    formTessellServiceCredsPayload(d.Get("creds")),
		MaintenanceWindow:        formTessellServiceMaintenanceWindow(d.Get("maintenance_window")),
		DeletionConfig:           formTessellServiceDeletionConfig(d.Get("deletion_config")),
		SnapshotConfiguration:    formSnapshotConfigurationPayload(d.Get("snapshot_configuration")),
		RPOPolicyConfig:          formRPOPolicyConfig(d.Get("rpo_policy_config")),
		EngineConfiguration:      formTessellServiceEngineConfigurationPayload(d.Get("engine_configuration")),
		Databases:                formCreateDatabasePayloadList(d.Get("databases")),
		IntegrationsConfig:       formTessellServiceIntegrationsPayload(d.Get("integrations_config")),
		Tags:                     formTessellTagList(d.Get("tags")),
	}

	return provisionServicePayloadFormed
}

func formPayloadForDeleteTessellServiceInstances(d *schema.ResourceData, remoteInstances *[]model.TessellServiceInstanceDTO) *model.DeleteTessellServiceInstancePayload {
	// parse tf instance input
	tfInstances := formatTfInputInstances(d)

	// create map for name of remote instances
	tfInstanceNames := make(map[string]struct{})
	for _, ri := range *tfInstances {
		tfInstanceNames[*ri.Name] = struct{}{}
	}

	// Get instanceIds from remote instances whose name is not present in tfInstances
	var instanceIds []string
	for _, inst := range *remoteInstances {
		if _, found := tfInstanceNames[*inst.Name]; !found {
			instanceIds = append(instanceIds, *inst.Id)
		}
	}
	deleteTessellServiceInstancePayloadFormed := model.DeleteTessellServiceInstancePayload{
		InstanceIds: &instanceIds,
	}

	return &deleteTessellServiceInstancePayloadFormed
}

func formPayloadForStartTessellService(d *schema.ResourceData) model.StartTessellServicePayload {
	startTessellServicePayloadFormed := model.StartTessellServicePayload{
		Comment: helper.GetStringPointer(d.Get("comment")),
	}

	return startTessellServicePayloadFormed
}

func formPayloadForStopTessellService(d *schema.ResourceData) model.StopTessellServicePayload {
	stopTessellServicePayloadFormed := model.StopTessellServicePayload{
		Comment: helper.GetStringPointer(d.Get("comment")),
	}

	return stopTessellServicePayloadFormed
}

func formPayloadForSwitchoverTessellService(d *schema.ResourceData, remoteInstances *[]model.TessellServiceInstanceDTO) *model.SwitchOverTessellServicePayload {
	// parse tf instance input
	tfInstances := formatTfInputInstances(d)

	// Get instance name of primary instance from input
	var tfPrimaryInstanceName string
	for _, tfInstance := range *tfInstances {
		if *tfInstance.Role == "primary" {
			tfPrimaryInstanceName = *tfInstance.Name
		}
	}

	// Get instanceId of input primary instance
	// And Switchover
	for _, rmInstance := range *remoteInstances {
		if *rmInstance.Name == tfPrimaryInstanceName {
			if *rmInstance.Role != "primary" {
				switchOverTessellServicePayloadFormed := model.SwitchOverTessellServicePayload{
					SwitchToInstanceId: rmInstance.Id,
				}
				return &switchOverTessellServicePayloadFormed
			} else { // if same instance is already primary
				return nil
			}
		}
	}

	return nil
}

func formPayloadForUpdateTessellService(d *schema.ResourceData) model.UpdateTessellServicePayload {
	updateTessellServicePayloadFormed := model.UpdateTessellServicePayload{
		Name:                     helper.GetStringPointer(d.Get("name")),
		Description:              helper.GetStringPointer(d.Get("description")),
		EnableDeletionProtection: helper.GetBoolPointer(d.Get("enable_deletion_protection")),
		EnableStopProtection:     helper.GetBoolPointer(d.Get("enable_stop_protection")),
		AutoMinorVersionUpdate:   helper.GetBoolPointer(d.Get("auto_minor_version_update")),
	}

	return updateTessellServicePayloadFormed
}

func formPayloadForUpdateTessellServiceCredentials(d *schema.ResourceData) model.ResetTessellServiceCredsPayload {
	resetTessellServiceCredsPayloadFormed := model.ResetTessellServiceCredsPayload{
		Creds: formResetTessellServiceCredsPayloadCredsList(d.Get("creds")),
	}

	return resetTessellServiceCredsPayloadFormed
}

func formAwsInfraConfig(awsInfraConfigRaw interface{}) *model.AwsInfraConfig {
	if awsInfraConfigRaw == nil || len(awsInfraConfigRaw.([]interface{})) == 0 {
		return nil
	}

	awsInfraConfigData := awsInfraConfigRaw.([]interface{})[0].(map[string]interface{})

	awsInfraConfigFormed := model.AwsInfraConfig{
		AwsCpuOptions: formAwsCpuOptions(awsInfraConfigData["aws_cpu_options"]),
	}

	return &awsInfraConfigFormed
}

func formAwsCpuOptions(awsCpuOptionsRaw interface{}) *model.AwsCpuOptions {
	if awsCpuOptionsRaw == nil || len(awsCpuOptionsRaw.([]interface{})) == 0 {
		return nil
	}

	awsCpuOptionsData := awsCpuOptionsRaw.([]interface{})[0].(map[string]interface{})

	awsCpuOptionsFormed := model.AwsCpuOptions{
		Vcpus: helper.GetIntPointer(awsCpuOptionsData["vcpus"]),
	}

	return &awsCpuOptionsFormed
}

func formAddDBServiceInstancePayloadList(tfInstancePayload *model.AddDBServiceInstancePayloadV2) *[]model.AddDBServiceInstancePayload {
	if tfInstancePayload == nil {
		return nil
	}
	newInstance := model.AddDBServiceInstancePayload{
		Name:             tfInstancePayload.Name,
		Role:             tfInstancePayload.Role,
		AvailabilityZone: tfInstancePayload.AvailabilityZone,
		ComputeId:        tfInstancePayload.ComputeId,
		StorageConfig:    tfInstancePayload.StorageConfig,
		PrivateSubnet:    tfInstancePayload.PrivateSubnet,
	}

	if tfInstancePayload.PrivateSubnet != nil {
		newInstance.PrivateSubnet = tfInstancePayload.PrivateSubnet
	}

	if tfInstancePayload.ComputeName != nil {
		newInstance.ComputeName = tfInstancePayload.ComputeName
	}

	if tfInstancePayload.Iops != nil && *tfInstancePayload.Iops != 0 {
		newInstance.Iops = tfInstancePayload.Iops
	}
	if tfInstancePayload.Throughput != nil && *tfInstancePayload.Throughput != 0 {
		newInstance.Throughput = tfInstancePayload.Throughput
	}

	if tfInstancePayload.ComputeConfig != nil {
		newInstance.ComputeConfig = tfInstancePayload.ComputeConfig
	}
	if tfInstancePayload.StorageConfig != nil {
		newInstance.StorageConfig = tfInstancePayload.StorageConfig
	}
	if tfInstancePayload.ArchiveStorageConfig != nil {
		newInstance.ArchiveStorageConfig = tfInstancePayload.ArchiveStorageConfig
	}

	InstancesListFormed := []model.AddDBServiceInstancePayload{
		newInstance,
	}

	return &InstancesListFormed
}

func formServiceInstanceEngineInfo(engineConfigurationRaw interface{}) *model.ServiceInstanceEngineInfo {
	if engineConfigurationRaw == nil || len(engineConfigurationRaw.([]interface{})) == 0 {
		return nil
	}

	engineConfigurationData := engineConfigurationRaw.([]interface{})[0].(map[string]interface{})

	serviceInstanceEngineInfoFormed := model.ServiceInstanceEngineInfo{
		OracleConfig: formServiceInstanceOracleEngineConfig(engineConfigurationData["oracle_config"]),
	}

	return &serviceInstanceEngineInfoFormed
}

func formServiceInstanceOracleEngineConfig(oracleConfigRaw interface{}) *model.ServiceInstanceOracleEngineConfig {
	if oracleConfigRaw == nil || len(oracleConfigRaw.([]interface{})) == 0 {
		return nil
	}

	oracleConfigData := oracleConfigRaw.([]interface{})[0].(map[string]interface{})

	serviceInstanceOracleEngineConfigFormed := model.ServiceInstanceOracleEngineConfig{
		AccessMode: helper.GetStringPointer(oracleConfigData["access_mode"]),
	}

	return &serviceInstanceOracleEngineConfigFormed
}

func formComputeConfigPayload(computeConfigPayloadRaw interface{}) *model.ComputeConfigPayload {
	if computeConfigPayloadRaw == nil || len(computeConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	computeConfigPayloadData := computeConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	computeConfigPayloadFormed := model.ComputeConfigPayload{
		Provider:      helper.GetStringPointer(computeConfigPayloadData["provider"]),
		ExadataConfig: formExadataComputeConfigPayload(computeConfigPayloadData["exadata_config"]),
	}

	return &computeConfigPayloadFormed
}

func formExadataComputeConfigPayload(exadataComputeConfigPayloadRaw interface{}) *model.ExadataComputeConfigPayload {
	if exadataComputeConfigPayloadRaw == nil || len(exadataComputeConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	exadataComputeConfigPayloadData := exadataComputeConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	exadataComputeConfigPayloadFormed := model.ExadataComputeConfigPayload{
		InfrastructureId: helper.GetStringPointer(exadataComputeConfigPayloadData["infrastructure_id"]),
		VmClusterId:      helper.GetStringPointer(exadataComputeConfigPayloadData["vm_cluster_id"]),
		ComputeId:        helper.GetStringPointer(exadataComputeConfigPayloadData["compute_id"]),
	}

	return &exadataComputeConfigPayloadFormed
}

func formStorageConfigPayload(storageConfigPayloadRaw interface{}) *model.StorageConfigPayload {
	if storageConfigPayloadRaw == nil || len(storageConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	storageConfigPayloadData := storageConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	storageConfigPayloadFormed := model.StorageConfigPayload{
		Provider:          helper.GetStringPointer(storageConfigPayloadData["provider"]),
		FsxNetAppConfig:   formFsxNetAppConfigPayload(storageConfigPayloadData["fsx_net_app_config"]),
		AzureNetAppConfig: formAzureNetAppConfigPayload(storageConfigPayloadData["azure_net_app_config"]),
	}

	return &storageConfigPayloadFormed
}

func formFsxNetAppConfigPayload(fsxNetAppConfigPayloadRaw interface{}) *model.FsxNetAppConfigPayload {
	if fsxNetAppConfigPayloadRaw == nil || len(fsxNetAppConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	fsxNetAppConfigPayloadData := fsxNetAppConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	fsxNetAppConfigPayloadFormed := model.FsxNetAppConfigPayload{
		FileSystemId: helper.GetStringPointer(fsxNetAppConfigPayloadData["file_system_id"]),
		SvmId:        helper.GetStringPointer(fsxNetAppConfigPayloadData["svm_id"]),
	}

	return &fsxNetAppConfigPayloadFormed
}

func formAzureNetAppConfigPayload(azureNetAppConfigPayloadRaw interface{}) *model.AzureNetAppConfigPayload {
	if azureNetAppConfigPayloadRaw == nil || len(azureNetAppConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	azureNetAppConfigPayloadData := azureNetAppConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	azureNetAppConfigPayloadFormed := model.AzureNetAppConfigPayload{
		AzureNetAppId:  helper.GetStringPointer(azureNetAppConfigPayloadData["azure_net_app_id"]),
		CapacityPoolId: helper.GetStringPointer(azureNetAppConfigPayloadData["capacity_pool_id"]),
		Configurations: formAzureNetAppConfigPayloadConfigurations(azureNetAppConfigPayloadData["configurations"]),
	}

	return &azureNetAppConfigPayloadFormed
}

func formAzureNetAppConfigPayloadConfigurations(azureNetAppConfigPayloadConfigurationsRaw interface{}) *model.AzureNetAppConfigPayloadConfigurations {
	if azureNetAppConfigPayloadConfigurationsRaw == nil || len(azureNetAppConfigPayloadConfigurationsRaw.([]interface{})) == 0 {
		return nil
	}

	azureNetAppConfigPayloadConfigurationsData := azureNetAppConfigPayloadConfigurationsRaw.([]interface{})[0].(map[string]interface{})

	azureNetAppConfigPayloadConfigurationsFormed := model.AzureNetAppConfigPayloadConfigurations{
		NetworkFeatures: helper.GetStringPointer(azureNetAppConfigPayloadConfigurationsData["network_features"]),
	}

	return &azureNetAppConfigPayloadConfigurationsFormed
}

func formCreateUpdateRefreshSchedulePayload(refreshScheduleRaw interface{}) *model.CreateUpdateRefreshSchedulePayload {
	if refreshScheduleRaw == nil || len(refreshScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	refreshScheduleData := refreshScheduleRaw.([]interface{})[0].(map[string]interface{})

	createUpdateRefreshSchedulePayloadFormed := model.CreateUpdateRefreshSchedulePayload{
		RefreshMode:  helper.GetStringPointer(refreshScheduleData["refresh_mode"]),
		ContentType:  helper.GetStringPointer(refreshScheduleData["content_type"]),
		ScriptInfo:   formPrePostScriptInfo(refreshScheduleData["script_info"]),
		ScheduleInfo: formRefreshScheduleInfo(refreshScheduleData["schedule_info"]),
	}

	return &createUpdateRefreshSchedulePayloadFormed
}

func formPrePostScriptInfo(scriptInfoRaw interface{}) *model.PrePostScriptInfo {
	if scriptInfoRaw == nil || len(scriptInfoRaw.([]interface{})) == 0 {
		return nil
	}

	scriptInfoData := scriptInfoRaw.([]interface{})[0].(map[string]interface{})

	prePostScriptInfoFormed := model.PrePostScriptInfo{
		PreScriptInfo:  formScriptInfoList(scriptInfoData["pre_script_info"]),
		PostScriptInfo: formScriptInfoList(scriptInfoData["post_script_info"]),
	}

	return &prePostScriptInfoFormed
}

func formScriptInfo(scriptInfoRaw interface{}) *model.ScriptInfo {
	if scriptInfoRaw == nil || len(scriptInfoRaw.([]interface{})) == 0 {
		return nil
	}

	scriptInfoData := scriptInfoRaw.([]interface{})[0].(map[string]interface{})

	scriptInfoFormed := model.ScriptInfo{
		ScriptId:      helper.GetStringPointer(scriptInfoData["script_id"]),
		ScriptVersion: helper.GetStringPointer(scriptInfoData["script_version"]),
	}

	return &scriptInfoFormed
}
func formScriptInfoList(scriptInfoListRaw interface{}) *[]model.ScriptInfo {
	if scriptInfoListRaw == nil || len(scriptInfoListRaw.([]interface{})) == 0 {
		return nil
	}

	ScriptInfoListFormed := make([]model.ScriptInfo, len(scriptInfoListRaw.([]interface{})))

	for i, scriptInfo := range scriptInfoListRaw.([]interface{}) {
		ScriptInfoListFormed[i] = *formScriptInfo(scriptInfo)
	}

	return &ScriptInfoListFormed
}
func formRefreshScheduleInfo(scheduleInfoRaw interface{}) *model.RefreshScheduleInfo {
	if scheduleInfoRaw == nil || len(scheduleInfoRaw.([]interface{})) == 0 {
		return nil
	}

	scheduleInfoData := scheduleInfoRaw.([]interface{})[0].(map[string]interface{})

	refreshScheduleInfoFormed := model.RefreshScheduleInfo{
		Recurring: formRefreshScheduleRecurrenceInfo(scheduleInfoData["recurring"]),
	}

	return &refreshScheduleInfoFormed
}

func formRefreshScheduleRecurrenceInfo(recurringRaw interface{}) *model.RefreshScheduleRecurrenceInfo {
	if recurringRaw == nil || len(recurringRaw.([]interface{})) == 0 {
		return nil
	}

	recurringData := recurringRaw.([]interface{})[0].(map[string]interface{})

	refreshScheduleRecurrenceInfoFormed := model.RefreshScheduleRecurrenceInfo{
		TriggerTime:    formTimeFormat(recurringData["trigger_time"]),
		DailySchedule:  formRefreshScheduleRecurrenceInfoDailySchedule(recurringData["daily_schedule"]),
		WeeklySchedule: formRefreshScheduleRecurrenceInfoWeeklySchedule(recurringData["weekly_schedule"]),
	}

	return &refreshScheduleRecurrenceInfoFormed
}

func formTimeFormat(timeFormatRaw interface{}) *model.TimeFormat {
	if timeFormatRaw == nil || len(timeFormatRaw.([]interface{})) == 0 {
		return nil
	}

	timeFormatData := timeFormatRaw.([]interface{})[0].(map[string]interface{})

	timeFormatFormed := model.TimeFormat{
		Hour:   helper.GetIntPointer(timeFormatData["hour"]),
		Minute: helper.GetIntPointer(timeFormatData["minute"]),
	}

	return &timeFormatFormed
}

func formRefreshScheduleRecurrenceInfoDailySchedule(dailyScheduleRaw interface{}) *model.RefreshScheduleRecurrenceInfoDailySchedule {
	if dailyScheduleRaw == nil || len(dailyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	dailyScheduleData := dailyScheduleRaw.([]interface{})[0].(map[string]interface{})

	refreshScheduleRecurrenceInfoDailyScheduleFormed := model.RefreshScheduleRecurrenceInfoDailySchedule{
		Enabled: helper.GetBoolPointer(dailyScheduleData["enabled"]),
	}

	return &refreshScheduleRecurrenceInfoDailyScheduleFormed
}

func formRefreshScheduleRecurrenceInfoWeeklySchedule(weeklyScheduleRaw interface{}) *model.RefreshScheduleRecurrenceInfoWeeklySchedule {
	if weeklyScheduleRaw == nil || len(weeklyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	weeklyScheduleData := weeklyScheduleRaw.([]interface{})[0].(map[string]interface{})

	refreshScheduleRecurrenceInfoWeeklyScheduleFormed := model.RefreshScheduleRecurrenceInfoWeeklySchedule{
		WeekDays: formWeekDayList(weeklyScheduleData["week_days"]),
	}

	return &refreshScheduleRecurrenceInfoWeeklyScheduleFormed
}

func formWeekDayList(weekDaysListRaw interface{}) *[]string {
	if weekDaysListRaw == nil || len(weekDaysListRaw.([]interface{})) == 0 {
		return nil
	}

	WeekDaysListFormed := make([]string, len(weekDaysListRaw.([]interface{})))

	for i, weekDay := range weekDaysListRaw.([]interface{}) {
		WeekDaysListFormed[i] = *helper.GetStringPointer(weekDay)
	}

	return &WeekDaysListFormed
}
func formProvisionInfraPayload(provisionInfraPayloadRaw interface{}) *model.ProvisionInfraPayload {
	if provisionInfraPayloadRaw == nil || len(provisionInfraPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	provisionInfraPayloadData := provisionInfraPayloadRaw.([]interface{})[0].(map[string]interface{})

	provisionInfraPayloadFormed := model.ProvisionInfraPayload{
		Cloud:                helper.GetStringPointer(provisionInfraPayloadData["cloud"]),
		Region:               helper.GetStringPointer(provisionInfraPayloadData["region"]),
		AvailabilityZone:     helper.GetStringPointer(provisionInfraPayloadData["availability_zone"]),
		VPC:                  helper.GetStringPointer(provisionInfraPayloadData["vpc"]),
		PrivateSubnet:        helper.GetStringPointer(provisionInfraPayloadData["private_subnet"]),
		EnableEncryption:     helper.GetBoolPointer(provisionInfraPayloadData["enable_encryption"]),
		EncryptionKey:        helper.GetStringPointer(provisionInfraPayloadData["encryption_key"]),
		ComputeType:          helper.GetStringPointer(provisionInfraPayloadData["compute_type"]),
		AwsInfraConfig:       formAwsInfraConfig(provisionInfraPayloadData["aws_infra_config"]),
		AdditionalStorage:    helper.GetIntPointer(provisionInfraPayloadData["additional_storage"]),
		EnableComputeSharing: helper.GetBoolPointer(provisionInfraPayloadData["enable_compute_sharing"]),
		ComputeNamePrefix:    helper.GetStringPointer(provisionInfraPayloadData["compute_name_prefix"]),
		Timezone:             helper.GetStringPointer(provisionInfraPayloadData["timezone"]),
		Computes:             formProvisionComputePayloadList(provisionInfraPayloadData["computes"]),
		Iops:                 helper.GetIntPointer(provisionInfraPayloadData["iops"]),
		Throughput:           helper.GetIntPointer(provisionInfraPayloadData["throughput"]),
	}

	return &provisionInfraPayloadFormed
}

func formProvisionComputePayload(provisionComputePayloadRaw interface{}) *model.ProvisionComputePayload {
	if provisionComputePayloadRaw == nil {
		return nil
	}

	provisionComputePayloadData := provisionComputePayloadRaw.(map[string]interface{})

	provisionComputePayloadFormed := model.ProvisionComputePayload{
		Name:                 helper.GetStringPointer(provisionComputePayloadData["name"]),
		InstanceGroupName:    helper.GetStringPointer(provisionComputePayloadData["instance_group_name"]),
		Region:               helper.GetStringPointer(provisionComputePayloadData["region"]),
		AvailabilityZone:     helper.GetStringPointer(provisionComputePayloadData["availability_zone"]),
		Role:                 helper.GetStringPointer(provisionComputePayloadData["role"]),
		VPC:                  helper.GetStringPointer(provisionComputePayloadData["vpc"]),
		PrivateSubnet:        helper.GetStringPointer(provisionComputePayloadData["private_subnet"]),
		ComputeType:          helper.GetStringPointer(provisionComputePayloadData["compute_type"]),
		ComputeId:            helper.GetStringPointer(provisionComputePayloadData["compute_id"]),
		Timezone:             helper.GetStringPointer(provisionComputePayloadData["timezone"]),
		ComputeConfig:        formComputeConfigPayload(provisionComputePayloadData["compute_config"]),
		StorageConfig:        formStorageConfigPayload(provisionComputePayloadData["storage_config"]),
		ArchiveStorageConfig: formStorageConfigPayload(provisionComputePayloadData["archive_storage_config"]),
	}

	if provisionComputePayloadData["compute_name"] != nil && provisionComputePayloadData["compute_name"] != "" {
		provisionComputePayloadFormed.ComputeName = helper.GetStringPointer(provisionComputePayloadData["compute_name"])
	}

	return &provisionComputePayloadFormed
}
func formProvisionComputePayloadList(provisionComputePayloadListRaw interface{}) *[]model.ProvisionComputePayload {
	if provisionComputePayloadListRaw == nil || len(provisionComputePayloadListRaw.([]interface{})) == 0 {
		return nil
	}

	ProvisionComputePayloadListFormed := make([]model.ProvisionComputePayload, len(provisionComputePayloadListRaw.([]interface{})))

	for i, provisionComputePayload := range provisionComputePayloadListRaw.([]interface{}) {
		ProvisionComputePayloadListFormed[i] = *formProvisionComputePayload(provisionComputePayload)
	}

	return &ProvisionComputePayloadListFormed
}
func formAddDBServiceInstancePayloadV2(addDBServiceInstancePayloadV2Raw interface{}) *model.AddDBServiceInstancePayloadV2 {
	if addDBServiceInstancePayloadV2Raw == nil {
		return nil
	}

	addDBServiceInstancePayloadV2Data := addDBServiceInstancePayloadV2Raw.(map[string]interface{})

	addDBServiceInstancePayloadV2Formed := model.AddDBServiceInstancePayloadV2{
		InstanceGroupName:    helper.GetStringPointer(addDBServiceInstancePayloadV2Data["instance_group_name"]),
		Name:                 helper.GetStringPointer(addDBServiceInstancePayloadV2Data["name"]),
		Region:               helper.GetStringPointer(addDBServiceInstancePayloadV2Data["region"]),
		VPC:                  helper.GetStringPointer(addDBServiceInstancePayloadV2Data["vpc"]),
		PrivateSubnet:        helper.GetStringPointer(addDBServiceInstancePayloadV2Data["private_subnet"]),
		ComputeType:          helper.GetStringPointer(addDBServiceInstancePayloadV2Data["compute_type"]),
		ComputeId:            helper.GetStringPointer(addDBServiceInstancePayloadV2Data["compute_id"]),
		EnablePerfInsights:   helper.GetBoolPointer(addDBServiceInstancePayloadV2Data["enable_perf_insights"]),
		AwsInfraConfig:       formAwsInfraConfig(addDBServiceInstancePayloadV2Data["aws_infra_config"]),
		Role:                 helper.GetStringPointer(addDBServiceInstancePayloadV2Data["role"]),
		AvailabilityZone:     helper.GetStringPointer(addDBServiceInstancePayloadV2Data["availability_zone"]),
		ComputeConfig:        formComputeConfigPayload(addDBServiceInstancePayloadV2Data["compute_config"]),
		StorageConfig:        formStorageConfigPayload(addDBServiceInstancePayloadV2Data["storage_config"]),
		ArchiveStorageConfig: formStorageConfigPayload(addDBServiceInstancePayloadV2Data["archive_storage_config"]),
	}

	if addDBServiceInstancePayloadV2Data["compute_name"] != nil && addDBServiceInstancePayloadV2Data["compute_name"] != "" {
		addDBServiceInstancePayloadV2Formed.ComputeName = helper.GetStringPointer(addDBServiceInstancePayloadV2Data["compute_name"])
	}

	return &addDBServiceInstancePayloadV2Formed
}
func formAddDBServiceInstancePayloadV2List(addDBServiceInstancePayloadV2ListRaw interface{}) *[]model.AddDBServiceInstancePayloadV2 {
	if addDBServiceInstancePayloadV2ListRaw == nil || len(addDBServiceInstancePayloadV2ListRaw.([]interface{})) == 0 {
		return nil
	}

	AddDBServiceInstancePayloadV2ListFormed := make([]model.AddDBServiceInstancePayloadV2, len(addDBServiceInstancePayloadV2ListRaw.([]interface{})))

	for i, addDBServiceInstancePayloadV2 := range addDBServiceInstancePayloadV2ListRaw.([]interface{}) {
		AddDBServiceInstancePayloadV2ListFormed[i] = *formAddDBServiceInstancePayloadV2(addDBServiceInstancePayloadV2)
	}

	return &AddDBServiceInstancePayloadV2ListFormed
}
func formTessellServiceConnectivityInfoPayload(tessellServiceConnectivityInfoPayloadRaw interface{}) *model.TessellServiceConnectivityInfoPayload {
	if tessellServiceConnectivityInfoPayloadRaw == nil || len(tessellServiceConnectivityInfoPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceConnectivityInfoPayloadData := tessellServiceConnectivityInfoPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceConnectivityInfoPayloadFormed := model.TessellServiceConnectivityInfoPayload{
		EnableSSL:          helper.GetBoolPointer(tessellServiceConnectivityInfoPayloadData["enable_ssl"]),
		DNSPrefix:          helper.GetStringPointer(tessellServiceConnectivityInfoPayloadData["dns_prefix"]),
		ServicePort:        helper.GetIntPointer(tessellServiceConnectivityInfoPayloadData["service_port"]),
		EnablePublicAccess: helper.GetBoolPointer(tessellServiceConnectivityInfoPayloadData["enable_public_access"]),
		AllowedIpAddresses: helper.InterfaceToStringSlice(tessellServiceConnectivityInfoPayloadData["allowed_ip_addresses"]),
	}

	return &tessellServiceConnectivityInfoPayloadFormed
}

func formTessellServiceCredsPayload(tessellServiceCredsPayloadRaw interface{}) *model.TessellServiceCredsPayload {
	if tessellServiceCredsPayloadRaw == nil || len(tessellServiceCredsPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceCredsPayloadData := tessellServiceCredsPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceCredsPayloadFormed := model.TessellServiceCredsPayload{
		MasterUser:     helper.GetStringPointer(tessellServiceCredsPayloadData["master_user"]),
		MasterPassword: helper.GetStringPointer(tessellServiceCredsPayloadData["master_password"]),
	}

	return &tessellServiceCredsPayloadFormed
}

func formTessellServiceMaintenanceWindow(tessellServiceMaintenanceWindowRaw interface{}) *model.TessellServiceMaintenanceWindow {
	if tessellServiceMaintenanceWindowRaw == nil || len(tessellServiceMaintenanceWindowRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceMaintenanceWindowData := tessellServiceMaintenanceWindowRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceMaintenanceWindowFormed := model.TessellServiceMaintenanceWindow{
		Day:      helper.GetStringPointer(tessellServiceMaintenanceWindowData["day"]),
		Time:     helper.GetStringPointer(tessellServiceMaintenanceWindowData["time"]),
		Duration: helper.GetIntPointer(tessellServiceMaintenanceWindowData["duration"]),
	}

	return &tessellServiceMaintenanceWindowFormed
}

func formTessellServiceDeletionConfig(tessellServiceDeletionConfigRaw interface{}) *model.TessellServiceDeletionConfig {
	if tessellServiceDeletionConfigRaw == nil || len(tessellServiceDeletionConfigRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceDeletionConfigData := tessellServiceDeletionConfigRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceDeletionConfigFormed := model.TessellServiceDeletionConfig{
		RetainAvailabilityMachine: helper.GetBoolPointer(tessellServiceDeletionConfigData["retain_availability_machine"]),
	}

	return &tessellServiceDeletionConfigFormed
}

func formSnapshotConfigurationPayload(snapshotConfigurationPayloadRaw interface{}) *model.SnapshotConfigurationPayload {
	if snapshotConfigurationPayloadRaw == nil || len(snapshotConfigurationPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	snapshotConfigurationPayloadData := snapshotConfigurationPayloadRaw.([]interface{})[0].(map[string]interface{})

	snapshotConfigurationPayloadFormed := model.SnapshotConfigurationPayload{
		SnapshotWindow:         formSnapshotConfigurationPayloadAllOfSnapshotWindow(snapshotConfigurationPayloadData["snapshot_window"]),
		SLA:                    helper.GetStringPointer(snapshotConfigurationPayloadData["sla"]),
		Schedule:               formScheduleInfo(snapshotConfigurationPayloadData["schedule"]),
		FullBackupSchedule:     formFullBackupSchedule(snapshotConfigurationPayloadData["full_backup_schedule"]),
		RetentionDays:          helper.GetIntPointer(snapshotConfigurationPayloadData["retention_days"]),
		IncludeTransactionLogs: helper.GetBoolPointer(snapshotConfigurationPayloadData["include_transaction_logs"]),
		SnapshotStartTime:      formTimeFormat(snapshotConfigurationPayloadData["snapshot_start_time"]),
	}

	return &snapshotConfigurationPayloadFormed
}

func formSnapshotConfigurationPayloadAllOfSnapshotWindow(snapshotConfigurationPayloadAllOfSnapshotWindowRaw interface{}) *model.SnapshotConfigurationPayloadAllOfSnapshotWindow {
	if snapshotConfigurationPayloadAllOfSnapshotWindowRaw == nil || len(snapshotConfigurationPayloadAllOfSnapshotWindowRaw.([]interface{})) == 0 {
		return nil
	}

	snapshotConfigurationPayloadAllOfSnapshotWindowData := snapshotConfigurationPayloadAllOfSnapshotWindowRaw.([]interface{})[0].(map[string]interface{})

	snapshotConfigurationPayloadAllOfSnapshotWindowFormed := model.SnapshotConfigurationPayloadAllOfSnapshotWindow{
		Time: helper.GetStringPointer(snapshotConfigurationPayloadAllOfSnapshotWindowData["time"]),
	}

	return &snapshotConfigurationPayloadAllOfSnapshotWindowFormed
}

func formScheduleInfo(scheduleInfoRaw interface{}) *model.ScheduleInfo {
	if scheduleInfoRaw == nil || len(scheduleInfoRaw.([]interface{})) == 0 {
		return nil
	}

	scheduleInfoData := scheduleInfoRaw.([]interface{})[0].(map[string]interface{})

	scheduleInfoFormed := model.ScheduleInfo{
		BackupStartTime: formTimeFormat(scheduleInfoData["backup_start_time"]),
		DailySchedule:   formDailySchedule(scheduleInfoData["daily_schedule"]),
		WeeklySchedule:  formWeeklySchedule(scheduleInfoData["weekly_schedule"]),
		MonthlySchedule: formMonthlySchedule(scheduleInfoData["monthly_schedule"]),
		YearlySchedule:  formYearlySchedule(scheduleInfoData["yearly_schedule"]),
	}

	return &scheduleInfoFormed
}

func formDailySchedule(dailyScheduleRaw interface{}) *model.DailySchedule {
	if dailyScheduleRaw == nil || len(dailyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	dailyScheduleData := dailyScheduleRaw.([]interface{})[0].(map[string]interface{})

	dailyScheduleFormed := model.DailySchedule{
		BackupsPerDay: helper.GetIntPointer(dailyScheduleData["backups_per_day"]),
	}

	return &dailyScheduleFormed
}

func formWeeklySchedule(weeklyScheduleRaw interface{}) *model.WeeklySchedule {
	if weeklyScheduleRaw == nil || len(weeklyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	weeklyScheduleData := weeklyScheduleRaw.([]interface{})[0].(map[string]interface{})

	weeklyScheduleFormed := model.WeeklySchedule{
		Days: helper.InterfaceToStringSlice(weeklyScheduleData["days"]),
	}

	return &weeklyScheduleFormed
}

func formMonthlySchedule(monthlyScheduleRaw interface{}) *model.MonthlySchedule {
	if monthlyScheduleRaw == nil || len(monthlyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	monthlyScheduleData := monthlyScheduleRaw.([]interface{})[0].(map[string]interface{})

	monthlyScheduleFormed := model.MonthlySchedule{
		CommonSchedule: formDatesForEachMonth(monthlyScheduleData["common_schedule"]),
	}

	return &monthlyScheduleFormed
}

func formDatesForEachMonth(datesForEachMonthRaw interface{}) *model.DatesForEachMonth {
	if datesForEachMonthRaw == nil || len(datesForEachMonthRaw.([]interface{})) == 0 {
		return nil
	}

	datesForEachMonthData := datesForEachMonthRaw.([]interface{})[0].(map[string]interface{})

	datesForEachMonthFormed := model.DatesForEachMonth{
		Dates:          helper.InterfaceToIntSlice(datesForEachMonthData["dates"]),
		LastDayOfMonth: helper.GetBoolPointer(datesForEachMonthData["last_day_of_month"]),
	}

	return &datesForEachMonthFormed
}

func formYearlySchedule(yearlyScheduleRaw interface{}) *model.YearlySchedule {
	if yearlyScheduleRaw == nil || len(yearlyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	yearlyScheduleData := yearlyScheduleRaw.([]interface{})[0].(map[string]interface{})

	yearlyScheduleFormed := model.YearlySchedule{
		CommonSchedule:        formCommonYearlySchedule(yearlyScheduleData["common_schedule"]),
		MonthSpecificSchedule: formMonthWiseDatesList(yearlyScheduleData["month_specific_schedule"]),
	}

	return &yearlyScheduleFormed
}

func formCommonYearlySchedule(commonYearlyScheduleRaw interface{}) *model.CommonYearlySchedule {
	if commonYearlyScheduleRaw == nil || len(commonYearlyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	commonYearlyScheduleData := commonYearlyScheduleRaw.([]interface{})[0].(map[string]interface{})

	commonYearlyScheduleFormed := model.CommonYearlySchedule{
		Dates:          helper.InterfaceToIntSlice(commonYearlyScheduleData["dates"]),
		LastDayOfMonth: helper.GetBoolPointer(commonYearlyScheduleData["last_day_of_month"]),
		Months:         helper.InterfaceToStringSlice(commonYearlyScheduleData["months"]),
	}

	return &commonYearlyScheduleFormed
}

func formMonthWiseDates(monthWiseDatesRaw interface{}) *model.MonthWiseDates {
	if monthWiseDatesRaw == nil {
		return nil
	}

	monthWiseDatesData := monthWiseDatesRaw.(map[string]interface{})

	monthWiseDatesFormed := model.MonthWiseDates{
		Month: helper.GetStringPointer(monthWiseDatesData["month"]),
		Dates: helper.InterfaceToIntSlice(monthWiseDatesData["dates"]),
	}

	return &monthWiseDatesFormed
}
func formMonthWiseDatesList(monthWiseDatesListRaw interface{}) *[]model.MonthWiseDates {
	if monthWiseDatesListRaw == nil || len(monthWiseDatesListRaw.([]interface{})) == 0 {
		return nil
	}

	MonthWiseDatesListFormed := make([]model.MonthWiseDates, len(monthWiseDatesListRaw.([]interface{})))

	for i, monthWiseDates := range monthWiseDatesListRaw.([]interface{}) {
		MonthWiseDatesListFormed[i] = *formMonthWiseDates(monthWiseDates)
	}

	return &MonthWiseDatesListFormed
}
func formFullBackupSchedule(fullBackupScheduleRaw interface{}) *model.FullBackupSchedule {
	if fullBackupScheduleRaw == nil || len(fullBackupScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	fullBackupScheduleData := fullBackupScheduleRaw.([]interface{})[0].(map[string]interface{})

	fullBackupScheduleFormed := model.FullBackupSchedule{
		StartTime:      formTimeFormat(fullBackupScheduleData["start_time"]),
		WeeklySchedule: formWeeklySchedule(fullBackupScheduleData["weekly_schedule"]),
	}

	return &fullBackupScheduleFormed
}

func formRPOPolicyConfig(rpoPolicyConfigRaw interface{}) *model.RPOPolicyConfig {
	if rpoPolicyConfigRaw == nil || len(rpoPolicyConfigRaw.([]interface{})) == 0 {
		return nil
	}

	rpoPolicyConfigData := rpoPolicyConfigRaw.([]interface{})[0].(map[string]interface{})

	rpoPolicyConfigFormed := model.RPOPolicyConfig{
		IncludeTransactionLogs: helper.GetBoolPointer(rpoPolicyConfigData["include_transaction_logs"]),
		EnableAutoSnapshot:     helper.GetBoolPointer(rpoPolicyConfigData["enable_auto_snapshot"]),
		StandardPolicy:         formStandardRPOPolicy(rpoPolicyConfigData["standard_policy"]),
		CustomPolicy:           formCustomRPOPolicy(rpoPolicyConfigData["custom_policy"]),
		FullBackupSchedule:     formFullBackupSchedule(rpoPolicyConfigData["full_backup_schedule"]),
		EnableAutoBackup:       helper.GetBoolPointer(rpoPolicyConfigData["enable_auto_backup"]),
		BackupRPOConfig:        formRPOPolicyConfigBackupRPOConfig(rpoPolicyConfigData["backup_rpo_config"]),
	}

	return &rpoPolicyConfigFormed
}

func formStandardRPOPolicy(standardRPOPolicyRaw interface{}) *model.StandardRPOPolicy {
	if standardRPOPolicyRaw == nil || len(standardRPOPolicyRaw.([]interface{})) == 0 {
		return nil
	}

	standardRPOPolicyData := standardRPOPolicyRaw.([]interface{})[0].(map[string]interface{})

	standardRPOPolicyFormed := model.StandardRPOPolicy{
		RetentionDays:          helper.GetIntPointer(standardRPOPolicyData["retention_days"]),
		IncludeTransactionLogs: helper.GetBoolPointer(standardRPOPolicyData["include_transaction_logs"]),
		SnapshotStartTime:      formTimeFormat(standardRPOPolicyData["snapshot_start_time"]),
	}

	return &standardRPOPolicyFormed
}

func formCustomRPOPolicy(customRPOPolicyRaw interface{}) *model.CustomRPOPolicy {
	if customRPOPolicyRaw == nil || len(customRPOPolicyRaw.([]interface{})) == 0 {
		return nil
	}

	customRPOPolicyData := customRPOPolicyRaw.([]interface{})[0].(map[string]interface{})

	customRPOPolicyFormed := model.CustomRPOPolicy{
		Name:     helper.GetStringPointer(customRPOPolicyData["name"]),
		Schedule: formScheduleInfo(customRPOPolicyData["schedule"]),
	}

	return &customRPOPolicyFormed
}

func formRPOPolicyConfigBackupRPOConfig(rpoPolicyConfigBackupRPOConfigRaw interface{}) *model.RPOPolicyConfigBackupRPOConfig {
	if rpoPolicyConfigBackupRPOConfigRaw == nil || len(rpoPolicyConfigBackupRPOConfigRaw.([]interface{})) == 0 {
		return nil
	}

	rpoPolicyConfigBackupRPOConfigData := rpoPolicyConfigBackupRPOConfigRaw.([]interface{})[0].(map[string]interface{})

	rpoPolicyConfigBackupRPOConfigFormed := model.RPOPolicyConfigBackupRPOConfig{
		FullBackupSchedule: formFullBackupSchedule(rpoPolicyConfigBackupRPOConfigData["full_backup_schedule"]),
		StandardPolicy:     formBackupStandardRPOPolicy(rpoPolicyConfigBackupRPOConfigData["standard_policy"]),
		CustomPolicy:       formBackupCustomRPOPolicy(rpoPolicyConfigBackupRPOConfigData["custom_policy"]),
	}

	return &rpoPolicyConfigBackupRPOConfigFormed
}

func formBackupStandardRPOPolicy(backupStandardRPOPolicyRaw interface{}) *model.BackupStandardRPOPolicy {
	if backupStandardRPOPolicyRaw == nil || len(backupStandardRPOPolicyRaw.([]interface{})) == 0 {
		return nil
	}

	backupStandardRPOPolicyData := backupStandardRPOPolicyRaw.([]interface{})[0].(map[string]interface{})

	backupStandardRPOPolicyFormed := model.BackupStandardRPOPolicy{
		RetentionDays:   helper.GetIntPointer(backupStandardRPOPolicyData["retention_days"]),
		BackupStartTime: formTimeFormat(backupStandardRPOPolicyData["backup_start_time"]),
	}

	return &backupStandardRPOPolicyFormed
}

func formBackupCustomRPOPolicy(backupCustomRPOPolicyRaw interface{}) *model.BackupCustomRPOPolicy {
	if backupCustomRPOPolicyRaw == nil || len(backupCustomRPOPolicyRaw.([]interface{})) == 0 {
		return nil
	}

	backupCustomRPOPolicyData := backupCustomRPOPolicyRaw.([]interface{})[0].(map[string]interface{})

	backupCustomRPOPolicyFormed := model.BackupCustomRPOPolicy{
		Name:     helper.GetStringPointer(backupCustomRPOPolicyData["name"]),
		Schedule: formScheduleInfo(backupCustomRPOPolicyData["schedule"]),
	}

	return &backupCustomRPOPolicyFormed
}

func formTessellServiceEngineConfigurationPayload(tessellServiceEngineConfigurationPayloadRaw interface{}) *model.TessellServiceEngineConfigurationPayload {
	if tessellServiceEngineConfigurationPayloadRaw == nil || len(tessellServiceEngineConfigurationPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceEngineConfigurationPayloadData := tessellServiceEngineConfigurationPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceEngineConfigurationPayloadFormed := model.TessellServiceEngineConfigurationPayload{
		PreScriptInfo:     formScriptInfo(tessellServiceEngineConfigurationPayloadData["pre_script_info"]),
		PostScriptInfo:    formScriptInfo(tessellServiceEngineConfigurationPayloadData["post_script_info"]),
		OracleConfig:      formOracleEngineConfigPayload(tessellServiceEngineConfigurationPayloadData["oracle_config"]),
		PostgresqlConfig:  formPostgresqlEngineConfigPayload(tessellServiceEngineConfigurationPayloadData["postgresql_config"]),
		MysqlConfig:       formMysqlEngineConfigPayload(tessellServiceEngineConfigurationPayloadData["mysql_config"]),
		SqlServerConfig:   formSqlServerEngineConfigPayload(tessellServiceEngineConfigurationPayloadData["sql_server_config"]),
		ApacheKafkaConfig: formApacheKafkaEngineConfigPayload(tessellServiceEngineConfigurationPayloadData["apache_kafka_config"]),
		MongoDBConfig:     formMongoDBEngineConfigPayload(tessellServiceEngineConfigurationPayloadData["mongodb_config"]),
		MilvusConfig:      formMilvusEngineConfigPayload(tessellServiceEngineConfigurationPayloadData["milvus_config"]),
		CollationConfig:   formDBEngineCollationConfig(tessellServiceEngineConfigurationPayloadData["collation_config"]),
	}

	if tessellServiceEngineConfigurationPayloadFormed.BackupUrl != nil {
		tessellServiceEngineConfigurationPayloadFormed.BackupUrl = helper.GetStringPointer(tessellServiceEngineConfigurationPayloadData["backup_url"])
	}

	if tessellServiceEngineConfigurationPayloadFormed.PreScriptInfo != nil {
		tessellServiceEngineConfigurationPayloadFormed.IgnorePreScriptFailure = helper.GetBoolPointer(tessellServiceEngineConfigurationPayloadData["ignore_pre_script_failure"])
	}

	if tessellServiceEngineConfigurationPayloadFormed.PostScriptInfo != nil {
		tessellServiceEngineConfigurationPayloadFormed.IgnorePostScriptFailure = helper.GetBoolPointer(tessellServiceEngineConfigurationPayloadData["ignore_post_script_failure"])
	}

	return &tessellServiceEngineConfigurationPayloadFormed
}

func formOracleEngineConfigPayload(oracleEngineConfigPayloadRaw interface{}) *model.OracleEngineConfigPayload {
	if oracleEngineConfigPayloadRaw == nil || len(oracleEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	oracleEngineConfigPayloadData := oracleEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	oracleEngineConfigPayloadFormed := model.OracleEngineConfigPayload{
		MultiTenant:          helper.GetBoolPointer(oracleEngineConfigPayloadData["multi_tenant"]),
		Sid:                  helper.GetStringPointer(oracleEngineConfigPayloadData["sid"]),
		ParameterProfileId:   helper.GetStringPointer(oracleEngineConfigPayloadData["parameter_profile_id"]),
		OptionsProfile:       helper.GetStringPointer(oracleEngineConfigPayloadData["options_profile"]),
		OptionProfileId:      helper.GetStringPointer(oracleEngineConfigPayloadData["option_profile_id"]),
		CharacterSet:         helper.GetStringPointer(oracleEngineConfigPayloadData["character_set"]),
		NationalCharacterSet: helper.GetStringPointer(oracleEngineConfigPayloadData["national_character_set"]),
		EnableArchiveMode:    helper.GetBoolPointer(oracleEngineConfigPayloadData["enable_archive_mode"]),
	}

	return &oracleEngineConfigPayloadFormed
}

func formPostgresqlEngineConfigPayload(postgresqlEngineConfigPayloadRaw interface{}) *model.PostgresqlEngineConfigPayload {
	if postgresqlEngineConfigPayloadRaw == nil || len(postgresqlEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	postgresqlEngineConfigPayloadData := postgresqlEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	postgresqlEngineConfigPayloadFormed := model.PostgresqlEngineConfigPayload{
		ParameterProfileId: helper.GetStringPointer(postgresqlEngineConfigPayloadData["parameter_profile_id"]),
		AdDomainId:         helper.GetStringPointer(postgresqlEngineConfigPayloadData["ad_domain_id"]),
		ProxyPort:          helper.GetIntPointer(postgresqlEngineConfigPayloadData["proxy_port"]),
		OptionsProfile:     helper.GetStringPointer(postgresqlEngineConfigPayloadData["options_profile"]),
		OptionProfileId:    helper.GetStringPointer(postgresqlEngineConfigPayloadData["option_profile_id"]),
	}

	return &postgresqlEngineConfigPayloadFormed
}

func formMysqlEngineConfigPayload(mysqlEngineConfigPayloadRaw interface{}) *model.MysqlEngineConfigPayload {
	if mysqlEngineConfigPayloadRaw == nil || len(mysqlEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	mysqlEngineConfigPayloadData := mysqlEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	mysqlEngineConfigPayloadFormed := model.MysqlEngineConfigPayload{
		ParameterProfileId: helper.GetStringPointer(mysqlEngineConfigPayloadData["parameter_profile_id"]),
		AdDomainId:         helper.GetStringPointer(mysqlEngineConfigPayloadData["ad_domain_id"]),
		OptionProfileId:    helper.GetStringPointer(mysqlEngineConfigPayloadData["option_profile_id"]),
	}

	return &mysqlEngineConfigPayloadFormed
}

func formSqlServerEngineConfigPayload(sqlServerEngineConfigPayloadRaw interface{}) *model.SqlServerEngineConfigPayload {
	if sqlServerEngineConfigPayloadRaw == nil || len(sqlServerEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	sqlServerEngineConfigPayloadData := sqlServerEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	sqlServerEngineConfigPayloadFormed := model.SqlServerEngineConfigPayload{
		ParameterProfileId:       helper.GetStringPointer(sqlServerEngineConfigPayloadData["parameter_profile_id"]),
		AdDomainId:               helper.GetStringPointer(sqlServerEngineConfigPayloadData["ad_domain_id"]),
		ServiceAccountCreds:      formCredentialsPayload(sqlServerEngineConfigPayloadData["service_account_creds"]),
		AgentServiceAccountCreds: formCredentialsPayload(sqlServerEngineConfigPayloadData["agent_service_account_creds"]),
		InstanceName:             helper.GetStringPointer(sqlServerEngineConfigPayloadData["instance_name"]),
	}

	return &sqlServerEngineConfigPayloadFormed
}

func formCredentialsPayload(credentialsPayloadRaw interface{}) *model.CredentialsPayload {
	if credentialsPayloadRaw == nil || len(credentialsPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	credentialsPayloadData := credentialsPayloadRaw.([]interface{})[0].(map[string]interface{})

	credentialsPayloadFormed := model.CredentialsPayload{
		User:     helper.GetStringPointer(credentialsPayloadData["user"]),
		Password: helper.GetStringPointer(credentialsPayloadData["password"]),
	}

	return &credentialsPayloadFormed
}

func formApacheKafkaEngineConfigPayload(apacheKafkaEngineConfigPayloadRaw interface{}) *model.ApacheKafkaEngineConfigPayload {
	if apacheKafkaEngineConfigPayloadRaw == nil || len(apacheKafkaEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	apacheKafkaEngineConfigPayloadData := apacheKafkaEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	apacheKafkaEngineConfigPayloadFormed := model.ApacheKafkaEngineConfigPayload{
		ParameterProfileId: helper.GetStringPointer(apacheKafkaEngineConfigPayloadData["parameter_profile_id"]),
	}

	return &apacheKafkaEngineConfigPayloadFormed
}

func formMongoDBEngineConfigPayload(mongoDBEngineConfigPayloadRaw interface{}) *model.MongoDBEngineConfigPayload {
	if mongoDBEngineConfigPayloadRaw == nil || len(mongoDBEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	mongoDBEngineConfigPayloadData := mongoDBEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	mongoDBEngineConfigPayloadFormed := model.MongoDBEngineConfigPayload{
		ClusterName:        helper.GetStringPointer(mongoDBEngineConfigPayloadData["cluster_name"]),
		ParameterProfileId: helper.GetStringPointer(mongoDBEngineConfigPayloadData["parameter_profile_id"]),
	}

	return &mongoDBEngineConfigPayloadFormed
}

func formMilvusEngineConfigPayload(milvusEngineConfigPayloadRaw interface{}) *model.MilvusEngineConfigPayload {
	if milvusEngineConfigPayloadRaw == nil || len(milvusEngineConfigPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	milvusEngineConfigPayloadData := milvusEngineConfigPayloadRaw.([]interface{})[0].(map[string]interface{})

	milvusEngineConfigPayloadFormed := model.MilvusEngineConfigPayload{
		ParameterProfileId: helper.GetStringPointer(milvusEngineConfigPayloadData["parameter_profile_id"]),
	}

	return &milvusEngineConfigPayloadFormed
}

func formDBEngineCollationConfig(dbEngineCollationConfigRaw interface{}) *model.DBEngineCollationConfig {
	if dbEngineCollationConfigRaw == nil || len(dbEngineCollationConfigRaw.([]interface{})) == 0 {
		return nil
	}

	dbEngineCollationConfigData := dbEngineCollationConfigRaw.([]interface{})[0].(map[string]interface{})

	dbEngineCollationConfigFormed := model.DBEngineCollationConfig{
		CollationName: helper.GetStringPointer(dbEngineCollationConfigData["collation_name"]),
	}

	return &dbEngineCollationConfigFormed
}

func formCreateDatabasePayload(createDatabasePayloadRaw interface{}) *model.CreateDatabasePayload {
	if createDatabasePayloadRaw == nil {
		return nil
	}

	createDatabasePayloadData := createDatabasePayloadRaw.(map[string]interface{})

	createDatabasePayloadFormed := model.CreateDatabasePayload{
		DatabaseName:          helper.GetStringPointer(createDatabasePayloadData["database_name"]),
		Description:           helper.GetStringPointer(createDatabasePayloadData["description"]),
		SourceDatabaseId:      helper.GetStringPointer(createDatabasePayloadData["source_database_id"]),
		DatabaseConfiguration: formCreateDatabasePayloadDatabaseConfiguration(createDatabasePayloadData["database_configuration"]),
		CollectionConfig:      formDBCollectionCreatePayload(createDatabasePayloadData["collection_config"]),
	}

	return &createDatabasePayloadFormed
}
func formCreateDatabasePayloadList(createDatabasePayloadListRaw interface{}) *[]model.CreateDatabasePayload {
	if createDatabasePayloadListRaw == nil || len(createDatabasePayloadListRaw.([]interface{})) == 0 {
		return nil
	}

	CreateDatabasePayloadListFormed := make([]model.CreateDatabasePayload, len(createDatabasePayloadListRaw.([]interface{})))

	for i, createDatabasePayload := range createDatabasePayloadListRaw.([]interface{}) {
		CreateDatabasePayloadListFormed[i] = *formCreateDatabasePayload(createDatabasePayload)
	}

	return &CreateDatabasePayloadListFormed
}
func formCreateDatabasePayloadDatabaseConfiguration(createDatabasePayloadDatabaseConfigurationRaw interface{}) *model.CreateDatabasePayloadDatabaseConfiguration {
	if createDatabasePayloadDatabaseConfigurationRaw == nil || len(createDatabasePayloadDatabaseConfigurationRaw.([]interface{})) == 0 {
		return nil
	}

	createDatabasePayloadDatabaseConfigurationData := createDatabasePayloadDatabaseConfigurationRaw.([]interface{})[0].(map[string]interface{})

	createDatabasePayloadDatabaseConfigurationFormed := model.CreateDatabasePayloadDatabaseConfiguration{
		OracleConfig:     formCreateOracleDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["oracle_config"]),
		PostgresqlConfig: formPostgresqlDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["postgresql_config"]),
		MysqlConfig:      formMysqlDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["mysql_config"]),
		SqlServerConfig:  formSqlServerDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["sql_server_config"]),
		MongoDBConfig:    formMongoDBDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["mongodb_config"]),
		MilvusConfig:     formMilvusDatabaseConfig(createDatabasePayloadDatabaseConfigurationData["milvus_config"]),
	}

	return &createDatabasePayloadDatabaseConfigurationFormed
}

func formCreateOracleDatabaseConfig(createOracleDatabaseConfigRaw interface{}) *model.CreateOracleDatabaseConfig {
	if createOracleDatabaseConfigRaw == nil || len(createOracleDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	createOracleDatabaseConfigData := createOracleDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	createOracleDatabaseConfigFormed := model.CreateOracleDatabaseConfig{
		ParameterProfileId: helper.GetStringPointer(createOracleDatabaseConfigData["parameter_profile_id"]),
		OptionsProfile:     helper.GetStringPointer(createOracleDatabaseConfigData["options_profile"]),
		OptionProfileId:    helper.GetStringPointer(createOracleDatabaseConfigData["option_profile_id"]),
		Username:           helper.GetStringPointer(createOracleDatabaseConfigData["username"]),
		Password:           helper.GetStringPointer(createOracleDatabaseConfigData["password"]),
	}

	return &createOracleDatabaseConfigFormed
}

func formPostgresqlDatabaseConfig(postgresqlDatabaseConfigRaw interface{}) *model.PostgresqlDatabaseConfig {
	if postgresqlDatabaseConfigRaw == nil || len(postgresqlDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	postgresqlDatabaseConfigData := postgresqlDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	postgresqlDatabaseConfigFormed := model.PostgresqlDatabaseConfig{
		ParameterProfileId: helper.GetStringPointer(postgresqlDatabaseConfigData["parameter_profile_id"]),
		OptionProfileId:    helper.GetStringPointer(postgresqlDatabaseConfigData["option_profile_id"]),
	}

	return &postgresqlDatabaseConfigFormed
}

func formMysqlDatabaseConfig(mysqlDatabaseConfigRaw interface{}) *model.MysqlDatabaseConfig {
	if mysqlDatabaseConfigRaw == nil || len(mysqlDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	mysqlDatabaseConfigData := mysqlDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	mysqlDatabaseConfigFormed := model.MysqlDatabaseConfig{
		ParameterProfileId: helper.GetStringPointer(mysqlDatabaseConfigData["parameter_profile_id"]),
		OptionProfileId:    helper.GetStringPointer(mysqlDatabaseConfigData["option_profile_id"]),
	}

	return &mysqlDatabaseConfigFormed
}

func formSqlServerDatabaseConfig(sqlServerDatabaseConfigRaw interface{}) *model.SqlServerDatabaseConfig {
	if sqlServerDatabaseConfigRaw == nil || len(sqlServerDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	sqlServerDatabaseConfigData := sqlServerDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	sqlServerDatabaseConfigFormed := model.SqlServerDatabaseConfig{
		ParameterProfileId: helper.GetStringPointer(sqlServerDatabaseConfigData["parameter_profile_id"]),
	}

	return &sqlServerDatabaseConfigFormed
}

func formMongoDBDatabaseConfig(mongoDBDatabaseConfigRaw interface{}) *model.MongoDBDatabaseConfig {
	if mongoDBDatabaseConfigRaw == nil || len(mongoDBDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	mongoDBDatabaseConfigData := mongoDBDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	mongoDBDatabaseConfigFormed := model.MongoDBDatabaseConfig{
		ParameterProfileId: helper.GetStringPointer(mongoDBDatabaseConfigData["parameter_profile_id"]),
	}

	return &mongoDBDatabaseConfigFormed
}

func formMilvusDatabaseConfig(milvusDatabaseConfigRaw interface{}) *model.MilvusDatabaseConfig {
	if milvusDatabaseConfigRaw == nil || len(milvusDatabaseConfigRaw.([]interface{})) == 0 {
		return nil
	}

	milvusDatabaseConfigData := milvusDatabaseConfigRaw.([]interface{})[0].(map[string]interface{})

	milvusDatabaseConfigFormed := model.MilvusDatabaseConfig{
		ParameterProfileId: helper.GetStringPointer(milvusDatabaseConfigData["parameter_profile_id"]),
	}

	return &milvusDatabaseConfigFormed
}

func formDBCollectionCreatePayload(dbCollectionCreatePayloadRaw interface{}) *model.DBCollectionCreatePayload {
	if dbCollectionCreatePayloadRaw == nil || len(dbCollectionCreatePayloadRaw.([]interface{})) == 0 {
		return nil
	}

	dbCollectionCreatePayloadData := dbCollectionCreatePayloadRaw.([]interface{})[0].(map[string]interface{})

	dbCollectionCreatePayloadFormed := model.DBCollectionCreatePayload{
		Name:                   helper.GetStringPointer(dbCollectionCreatePayloadData["name"]),
		MilvusCollectionConfig: formMilvusCreateCollectionConfig(dbCollectionCreatePayloadData["milvus_collection_config"]),
	}

	return &dbCollectionCreatePayloadFormed
}

func formMilvusCreateCollectionConfig(milvusCreateCollectionConfigRaw interface{}) *model.MilvusCreateCollectionConfig {
	if milvusCreateCollectionConfigRaw == nil || len(milvusCreateCollectionConfigRaw.([]interface{})) == 0 {
		return nil
	}

	milvusCreateCollectionConfigData := milvusCreateCollectionConfigRaw.([]interface{})[0].(map[string]interface{})

	milvusCreateCollectionConfigFormed := model.MilvusCreateCollectionConfig{
		Name:          helper.GetStringPointer(milvusCreateCollectionConfigData["name"]),
		ShardNums:     helper.GetIntPointer(milvusCreateCollectionConfigData["shard_nums"]),
		NumPartitions: helper.GetIntPointer(milvusCreateCollectionConfigData["num_partitions"]),
		Schema:        formMilvusCreateCollectionConfigSchema(milvusCreateCollectionConfigData["schema"]),
	}

	return &milvusCreateCollectionConfigFormed
}

func formMilvusCreateCollectionConfigSchema(milvusCreateCollectionConfigSchemaRaw interface{}) *model.MilvusCreateCollectionConfigSchema {
	if milvusCreateCollectionConfigSchemaRaw == nil || len(milvusCreateCollectionConfigSchemaRaw.([]interface{})) == 0 {
		return nil
	}

	milvusCreateCollectionConfigSchemaData := milvusCreateCollectionConfigSchemaRaw.([]interface{})[0].(map[string]interface{})

	milvusCreateCollectionConfigSchemaFormed := model.MilvusCreateCollectionConfigSchema{
		Description:        helper.GetStringPointer(milvusCreateCollectionConfigSchemaData["description"]),
		EnableDynamicField: helper.GetBoolPointer(milvusCreateCollectionConfigSchemaData["enable_dynamic_field"]),
		Fields:             formMilvusCreateCollectionFieldList(milvusCreateCollectionConfigSchemaData["fields"]),
	}

	return &milvusCreateCollectionConfigSchemaFormed
}

func formMilvusCreateCollectionField(milvusCreateCollectionFieldRaw interface{}) *model.MilvusCreateCollectionField {
	if milvusCreateCollectionFieldRaw == nil {
		return nil
	}

	milvusCreateCollectionFieldData := milvusCreateCollectionFieldRaw.(map[string]interface{})

	milvusCreateCollectionFieldFormed := model.MilvusCreateCollectionField{
		Name:               helper.GetStringPointer(milvusCreateCollectionFieldData["name"]),
		Description:        helper.GetStringPointer(milvusCreateCollectionFieldData["description"]),
		Dtype:              helper.GetStringPointer(milvusCreateCollectionFieldData["dtype"]),
		IsPrimary:          helper.GetBoolPointer(milvusCreateCollectionFieldData["is_primary"]),
		AutoId:             helper.GetBoolPointer(milvusCreateCollectionFieldData["auto_id"]),
		DefaultValue:       helper.GetStringPointer(milvusCreateCollectionFieldData["default_value"]),
		IsPartitionKey:     helper.GetBoolPointer(milvusCreateCollectionFieldData["is_partition_key"]),
		MaxLength:          helper.GetIntPointer(milvusCreateCollectionFieldData["max_length"]),
		Dim:                helper.GetIntPointer(milvusCreateCollectionFieldData["dim"]),
		IndexCreatePayload: formDBCollectionIndexPayload(milvusCreateCollectionFieldData["index_create_payload"]),
	}

	return &milvusCreateCollectionFieldFormed
}
func formMilvusCreateCollectionFieldList(milvusCreateCollectionFieldListRaw interface{}) *[]model.MilvusCreateCollectionField {
	if milvusCreateCollectionFieldListRaw == nil || len(milvusCreateCollectionFieldListRaw.([]interface{})) == 0 {
		return nil
	}

	MilvusCreateCollectionFieldListFormed := make([]model.MilvusCreateCollectionField, len(milvusCreateCollectionFieldListRaw.([]interface{})))

	for i, milvusCreateCollectionField := range milvusCreateCollectionFieldListRaw.([]interface{}) {
		MilvusCreateCollectionFieldListFormed[i] = *formMilvusCreateCollectionField(milvusCreateCollectionField)
	}

	return &MilvusCreateCollectionFieldListFormed
}
func formDBCollectionIndexPayload(dbCollectionIndexPayloadRaw interface{}) *model.DBCollectionIndexPayload {
	if dbCollectionIndexPayloadRaw == nil || len(dbCollectionIndexPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	dbCollectionIndexPayloadData := dbCollectionIndexPayloadRaw.([]interface{})[0].(map[string]interface{})

	dbCollectionIndexPayloadFormed := model.DBCollectionIndexPayload{
		FieldName:         helper.GetStringPointer(dbCollectionIndexPayloadData["field_name"]),
		IndexName:         helper.GetStringPointer(dbCollectionIndexPayloadData["index_name"]),
		MilvusIndexConfig: formMilvusIndexConfig(dbCollectionIndexPayloadData["milvus_index_config"]),
	}

	return &dbCollectionIndexPayloadFormed
}

func formMilvusIndexConfig(milvusIndexConfigRaw interface{}) *model.MilvusIndexConfig {
	if milvusIndexConfigRaw == nil || len(milvusIndexConfigRaw.([]interface{})) == 0 {
		return nil
	}

	milvusIndexConfigData := milvusIndexConfigRaw.([]interface{})[0].(map[string]interface{})

	milvusIndexConfigFormed := model.MilvusIndexConfig{
		IndexSpecifications: formIndexSpecificationList(milvusIndexConfigData["index_specifications"]),
	}

	return &milvusIndexConfigFormed
}

func formIndexSpecification(indexSpecificationRaw interface{}) *model.IndexSpecification {
	if indexSpecificationRaw == nil {
		return nil
	}

	indexSpecificationData := indexSpecificationRaw.(map[string]interface{})

	indexSpecificationFormed := model.IndexSpecification{
		IndexType:  helper.GetStringPointer(indexSpecificationData["index_type"]),
		MetricType: helper.GetStringPointer(indexSpecificationData["metric_type"]),
		Parameters: formMilvusIndexParameters(indexSpecificationData["parameters"]),
	}

	return &indexSpecificationFormed
}
func formIndexSpecificationList(indexSpecificationListRaw interface{}) *[]model.IndexSpecification {
	if indexSpecificationListRaw == nil || len(indexSpecificationListRaw.([]interface{})) == 0 {
		return nil
	}

	IndexSpecificationListFormed := make([]model.IndexSpecification, len(indexSpecificationListRaw.([]interface{})))

	for i, indexSpecification := range indexSpecificationListRaw.([]interface{}) {
		IndexSpecificationListFormed[i] = *formIndexSpecification(indexSpecification)
	}

	return &IndexSpecificationListFormed
}
func formMilvusIndexParameters(milvusIndexParametersRaw interface{}) *model.MilvusIndexParameters {
	if milvusIndexParametersRaw == nil || len(milvusIndexParametersRaw.([]interface{})) == 0 {
		return nil
	}

	milvusIndexParametersData := milvusIndexParametersRaw.([]interface{})[0].(map[string]interface{})

	milvusIndexParametersFormed := model.MilvusIndexParameters{
		M:               helper.GetIntPointer(milvusIndexParametersData["m"]),
		Nlist:           helper.GetIntPointer(milvusIndexParametersData["nlist"]),
		Efconstructions: helper.GetIntPointer(milvusIndexParametersData["efconstructions"]),
		Ntrees:          helper.GetIntPointer(milvusIndexParametersData["ntrees"]),
	}

	return &milvusIndexParametersFormed
}

func formTessellServiceIntegrationsPayload(tessellServiceIntegrationsPayloadRaw interface{}) *model.TessellServiceIntegrationsPayload {
	if tessellServiceIntegrationsPayloadRaw == nil || len(tessellServiceIntegrationsPayloadRaw.([]interface{})) == 0 {
		return nil
	}

	tessellServiceIntegrationsPayloadData := tessellServiceIntegrationsPayloadRaw.([]interface{})[0].(map[string]interface{})

	tessellServiceIntegrationsPayloadFormed := model.TessellServiceIntegrationsPayload{
		Integrations: helper.InterfaceToStringSlice(tessellServiceIntegrationsPayloadData["integrations"]),
	}

	return &tessellServiceIntegrationsPayloadFormed
}

func formTessellTag(tessellTagRaw interface{}) *model.TessellTag {
	if tessellTagRaw == nil {
		return nil
	}

	tessellTagData := tessellTagRaw.(map[string]interface{})

	tessellTagFormed := model.TessellTag{
		Name:  helper.GetStringPointer(tessellTagData["name"]),
		Value: helper.GetStringPointer(tessellTagData["value"]),
	}

	return &tessellTagFormed
}
func formTessellTagList(tessellTagListRaw interface{}) *[]model.TessellTag {
	if tessellTagListRaw == nil || len(tessellTagListRaw.([]interface{})) == 0 {
		return nil
	}

	TessellTagListFormed := make([]model.TessellTag, len(tessellTagListRaw.([]interface{})))

	for i, tessellTag := range tessellTagListRaw.([]interface{}) {
		TessellTagListFormed[i] = *formTessellTag(tessellTag)
	}

	return &TessellTagListFormed
}
func formResetTessellServiceCredsPayloadCreds(credsRaw interface{}) *model.ResetTessellServiceCredsPayloadCreds {
	if credsRaw == nil {
		return nil
	}

	credsData := credsRaw.(map[string]interface{})

	resetTessellServiceCredsPayloadCredsFormed := model.ResetTessellServiceCredsPayloadCreds{
		UserName:    helper.GetStringPointer(credsData["master_user"]),
		NewPassword: helper.GetStringPointer(credsData["master_password"]),
	}

	return &resetTessellServiceCredsPayloadCredsFormed
}
func formResetTessellServiceCredsPayloadCredsList(credsListRaw interface{}) *[]model.ResetTessellServiceCredsPayloadCreds {
	if credsListRaw == nil || len(credsListRaw.([]interface{})) == 0 {
		return nil
	}

	CredsListFormed := make([]model.ResetTessellServiceCredsPayloadCreds, len(credsListRaw.([]interface{})))

	for i, creds := range credsListRaw.([]interface{}) {
		CredsListFormed[i] = *formResetTessellServiceCredsPayloadCreds(creds)
	}

	return &CredsListFormed
}
