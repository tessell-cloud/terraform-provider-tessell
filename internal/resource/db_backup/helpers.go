package db_backup

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, databaseBackup *model.DatabaseBackup) error {

	if err := d.Set("id", databaseBackup.Id); err != nil {
		return err
	}

	if err := d.Set("name", databaseBackup.Name); err != nil {
		return err
	}

	if err := d.Set("backup_time", databaseBackup.BackupTime); err != nil {
		return err
	}

	if err := d.Set("status", databaseBackup.Status); err != nil {
		return err
	}

	if err := d.Set("size", databaseBackup.Size); err != nil {
		return err
	}

	if err := d.Set("manual", databaseBackup.Manual); err != nil {
		return err
	}

	if err := d.Set("is_incremental", databaseBackup.IsIncremental); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseCloudRegionInfoListWithResData(databaseBackup.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("availability_config", parseSnapshotAvailabilityConfigListWithResData(databaseBackup.AvailabilityConfig, d)); err != nil {
		return err
	}

	if err := d.Set("databases", parseBackupDatabaseInfoListWithResData(databaseBackup.Databases, d)); err != nil {
		return err
	}

	if err := d.Set("backup_source", databaseBackup.BackupSource); err != nil {
		return err
	}

	if err := d.Set("backup_info", parseBackupSourceInfoWithResData(databaseBackup.BackupInfo, d)); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseDatabaseBackupSharedWithWithResData(databaseBackup.SharedWith, d)); err != nil {
		return err
	}

	if err := d.Set("download_url_status", databaseBackup.DownloadUrlStatus); err != nil {
		return err
	}

	return nil
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

func parseSnapshotAvailabilityConfigListWithResData(availabilityConfig *[]model.SnapshotAvailabilityConfig, d *schema.ResourceData) []interface{} {
	if availabilityConfig == nil {
		return nil
	}
	snapshotAvailabilityConfigList := make([]interface{}, 0)

	if availabilityConfig != nil {
		snapshotAvailabilityConfigList = make([]interface{}, len(*availabilityConfig))
		for i, snapshotAvailabilityConfigItem := range *availabilityConfig {
			snapshotAvailabilityConfigList[i] = parseSnapshotAvailabilityConfig(&snapshotAvailabilityConfigItem)
		}
	}

	return snapshotAvailabilityConfigList
}

func parseSnapshotAvailabilityConfigList(availabilityConfig *[]model.SnapshotAvailabilityConfig) []interface{} {
	if availabilityConfig == nil {
		return nil
	}
	snapshotAvailabilityConfigList := make([]interface{}, 0)

	if availabilityConfig != nil {
		snapshotAvailabilityConfigList = make([]interface{}, len(*availabilityConfig))
		for i, snapshotAvailabilityConfigItem := range *availabilityConfig {
			snapshotAvailabilityConfigList[i] = parseSnapshotAvailabilityConfig(&snapshotAvailabilityConfigItem)
		}
	}

	return snapshotAvailabilityConfigList
}

func parseSnapshotAvailabilityConfig(availabilityConfig *model.SnapshotAvailabilityConfig) interface{} {
	if availabilityConfig == nil {
		return nil
	}
	parsedAvailabilityConfig := make(map[string]interface{})
	parsedAvailabilityConfig["availability_configured_manually"] = availabilityConfig.AvailabilityConfiguredManually
	parsedAvailabilityConfig["dap_id"] = availabilityConfig.DAPId

	var cloudAvailabilityConfig *[]model.SnapshotCloudAvailabilityInfo
	if availabilityConfig.CloudAvailabilityConfig != cloudAvailabilityConfig {
		parsedAvailabilityConfig["cloud_availability_config"] = parseSnapshotCloudAvailabilityInfoList(availabilityConfig.CloudAvailabilityConfig)
	}

	return parsedAvailabilityConfig
}

func parseSnapshotCloudAvailabilityInfoList(snapshotCloudAvailabilityInfo *[]model.SnapshotCloudAvailabilityInfo) []interface{} {
	if snapshotCloudAvailabilityInfo == nil {
		return nil
	}
	snapshotCloudAvailabilityInfoList := make([]interface{}, 0)

	if snapshotCloudAvailabilityInfo != nil {
		snapshotCloudAvailabilityInfoList = make([]interface{}, len(*snapshotCloudAvailabilityInfo))
		for i, snapshotCloudAvailabilityInfoItem := range *snapshotCloudAvailabilityInfo {
			snapshotCloudAvailabilityInfoList[i] = parseSnapshotCloudAvailabilityInfo(&snapshotCloudAvailabilityInfoItem)
		}
	}

	return snapshotCloudAvailabilityInfoList
}

func parseSnapshotCloudAvailabilityInfo(snapshotCloudAvailabilityInfo *model.SnapshotCloudAvailabilityInfo) interface{} {
	if snapshotCloudAvailabilityInfo == nil {
		return nil
	}
	parsedSnapshotCloudAvailabilityInfo := make(map[string]interface{})
	parsedSnapshotCloudAvailabilityInfo["cloud"] = snapshotCloudAvailabilityInfo.Cloud

	var regions *[]model.SnapshotRegionAvailability
	if snapshotCloudAvailabilityInfo.Regions != regions {
		parsedSnapshotCloudAvailabilityInfo["regions"] = parseSnapshotRegionAvailabilityList(snapshotCloudAvailabilityInfo.Regions)
	}

	return parsedSnapshotCloudAvailabilityInfo
}

func parseSnapshotRegionAvailabilityList(snapshotRegionAvailability *[]model.SnapshotRegionAvailability) []interface{} {
	if snapshotRegionAvailability == nil {
		return nil
	}
	snapshotRegionAvailabilityList := make([]interface{}, 0)

	if snapshotRegionAvailability != nil {
		snapshotRegionAvailabilityList = make([]interface{}, len(*snapshotRegionAvailability))
		for i, snapshotRegionAvailabilityItem := range *snapshotRegionAvailability {
			snapshotRegionAvailabilityList[i] = parseSnapshotRegionAvailability(&snapshotRegionAvailabilityItem)
		}
	}

	return snapshotRegionAvailabilityList
}

func parseSnapshotRegionAvailability(snapshotRegionAvailability *model.SnapshotRegionAvailability) interface{} {
	if snapshotRegionAvailability == nil {
		return nil
	}
	parsedSnapshotRegionAvailability := make(map[string]interface{})
	parsedSnapshotRegionAvailability["region"] = snapshotRegionAvailability.Region
	parsedSnapshotRegionAvailability["status"] = snapshotRegionAvailability.Status

	return parsedSnapshotRegionAvailability
}

func parseBackupDatabaseInfoListWithResData(databases *[]model.BackupDatabaseInfo, d *schema.ResourceData) []interface{} {
	if databases == nil {
		return nil
	}
	backupDatabaseInfoList := make([]interface{}, 0)

	if databases != nil {
		backupDatabaseInfoList = make([]interface{}, len(*databases))
		for i, backupDatabaseInfoItem := range *databases {
			backupDatabaseInfoList[i] = parseBackupDatabaseInfo(&backupDatabaseInfoItem)
		}
	}

	return backupDatabaseInfoList
}

func parseBackupDatabaseInfoList(databases *[]model.BackupDatabaseInfo) []interface{} {
	if databases == nil {
		return nil
	}
	backupDatabaseInfoList := make([]interface{}, 0)

	if databases != nil {
		backupDatabaseInfoList = make([]interface{}, len(*databases))
		for i, backupDatabaseInfoItem := range *databases {
			backupDatabaseInfoList[i] = parseBackupDatabaseInfo(&backupDatabaseInfoItem)
		}
	}

	return backupDatabaseInfoList
}

func parseBackupDatabaseInfo(databases *model.BackupDatabaseInfo) interface{} {
	if databases == nil {
		return nil
	}
	parsedDatabases := make(map[string]interface{})
	parsedDatabases["id"] = databases.Id
	parsedDatabases["name"] = databases.Name
	parsedDatabases["status"] = databases.Status

	return parsedDatabases
}

func parseBackupSourceInfoWithResData(backupInfo *model.BackupSourceInfo, d *schema.ResourceData) []interface{} {
	if backupInfo == nil {
		return nil
	}
	parsedBackupInfo := make(map[string]interface{})
	if d.Get("backup_info") != nil {
		backupInfoResourceData := d.Get("backup_info").([]interface{})
		if len(backupInfoResourceData) > 0 {
			parsedBackupInfo = (backupInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedBackupInfo["source_snapshot_id"] = backupInfo.SourceSnapshotId
	parsedBackupInfo["snapshot_name"] = backupInfo.SnapshotName
	parsedBackupInfo["snapshot_time"] = backupInfo.SnapshotTime

	return []interface{}{parsedBackupInfo}
}

func parseBackupSourceInfo(backupInfo *model.BackupSourceInfo) interface{} {
	if backupInfo == nil {
		return nil
	}
	parsedBackupInfo := make(map[string]interface{})
	parsedBackupInfo["source_snapshot_id"] = backupInfo.SourceSnapshotId
	parsedBackupInfo["snapshot_name"] = backupInfo.SnapshotName
	parsedBackupInfo["snapshot_time"] = backupInfo.SnapshotTime

	return parsedBackupInfo
}

func parseDatabaseBackupSharedWithWithResData(sharedWith *model.DatabaseBackupSharedWith, d *schema.ResourceData) []interface{} {
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

	var users *[]model.BackupUserInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseBackupUserInfoList(sharedWith.Users)
	}

	return []interface{}{parsedSharedWith}
}

func parseDatabaseBackupSharedWith(sharedWith *model.DatabaseBackupSharedWith) interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})

	var users *[]model.BackupUserInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseBackupUserInfoList(sharedWith.Users)
	}

	return parsedSharedWith
}

func parseBackupUserInfoList(backupUserInfo *[]model.BackupUserInfo) []interface{} {
	if backupUserInfo == nil {
		return nil
	}
	backupUserInfoList := make([]interface{}, 0)

	if backupUserInfo != nil {
		backupUserInfoList = make([]interface{}, len(*backupUserInfo))
		for i, backupUserInfoItem := range *backupUserInfo {
			backupUserInfoList[i] = parseBackupUserInfo(&backupUserInfoItem)
		}
	}

	return backupUserInfoList
}

func parseBackupUserInfo(backupUserInfo *model.BackupUserInfo) interface{} {
	if backupUserInfo == nil {
		return nil
	}
	parsedBackupUserInfo := make(map[string]interface{})
	parsedBackupUserInfo["user_email"] = backupUserInfo.UserEmail
	parsedBackupUserInfo["download_url_status"] = backupUserInfo.DownloadUrlStatus

	var expiryConfig *model.ExpiryConfig
	if backupUserInfo.ExpiryConfig != expiryConfig {
		parsedBackupUserInfo["expiry_config"] = []interface{}{parseExpiryConfig(backupUserInfo.ExpiryConfig)}
	}

	return parsedBackupUserInfo
}

func parseExpiryConfig(expiryConfig *model.ExpiryConfig) interface{} {
	if expiryConfig == nil {
		return nil
	}
	parsedExpiryConfig := make(map[string]interface{})
	parsedExpiryConfig["expire_at"] = expiryConfig.ExpireAt

	return parsedExpiryConfig
}
