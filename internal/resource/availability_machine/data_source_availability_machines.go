package availability_machine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func DataSourceAvailabilityMachines() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceAvailabilityMachinesRead,

		Schema: map[string]*schema.Schema{
			"availability_machines": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "ID of the Availability Machine",
							Computed:    true,
						},
						"tessell_service_id": {
							Type:        schema.TypeString,
							Description: "ID of the DB Service that is associated with the Availability Machine",
							Computed:    true,
						},
						"service_name": {
							Type:        schema.TypeString,
							Description: "Name of the DB Service that is associated with the Availability Machine",
							Computed:    true,
						},
						"tenant": {
							Type:        schema.TypeString,
							Description: "ID of the tenant under which this Availability Machine is effective",
							Computed:    true,
						},
						"subscription": {
							Type:        schema.TypeString,
							Description: "Name of the subscription under which the associated DB Service is hosted",
							Computed:    true,
						},
						"engine_type": {
							Type:        schema.TypeString,
							Description: "Database Engine Type",
							Computed:    true,
						},
						"data_ingestion_status": {
							Type:        schema.TypeString,
							Description: "Availability Machine's data ingestion status",
							Computed:    true,
						},
						"user_id": {
							Type:        schema.TypeString,
							Description: "User details representing the owner for the Availability Machine",
							Computed:    true,
						},
						"owner": {
							Type:        schema.TypeString,
							Description: "User details representing the owner for the Availability Machine",
							Computed:    true,
						},
						"logged_in_user_role": {
							Type:        schema.TypeString,
							Description: "The role of the logged in user for accessing this Availability Machine",
							Computed:    true,
						},
						"shared_with": {
							Type:        schema.TypeList,
							Description: "Tessell Entity ACL Sharing Info",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"email_id": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"role": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"cloud_availability": {
							Type:        schema.TypeList,
							Description: "Availability Machine manages data across multiple regions within a cloud. This sections provides information about the cloud and regions where this Availability Machine is managing the data.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"regions": {
										Type:        schema.TypeList,
										Description: "The regions details",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": {
													Type:        schema.TypeString,
													Description: "The cloud region name",
													Computed:    true,
												},
												"availability_zones": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"rpo_sla": {
							Type:        schema.TypeList,
							Description: "This is a definition for Tessell Availability Machine's sla and schedule details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availability_machine_id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"availability_machine": {
										Type:        schema.TypeString,
										Description: "Associated Availability Machine Name",
										Computed:    true,
									},
									"topology": {
										Type:        schema.TypeList,
										Description: "The availability location details: cloudAccount to region",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"cloud_type": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"region": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"availability_zones": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"rpo_sla_status": {
										Type:        schema.TypeString,
										Description: "The availability status",
										Computed:    true,
									},
									"sla": {
										Type:        schema.TypeString,
										Description: "Associated SLA",
										Computed:    true,
									},
									"sla_retention_info": {
										Type:        schema.TypeList,
										Description: "Retention details of the data that is being governed by the SLA",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pitr": {
													Type:        schema.TypeInt,
													Description: "Retention time (in days) for Point-In-Time recoverability",
													Computed:    true,
												},
												"daily": {
													Type:        schema.TypeInt,
													Description: "Retention time (in days) to retain daily snapshots",
													Computed:    true,
												},
												"weekly": {
													Type:        schema.TypeInt,
													Description: "Retention time (number of weeks) to retain weekly snapshots",
													Computed:    true,
												},
												"monthly": {
													Type:        schema.TypeInt,
													Description: "Retention time (number of months) to retain monthly snapshots",
													Computed:    true,
												},
												"yearly": {
													Type:        schema.TypeInt,
													Description: "Retention time (number of years) to retain yearly snapshots",
													Computed:    true,
												},
											},
										},
									},
									"schedule": {
										Type:        schema.TypeList,
										Description: "Schedule Information",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"backup_start_time": {
													Type:        schema.TypeList,
													Description: "Clock time format value in hour and minute.",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"hour": {
																Type:        schema.TypeInt,
																Description: "",
																Computed:    true,
															},
															"minute": {
																Type:        schema.TypeInt,
																Description: "",
																Computed:    true,
															},
														},
													},
												},
												"daily_schedule": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"backups_per_day": {
																Type:        schema.TypeInt,
																Description: "The number of backups to be captured per day.",
																Computed:    true,
															},
														},
													},
												},
												"weekly_schedule": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"days": {
																Type:        schema.TypeList,
																Description: "Days in a week to retain weekly backups for",
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"monthly_schedule": {
													Type:        schema.TypeList,
													Description: "Definition for taking month specific schedule.",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"common_schedule": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dates": {
																			Type:        schema.TypeList,
																			Description: "Dates in a month to retain monthly backups",
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeInt,
																			},
																		},
																		"last_day_of_month": {
																			Type:        schema.TypeBool,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"yearly_schedule": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"common_schedule": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dates": {
																			Type:        schema.TypeList,
																			Description: "Dates in a month to retain monthly backups",
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeInt,
																			},
																		},
																		"last_day_of_month": {
																			Type:        schema.TypeBool,
																			Description: "",
																			Computed:    true,
																		},
																		"months": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},
															"month_specific_schedule": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"month": {
																			Type:        schema.TypeString,
																			Description: "Name of a month",
																			Computed:    true,
																		},
																		"dates": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeInt,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"daps": {
							Type:        schema.TypeList,
							Description: "The Access Policies (DAP) that have configured for this Availability Machine",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "ID of the Access Policy",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name of the Access Policy",
										Computed:    true,
									},
									"availability_machine_id": {
										Type:        schema.TypeString,
										Description: "ID of the Availability Machine",
										Computed:    true,
									},
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "ID of the associated DB Service",
										Computed:    true,
									},
									"service_name": {
										Type:        schema.TypeString,
										Description: "Name of the associated DB Service",
										Computed:    true,
									},
									"engine_type": {
										Type:        schema.TypeString,
										Description: "Database engine type of the associated DB Service",
										Computed:    true,
									},
									"content_type": {
										Type:        schema.TypeString,
										Description: "Content Type for the Data Access Policy",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Database Access Policy Status",
										Computed:    true,
									},
									"content_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"as_is_content": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"automated": {
																Type:        schema.TypeBool,
																Description: "Share the automated as-is snapshots. This is exclusive with manual specification.",
																Computed:    true,
															},
															"manual": {
																Type:        schema.TypeList,
																Description: "The list of snapshots that are to be shared as part of this access policy",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:        schema.TypeString,
																			Description: "The DB Service snapshot id",
																			Computed:    true,
																		},
																		"name": {
																			Type:        schema.TypeString,
																			Description: "The DB Service snapshot name",
																			Computed:    true,
																		},
																		"creation_time": {
																			Type:        schema.TypeString,
																			Description: "DB Service snapshot capture time",
																			Computed:    true,
																		},
																		"shared_at": {
																			Type:        schema.TypeString,
																			Description: "The timestamp when the snapshot was added to DAP for sharing",
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"sanitized_content": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"automated": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"sanitization_schedule_id": {
																			Type:        schema.TypeString,
																			Description: "Id of the sanitization schedule to process automated backups, required only if contentType = Sanitized.",
																			Computed:    true,
																		},
																	},
																},
															},
															"manual": {
																Type:        schema.TypeList,
																Description: "The list of sanitized snapshots that are to be shared as part of this access policy",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:        schema.TypeString,
																			Description: "The DB Service snapshot id",
																			Computed:    true,
																		},
																		"name": {
																			Type:        schema.TypeString,
																			Description: "The DB Service snapshot name",
																			Computed:    true,
																		},
																		"creation_time": {
																			Type:        schema.TypeString,
																			Description: "DB Service snapshot capture time",
																			Computed:    true,
																		},
																		"shared_at": {
																			Type:        schema.TypeString,
																			Description: "The timestamp when the snapshot was added to DAP for sharing",
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"backup_content": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"automated": {
																Type:        schema.TypeBool,
																Description: "Share the automated backups. This is exclusive with manual specification.",
																Computed:    true,
															},
															"manual": {
																Type:        schema.TypeList,
																Description: "The list of backups that are to be shared as part of this access policy",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:        schema.TypeString,
																			Description: "The DB Service snapshot id",
																			Computed:    true,
																		},
																		"name": {
																			Type:        schema.TypeString,
																			Description: "The DB Service snapshot name",
																			Computed:    true,
																		},
																		"creation_time": {
																			Type:        schema.TypeString,
																			Description: "DB Service snapshot capture time",
																			Computed:    true,
																		},
																		"shared_at": {
																			Type:        schema.TypeString,
																			Description: "The timestamp when the snapshot was added to DAP for sharing",
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"cloud_availability": {
										Type:        schema.TypeList,
										Description: "The cloud and region information where the data is being managed by this Access Policy",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cloud": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"regions": {
													Type:        schema.TypeList,
													Description: "The regions details",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"region": {
																Type:        schema.TypeString,
																Description: "The cloud region name",
																Computed:    true,
															},
															"availability_zones": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"data_access_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"daily_backups": {
													Type:        schema.TypeInt,
													Description: "Number of daily backups to replicate",
													Computed:    true,
												},
											},
										},
									},
									"owner": {
										Type:        schema.TypeString,
										Description: "Owner of the Access Policy",
										Computed:    true,
									},
									"logged_in_user_role": {
										Type:        schema.TypeString,
										Description: "The role of the logged in user for accessing the Availability Machine",
										Computed:    true,
									},
									"shared_with": {
										Type:        schema.TypeList,
										Description: "Tessell Entity ACL Sharing Info",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"users": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"email_id": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"role": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"date_created": {
										Type:        schema.TypeString,
										Description: "Timestamp when this Access Policy was created at",
										Computed:    true,
									},
									"date_modified": {
										Type:        schema.TypeString,
										Description: "Timestamp when this Access Policy was last updated at",
										Computed:    true,
									},
								},
							},
						},
						"clones": {
							Type:        schema.TypeList,
							Description: "The clone DB Services that have been created using contents (snapshots, Sanitized Snapshots, PITR, backups) from this Availability Machine",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name of the clone database",
										Computed:    true,
									},
									"subscription": {
										Type:        schema.TypeString,
										Description: "Clone's subscription name",
										Computed:    true,
									},
									"compute_type": {
										Type:        schema.TypeString,
										Description: "Clone's compute type",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Status of the clone database",
										Computed:    true,
									},
									"cloud_availability": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cloud": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"regions": {
													Type:        schema.TypeList,
													Description: "The regions details",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"region": {
																Type:        schema.TypeString,
																Description: "The cloud region name",
																Computed:    true,
															},
															"availability_zones": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"clone_info": {
										Type:        schema.TypeMap,
										Description: "Miscellaneous information",
										Computed:    true,
									},
									"owner": {
										Type:        schema.TypeString,
										Description: "The user who created database clone",
										Computed:    true,
									},
									"date_created": {
										Type:        schema.TypeString,
										Description: "Timestamp when the entity was created",
										Computed:    true,
									},
								},
							},
						},
						"date_created": {
							Type:        schema.TypeString,
							Description: "The timestamp when the Availability Machine was incarnated",
							Computed:    true,
						},
						"date_modified": {
							Type:        schema.TypeString,
							Description: "The timestamp when the Availability Machine was last updated",
							Computed:    true,
						},
						"tsm": {
							Type:        schema.TypeBool,
							Description: "Specify whether the associated DB Service is created using TSM compute type",
							Computed:    true,
						},
						"backup_download_config": {
							Type:        schema.TypeList,
							Description: "This is a definition for backup download config",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_backup_downloads_for_all_users": {
										Type:        schema.TypeBool,
										Description: "Allow all users to download the backup, if false only owner/co-owner(s) will be allowed",
										Computed:    true,
									},
									"allow_backup_downloads": {
										Type:        schema.TypeBool,
										Description: "Allow download of the backup for owner/co-owner of the AM",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Availability Machine",
				Optional:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "status",
				Optional:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "Availaility Machine's engine-types",
				Optional:    true,
			},
			"owners": {
				Type:        schema.TypeList,
				Description: "List of Email Addresses for entity or resource owners",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"load_acls": {
				Type:        schema.TypeBool,
				Description: "Load ACL information",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func dataSourceAvailabilityMachinesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	engineType := d.Get("engine_type").(string)
	owners := *helper.InterfaceToStringSlice(d.Get("owners"))
	loadAcls := d.Get("load_acls").(bool)
	status := d.Get("status").(string)

	response, _, err := client.GetAvailabilityMachines(name, status, engineType, loadAcls, owners)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("AvailabilityMachineList")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, AvailabilityMachineList *[]model.TessellDMMServiceConsumerDTO) error {
	parsedAvailabilityMachineList := make([]interface{}, 0)

	if AvailabilityMachineList != nil {
		parsedAvailabilityMachineList = make([]interface{}, len(*AvailabilityMachineList))
		for i, AvailabilityMachine := range *AvailabilityMachineList {
			parsedAvailabilityMachineList[i] = map[string]interface{}{
				"id":                     AvailabilityMachine.Id,
				"tessell_service_id":     AvailabilityMachine.TessellServiceId,
				"service_name":           AvailabilityMachine.ServiceName,
				"tenant":                 AvailabilityMachine.Tenant,
				"subscription":           AvailabilityMachine.Subscription,
				"engine_type":            AvailabilityMachine.EngineType,
				"data_ingestion_status":  AvailabilityMachine.DataIngestionStatus,
				"user_id":                AvailabilityMachine.UserId,
				"owner":                  AvailabilityMachine.Owner,
				"logged_in_user_role":    AvailabilityMachine.LoggedInUserRole,
				"shared_with":            []interface{}{parseEntityAclSharingInfo(AvailabilityMachine.SharedWith)},
				"cloud_availability":     parseCloudRegionInfoList(AvailabilityMachine.CloudAvailability),
				"rpo_sla":                []interface{}{parseTessellDMMAvailabilityServiceView(AvailabilityMachine.RPOSLA)},
				"daps":                   parseTessellDAPServiceDTOList(AvailabilityMachine.DAPs),
				"clones":                 parseTessellCloneSummaryInfoList(AvailabilityMachine.Clones),
				"date_created":           AvailabilityMachine.DateCreated,
				"date_modified":          AvailabilityMachine.DateModified,
				"tsm":                    AvailabilityMachine.Tsm,
				"backup_download_config": []interface{}{parseBackupDownloadConfig(AvailabilityMachine.BackupDownloadConfig)},
			}
		}
	}

	if err := d.Set("availability_machines", parsedAvailabilityMachineList); err != nil {
		return err
	}
	return nil
}
