package db_service_start_stop_schedule

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, startStopScheduleDTO *model.StartStopScheduleDTO) error {

	if err := d.Set("id", startStopScheduleDTO.Id); err != nil {
		return err
	}

	if err := d.Set("name", startStopScheduleDTO.Name); err != nil {
		return err
	}

	if err := d.Set("description", startStopScheduleDTO.Description); err != nil {
		return err
	}

	if err := d.Set("service_id", startStopScheduleDTO.ServiceId); err != nil {
		return err
	}

	if err := d.Set("status", startStopScheduleDTO.Status); err != nil {
		return err
	}

	if err := d.Set("schedule_info", parseStartStopScheduleInfoWithResData(startStopScheduleDTO.ScheduleInfo, d)); err != nil {
		return err
	}

	if err := d.Set("metadata", parseStartStopScheduleMetadataWithResData(startStopScheduleDTO.Metadata, d)); err != nil {
		return err
	}

	if err := d.Set("date_created", startStopScheduleDTO.DateCreated); err != nil {
		return err
	}

	if err := d.Set("date_modified", startStopScheduleDTO.DateModified); err != nil {
		return err
	}

	if err := d.Set("last_run", startStopScheduleDTO.LastRun); err != nil {
		return err
	}

	return nil
}

func parseStartStopScheduleInfoWithResData(scheduleInfo *model.StartStopScheduleInfo, d *schema.ResourceData) []interface{} {
	if scheduleInfo == nil {
		return nil
	}
	parsedScheduleInfo := make(map[string]interface{})
	if d.Get("schedule_info") != nil {
		scheduleInfoResourceData := d.Get("schedule_info").([]interface{})
		if len(scheduleInfoResourceData) > 0 {
			parsedScheduleInfo = (scheduleInfoResourceData[0]).(map[string]interface{})
		}
	}

	var oneTime *model.StartStopOneTimeSchedule
	if scheduleInfo.OneTime != oneTime {
		parsedScheduleInfo["one_time"] = []interface{}{parseStartStopOneTimeSchedule(scheduleInfo.OneTime)}
	}

	var recurring *model.StartStopRecurringSchedule
	if scheduleInfo.Recurring != recurring {
		parsedScheduleInfo["recurring"] = []interface{}{parseStartStopRecurringSchedule(scheduleInfo.Recurring)}
	}

	return []interface{}{parsedScheduleInfo}
}

func parseStartStopOneTimeSchedule(oneTime *model.StartStopOneTimeSchedule) interface{} {
	if oneTime == nil {
		return nil
	}
	parsedOneTime := make(map[string]interface{})
	parsedOneTime["db_service_start_at"] = oneTime.DBServiceStartAt
	parsedOneTime["db_service_stop_at"] = oneTime.DBServiceStopAt

	return parsedOneTime
}

func parseStartStopRecurringSchedule(recurring *model.StartStopRecurringSchedule) interface{} {
	if recurring == nil {
		return nil
	}
	parsedRecurring := make(map[string]interface{})
	parsedRecurring["schedule_start_date"] = recurring.ScheduleStartDate
	parsedRecurring["db_service_start_at"] = recurring.DBServiceStartAt
	parsedRecurring["db_service_stop_at"] = recurring.DBServiceStopAt

	parsedRecurring["daily_schedule"] = recurring.DailySchedule

	var scheduleExpiry *model.StartStopRecurringScheduleExpiry
	if recurring.ScheduleExpiry != scheduleExpiry {
		parsedRecurring["schedule_expiry"] = []interface{}{parseStartStopRecurringScheduleExpiry(recurring.ScheduleExpiry)}
	}

	var weeklySchedule *model.StartStopWeeklySchedule
	if recurring.WeeklySchedule != weeklySchedule {
		parsedRecurring["weekly_schedule"] = []interface{}{parseStartStopWeeklySchedule(recurring.WeeklySchedule)}
	}

	return parsedRecurring
}

func parseStartStopRecurringScheduleExpiry(scheduleExpiry *model.StartStopRecurringScheduleExpiry) interface{} {
	if scheduleExpiry == nil {
		return nil
	}
	parsedScheduleExpiry := make(map[string]interface{})
	parsedScheduleExpiry["on"] = scheduleExpiry.On
	parsedScheduleExpiry["after_occurrences"] = scheduleExpiry.AfterOccurrences
	parsedScheduleExpiry["never"] = scheduleExpiry.Never

	return parsedScheduleExpiry
}

func parseStartStopWeeklySchedule(weeklySchedule *model.StartStopWeeklySchedule) interface{} {
	if weeklySchedule == nil {
		return nil
	}
	parsedWeeklySchedule := make(map[string]interface{})
	parsedWeeklySchedule["days"] = weeklySchedule.Days

	return parsedWeeklySchedule
}

func parseStartStopScheduleMetadataWithResData(metadata *model.StartStopScheduleMetadata, d *schema.ResourceData) []interface{} {
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
	parsedMetadata["schedule_counter"] = metadata.ScheduleCounter

	return []interface{}{parsedMetadata}
}

func formPayloadForCreateServiceStartStopSchedule(d *schema.ResourceData) model.CreateStartStopSchedulePayload {
	createStartStopSchedulePayloadFormed := model.CreateStartStopSchedulePayload{
		Name:         helper.GetStringPointer(d.Get("name")),
		Description:  helper.GetStringPointer(d.Get("description")),
		ScheduleInfo: formStartStopScheduleInfo(d.Get("schedule_info")),
	}

	return createStartStopSchedulePayloadFormed
}

func formPayloadForUpdateServiceStartStopSchedule(d *schema.ResourceData) model.UpdateStartStopSchedulePayload {
	updateStartStopSchedulePayloadFormed := model.UpdateStartStopSchedulePayload{
		Name:         helper.GetStringPointer(d.Get("name")),
		Description:  helper.GetStringPointer(d.Get("description")),
		ScheduleInfo: formStartStopScheduleInfo(d.Get("schedule_info")),
	}

	return updateStartStopSchedulePayloadFormed
}

func formStartStopScheduleInfo(startStopScheduleInfoRaw interface{}) *model.StartStopScheduleInfo {
	if startStopScheduleInfoRaw == nil || len(startStopScheduleInfoRaw.([]interface{})) == 0 {
		return nil
	}

	startStopScheduleInfoData := startStopScheduleInfoRaw.([]interface{})[0].(map[string]interface{})

	startStopScheduleInfoFormed := model.StartStopScheduleInfo{
		OneTime:   formStartStopOneTimeSchedule(startStopScheduleInfoData["one_time"]),
		Recurring: formStartStopRecurringSchedule(startStopScheduleInfoData["recurring"]),
	}

	return &startStopScheduleInfoFormed
}

func formStartStopOneTimeSchedule(startStopOneTimeScheduleRaw interface{}) *model.StartStopOneTimeSchedule {
	if startStopOneTimeScheduleRaw == nil || len(startStopOneTimeScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	startStopOneTimeScheduleData := startStopOneTimeScheduleRaw.([]interface{})[0].(map[string]interface{})

	startStopOneTimeScheduleFormed := model.StartStopOneTimeSchedule{
		DBServiceStartAt: helper.GetStringPointer(startStopOneTimeScheduleData["db_service_start_at"]),
		DBServiceStopAt:  helper.GetStringPointer(startStopOneTimeScheduleData["db_service_stop_at"]),
	}

	return &startStopOneTimeScheduleFormed
}

func formStartStopRecurringSchedule(startStopRecurringScheduleRaw interface{}) *model.StartStopRecurringSchedule {
	if startStopRecurringScheduleRaw == nil || len(startStopRecurringScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	startStopRecurringScheduleData := startStopRecurringScheduleRaw.([]interface{})[0].(map[string]interface{})

	startStopRecurringScheduleFormed := model.StartStopRecurringSchedule{
		ScheduleStartDate: helper.GetStringPointer(startStopRecurringScheduleData["schedule_start_date"]),
		DBServiceStartAt:  helper.GetStringPointer(startStopRecurringScheduleData["db_service_start_at"]),
		DBServiceStopAt:   helper.GetStringPointer(startStopRecurringScheduleData["db_service_stop_at"]),
		ScheduleExpiry:    formStartStopRecurringScheduleExpiry(startStopRecurringScheduleData["schedule_expiry"]),
		DailySchedule:     helper.GetBoolPointer(startStopRecurringScheduleData["daily_schedule"]),
		WeeklySchedule:    formStartStopWeeklySchedule(startStopRecurringScheduleData["weekly_schedule"]),
	}

	return &startStopRecurringScheduleFormed
}

func formStartStopRecurringScheduleExpiry(startStopRecurringScheduleExpiryRaw interface{}) *model.StartStopRecurringScheduleExpiry {
	if startStopRecurringScheduleExpiryRaw == nil || len(startStopRecurringScheduleExpiryRaw.([]interface{})) == 0 {
		return nil
	}

	startStopRecurringScheduleExpiryData := startStopRecurringScheduleExpiryRaw.([]interface{})[0].(map[string]interface{})

	startStopRecurringScheduleExpiryFormed := model.StartStopRecurringScheduleExpiry{
		On:               helper.GetStringPointer(startStopRecurringScheduleExpiryData["on"]),
		AfterOccurrences: helper.GetIntPointer(startStopRecurringScheduleExpiryData["after_occurrences"]),
		Never:            helper.GetBoolPointer(startStopRecurringScheduleExpiryData["never"]),
	}

	return &startStopRecurringScheduleExpiryFormed
}

func formStartStopWeeklySchedule(startStopWeeklyScheduleRaw interface{}) *model.StartStopWeeklySchedule {
	if startStopWeeklyScheduleRaw == nil || len(startStopWeeklyScheduleRaw.([]interface{})) == 0 {
		return nil
	}

	startStopWeeklyScheduleData := startStopWeeklyScheduleRaw.([]interface{})[0].(map[string]interface{})

	startStopWeeklyScheduleFormed := model.StartStopWeeklySchedule{
		Days: helper.InterfaceToStringSlice(startStopWeeklyScheduleData["days"]),
	}

	return &startStopWeeklyScheduleFormed
}
