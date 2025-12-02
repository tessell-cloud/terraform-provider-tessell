package db_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func DataSourceDBServices() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDBServicesRead,

		Schema: map[string]*schema.Schema{
			"db_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Tessell generated UUID for the DB Service. This is the unique identifier for the DB Service.",
							Computed:    true,
						},
						"availability_machine_id": {
							Type:        schema.TypeString,
							Description: "Unique ID of the associated Availability Machine",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the DB Service",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "User specified description for the DB Service",
							Computed:    true,
						},
						"engine_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"topology": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"num_of_instances": {
							Type:        schema.TypeInt,
							Description: "This specifies the number of instances (nodes) that are created for the DB Service",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"context_info": {
							Type:        schema.TypeList,
							Description: "Provide more context of DB Service state",
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
						"license_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"auto_minor_version_update": {
							Type:        schema.TypeBool,
							Description: "This field specifies whether to automatically update minor version for the DB Service",
							Computed:    true,
						},
						"enable_deletion_protection": {
							Type:        schema.TypeBool,
							Description: "This field specifies whether to enable deletion protection for the DB Service. If this is enabled, the deletion for the DB Service would be disallowed until this setting is disabled.",
							Computed:    true,
						},
						"enable_stop_protection": {
							Type:        schema.TypeBool,
							Description: "This field specifies whether to enable stop protection for the DB Service. If this is enabled, the stop for the DB Service would be disallowed until this setting is disabled.",
							Computed:    true,
						},
						"enable_perf_insights": {
							Type:        schema.TypeBool,
							Description: "This field specifies whether to enable performance insights for the DB Service.",
							Computed:    true,
						},
						"edition": {
							Type:        schema.TypeString,
							Description: "Edition of the software image that has been used to create the DB Service (e.g. COMMUNITY/ENTERPRISE etc)",
							Computed:    true,
						},
						"software_image": {
							Type:        schema.TypeString,
							Description: "The software image that has been used to create the DB Service",
							Computed:    true,
						},
						"software_image_version": {
							Type:        schema.TypeString,
							Description: "The software image version that is used to create the DB Service",
							Computed:    true,
						},
						"software_image_version_family": {
							Type:        schema.TypeString,
							Description: "The software image version family that is used to create the DB Service",
							Computed:    true,
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Description: "The tenant identifier under which the DB Service is created",
							Computed:    true,
						},
						"subscription": {
							Type:        schema.TypeString,
							Description: "The Tessell Subscription under which this DB Service is created",
							Computed:    true,
						},
						"owner": {
							Type:        schema.TypeString,
							Description: "This field specifies who is the owner for the DB Service",
							Computed:    true,
						},
						"logged_in_user_role": {
							Type:        schema.TypeString,
							Description: "This field specifies access role on the DB Service for the currently logged in user",
							Computed:    true,
						},
						"date_created": {
							Type:        schema.TypeString,
							Description: "This field specifies the timestamp when the DB Service was created at",
							Computed:    true,
						},
						"started_at": {
							Type:        schema.TypeString,
							Description: "This field specifies the timestamp when the DB Service was last started at",
							Computed:    true,
						},
						"stopped_at": {
							Type:        schema.TypeString,
							Description: "This field specifies the timestamp when the DB Service was last stopped at",
							Computed:    true,
						},
						"cloned_from_info": {
							Type:        schema.TypeList,
							Description: "If the DB Service is created as a clone from some other DB Service, this section describes the parent DB Service and cloning details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"clone_type": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"content_type": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "The DB Service ID using which this DB Service clone is created",
										Computed:    true,
									},
									"availability_machine_id": {
										Type:        schema.TypeString,
										Description: "The Availability Machine ID using which this DB Service clone is created",
										Computed:    true,
									},
									"tessell_service": {
										Type:        schema.TypeString,
										Description: "The DB Service name using which this DB Service clone is created",
										Computed:    true,
									},
									"availability_machine": {
										Type:        schema.TypeString,
										Description: "The Availability Machine name using which this DB Service clone is created",
										Computed:    true,
									},
									"snapshot_name": {
										Type:        schema.TypeString,
										Description: "The snapshot using which this DB Service clone is created",
										Computed:    true,
									},
									"snapshot_id": {
										Type:        schema.TypeString,
										Description: "The snapshot ID using which this DB Service clone is created",
										Computed:    true,
									},
									"snapshot_time": {
										Type:        schema.TypeString,
										Description: "DB Service snapshot capture time",
										Computed:    true,
									},
									"pitr_time": {
										Type:        schema.TypeString,
										Description: "If the database was created using a Point-In-Time mechanism, it specifies the timestamp in UTC",
										Computed:    true,
									},
									"maximum_recoverability": {
										Type:        schema.TypeBool,
										Description: "If the service was created using a maximum recoverability from the parent service",
										Computed:    true,
									},
									"storage_provider": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"refresh_info": {
							Type:        schema.TypeList,
							Description: "Service refresh details",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content_type": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"snapshot_name": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"snapshot_time": {
										Type:        schema.TypeString,
										Description: "Time at which snapshot is created.",
										Computed:    true,
									},
									"pitr": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"script_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pre_script_info": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"script_id": {
																Type:        schema.TypeString,
																Description: "The Tessell Script ID",
																Computed:    true,
															},
															"script_version": {
																Type:        schema.TypeString,
																Description: "The Tessell Script version",
																Computed:    true,
															},
														},
													},
												},
												"post_script_info": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"script_id": {
																Type:        schema.TypeString,
																Description: "The Tessell Script ID",
																Computed:    true,
															},
															"script_version": {
																Type:        schema.TypeString,
																Description: "The Tessell Script version",
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"schedule_id": {
										Type:        schema.TypeString,
										Description: "If refreshed using schedule then schedule id, else null",
										Computed:    true,
									},
									"last_successful_refresh_time": {
										Type:        schema.TypeString,
										Description: "Time at which refresh would be successfully completed.",
										Computed:    true,
									},
								},
							},
						},
						"service_connectivity": {
							Type:        schema.TypeList,
							Description: "DB Service's connectivity information",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_ssl": {
										Type:        schema.TypeBool,
										Description: "Specify whether to enable SSL to the DB Service, default false",
										Computed:    true,
									},
									"ca_cert_id": {
										Type:        schema.TypeString,
										Description: "The CA certificate ID associated with the DB Service",
										Computed:    true,
									},
									"dns_prefix": {
										Type:        schema.TypeString,
										Description: "DNS Prefix associated with the DB Service",
										Computed:    true,
									},
									"service_port": {
										Type:        schema.TypeInt,
										Description: "The connection port for the DB Service",
										Computed:    true,
									},
									"enable_public_access": {
										Type:        schema.TypeBool,
										Description: "Specify whether to enable public access to the DB Service, default false",
										Computed:    true,
									},
									"allowed_ip_addresses": {
										Type:        schema.TypeList,
										Description: "The list of allowed ipv4 addresses that can connect to the DB Service",
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"connect_strings": {
										Type:        schema.TypeList,
										Description: "The list of connect strings for the DB Service",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"usage_type": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"connect_descriptor": {
													Type:        schema.TypeString,
													Description: "The connection description for the DB Service",
													Computed:    true,
												},
												"endpoint": {
													Type:        schema.TypeString,
													Description: "The connection end point for the DB Service",
													Computed:    true,
												},
												"master_user": {
													Type:        schema.TypeString,
													Description: "The master user name for the DB Service",
													Computed:    true,
												},
												"service_port": {
													Type:        schema.TypeInt,
													Description: "The connection port for the DB Service",
													Computed:    true,
												},
											},
										},
									},
									"private_link": {
										Type:        schema.TypeList,
										Description: "The interface endpoint or Gateway Load Balancer endpoint to connect to your DB service.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"status": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"service_principals": {
													Type:        schema.TypeList,
													Description: "The list of AWS account principals that are currently enabled",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"endpoint_service_name": {
													Type:        schema.TypeString,
													Description: "The configured endpoint as a result of configuring the service-principals",
													Computed:    true,
												},
												"client_azure_subscription_ids": {
													Type:        schema.TypeList,
													Description: "The list of Azure subscription Ids",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"private_link_service_alias": {
													Type:        schema.TypeString,
													Description: "The Azure private link service alias",
													Computed:    true,
												},
											},
										},
									},
									"computes_connectivity": {
										Type:        schema.TypeList,
										Description: "The Genie endpoint to connect to your DB service.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"compute_resource_id": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"port_access_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"port": {
																Type:        schema.TypeInt,
																Description: "The connection port for the DB Service",
																Computed:    true,
															},
															"enable_public_access": {
																Type:        schema.TypeBool,
																Description: "Enable public access to database (true/false)",
																Computed:    true,
															},
															"allowed_ip_addresses": {
																Type:        schema.TypeList,
																Description: "Set allowed IP address if enablePublicAccess is true.",
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
									"update_in_progress_info": {
										Type:        schema.TypeList,
										Description: "DB Service connectivity update-in-progress details",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_prefix": {
													Type:        schema.TypeString,
													Description: "The DNS prefix associated with the DB Service",
													Computed:    true,
												},
												"enable_public_access": {
													Type:        schema.TypeBool,
													Description: "Specify whether to enable public access to the DB Service, default false",
													Computed:    true,
												},
												"allowed_ip_addresses": {
													Type:        schema.TypeList,
													Description: "The list of allowed ipv4 addresses that can connect to the DB Service",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"private_link": {
													Type:        schema.TypeList,
													Description: "The interface endpoint or Gateway Load Balancer endpoint to connect to your DB service.",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"service_principals": {
																Type:        schema.TypeList,
																Description: "The list of AWS account principals that are currently enabled",
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"client_azure_subscription_ids": {
																Type:        schema.TypeList,
																Description: "The list of Azure subscription Ids",
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"computes_connectivity": {
													Type:        schema.TypeList,
													Description: "The Genie endpoint to connect to your DB service.",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"compute_resource_id": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"port_access_config": {
																Type:        schema.TypeList,
																Description: "",
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"port": {
																			Type:        schema.TypeInt,
																			Description: "The connection port for the DB Service",
																			Computed:    true,
																		},
																		"enable_public_access": {
																			Type:        schema.TypeBool,
																			Description: "Enable public access to database (true/false)",
																			Computed:    true,
																		},
																		"allowed_ip_addresses": {
																			Type:        schema.TypeList,
																			Description: "Set allowed IP address if enablePublicAccess is true.",
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
											},
										},
									},
								},
							},
						},
						"tessell_genie_status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"infrastructure": {
							Type:        schema.TypeList,
							Description: "This field contains DB Service's infrastructure related information, like, where the service is hosted - cloud, region; what compute shape, or network is is configured with.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud": {
										Type:        schema.TypeString,
										Description: "The cloud-type in which the DB Service is provisioned (ex. aws, azure)",
										Computed:    true,
									},
									"region": {
										Type:        schema.TypeString,
										Description: "The region in which the DB Service provisioned",
										Computed:    true,
									},
									"availability_zone": {
										Type:        schema.TypeString,
										Description: "The availability-zone in which the DB Service is provisioned",
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
									"vpc": {
										Type:        schema.TypeString,
										Description: "The VPC to be used for provisioning the DB Service",
										Computed:    true,
									},
									"enable_encryption": {
										Type:        schema.TypeBool,
										Description: "",
										Computed:    true,
									},
									"encryption_key": {
										Type:        schema.TypeString,
										Description: "The encryption key name which is used to encrypt the data at rest",
										Computed:    true,
									},
									"compute_type": {
										Type:        schema.TypeString,
										Description: "The compute-type to be used for provisioning the DB Service",
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
									"enable_compute_sharing": {
										Type:        schema.TypeBool,
										Description: "Specify if the computes should be shared across DB Services",
										Computed:    true,
									},
									"iops": {
										Type:        schema.TypeInt,
										Description: "IOPS requested for the DB Service",
										Computed:    true,
									},
									"throughput": {
										Type:        schema.TypeInt,
										Description: "throughput requested for the DB Service",
										Computed:    true,
									},
									"storage": {
										Type:        schema.TypeInt,
										Description: "The storage (in bytes) that has been provisioned for the DB Service",
										Computed:    true,
									},
									"additional_storage": {
										Type:        schema.TypeInt,
										Description: "Storage in bytes that is over and above the storage included with compute. This is maintained for backward compatibility and would be deprecated soon.",
										Computed:    true,
									},
									"timezone": {
										Type:        schema.TypeString,
										Description: "The timezone detail",
										Computed:    true,
									},
									"multi_disk": {
										Type:        schema.TypeBool,
										Description: "Specify whether the DB service uses multiple data disks",
										Computed:    true,
									},
									"storage_provider": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
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
												"azure_net_app_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
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
												"azure_net_app_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
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
									"compute_provider": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"maintenance_window": {
							Type:        schema.TypeList,
							Description: "This field details the DB Service maintenance related details.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"day": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"time": {
										Type:        schema.TypeString,
										Description: "Time value in (hh:mm) format. ex. '02:00'",
										Computed:    true,
									},
									"duration": {
										Type:        schema.TypeInt,
										Description: "The duration during which the maintenance window will be allowed to trigger",
										Computed:    true,
									},
								},
							},
						},
						"engine_configuration": {
							Type:        schema.TypeList,
							Description: "This field details the DB Service engine configuration details like - parameter profile, or options profile (if applicable) are used to configure the DB Service.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oracle_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"multi_tenant": {
													Type:        schema.TypeBool,
													Description: "Specify whether the DB Service is multi-tenant.",
													Computed:    true,
												},
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile id for the database",
													Computed:    true,
												},
												"options_profile": {
													Type:        schema.TypeString,
													Description: "The options profile for the database",
													Computed:    true,
												},
												"option_profile_id": {
													Type:        schema.TypeString,
													Description: "The options profile for the database",
													Computed:    true,
												},
												"sid": {
													Type:        schema.TypeString,
													Description: "SID for oracle database",
													Computed:    true,
												},
												"character_set": {
													Type:        schema.TypeString,
													Description: "The character-set for the database",
													Computed:    true,
												},
												"national_character_set": {
													Type:        schema.TypeString,
													Description: "The national-character-set for the database",
													Computed:    true,
												},
												"enable_archive_mode": {
													Type:        schema.TypeBool,
													Description: "To explicitly enable archive mode, when PITR is disabled",
													Computed:    true,
												},
											},
										},
									},
									"postgresql_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Computed:    true,
												},
												"ad_domain_id": {
													Type:        schema.TypeString,
													Description: "Active Directory Domain ID",
													Computed:    true,
												},
												"proxy_port": {
													Type:        schema.TypeInt,
													Description: "",
													Computed:    true,
												},
												"option_profile_name": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"option_profile_id": {
													Type:        schema.TypeString,
													Description: "The options profile for the database",
													Computed:    true,
												},
											},
										},
									},
									"mysql_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Computed:    true,
												},
												"ad_domain_id": {
													Type:        schema.TypeString,
													Description: "Active Directory Domain ID",
													Computed:    true,
												},
												"option_profile_id": {
													Type:        schema.TypeString,
													Description: "The options profile for the database",
													Computed:    true,
												},
											},
										},
									},
									"sql_server_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Computed:    true,
												},
												"ad_domain_id": {
													Type:        schema.TypeString,
													Description: "Active Directory Domain ID",
													Computed:    true,
												},
												"service_account_user": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"agent_service_account_user": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"instance_name": {
													Type:        schema.TypeString,
													Description: "The named instance for SQL Server database (max 16 characters as per SQL Server limitation)",
													Computed:    true,
												},
											},
										},
									},
									"apache_kafka_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile id for the database",
													Computed:    true,
												},
											},
										},
									},
									"mongodb_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cluster_name": {
													Type:        schema.TypeString,
													Description: "The MongoDB Cluster name",
													Computed:    true,
												},
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Computed:    true,
												},
											},
										},
									},
									"milvus_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameter_profile_id": {
													Type:        schema.TypeString,
													Description: "The parameter profile ID for the database",
													Computed:    true,
												},
											},
										},
									},
									"pre_script_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"script_id": {
													Type:        schema.TypeString,
													Description: "The Tessell Script ID",
													Computed:    true,
												},
												"script_version": {
													Type:        schema.TypeString,
													Description: "The Tessell Script version",
													Computed:    true,
												},
											},
										},
									},
									"post_script_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"script_id": {
													Type:        schema.TypeString,
													Description: "The Tessell Script ID",
													Computed:    true,
												},
												"script_version": {
													Type:        schema.TypeString,
													Description: "The Tessell Script version",
													Computed:    true,
												},
											},
										},
									},
									"collation_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"collation_name": {
													Type:        schema.TypeString,
													Description: "Collation name for the database",
													Computed:    true,
												},
											},
										},
									},
									"backup_url": {
										Type:        schema.TypeString,
										Description: "The URL where the backup is stored",
										Computed:    true,
									},
									"ignore_post_script_failure": {
										Type:        schema.TypeBool,
										Description: "",
										Computed:    true,
									},
									"ignore_pre_script_failure": {
										Type:        schema.TypeBool,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"integrations_config": {
							Type:        schema.TypeList,
							Description: "Tessell provides support to integrate third party softwares with DB Services. This fields details the information about what third-party integrations are configured with the DB Service.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integrations": {
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
						"deletion_config": {
							Type:        schema.TypeList,
							Description: "If the DB Service is to be deleted, this config would be honoured if no preference is provided during deleting the service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retain_availability_machine": {
										Type:        schema.TypeBool,
										Description: "If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained",
										Computed:    true,
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "The tags associated with the DB Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Case sensitive, tag name",
										Computed:    true,
									},
									"value": {
										Type:        schema.TypeString,
										Description: "Case sensitive, tag value",
										Computed:    true,
									},
								},
							},
						},
						"updates_in_progress": {
							Type:        schema.TypeList,
							Description: "The updates that are in progress for this resource",
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
						"instances": {
							Type:        schema.TypeList,
							Description: "The instances (nodes) associated with this DB Service",
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
									"instance_group_name": {
										Type:        schema.TypeString,
										Description: "Name of the instance group",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "DB Service instance type",
										Computed:    true,
									},
									"role": {
										Type:        schema.TypeString,
										Description: "DB Service instance role",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "DB Service instance status",
										Computed:    true,
									},
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "DB Service Instance's associated DB Service id",
										Computed:    true,
									},
									"cloud": {
										Type:        schema.TypeString,
										Description: "DB Service Instance's cloud type",
										Computed:    true,
									},
									"region": {
										Type:        schema.TypeString,
										Description: "DB Service Instance's cloud region",
										Computed:    true,
									},
									"availability_zone": {
										Type:        schema.TypeString,
										Description: "DB Service Instance's cloud availability zone",
										Computed:    true,
									},
									"instance_group_id": {
										Type:        schema.TypeString,
										Description: "The instance groupd Id",
										Computed:    true,
									},
									"compute_type": {
										Type:        schema.TypeString,
										Description: "The compute used for creation of the Tessell Service Instance",
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
									"compute_id": {
										Type:        schema.TypeString,
										Description: "The associated compute identifier",
										Computed:    true,
									},
									"compute_name": {
										Type:        schema.TypeString,
										Description: "The associated compute name",
										Computed:    true,
									},
									"storage": {
										Type:        schema.TypeInt,
										Description: "The storage (in bytes) that has been provisioned for the DB Service instance.",
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
									"enable_perf_insights": {
										Type:        schema.TypeBool,
										Description: "",
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
												"version": {
													Type:        schema.TypeString,
													Description: "The version of the option profile associated with the instance",
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
									"vpc": {
										Type:        schema.TypeString,
										Description: "The VPC used for creation of the DB Service Instance",
										Computed:    true,
									},
									"public_subnet": {
										Type:        schema.TypeString,
										Description: "The public subnet used for creation of the DB Service Instance",
										Computed:    true,
									},
									"private_subnet": {
										Type:        schema.TypeString,
										Description: "The private subnet used for creation of the DB Service Instance",
										Computed:    true,
									},
									"encryption_key": {
										Type:        schema.TypeString,
										Description: "The encryption key name which is used to encrypt the data at rest",
										Computed:    true,
									},
									"software_image": {
										Type:        schema.TypeString,
										Description: "Software Image to be used to create the instance",
										Computed:    true,
									},
									"software_image_version": {
										Type:        schema.TypeString,
										Description: "Software Image Version to be used to create the instance",
										Computed:    true,
									},
									"date_created": {
										Type:        schema.TypeString,
										Description: "Timestamp when the entity was created",
										Computed:    true,
									},
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
									"updates_in_progress": {
										Type:        schema.TypeList,
										Description: "The updates that are in progress for this resource",
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
												"volume_type": {
													Type:        schema.TypeString,
													Description: "Data disk volume type",
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
												"volume_type": {
													Type:        schema.TypeString,
													Description: "Data disk volume type",
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
									"private_link_info": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"status": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"endpoint_service_name": {
													Type:        schema.TypeString,
													Description: "The configured endpoint as a result of configuring the service-principals",
													Computed:    true,
												},
												"private_link_service_alias": {
													Type:        schema.TypeString,
													Description: "The Azure private link service alias",
													Computed:    true,
												},
												"service_principals": {
													Type:        schema.TypeList,
													Description: "The list of AWS account principals that are currently enabled. This is only applicable for DB Services hosted on AWS.",
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"client_azure_subscription_ids": {
													Type:        schema.TypeList,
													Description: "The list of Azure subscription Ids. This is only applicable for DB Services hosted on AZURE.",
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
						"databases": {
							Type:        schema.TypeList,
							Description: "This field details about the databases that are created under this DB Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "Tessell generated UUID for the database",
										Computed:    true,
									},
									"database_name": {
										Type:        schema.TypeString,
										Description: "Database name",
										Computed:    true,
									},
									"description": {
										Type:        schema.TypeString,
										Description: "Database description",
										Computed:    true,
									},
									"tessell_service_id": {
										Type:        schema.TypeString,
										Description: "Associated DB Service ID",
										Computed:    true,
									},
									"engine_type": {
										Type:        schema.TypeString,
										Description: "Database Engine Type",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Database status",
										Computed:    true,
									},
									"date_created": {
										Type:        schema.TypeString,
										Description: "Timestamp when the entity was created",
										Computed:    true,
									},
									"tessell_created": {
										Type:        schema.TypeBool,
										Description: "Database created from Tessell platform",
										Computed:    true,
									},
									"cloned_from_info": {
										Type:        schema.TypeList,
										Description: "If a database is created as a clone from some other DB Service's database, this section describes the original database details",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"database_id": {
													Type:        schema.TypeString,
													Description: "The original database Id using which this database clone is created",
													Computed:    true,
												},
											},
										},
									},
									"database_configuration": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"oracle_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile_id": {
																Type:        schema.TypeString,
																Description: "The parameter profile id for the database",
																Computed:    true,
															},
															"options_profile": {
																Type:        schema.TypeString,
																Description: "The options profile for the database",
																Computed:    true,
															},
															"username": {
																Type:        schema.TypeString,
																Description: "Username for the oracle database",
																Computed:    true,
															},
															"option_profile_id": {
																Type:        schema.TypeString,
																Description: "The option profile id for the database",
																Computed:    true,
															},
														},
													},
												},
												"postgresql_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile_id": {
																Type:        schema.TypeString,
																Description: "The parameter profile ID for the database",
																Computed:    true,
															},
															"option_profile_id": {
																Type:        schema.TypeString,
																Description: "The options profile for the database",
																Computed:    true,
															},
														},
													},
												},
												"mysql_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile_id": {
																Type:        schema.TypeString,
																Description: "The parameter profile ID for the database",
																Computed:    true,
															},
															"option_profile_id": {
																Type:        schema.TypeString,
																Description: "The options profile for the database",
																Computed:    true,
															},
														},
													},
												},
												"sql_server_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile_id": {
																Type:        schema.TypeString,
																Description: "The parameter profile ID for the database",
																Computed:    true,
															},
														},
													},
												},
												"mongodb_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile_id": {
																Type:        schema.TypeString,
																Description: "The parameter profile ID for the database",
																Computed:    true,
															},
														},
													},
												},
												"milvus_config": {
													Type:        schema.TypeList,
													Description: "",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"parameter_profile_id": {
																Type:        schema.TypeString,
																Description: "The parameter profile ID for the database",
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
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
								},
							},
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
						"deletion_schedule": {
							Type:        schema.TypeList,
							Description: "Details of the deletion schedule on a DB Service",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"delete_at": {
										Type:        schema.TypeString,
										Description: "Time at which the DB Service should be deleted at",
										Computed:    true,
									},
									"deletion_config": {
										Type:        schema.TypeList,
										Description: "If the DB Service is to be deleted, this config would be honoured if no preference is provided during deleting the service",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"retain_availability_machine": {
													Type:        schema.TypeBool,
													Description: "If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"upcoming_scheduled_actions": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_stop": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:        schema.TypeString,
													Description: "Action which can be either start/stop",
													Computed:    true,
												},
												"at": {
													Type:        schema.TypeString,
													Description: "The time at which the specified action is to be performed",
													Computed:    true,
												},
											},
										},
									},
									"patch": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"at": {
													Type:        schema.TypeString,
													Description: "The updated time at which the specified action is to be performed",
													Computed:    true,
												},
												"message": {
													Type:        schema.TypeString,
													Description: "Details for the update",
													Computed:    true,
												},
											},
										},
									},
									"delete": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"at": {
													Type:        schema.TypeString,
													Description: "The scheduled time for the action to be deleted",
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
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the DB Service",
				Optional:    true,
			},
			"statuses": {
				Type:        schema.TypeList,
				Description: "statuses",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"engine_types": {
				Type:        schema.TypeList,
				Description: "DB Service's engine-types",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cloned_from_service_id": {
				Type:        schema.TypeString,
				Description: "The ID of the DB Service from which the services are cloned",
				Optional:    true,
			},
			"cloned_from_availability_machine_id": {
				Type:        schema.TypeString,
				Description: "The id of the Availability Machine from which the services are cloned",
				Optional:    true,
			},
			"load_instances": {
				Type:        schema.TypeBool,
				Description: "Load the instances that are part of the DB Service",
				Optional:    true,
				Default:     true,
			},
			"load_databases": {
				Type:        schema.TypeBool,
				Description: "Load the databases that are part of the DB Service",
				Optional:    true,
				Default:     true,
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

func dataSourceDBServicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	loadDatabases := d.Get("load_databases").(bool)
	clonedFromAvailabilityMachineId := d.Get("cloned_from_availability_machine_id").(string)
	loadAcls := d.Get("load_acls").(bool)
	name := d.Get("name").(string)
	statuses := *helper.InterfaceToStringSlice(d.Get("statuses"))
	owners := *helper.InterfaceToStringSlice(d.Get("owners"))
	engineTypes := *helper.InterfaceToStringSlice(d.Get("engine_types"))
	clonedFromServiceId := d.Get("cloned_from_service_id").(string)
	loadInstances := d.Get("load_instances").(bool)

	response, _, err := client.GetTessellServices(name, statuses, engineTypes, clonedFromServiceId, clonedFromAvailabilityMachineId, loadInstances, loadDatabases, owners, loadAcls)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("DBServiceList")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, DBServiceList *[]model.TessellServiceDTO) error {
	parsedDBServiceList := make([]interface{}, 0)

	if DBServiceList != nil {
		parsedDBServiceList = make([]interface{}, len(*DBServiceList))
		for i, DBService := range *DBServiceList {
			parsedDBServiceList[i] = map[string]interface{}{
				"id":                            DBService.Id,
				"availability_machine_id":       DBService.AvailabilityMachineId,
				"name":                          DBService.Name,
				"description":                   DBService.Description,
				"engine_type":                   DBService.EngineType,
				"topology":                      DBService.Topology,
				"num_of_instances":              DBService.NumOfInstances,
				"status":                        DBService.Status,
				"context_info":                  []interface{}{parseTessellServiceContextInfo(DBService.ContextInfo)},
				"license_type":                  DBService.LicenseType,
				"auto_minor_version_update":     DBService.AutoMinorVersionUpdate,
				"enable_deletion_protection":    DBService.EnableDeletionProtection,
				"enable_stop_protection":        DBService.EnableStopProtection,
				"enable_perf_insights":          DBService.EnablePerfInsights,
				"edition":                       DBService.Edition,
				"software_image":                DBService.SoftwareImage,
				"software_image_version":        DBService.SoftwareImageVersion,
				"software_image_version_family": DBService.SoftwareImageVersionFamily,
				"tenant_id":                     DBService.TenantId,
				"subscription":                  DBService.Subscription,
				"owner":                         DBService.Owner,
				"logged_in_user_role":           DBService.LoggedInUserRole,
				"date_created":                  DBService.DateCreated,
				"started_at":                    DBService.StartedAt,
				"stopped_at":                    DBService.StoppedAt,
				"cloned_from_info":              []interface{}{parseTessellServiceClonedFromInfo(DBService.ClonedFromInfo)},
				"refresh_info":                  []interface{}{parseRefreshServiceInfo(DBService.RefreshInfo)},
				"service_connectivity":          []interface{}{parseTessellServiceConnectivityInfo(DBService.ServiceConnectivity)},
				"tessell_genie_status":          DBService.TessellGenieStatus,
				"infrastructure":                []interface{}{parseTessellServiceInfrastructureInfo(DBService.Infrastructure)},
				"maintenance_window":            []interface{}{parseTessellServiceMaintenanceWindow(DBService.MaintenanceWindow)},
				"engine_configuration":          []interface{}{parseTessellServiceEngineInfo(DBService.EngineConfiguration)},
				"integrations_config":           []interface{}{parseTessellServiceIntegrationsInfo(DBService.IntegrationsConfig)},
				"deletion_config":               []interface{}{parseTessellServiceDeletionConfig(DBService.DeletionConfig)},
				"tags":                          parseTessellTagList(DBService.Tags),
				"updates_in_progress":           parseTessellResourceUpdateInfoList(DBService.UpdatesInProgress),
				"instances":                     parseTessellServiceInstanceDTOList(DBService.Instances),
				"databases":                     parseTessellDatabaseDTOList(DBService.Databases),
				"shared_with":                   []interface{}{parseEntityAclSharingInfo(DBService.SharedWith)},
				"deletion_schedule":             []interface{}{parseDeletionScheduleDTO(DBService.DeletionSchedule)},
				"upcoming_scheduled_actions":    []interface{}{parseServiceUpcomingScheduledActions(DBService.UpcomingScheduledActions)},
			}
		}
	}

	if err := d.Set("db_services", parsedDBServiceList); err != nil {
		return err
	}
	return nil
}
