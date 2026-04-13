package db_service_delete_schedule

import (
	//"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

// suppressRFC3339Diff suppresses diff when two timestamps represent the same point in time
// regardless of formatting differences (e.g. "Z" vs "+00:00", or milliseconds).
func suppressRFC3339Diff(old, new string) bool {
	if old == new {
		return true
	}
	formats := []string{time.RFC3339, time.RFC3339Nano, "2006-01-02T15:04:05.000Z07:00"}
	var oldT, newT time.Time
	var err error
	for _, f := range formats {
		oldT, err = time.Parse(f, old)
		if err == nil {
			break
		}
	}
	for _, f := range formats {
		newT, err = time.Parse(f, new)
		if err == nil {
			break
		}
	}
	return oldT.Equal(newT)
}

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
