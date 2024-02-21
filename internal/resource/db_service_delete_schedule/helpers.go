package db_service_delete_schedule

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, deletionScheduleDTO *model.DeletionScheduleDTO) error {

	if err := d.Set("id", deletionScheduleDTO.Id); err != nil {
		return err
	}

	if err := d.Set("delete_at", deletionScheduleDTO.DeleteAt); err != nil {
		return err
	}

	if err := d.Set("deletion_config", parseTessellServiceDeletionConfigWithResData(deletionScheduleDTO.DeletionConfig, d)); err != nil {
		return err
	}

	return nil
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

func formPayloadForCreateServiceDeletionSchedule(d *schema.ResourceData) model.DeletionSchedulePayload {
	deletionSchedulePayloadFormed := model.DeletionSchedulePayload{
		DeleteAt:       helper.GetStringPointer(d.Get("delete_at")),
		DeletionConfig: formTessellServiceDeletionConfig(d.Get("deletion_config")),
	}

	return deletionSchedulePayloadFormed
}

func formPayloadForUpdateServiceDeletionScheduleTFP(d *schema.ResourceData) model.DeletionSchedulePayload {
	deletionSchedulePayloadFormed := model.DeletionSchedulePayload{
		DeleteAt:       helper.GetStringPointer(d.Get("delete_at")),
		DeletionConfig: formTessellServiceDeletionConfig(d.Get("deletion_config")),
	}

	return deletionSchedulePayloadFormed
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
