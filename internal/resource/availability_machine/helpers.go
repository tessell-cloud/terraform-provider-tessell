package availability_machine

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellDMMServiceConsumerDTO *model.TessellDMMServiceConsumerDTO) error {

	if err := d.Set("id", tessellDMMServiceConsumerDTO.Id); err != nil {
		return err
	}

	if err := d.Set("tessell_service_id", tessellDMMServiceConsumerDTO.TessellServiceId); err != nil {
		return err
	}

	if err := d.Set("service_name", tessellDMMServiceConsumerDTO.ServiceName); err != nil {
		return err
	}

	if err := d.Set("tenant", tessellDMMServiceConsumerDTO.Tenant); err != nil {
		return err
	}

	if err := d.Set("subscription", tessellDMMServiceConsumerDTO.Subscription); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellDMMServiceConsumerDTO.EngineType); err != nil {
		return err
	}

	if err := d.Set("data_ingestion_status", tessellDMMServiceConsumerDTO.DataIngestionStatus); err != nil {
		return err
	}

	if err := d.Set("user_id", tessellDMMServiceConsumerDTO.UserId); err != nil {
		return err
	}

	if err := d.Set("owner", tessellDMMServiceConsumerDTO.Owner); err != nil {
		return err
	}

	if err := d.Set("logged_in_user_role", tessellDMMServiceConsumerDTO.LoggedInUserRole); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingInfoWithResData(tessellDMMServiceConsumerDTO.SharedWith, d)); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseCloudRegionInfoListWithResData(tessellDMMServiceConsumerDTO.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("rpo_sla", parseTessellDMMAvailabilityServiceViewWithResData(tessellDMMServiceConsumerDTO.RPOSLA, d)); err != nil {
		return err
	}

	if err := d.Set("daps", parseTessellDAPServiceDTOListWithResData(tessellDMMServiceConsumerDTO.DAPs, d)); err != nil {
		return err
	}

	if err := d.Set("clones", parseTessellCloneSummaryInfoListWithResData(tessellDMMServiceConsumerDTO.Clones, d)); err != nil {
		return err
	}

	if err := d.Set("date_created", tessellDMMServiceConsumerDTO.DateCreated); err != nil {
		return err
	}

	if err := d.Set("date_modified", tessellDMMServiceConsumerDTO.DateModified); err != nil {
		return err
	}

	if err := d.Set("backup_download_config", parseBackupDownloadConfigWithResData(tessellDMMServiceConsumerDTO.BackupDownloadConfig, d)); err != nil {
		return err
	}

	return nil
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

func parseCloudRegionInfoListWithResData(cloudAvailability *[]model.CloudRegionInfo, d *schema.ResourceData) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	cloudRegionInfoList := make([]interface{}, 0)

	if cloudAvailability != nil {
		cloudRegionInfoList = make([]interface{}, len(*cloudAvailability))
		for i, cloudRegionInfoItem := range *cloudAvailability {
			cloudRegionInfoList[i] = parseCloudRegionInfo(&cloudRegionInfoItem)
		}
	}

	return cloudRegionInfoList
}

func parseCloudRegionInfoList(cloudAvailability *[]model.CloudRegionInfo) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	cloudRegionInfoList := make([]interface{}, 0)

	if cloudAvailability != nil {
		cloudRegionInfoList = make([]interface{}, len(*cloudAvailability))
		for i, cloudRegionInfoItem := range *cloudAvailability {
			cloudRegionInfoList[i] = parseCloudRegionInfo(&cloudRegionInfoItem)
		}
	}

	return cloudRegionInfoList
}

func parseCloudRegionInfo(cloudAvailability *model.CloudRegionInfo) interface{} {
	if cloudAvailability == nil {
		return nil
	}
	parsedCloudAvailability := make(map[string]interface{})
	parsedCloudAvailability["cloud"] = cloudAvailability.Cloud

	var regions *[]model.RegionInfo
	if cloudAvailability.Regions != regions {
		parsedCloudAvailability["regions"] = parseRegionInfoList(cloudAvailability.Regions)
	}

	return parsedCloudAvailability
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

func parseTessellDMMAvailabilityServiceViewWithResData(rpoSla *model.TessellDMMAvailabilityServiceView, d *schema.ResourceData) []interface{} {
	if rpoSla == nil {
		return nil
	}
	parsedRpoSla := make(map[string]interface{})
	if d.Get("rpo_sla") != nil {
		rpoSlaResourceData := d.Get("rpo_sla").([]interface{})
		if len(rpoSlaResourceData) > 0 {
			parsedRpoSla = (rpoSlaResourceData[0]).(map[string]interface{})
		}
	}
	parsedRpoSla["availability_machine_id"] = rpoSla.AvailabilityMachineId
	parsedRpoSla["availability_machine"] = rpoSla.AvailabilityMachine

	parsedRpoSla["rpo_sla_status"] = rpoSla.RPOSLAStatus
	parsedRpoSla["sla"] = rpoSla.SLA

	var topology *[]model.DBServiceTopology
	if rpoSla.Topology != topology {
		parsedRpoSla["topology"] = parseDBServiceTopologyList(rpoSla.Topology)
	}

	var schedule *model.ScheduleInfo
	if rpoSla.Schedule != schedule {
		parsedRpoSla["schedule"] = []interface{}{parseScheduleInfo(rpoSla.Schedule)}
	}

	return []interface{}{parsedRpoSla}
}

func parseTessellDMMAvailabilityServiceView(rpoSla *model.TessellDMMAvailabilityServiceView) interface{} {
	if rpoSla == nil {
		return nil
	}
	parsedRpoSla := make(map[string]interface{})
	parsedRpoSla["availability_machine_id"] = rpoSla.AvailabilityMachineId
	parsedRpoSla["availability_machine"] = rpoSla.AvailabilityMachine

	parsedRpoSla["rpo_sla_status"] = rpoSla.RPOSLAStatus
	parsedRpoSla["sla"] = rpoSla.SLA

	var topology *[]model.DBServiceTopology
	if rpoSla.Topology != topology {
		parsedRpoSla["topology"] = parseDBServiceTopologyList(rpoSla.Topology)
	}

	var schedule *model.ScheduleInfo
	if rpoSla.Schedule != schedule {
		parsedRpoSla["schedule"] = []interface{}{parseScheduleInfo(rpoSla.Schedule)}
	}

	return parsedRpoSla
}

func parseDBServiceTopologyList(dbServiceTopology *[]model.DBServiceTopology) []interface{} {
	if dbServiceTopology == nil {
		return nil
	}
	dbServiceTopologyList := make([]interface{}, 0)

	if dbServiceTopology != nil {
		dbServiceTopologyList = make([]interface{}, len(*dbServiceTopology))
		for i, dbServiceTopologyItem := range *dbServiceTopology {
			dbServiceTopologyList[i] = parseDBServiceTopology(&dbServiceTopologyItem)
		}
	}

	return dbServiceTopologyList
}

func parseDBServiceTopology(dbServiceTopology *model.DBServiceTopology) interface{} {
	if dbServiceTopology == nil {
		return nil
	}
	parsedDbServiceTopology := make(map[string]interface{})
	parsedDbServiceTopology["type"] = dbServiceTopology.Type
	parsedDbServiceTopology["cloud_type"] = dbServiceTopology.CloudType
	parsedDbServiceTopology["region"] = dbServiceTopology.Region
	parsedDbServiceTopology["availability_zones"] = dbServiceTopology.AvailabilityZones

	return parsedDbServiceTopology
}

func parseScheduleInfo(scheduleInfo *model.ScheduleInfo) interface{} {
	if scheduleInfo == nil {
		return nil
	}
	parsedScheduleInfo := make(map[string]interface{})

	var backupStartTime *model.ScheduleTimeFormat
	if scheduleInfo.BackupStartTime != backupStartTime {
		parsedScheduleInfo["backup_start_time"] = []interface{}{parseScheduleTimeFormat(scheduleInfo.BackupStartTime)}
	}

	var dailySchedule *model.DailySchedule
	if scheduleInfo.DailySchedule != dailySchedule {
		parsedScheduleInfo["daily_schedule"] = []interface{}{parseDailySchedule(scheduleInfo.DailySchedule)}
	}

	return parsedScheduleInfo
}

func parseScheduleTimeFormatList(scheduleTimeFormat *[]model.ScheduleTimeFormat) []interface{} {
	if scheduleTimeFormat == nil {
		return nil
	}
	scheduleTimeFormatList := make([]interface{}, 0)

	if scheduleTimeFormat != nil {
		scheduleTimeFormatList = make([]interface{}, len(*scheduleTimeFormat))
		for i, scheduleTimeFormatItem := range *scheduleTimeFormat {
			scheduleTimeFormatList[i] = parseScheduleTimeFormat(&scheduleTimeFormatItem)
		}
	}

	return scheduleTimeFormatList
}

func parseScheduleTimeFormat(scheduleTimeFormat *model.ScheduleTimeFormat) interface{} {
	if scheduleTimeFormat == nil {
		return nil
	}
	parsedScheduleTimeFormat := make(map[string]interface{})
	parsedScheduleTimeFormat["hour"] = scheduleTimeFormat.Hour
	parsedScheduleTimeFormat["minute"] = scheduleTimeFormat.Minute

	return parsedScheduleTimeFormat
}

func parseDailySchedule(dailySchedule *model.DailySchedule) interface{} {
	if dailySchedule == nil {
		return nil
	}
	parsedDailySchedule := make(map[string]interface{})
	parsedDailySchedule["backups_per_day"] = dailySchedule.BackupsPerDay

	var backupStartTimes *[]model.ScheduleTimeFormat
	if dailySchedule.BackupStartTimes != backupStartTimes {
		parsedDailySchedule["backup_start_times"] = parseScheduleTimeFormatList(dailySchedule.BackupStartTimes)
	}

	return parsedDailySchedule
}

func parseTessellDAPServiceDTOListWithResData(daps *[]model.TessellDAPServiceDTO, d *schema.ResourceData) []interface{} {
	if daps == nil {
		return nil
	}
	tessellDAPServiceDTOList := make([]interface{}, 0)

	if daps != nil {
		tessellDAPServiceDTOList = make([]interface{}, len(*daps))
		for i, tessellDAPServiceDTOItem := range *daps {
			tessellDAPServiceDTOList[i] = parseTessellDAPServiceDTO(&tessellDAPServiceDTOItem)
		}
	}

	return tessellDAPServiceDTOList
}

func parseTessellDAPServiceDTOList(daps *[]model.TessellDAPServiceDTO) []interface{} {
	if daps == nil {
		return nil
	}
	tessellDAPServiceDTOList := make([]interface{}, 0)

	if daps != nil {
		tessellDAPServiceDTOList = make([]interface{}, len(*daps))
		for i, tessellDAPServiceDTOItem := range *daps {
			tessellDAPServiceDTOList[i] = parseTessellDAPServiceDTO(&tessellDAPServiceDTOItem)
		}
	}

	return tessellDAPServiceDTOList
}

func parseTessellDAPServiceDTO(daps *model.TessellDAPServiceDTO) interface{} {
	if daps == nil {
		return nil
	}
	parsedDaps := make(map[string]interface{})
	parsedDaps["id"] = daps.Id
	parsedDaps["name"] = daps.Name
	parsedDaps["availability_machine_id"] = daps.AvailabilityMachineId
	parsedDaps["tessell_service_id"] = daps.TessellServiceId
	parsedDaps["service_name"] = daps.ServiceName
	parsedDaps["engine_type"] = daps.EngineType
	parsedDaps["content_type"] = daps.ContentType
	parsedDaps["status"] = daps.Status

	parsedDaps["owner"] = daps.Owner
	parsedDaps["logged_in_user_role"] = daps.LoggedInUserRole

	parsedDaps["date_created"] = daps.DateCreated
	parsedDaps["date_modified"] = daps.DateModified

	var contentInfo *model.DAPContentInfo
	if daps.ContentInfo != contentInfo {
		parsedDaps["content_info"] = []interface{}{parseDAPContentInfo(daps.ContentInfo)}
	}

	var cloudAvailability *[]model.CloudRegionInfo
	if daps.CloudAvailability != cloudAvailability {
		parsedDaps["cloud_availability"] = parseCloudRegionInfoList(daps.CloudAvailability)
	}

	var dataAccessConfig *model.RetentionAndScheduleInfo
	if daps.DataAccessConfig != dataAccessConfig {
		parsedDaps["data_access_config"] = []interface{}{parseRetentionAndScheduleInfo(daps.DataAccessConfig)}
	}

	var sharedWith *model.EntityAclSharingInfo
	if daps.SharedWith != sharedWith {
		parsedDaps["shared_with"] = []interface{}{parseEntityAclSharingInfo(daps.SharedWith)}
	}

	return parsedDaps
}

func parseDAPContentInfo(dapContentInfo *model.DAPContentInfo) interface{} {
	if dapContentInfo == nil {
		return nil
	}
	parsedDapContentInfo := make(map[string]interface{})

	var asIsContent *model.AsIsDAPContent
	if dapContentInfo.AsIsContent != asIsContent {
		parsedDapContentInfo["as_is_content"] = []interface{}{parseAsIsDAPContent(dapContentInfo.AsIsContent)}
	}

	var sanitizedContent *model.SanitizationDAPContent
	if dapContentInfo.SanitizedContent != sanitizedContent {
		parsedDapContentInfo["sanitized_content"] = []interface{}{parseSanitizationDAPContent(dapContentInfo.SanitizedContent)}
	}

	var backupContent *model.BackupDAPContent
	if dapContentInfo.BackupContent != backupContent {
		parsedDapContentInfo["backup_content"] = []interface{}{parseBackupDAPContent(dapContentInfo.BackupContent)}
	}

	return parsedDapContentInfo
}

func parseAsIsDAPContent(asIsDapContent *model.AsIsDAPContent) interface{} {
	if asIsDapContent == nil {
		return nil
	}
	parsedAsIsDapContent := make(map[string]interface{})
	parsedAsIsDapContent["automated"] = asIsDapContent.Automated

	var manual *[]model.DAPManualInfo
	if asIsDapContent.Manual != manual {
		parsedAsIsDapContent["manual"] = parseDAPManualInfoList(asIsDapContent.Manual)
	}

	return parsedAsIsDapContent
}

func parseDAPManualInfoList(dapManualInfo *[]model.DAPManualInfo) []interface{} {
	if dapManualInfo == nil {
		return nil
	}
	dapManualInfoList := make([]interface{}, 0)

	if dapManualInfo != nil {
		dapManualInfoList = make([]interface{}, len(*dapManualInfo))
		for i, dapManualInfoItem := range *dapManualInfo {
			dapManualInfoList[i] = parseDAPManualInfo(&dapManualInfoItem)
		}
	}

	return dapManualInfoList
}

func parseDAPManualInfo(dapManualInfo *model.DAPManualInfo) interface{} {
	if dapManualInfo == nil {
		return nil
	}
	parsedDapManualInfo := make(map[string]interface{})
	parsedDapManualInfo["id"] = dapManualInfo.Id
	parsedDapManualInfo["name"] = dapManualInfo.Name
	parsedDapManualInfo["creation_time"] = dapManualInfo.CreationTime
	parsedDapManualInfo["shared_at"] = dapManualInfo.SharedAt

	return parsedDapManualInfo
}

func parseSanitizationDAPContent(sanitizationDapContent *model.SanitizationDAPContent) interface{} {
	if sanitizationDapContent == nil {
		return nil
	}
	parsedSanitizationDapContent := make(map[string]interface{})

	var automated *model.SanitizationDAPContentAutomated
	if sanitizationDapContent.Automated != automated {
		parsedSanitizationDapContent["automated"] = []interface{}{parseSanitizationDAPContentAutomated(sanitizationDapContent.Automated)}
	}

	var manual *[]model.DAPManualInfo
	if sanitizationDapContent.Manual != manual {
		parsedSanitizationDapContent["manual"] = parseDAPManualInfoList(sanitizationDapContent.Manual)
	}

	return parsedSanitizationDapContent
}

func parseSanitizationDAPContentAutomated(sanitizationDapContent_automated *model.SanitizationDAPContentAutomated) interface{} {
	if sanitizationDapContent_automated == nil {
		return nil
	}
	parsedSanitizationDapContent_automated := make(map[string]interface{})
	parsedSanitizationDapContent_automated["sanitization_schedule_id"] = sanitizationDapContent_automated.SanitizationScheduleId

	return parsedSanitizationDapContent_automated
}

func parseBackupDAPContent(backupDapContent *model.BackupDAPContent) interface{} {
	if backupDapContent == nil {
		return nil
	}
	parsedBackupDapContent := make(map[string]interface{})
	parsedBackupDapContent["automated"] = backupDapContent.Automated

	var manual *[]model.DAPManualInfo
	if backupDapContent.Manual != manual {
		parsedBackupDapContent["manual"] = parseDAPManualInfoList(backupDapContent.Manual)
	}

	return parsedBackupDapContent
}

func parseRetentionAndScheduleInfo(retentionAndScheduleInfo *model.RetentionAndScheduleInfo) interface{} {
	if retentionAndScheduleInfo == nil {
		return nil
	}
	parsedRetentionAndScheduleInfo := make(map[string]interface{})
	parsedRetentionAndScheduleInfo["daily_backups"] = retentionAndScheduleInfo.DailyBackups

	return parsedRetentionAndScheduleInfo
}

func parseTessellCloneSummaryInfoListWithResData(clones *[]model.TessellCloneSummaryInfo, d *schema.ResourceData) []interface{} {
	if clones == nil {
		return nil
	}
	tessellCloneSummaryInfoList := make([]interface{}, 0)

	if clones != nil {
		tessellCloneSummaryInfoList = make([]interface{}, len(*clones))
		for i, tessellCloneSummaryInfoItem := range *clones {
			tessellCloneSummaryInfoList[i] = parseTessellCloneSummaryInfo(&tessellCloneSummaryInfoItem)
		}
	}

	return tessellCloneSummaryInfoList
}

func parseTessellCloneSummaryInfoList(clones *[]model.TessellCloneSummaryInfo) []interface{} {
	if clones == nil {
		return nil
	}
	tessellCloneSummaryInfoList := make([]interface{}, 0)

	if clones != nil {
		tessellCloneSummaryInfoList = make([]interface{}, len(*clones))
		for i, tessellCloneSummaryInfoItem := range *clones {
			tessellCloneSummaryInfoList[i] = parseTessellCloneSummaryInfo(&tessellCloneSummaryInfoItem)
		}
	}

	return tessellCloneSummaryInfoList
}

func parseTessellCloneSummaryInfo(clones *model.TessellCloneSummaryInfo) interface{} {
	if clones == nil {
		return nil
	}
	parsedClones := make(map[string]interface{})
	parsedClones["id"] = clones.Id
	parsedClones["name"] = clones.Name
	parsedClones["subscription"] = clones.Subscription
	parsedClones["compute_type"] = clones.ComputeType
	parsedClones["status"] = clones.Status

	parsedClones["clone_info"] = clones.CloneInfo
	parsedClones["owner"] = clones.Owner
	parsedClones["date_created"] = clones.DateCreated

	var cloudAvailability *[]model.CloudRegionInfo
	if clones.CloudAvailability != cloudAvailability {
		parsedClones["cloud_availability"] = parseCloudRegionInfoList(clones.CloudAvailability)
	}

	return parsedClones
}

func parseBackupDownloadConfigWithResData(backupDownloadConfig *model.BackupDownloadConfig, d *schema.ResourceData) []interface{} {
	if backupDownloadConfig == nil {
		return nil
	}
	parsedBackupDownloadConfig := make(map[string]interface{})
	if d.Get("backup_download_config") != nil {
		backupDownloadConfigResourceData := d.Get("backup_download_config").([]interface{})
		if len(backupDownloadConfigResourceData) > 0 {
			parsedBackupDownloadConfig = (backupDownloadConfigResourceData[0]).(map[string]interface{})
		}
	}
	parsedBackupDownloadConfig["allow_backup_downloads_for_all_users"] = backupDownloadConfig.AllowBackupDownloadsForAllUsers
	parsedBackupDownloadConfig["allow_backup_downloads"] = backupDownloadConfig.AllowBackupDownloads

	return []interface{}{parsedBackupDownloadConfig}
}

func parseBackupDownloadConfig(backupDownloadConfig *model.BackupDownloadConfig) interface{} {
	if backupDownloadConfig == nil {
		return nil
	}
	parsedBackupDownloadConfig := make(map[string]interface{})
	parsedBackupDownloadConfig["allow_backup_downloads_for_all_users"] = backupDownloadConfig.AllowBackupDownloadsForAllUsers
	parsedBackupDownloadConfig["allow_backup_downloads"] = backupDownloadConfig.AllowBackupDownloads

	return parsedBackupDownloadConfig
}
