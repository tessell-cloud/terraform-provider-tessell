package model

type StartStopScheduleInfo struct {
	OneTime   *StartStopOneTimeSchedule   `json:"oneTime,omitempty"`
	Recurring *StartStopRecurringSchedule `json:"recurring,omitempty"`
}

type StartStopOneTimeSchedule struct {
	DBServiceStartAt *string `json:"dbServiceStartAt,omitempty"` // Time at which the DB Service should be started at
	DBServiceStopAt  *string `json:"dbServiceStopAt,omitempty"`  // Time at which the DB Service should be stopped at
}

type StartStopRecurringSchedule struct {
	ScheduleStartDate *string                           `json:"scheduleStartDate,omitempty"` // Date from which the given recurring schedule would be applicable from
	DBServiceStartAt  *string                           `json:"dbServiceStartAt,omitempty"`  // Time at which the DB Service should be started at
	DBServiceStopAt   *string                           `json:"dbServiceStopAt,omitempty"`   // Time at which the DB Service should be stopped at
	ScheduleExpiry    *StartStopRecurringScheduleExpiry `json:"scheduleExpiry,omitempty"`
	DailySchedule     *bool                             `json:"dailySchedule,omitempty"` // Whether the given schedule is a daily schedule i.e. a schedule which is executed daily
	WeeklySchedule    *StartStopWeeklySchedule          `json:"weeklySchedule,omitempty"`
}

type StartStopRecurringScheduleExpiry struct {
	On               *string `json:"on,omitempty"`               // Date after which the schedule would not be applicable
	AfterOccurrences *int    `json:"afterOccurrences,omitempty"` // Number of occurrences which the schedule would not be applicable
	Never            *bool   `json:"never,omitempty"`            // If set to True, the schedule will be applicable forever
}

type StartStopWeeklySchedule struct {
	Days *[]string `json:"days,omitempty"` // Days of the week on which the recurring start/stop schedule would be applicable for the DB Service
}

type StartStopScheduleMetadata struct {
	ScheduleCounter *int `json:"scheduleCounter,omitempty"`
}

type StartStopScheduleDTO struct {
	Id           *string                    `json:"id,omitempty"`   // The ID of the schedule
	Name         *string                    `json:"name,omitempty"` // Name of the schedule
	Description  *string                    `json:"description,omitempty"`
	ServiceId    *string                    `json:"serviceId"`        // The ID of the DB Service
	Status       *string                    `json:"status,omitempty"` // StartStopScheduleStatus
	ScheduleInfo *StartStopScheduleInfo     `json:"scheduleInfo"`
	Metadata     *StartStopScheduleMetadata `json:"metadata,omitempty"`
	DateCreated  *string                    `json:"dateCreated,omitempty"`  // Timestamp when the schedule was created
	DateModified *string                    `json:"dateModified,omitempty"` // Timestamp when the schedule was last modified
	LastRun      *string                    `json:"lastRun,omitempty"`      // The date-time at which this schedule was last executed
}

type CreateStartStopSchedulePayload struct {
	Name         *string                `json:"name,omitempty"`
	Description  *string                `json:"description,omitempty"`
	ScheduleInfo *StartStopScheduleInfo `json:"scheduleInfo,omitempty"`
}

type UpdateStartStopSchedulePayload struct {
	Name         *string                `json:"name,omitempty"`
	Description  *string                `json:"description,omitempty"`
	ScheduleInfo *StartStopScheduleInfo `json:"scheduleInfo,omitempty"`
}
