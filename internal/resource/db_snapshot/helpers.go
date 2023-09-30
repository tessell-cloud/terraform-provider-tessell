package db_snapshot

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, databaseSnapshot *model.DatabaseSnapshot) error {

	if err := d.Set("id", databaseSnapshot.Id); err != nil {
		return err
	}

	if err := d.Set("name", databaseSnapshot.Name); err != nil {
		return err
	}

	if err := d.Set("description", databaseSnapshot.Description); err != nil {
		return err
	}

	if err := d.Set("snapshot_time", databaseSnapshot.SnapshotTime); err != nil {
		return err
	}

	if err := d.Set("status", databaseSnapshot.Status); err != nil {
		return err
	}

	if err := d.Set("size", databaseSnapshot.Size); err != nil {
		return err
	}

	if err := d.Set("manual", databaseSnapshot.Manual); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseDatabaseSnapshotCloudRegionInfoListWithResData(databaseSnapshot.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("availability_config", parseSnapshotAvailabilityConfigListWithResData(databaseSnapshot.AvailabilityConfig, d)); err != nil {
		return err
	}

	if err := d.Set("databases", parseBackupDatabaseInfoListWithResData(databaseSnapshot.Databases, d)); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingSummaryInfoWithResData(databaseSnapshot.SharedWith, d)); err != nil {
		return err
	}

	if err := d.Set("backup_status", databaseSnapshot.BackupStatus); err != nil {
		return err
	}

	return nil
}

func parseDatabaseSnapshotCloudRegionInfoListWithResData(cloudAvailability *[]model.DatabaseSnapshotCloudRegionInfo, d *schema.ResourceData) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	databaseSnapshotCloudRegionInfoList := make([]interface{}, 0)

	if cloudAvailability != nil {
		databaseSnapshotCloudRegionInfoList = make([]interface{}, len(*cloudAvailability))
		for i, databaseSnapshotCloudRegionInfoItem := range *cloudAvailability {
			databaseSnapshotCloudRegionInfoList[i] = parseDatabaseSnapshotCloudRegionInfo(&databaseSnapshotCloudRegionInfoItem)
		}
	}

	return databaseSnapshotCloudRegionInfoList
}

func parseDatabaseSnapshotCloudRegionInfoList(cloudAvailability *[]model.DatabaseSnapshotCloudRegionInfo) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	databaseSnapshotCloudRegionInfoList := make([]interface{}, 0)

	if cloudAvailability != nil {
		databaseSnapshotCloudRegionInfoList = make([]interface{}, len(*cloudAvailability))
		for i, databaseSnapshotCloudRegionInfoItem := range *cloudAvailability {
			databaseSnapshotCloudRegionInfoList[i] = parseDatabaseSnapshotCloudRegionInfo(&databaseSnapshotCloudRegionInfoItem)
		}
	}

	return databaseSnapshotCloudRegionInfoList
}

func parseDatabaseSnapshotCloudRegionInfo(cloudAvailability *model.DatabaseSnapshotCloudRegionInfo) interface{} {
	if cloudAvailability == nil {
		return nil
	}
	parsedCloudAvailability := make(map[string]interface{})
	parsedCloudAvailability["cloud"] = cloudAvailability.Cloud

	var regions *[]model.DatabaseSnapshotRegionInfo
	if cloudAvailability.Regions != regions {
		parsedCloudAvailability["regions"] = parseDatabaseSnapshotRegionInfoList(cloudAvailability.Regions)
	}

	return parsedCloudAvailability
}

func parseDatabaseSnapshotRegionInfoList(databaseSnapshotRegionInfo *[]model.DatabaseSnapshotRegionInfo) []interface{} {
	if databaseSnapshotRegionInfo == nil {
		return nil
	}
	databaseSnapshotRegionInfoList := make([]interface{}, 0)

	if databaseSnapshotRegionInfo != nil {
		databaseSnapshotRegionInfoList = make([]interface{}, len(*databaseSnapshotRegionInfo))
		for i, databaseSnapshotRegionInfoItem := range *databaseSnapshotRegionInfo {
			databaseSnapshotRegionInfoList[i] = parseDatabaseSnapshotRegionInfo(&databaseSnapshotRegionInfoItem)
		}
	}

	return databaseSnapshotRegionInfoList
}

func parseDatabaseSnapshotRegionInfo(databaseSnapshotRegionInfo *model.DatabaseSnapshotRegionInfo) interface{} {
	if databaseSnapshotRegionInfo == nil {
		return nil
	}
	parsedDatabaseSnapshotRegionInfo := make(map[string]interface{})
	parsedDatabaseSnapshotRegionInfo["region"] = databaseSnapshotRegionInfo.Region
	parsedDatabaseSnapshotRegionInfo["status"] = databaseSnapshotRegionInfo.Status

	return parsedDatabaseSnapshotRegionInfo
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

func parseEntityAclSharingSummaryInfoWithResData(sharedWith *model.EntityAclSharingSummaryInfo, d *schema.ResourceData) []interface{} {
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
	parsedSharedWith["users"] = sharedWith.Users

	return []interface{}{parsedSharedWith}
}

func parseEntityAclSharingSummaryInfo(sharedWith *model.EntityAclSharingSummaryInfo) interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})
	parsedSharedWith["users"] = sharedWith.Users

	return parsedSharedWith
}

func formPayloadForCreateDatabaseSnapshotRequest(d *schema.ResourceData) model.CreateDatabaseSnapshotTaskPayload {
	createDatabaseSnapshotTaskPayloadFormed := model.CreateDatabaseSnapshotTaskPayload{
		Name:        helper.GetStringPointer(d.Get("name")),
		Description: helper.GetStringPointer(d.Get("description")),
	}

	return createDatabaseSnapshotTaskPayloadFormed
}
