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
						"rpo_policy": {
							Type:        schema.TypeList,
							Description: "This is the definition for RPO Policy details for Tessell DB Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include_transaction_logs": {
										Type:        schema.TypeBool,
										Description: "Determines whether transaction logs should be retained to enable Point-In-Time Recovery (PITR) functionality",
										Computed:    true,
									},
									"enable_auto_snapshot": {
										Type:        schema.TypeBool,
										Description: "Specify whether system will take automatic snapshots",
										Computed:    true,
									},
									"standard_policy": {
										Type:        schema.TypeList,
										Description: "This is the definition of Standard RPO Policy for Snapshot for Tessell DB Service",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"retention_days": {
													Type:        schema.TypeInt,
													Description: "Number of days for which the snapshot of DB Service would be retained",
													Computed:    true,
												},
												"include_transaction_logs": {
													Type:        schema.TypeBool,
													Description: "Determines whether transaction logs should be retained to enable Point-In-Time Recovery (PITR) functionality",
													Computed:    true,
												},
												"snapshot_start_time": {
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
											},
										},
									},
									"custom_policy": {
										Type:        schema.TypeList,
										Description: "This is the definition of Custom RPO Policy for Snapshot for Tessell DB Service",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "Custom RPO policy name",
													Computed:    true,
												},
												"schedule": {
													Type:        schema.TypeList,
													Description: "",
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
									"full_backup_schedule": {
										Type:        schema.TypeList,
										Description: "The schedule at which full backups would be triggered",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"start_time": {
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
											},
										},
									},
									"enable_auto_backup": {
										Type:        schema.TypeBool,
										Description: "Specify whether system will take automatic backups",
										Computed:    true,
									},
									"backup_rpo_config": {
										Type:        schema.TypeList,
										Description: "Config for the Native Backup RPO policies",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"full_backup_schedule": {
													Type:        schema.TypeList,
													Description: "The schedule at which full backups would be triggered",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"start_time": {
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
														},
													},
												},
												"standard_policy": {
													Type:        schema.TypeList,
													Description: "This is the definition of Standard RPO Policy for Backup for Tessell DB Service",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"retention_days": {
																Type:        schema.TypeInt,
																Description: "Number of days for which the backup of DB Service would be retained",
																Computed:    true,
															},
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
														},
													},
												},
												"custom_policy": {
													Type:        schema.TypeList,
													Description: "This is the definition of Custom RPO Policy for Backup for Tessell DB Service",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:        schema.TypeString,
																Description: "Custom RPO policy name",
																Computed:    true,
															},
															"schedule": {
																Type:        schema.TypeList,
																Description: "",
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
									"data_access_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pitr": {
													Type:        schema.TypeInt,
													Description: "Retention time (in days) for Point-In-Time recoverability",
													Computed:    true,
												},
												"daily_backups": {
													Type:        schema.TypeInt,
													Description: "Retention time (in days) to retain daily snapshots",
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
									"subscriptions_cloud_locations_and_key": {
										Type:        schema.TypeList,
										Description: "The subscription, cloud and region information along with encryption key and user info for DAP",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"subscription_name": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"cloud_region_and_key": {
													Type:        schema.TypeMap,
													Description: "",
													Computed:    true,
												},
												"users": {
													Type:        schema.TypeList,
													Description: "List of users email id who have access to the data/content managed by this Access Policy",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
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
									"instances": {
										Type:        schema.TypeList,
										Description: "Instances associated with this DB Service",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Description: "Tessell generated UUID for the DB Service Instance",
													Computed:    true,
												},
												"name": {
													Type:        schema.TypeString,
													Description: "Name of the DB Service Instance",
													Computed:    true,
												},
												"compute_name": {
													Type:        schema.TypeString,
													Description: "compute-name of the DB Service Instance on Cloud",
													Computed:    true,
												},
												"description": {
													Type:        schema.TypeString,
													Description: "DB Service Instance description",
													Computed:    true,
												},
												"tessell_service_id": {
													Type:        schema.TypeString,
													Description: "DB Service Instance's associated DB Service ID",
													Computed:    true,
												},
												"compute_resource_id": {
													Type:        schema.TypeString,
													Description: "Associated compute resource ID",
													Computed:    true,
												},
												"cloud_location_id": {
													Type:        schema.TypeString,
													Description: "DB Service Instance's cloud location",
													Computed:    true,
												},
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "Parameter Profile linked with the DB service instance",
													Computed:    true,
												},
												"cloud_account_id": {
													Type:        schema.TypeString,
													Description: "The cloud account on which the instance is hosted",
													Computed:    true,
												},
												"instance_group_id": {
													Type:        schema.TypeString,
													Description: "The instance group Id",
													Computed:    true,
												},
												"type": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"role": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"user_visible_role": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"status": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"plugin_status": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"connection_info": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"connect_string": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"connect_descriptor": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"master_user": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"endpoint": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"service_port": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"end_points": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"endpoint": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"labels": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"data": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"data": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
															},
														},
													},
												},
												"generic_info": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"part_of_initial_primary_set": {
																Type:        schema.TypeBool,
																Description: "",
																Computed:    true,
															},
															"encryption_key": {
																Type:        schema.TypeString,
																Description: "The encryption key name which is used to encrypt the data at rest",
																Computed:    true,
															},
															"encryption_key_id": {
																Type:        schema.TypeString,
																Description: "The encryption key id which is used to encrypt the data at rest",
																Computed:    true,
															},
															"server_cert_id": {
																Type:        schema.TypeString,
																Description: "The CA certificate id which is configured for this instance",
																Computed:    true,
															},
															"vpc": {
																Type:        schema.TypeString,
																Description: "The VPC to be used for provisioning the instance",
																Computed:    true,
															},
															"vpc_id": {
																Type:        schema.TypeString,
																Description: "The VPC Id which is used for provisioning the instance",
																Computed:    true,
															},
															"public_subnet": {
																Type:        schema.TypeString,
																Description: "The public subnet used for provisioning the instance",
																Computed:    true,
															},
															"public_subnet_id": {
																Type:        schema.TypeString,
																Description: "The public subnet Id which is used for provisioning the instance",
																Computed:    true,
															},
															"private_subnet": {
																Type:        schema.TypeString,
																Description: "The private subnet used for provisioning the instance",
																Computed:    true,
															},
															"private_subnet_id": {
																Type:        schema.TypeString,
																Description: "The private subnet Id which is used for provisioning the instance",
																Computed:    true,
															},
															"network_profile_id": {
																Type:        schema.TypeString,
																Description: "The network-profile-id which is used for provisioning the instance",
																Computed:    true,
															},
															"compute_type": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"compute_id": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"aws_infra_config": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aws_cpu_options": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"vcpus": {
																						Type:        schema.TypeInt,
																						Description: "Number of vcpus for aws cpu options",
																						Computed:    true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"base_storage": {
																Type:        schema.TypeInt,
																Description: "The base storage (in bytes) that has been provisioned for the DB Service instance.",
																Computed:    true,
															},
															"additional_storage": {
																Type:        schema.TypeInt,
																Description: "The additional storage (in bytes) to be provisioned for the DB Service instance. This is in addition to what is specified in the compute type.",
																Computed:    true,
															},
															"allocated_storage": {
																Type:        schema.TypeInt,
																Description: "The actual storage (in bytes) that has been provisioned for the DB Service instance.",
																Computed:    true,
															},
															"max_memory": {
																Type:        schema.TypeInt,
																Description: "The allocated max memory (in bytes) for this instance",
																Computed:    true,
															},
															"software_image": {
																Type:        schema.TypeString,
																Description: "The software-image-name which is used for provisioning this instance",
																Computed:    true,
															},
															"software_image_version": {
																Type:        schema.TypeString,
																Description: "The software-image-version-name which is used for provisioning this instance",
																Computed:    true,
															},
															"software_image_id": {
																Type:        schema.TypeString,
																Description: "The software-image-id which is used for provisioning this instance",
																Computed:    true,
															},
															"software_image_version_id": {
																Type:        schema.TypeString,
																Description: "The software-image-version-id which is used for provisioning this instance",
																Computed:    true,
															},
															"data_volume_iops": {
																Type:        schema.TypeInt,
																Description: "",
																Computed:    true,
															},
															"throughput": {
																Type:        schema.TypeInt,
																Description: "Throughput requested for this DB Service instance",
																Computed:    true,
															},
															"multi_disk": {
																Type:        schema.TypeBool,
																Description: "Specify whether the DB service uses multiple data disks",
																Computed:    true,
															},
															"parameter_profile": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:        schema.TypeString,
																			Description: "Tessell generated UUID for the the parameter profile",
																			Computed:    true,
																		},
																		"name": {
																			Type:        schema.TypeString,
																			Description: "The name used to identify the parameter profile",
																			Computed:    true,
																		},
																		"version": {
																			Type:        schema.TypeString,
																			Description: "The version of the parameter profile associated with the instance",
																			Computed:    true,
																		},
																		"status": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"option_profile": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:        schema.TypeString,
																			Description: "Tessell generated UUID for the the option profile",
																			Computed:    true,
																		},
																		"name": {
																			Type:        schema.TypeString,
																			Description: "The name used to identify the option profile",
																			Computed:    true,
																		},
																		"status": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"parameter_profile_id": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"sync_mode": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"engine_configuration": {
																Type:        schema.TypeList,
																Description: "This field details the DB Service Instance engine configuration details like - access mode",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"oracle_config": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_mode": {
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
															"compute_config": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"provider": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"exadata_config": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"infrastructure_id": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"infrastructure_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"vm_cluster_id": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"vm_cluster_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"vcpus": {
																						Type:        schema.TypeInt,
																						Description: "",
																						Computed:    true,
																					},
																					"memory": {
																						Type:        schema.TypeInt,
																						Description: "",
																						Computed:    true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"storage_config": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"provider": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"fsx_net_app_config": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"file_system_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"svm_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"volume_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"file_system_id": {
																						Type:        schema.TypeString,
																						Description: "File System Id of the FSx NetApp registered with Tessell",
																						Computed:    true,
																					},
																					"svm_id": {
																						Type:        schema.TypeString,
																						Description: "Storage Virtual Machine Id of the FSx NetApp registered with Tessell",
																						Computed:    true,
																					},
																				},
																			},
																		},
																		"azure_net_app_config": {
																			Type:        schema.TypeList,
																			Description: "Service instance level Azure NetApp config",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"azure_net_app_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"capacity_pool_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"volume_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"azure_net_app_id": {
																						Type:        schema.TypeString,
																						Description: "Azure NetApp Id registered with Tessell",
																						Computed:    true,
																					},
																					"capacity_pool_id": {
																						Type:        schema.TypeString,
																						Description: "Capacity Pool Id of the Azure NetApp registered with Tessell",
																						Computed:    true,
																					},
																					"delegated_subnet_id": {
																						Type:        schema.TypeString,
																						Description: "Delegated Subnet name registered with Tessell for the Azure NetApp volume",
																						Computed:    true,
																					},
																					"delegated_subnet_name": {
																						Type:        schema.TypeString,
																						Description: "Delegated Subnet Id registered with Tessell for the Azure NetApp volume",
																						Computed:    true,
																					},
																					"encryption_key_info": {
																						Type:        schema.TypeList,
																						Description: "Details of encryption key",
																						Computed:    true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"id": {
																									Type:        schema.TypeString,
																									Description: "Id of the encryption key",
																									Computed:    true,
																								},
																								"name": {
																									Type:        schema.TypeString,
																									Description: "name of the encryption key",
																									Computed:    true,
																								},
																								"key_vault_cloud_resource_id": {
																									Type:        schema.TypeString,
																									Description: "name of the encryption key vault in cloud",
																									Computed:    true,
																								},
																								"key_source": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																							},
																						},
																					},
																					"network_features": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"service_level": {
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
															"archive_storage_config": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"provider": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"fsx_net_app_config": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"file_system_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"svm_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"volume_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"file_system_id": {
																						Type:        schema.TypeString,
																						Description: "File System Id of the FSx NetApp registered with Tessell",
																						Computed:    true,
																					},
																					"svm_id": {
																						Type:        schema.TypeString,
																						Description: "Storage Virtual Machine Id of the FSx NetApp registered with Tessell",
																						Computed:    true,
																					},
																				},
																			},
																		},
																		"azure_net_app_config": {
																			Type:        schema.TypeList,
																			Description: "Service instance level Azure NetApp config",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"azure_net_app_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"capacity_pool_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"volume_name": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"azure_net_app_id": {
																						Type:        schema.TypeString,
																						Description: "Azure NetApp Id registered with Tessell",
																						Computed:    true,
																					},
																					"capacity_pool_id": {
																						Type:        schema.TypeString,
																						Description: "Capacity Pool Id of the Azure NetApp registered with Tessell",
																						Computed:    true,
																					},
																					"delegated_subnet_id": {
																						Type:        schema.TypeString,
																						Description: "Delegated Subnet name registered with Tessell for the Azure NetApp volume",
																						Computed:    true,
																					},
																					"delegated_subnet_name": {
																						Type:        schema.TypeString,
																						Description: "Delegated Subnet Id registered with Tessell for the Azure NetApp volume",
																						Computed:    true,
																					},
																					"encryption_key_info": {
																						Type:        schema.TypeList,
																						Description: "Details of encryption key",
																						Computed:    true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"id": {
																									Type:        schema.TypeString,
																									Description: "Id of the encryption key",
																									Computed:    true,
																								},
																								"name": {
																									Type:        schema.TypeString,
																									Description: "name of the encryption key",
																									Computed:    true,
																								},
																								"key_vault_cloud_resource_id": {
																									Type:        schema.TypeString,
																									Description: "name of the encryption key vault in cloud",
																									Computed:    true,
																								},
																								"key_source": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																							},
																						},
																					},
																					"network_features": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"service_level": {
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
														},
													},
												},
												"license_info": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"licenses": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"license_id": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"lock_hash": {
																			Type:        schema.TypeString,
																			Description: "Acquired licenses lock-hash",
																			Computed:    true,
																		},
																		"quantity": {
																			Type:        schema.TypeFloat,
																			Description: "quantity of acquired license",
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"monitoring_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"perf_insights": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"perf_insights_enabled": {
																			Type:        schema.TypeBool,
																			Description: "",
																			Computed:    true,
																		},
																		"monitoring_deployment_id": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"status": {
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
													Description: "Timestamp when the entity was created",
													Computed:    true,
												},
												"date_modified": {
													Type:        schema.TypeString,
													Description: "Timestamp when the entity was last modified, either by system or by user",
													Computed:    true,
												},
												"date_modified_by_user": {
													Type:        schema.TypeString,
													Description: "Timestamp when the entity was last modified by the user",
													Computed:    true,
												},
												"metadata": {
													Type:        schema.TypeList,
													Description: "DB Service Instance's metadata information",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"instance_group_name": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"deletion_attempts": {
																Type:        schema.TypeInt,
																Description: "",
																Computed:    true,
															},
															"last_deletion_dispatch_time": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"last_resize_dispatch_time": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"last_storage_resize_dispatch_time": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"add_replica_context_id": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"data": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
															},
															"is_dcr_enabled": {
																Type:        schema.TypeBool,
																Description: "",
																Computed:    true,
															},
														},
													},
												},
												"driver_info": {
													Type:        schema.TypeList,
													Description: "DB Service Instance's metadata information for plugin use",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"data": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
															},
														},
													},
												},
												"updates_in_progress_info": {
													Type:        schema.TypeList,
													Description: "DB Service Instance's in progress updates",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"infra": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"resource_update_info": {
																			Type:        schema.TypeList,
																			Description: "In progress update information for a Tessell Resource",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"update_type": {
																						Type:        schema.TypeString,
																						Description: "Type of the update",
																						Computed:    true,
																					},
																					"reference_id": {
																						Type:        schema.TypeString,
																						Description: "The reference-id of the update request",
																						Computed:    true,
																					},
																					"submitted_at": {
																						Type:        schema.TypeString,
																						Description: "Timestamp when the resource update was requested",
																						Computed:    true,
																					},
																					"update_info": {
																						Type:        schema.TypeMap,
																						Description: "The specific details for a Tessell resource that are being updated",
																						Computed:    true,
																					},
																				},
																			},
																		},
																		"infra_update_info": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"compute_type": {
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
														},
													},
												},
												"last_started_at": {
													Type:        schema.TypeString,
													Description: "Timestamp when the service instance was last started at",
													Computed:    true,
												},
												"last_stopped_at": {
													Type:        schema.TypeString,
													Description: "Timestamp when the Service Instance was last stopped at",
													Computed:    true,
												},
												"last_degraded_at": {
													Type:        schema.TypeString,
													Description: "Timestamp when the Service Instance was DEGRADED",
													Computed:    true,
												},
												"deleted_for_user_at": {
													Type:        schema.TypeString,
													Description: "Timestamp when the service instance was marked 'delete for user'.",
													Computed:    true,
												},
												"is_consumable": {
													Type:        schema.TypeBool,
													Description: "Whether the service instance is consumable for purposes like billing",
													Computed:    true,
												},
												"tessell_agent_lcm_info": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"compute_resource_info": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
															},
															"service_info": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
															},
															"instance_info": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
															},
															"data": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
															},
														},
													},
												},
												"compute_resource": {
													Type:        schema.TypeList,
													Description: "This is a definition for Tessell Compute Resource Object",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeString,
																Description: "Tessell generated UUID for the entity",
																Computed:    true,
															},
															"name": {
																Type:        schema.TypeString,
																Description: "Name of the entity",
																Computed:    true,
															},
															"description": {
																Type:        schema.TypeString,
																Description: "Compute Resource description",
																Computed:    true,
															},
															"tenant_id": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"subscription_id": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"engine_type": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"status": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"tsm": {
																Type:        schema.TypeBool,
																Description: "",
																Computed:    true,
															},
															"cloud_status": {
																Type:        schema.TypeString,
																Description: "Compute Resource's status in the cloud",
																Computed:    true,
															},
															"compute_sharing_enabled": {
																Type:        schema.TypeBool,
																Description: "Whether the Compute Resource is shared across multiple DB Services",
																Computed:    true,
															},
															"cloud_account_id": {
																Type:        schema.TypeString,
																Description: "Compute Resource's Tessell cloud account identifier",
																Computed:    true,
															},
															"cloud_location": {
																Type:        schema.TypeString,
																Description: "Compute Resource's location in the cloud",
																Computed:    true,
															},
															"cloud_resource_id": {
																Type:        schema.TypeString,
																Description: "Compute Resource's cloud identifier",
																Computed:    true,
															},
															"type": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"machine_type": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"os_info": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"image": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"software_image_id": {
																Type:        schema.TypeString,
																Description: "Compute Resource's Software Image Id",
																Computed:    true,
															},
															"software_image_version_id": {
																Type:        schema.TypeString,
																Description: "Compute Resource's Software Image Version Id",
																Computed:    true,
															},
															"network_profile_id": {
																Type:        schema.TypeString,
																Description: "Compute Resource's Network Profile Id",
																Computed:    true,
															},
															"compute_type_id": {
																Type:        schema.TypeString,
																Description: "Compute Resource's compute type Id",
																Computed:    true,
															},
															"user_id": {
																Type:        schema.TypeString,
																Description: "Compute Resource's user id",
																Computed:    true,
															},
															"owner": {
																Type:        schema.TypeString,
																Description: "Compute resource's owner email address",
																Computed:    true,
															},
															"date_created": {
																Type:        schema.TypeString,
																Description: "Timestamp when the entity was created",
																Computed:    true,
															},
															"date_modified": {
																Type:        schema.TypeString,
																Description: "Timestamp when the entity was last modified",
																Computed:    true,
															},
															"timezone": {
																Type:        schema.TypeString,
																Description: "The timezone detail",
																Computed:    true,
															},
															"machine_fqdn_info": {
																Type:        schema.TypeList,
																Description: "Compute Resource's machineFqdnInfo",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"data": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"ip_address_info": {
																Type:        schema.TypeList,
																Description: "Compute Resource's IP Address details",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"data": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																		"ip_addresses": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"metadata": {
																Type:        schema.TypeList,
																Description: "Compute Resource's metadata information",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"subscription": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"compute_type": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"vpc": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"private_subnet_id": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"private_subnet": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"public_subnet_id": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"public_subnet": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"encryption_key_id": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"encryption_key": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"connectivity_info": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"service_port": {
																						Type:        schema.TypeInt,
																						Description: "The connection port for the DB Service",
																						Computed:    true,
																					},
																				},
																			},
																		},
																		"os_timezone": {
																			Type:        schema.TypeString,
																			Description: "The timezone detail",
																			Computed:    true,
																		},
																		"dbserver_info": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"first_provisioned_dbservice_id": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"db_service_ids": {
																						Type:        schema.TypeList,
																						Description: "The list of DB Service ids that are hosted on this compute resource",
																						Computed:    true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"enable_public_access": {
																						Type:        schema.TypeBool,
																						Description: "",
																						Computed:    true,
																					},
																					"enable_ssl": {
																						Type:        schema.TypeBool,
																						Description: "",
																						Computed:    true,
																					},
																					"software_image_info": {
																						Type:        schema.TypeList,
																						Description: "",
																						Computed:    true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"software_image": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"software_image_id": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"software_image_versions": {
																									Type:        schema.TypeList,
																									Description: "",
																									Computed:    true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"software_image_version": {
																												Type:        schema.TypeString,
																												Description: "",
																												Computed:    true,
																											},
																											"software_image_version_id": {
																												Type:        schema.TypeString,
																												Description: "",
																												Computed:    true,
																											},
																											"supported": {
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
																				},
																			},
																		},
																		"aws_infra_config": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"aws_cpu_options": {
																						Type:        schema.TypeList,
																						Description: "",
																						Computed:    true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"vcpus": {
																									Type:        schema.TypeInt,
																									Description: "Number of vcpus for aws cpu options",
																									Computed:    true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"license_info": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"acquirer_id": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"license_id": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"lock_hash": {
																						Type:        schema.TypeString,
																						Description: "Acquired licenses lock-hash",
																						Computed:    true,
																					},
																					"quantity": {
																						Type:        schema.TypeFloat,
																						Description: "quantity of acquired license",
																						Computed:    true,
																					},
																				},
																			},
																		},
																		"data": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																		"is_azure_monitor_agent_installed": {
																			Type:        schema.TypeBool,
																			Description: "",
																			Computed:    true,
																		},
																		"node_index": {
																			Type:        schema.TypeInt,
																			Description: "",
																			Computed:    true,
																		},
																		"use_azure_monitor_agent": {
																			Type:        schema.TypeBool,
																			Description: "",
																			Computed:    true,
																		},
																		"compute_config": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"provider": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"exadata_config": {
																						Type:        schema.TypeList,
																						Description: "",
																						Computed:    true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"infrastructure_id": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"infrastructure_name": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"vm_cluster_id": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"vm_cluster_name": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"ocpu": {
																									Type:        schema.TypeInt,
																									Description: "",
																									Computed:    true,
																								},
																								"memory_in_gbs": {
																									Type:        schema.TypeInt,
																									Description: "",
																									Computed:    true,
																								},
																								"floating_ip_address": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"db_server": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"private_ip_address": {
																									Type:        schema.TypeString,
																									Description: "",
																									Computed:    true,
																								},
																								"local_storage_in_gbs": {
																									Type:        schema.TypeInt,
																									Description: "",
																									Computed:    true,
																								},
																								"dns_name": {
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
																	},
																},
															},
															"driver_info": {
																Type:        schema.TypeList,
																Description: "Compute Resource's metadata information for driver use",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"data": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"tessell_stack_info": {
																Type:        schema.TypeList,
																Description: "Tessell Stack information that is present on this Compute Resource",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"data": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"tessell_agent_lcm_info": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"compute_resource_info": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																		"service_info": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																		"instance_info": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																		"data": {
																			Type:        schema.TypeList,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"tessell_endpoint_migration_info": {
																Type:        schema.TypeList,
																Description: "Compute Resource's Endpoint Migration information",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"validate_common_endpoint": {
																			Type:        schema.TypeBool,
																			Description: "",
																			Computed:    true,
																		},
																		"migrate_to_common_endpoint": {
																			Type:        schema.TypeBool,
																			Description: "",
																			Computed:    true,
																		},
																		"common_endpoint_successful_validation_time": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"common_endpoint_successful_migration_time": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"internal": {
																Type:        schema.TypeBool,
																Description: "Whether the Compute Resource is created for internal usage",
																Computed:    true,
															},
															"context_info": {
																Type:        schema.TypeList,
																Description: "Provide more context of DB Server state",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"sub_status": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																		"description": {
																			Type:        schema.TypeString,
																			Description: "",
																			Computed:    true,
																		},
																	},
																},
															},
															"action_metadata": {
																Type:        schema.TypeList,
																Description: "DB compute resource action metadata information",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"last_action_metadata": {
																			Type:        schema.TypeList,
																			Description: "Records last action metadata of a compute resource",
																			Computed:    true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"reference_id": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"context_id": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																					"action_type": {
																						Type:        schema.TypeString,
																						Description: "",
																						Computed:    true,
																					},
																				},
																			},
																		},
																		"last_resize_dispatch_time": {
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
											},
										},
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
						"storage_config": {
							Type:        schema.TypeList,
							Description: "The storage details to be provisioned.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"provider": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"fsx_net_app_config": {
										Type:        schema.TypeList,
										Description: "The FSx NetApp details to be provisioned",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"file_system_id": {
													Type:        schema.TypeString,
													Description: "File System Id of the FSx NetApp registered with Tessell",
													Computed:    true,
												},
												"svm_id": {
													Type:        schema.TypeString,
													Description: "Storage Virtual Machine Id of the FSx NetApp registered with Tessell",
													Computed:    true,
												},
											},
										},
									},
									"azure_net_app_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"azure_net_app_id": {
													Type:        schema.TypeString,
													Description: "Azure NetApp Id registered with Tessell",
													Computed:    true,
												},
												"capacity_pool_id": {
													Type:        schema.TypeString,
													Description: "Capacity pool Id of the Azure NetApp registered with Tessell",
													Computed:    true,
												},
												"configurations": {
													Type:        schema.TypeList,
													Description: "Azure NetApp configurations",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_features": {
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

func setDataSourceValues(d *schema.ResourceData, AvailabilityMachineList *[]model.DMMConsumerView) error {
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
				"topology":               parseDBServiceTopologyList(AvailabilityMachine.Topology),
				"rpo_policy":             []interface{}{parseRPOPolicyConfig(AvailabilityMachine.RPOPolicy)},
				"daps":                   parseTessellDAPServiceDTOList(AvailabilityMachine.DAPs),
				"clones":                 parseTessellCloneSummaryInfoList(AvailabilityMachine.Clones),
				"date_created":           AvailabilityMachine.DateCreated,
				"date_modified":          AvailabilityMachine.DateModified,
				"tsm":                    AvailabilityMachine.Tsm,
				"backup_download_config": []interface{}{parseBackupDownloadConfig(AvailabilityMachine.BackupDownloadConfig)},
				"storage_config":         []interface{}{parseStorageConfigPayload(AvailabilityMachine.StorageConfig)},
			}
		}
	}

	if err := d.Set("availability_machines", parsedAvailabilityMachineList); err != nil {
		return err
	}
	return nil
}
