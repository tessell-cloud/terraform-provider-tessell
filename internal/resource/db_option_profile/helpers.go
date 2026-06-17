package db_option_profile

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellDatabaseOptionProfileConsumptionDTO *model.TessellDatabaseOptionProfileConsumptionDTO) error {

	if err := d.Set("description", tessellDatabaseOptionProfileConsumptionDTO.Description); err != nil {
		return err
	}

	if err := d.Set("driver_info", parseDatabaseOptionProfileDriverInfoWithResData(tessellDatabaseOptionProfileConsumptionDTO.DriverInfo, d)); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellDatabaseOptionProfileConsumptionDTO.EngineType); err != nil {
		return err
	}

	if err := d.Set("id", tessellDatabaseOptionProfileConsumptionDTO.Id); err != nil {
		return err
	}

	if err := d.Set("option_type_id", tessellDatabaseOptionProfileConsumptionDTO.OptionTypeId); err != nil {
		return err
	}

	if err := d.Set("metadata", parseDatabaseOptionProfileMetadataWithResData(tessellDatabaseOptionProfileConsumptionDTO.Metadata, d)); err != nil {
		return err
	}

	if err := d.Set("name", tessellDatabaseOptionProfileConsumptionDTO.Name); err != nil {
		return err
	}

	if err := d.Set("maturity_status", tessellDatabaseOptionProfileConsumptionDTO.MaturityStatus); err != nil {
		return err
	}

	if err := d.Set("options", parseDatabaseProfileOptionTypeListWithResData(tessellDatabaseOptionProfileConsumptionDTO.Options, d)); err != nil {
		return err
	}

	if err := d.Set("owner", tessellDatabaseOptionProfileConsumptionDTO.Owner); err != nil {
		return err
	}

	if err := d.Set("tenant_id", tessellDatabaseOptionProfileConsumptionDTO.TenantId); err != nil {
		return err
	}

	if err := d.Set("version", tessellDatabaseOptionProfileConsumptionDTO.Version); err != nil {
		return err
	}

	return nil
}

func parseDatabaseOptionProfileDriverInfoWithResData(driverInfo *model.DatabaseOptionProfileDriverInfo, d *schema.ResourceData) []interface{} {
	if driverInfo == nil {
		return nil
	}
	parsedDriverInfo := make(map[string]interface{})
	if d.Get("driver_info") != nil {
		driverInfoResourceData := d.Get("driver_info").([]interface{})
		if len(driverInfoResourceData) > 0 {
			parsedDriverInfo = (driverInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedDriverInfo["data"] = driverInfo.Data

	return []interface{}{parsedDriverInfo}
}

func parseDatabaseOptionProfileDriverInfo(driverInfo *model.DatabaseOptionProfileDriverInfo) interface{} {
	if driverInfo == nil {
		return nil
	}
	parsedDriverInfo := make(map[string]interface{})
	parsedDriverInfo["data"] = driverInfo.Data

	return parsedDriverInfo
}

func parseDatabaseOptionProfileMetadataWithResData(metadata *model.DatabaseOptionProfileMetadata, d *schema.ResourceData) []interface{} {
	if metadata == nil {
		return nil
	}
	parsedMetadata := make(map[string]interface{})
	if d.Get("metadata") != nil {
		metadataResourceData := d.Get("metadata").([]interface{})
		if len(metadataResourceData) > 0 {
			parsedMetadata = (metadataResourceData[0]).(map[string]interface{})
		}
	}
	parsedMetadata["data"] = metadata.Data

	return []interface{}{parsedMetadata}
}

func parseDatabaseOptionProfileMetadata(metadata *model.DatabaseOptionProfileMetadata) interface{} {
	if metadata == nil {
		return nil
	}
	parsedMetadata := make(map[string]interface{})
	parsedMetadata["data"] = metadata.Data

	return parsedMetadata
}

func parseDatabaseProfileOptionTypeListWithResData(options *[]model.DatabaseProfileOptionType, d *schema.ResourceData) []interface{} {
	if options == nil {
		return nil
	}
	databaseProfileOptionTypeList := make([]interface{}, 0)

	if options != nil {
		databaseProfileOptionTypeList = make([]interface{}, len(*options))
		for i, databaseProfileOptionTypeItem := range *options {
			databaseProfileOptionTypeList[i] = parseDatabaseProfileOptionType(&databaseProfileOptionTypeItem)
		}
	}

	return databaseProfileOptionTypeList
}

func parseDatabaseProfileOptionTypeList(options *[]model.DatabaseProfileOptionType) []interface{} {
	if options == nil {
		return nil
	}
	databaseProfileOptionTypeList := make([]interface{}, 0)

	if options != nil {
		databaseProfileOptionTypeList = make([]interface{}, len(*options))
		for i, databaseProfileOptionTypeItem := range *options {
			databaseProfileOptionTypeList[i] = parseDatabaseProfileOptionType(&databaseProfileOptionTypeItem)
		}
	}

	return databaseProfileOptionTypeList
}

func parseDatabaseProfileOptionType(options *model.DatabaseProfileOptionType) interface{} {
	if options == nil {
		return nil
	}
	parsedOptions := make(map[string]interface{})
	parsedOptions["apply_now"] = options.ApplyNow
	parsedOptions["name"] = options.Name

	var optionSettings *[]model.DatabaseProfileOptionSettingType
	if options.OptionSettings != optionSettings {
		parsedOptions["option_settings"] = parseDatabaseProfileOptionSettingTypeList(options.OptionSettings)
	}

	return parsedOptions
}

func parseDatabaseProfileOptionSettingTypeList(databaseProfileOptionSettingType *[]model.DatabaseProfileOptionSettingType) []interface{} {
	if databaseProfileOptionSettingType == nil {
		return nil
	}
	databaseProfileOptionSettingTypeList := make([]interface{}, 0)

	if databaseProfileOptionSettingType != nil {
		databaseProfileOptionSettingTypeList = make([]interface{}, len(*databaseProfileOptionSettingType))
		for i, databaseProfileOptionSettingTypeItem := range *databaseProfileOptionSettingType {
			databaseProfileOptionSettingTypeList[i] = parseDatabaseProfileOptionSettingType(&databaseProfileOptionSettingTypeItem)
		}
	}

	return databaseProfileOptionSettingTypeList
}

func parseDatabaseProfileOptionSettingType(databaseProfileOptionSettingType *model.DatabaseProfileOptionSettingType) interface{} {
	if databaseProfileOptionSettingType == nil {
		return nil
	}
	parsedDatabaseProfileOptionSettingType := make(map[string]interface{})
	parsedDatabaseProfileOptionSettingType["name"] = databaseProfileOptionSettingType.Name
	parsedDatabaseProfileOptionSettingType["value"] = databaseProfileOptionSettingType.Value
	parsedDatabaseProfileOptionSettingType["is_global"] = databaseProfileOptionSettingType.IsGlobal

	return parsedDatabaseProfileOptionSettingType
}
