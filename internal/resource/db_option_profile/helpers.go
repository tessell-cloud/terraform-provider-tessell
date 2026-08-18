package db_option_profile

import (
	"terraform-provider-tessell/internal/model"
)

func parseDatabaseOptionProfileDriverInfo(driverInfo *model.DatabaseOptionProfileDriverInfo) []interface{} {
	if driverInfo == nil {
		return []interface{}{}
	}
	parsedDriverInfo := make(map[string]interface{})
	parsedDriverInfo["data"] = driverInfo.Data

	return []interface{}{parsedDriverInfo}
}

func parseDatabaseOptionProfileMetadata(metadata *model.DatabaseOptionProfileMetadata) []interface{} {
	if metadata == nil {
		return []interface{}{}
	}
	parsedMetadata := make(map[string]interface{})
	parsedMetadata["data"] = metadata.Data

	return []interface{}{parsedMetadata}
}

func parseDatabaseProfileOptionTypeList(options *[]model.DatabaseProfileOptionType) []interface{} {
	if options == nil {
		return nil
	}
	databaseProfileOptionTypeList := make([]interface{}, len(*options))
	for i, databaseProfileOptionTypeItem := range *options {
		databaseProfileOptionTypeList[i] = parseDatabaseProfileOptionType(&databaseProfileOptionTypeItem)
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

	if options.OptionSettings != nil {
		parsedOptions["option_settings"] = parseDatabaseProfileOptionSettingTypeList(options.OptionSettings)
	}

	return parsedOptions
}

func parseDatabaseProfileOptionSettingTypeList(databaseProfileOptionSettingType *[]model.DatabaseProfileOptionSettingType) []interface{} {
	if databaseProfileOptionSettingType == nil {
		return nil
	}
	databaseProfileOptionSettingTypeList := make([]interface{}, len(*databaseProfileOptionSettingType))
	for i, databaseProfileOptionSettingTypeItem := range *databaseProfileOptionSettingType {
		databaseProfileOptionSettingTypeList[i] = parseDatabaseProfileOptionSettingType(&databaseProfileOptionSettingTypeItem)
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
