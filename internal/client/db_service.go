package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (c *Client) AddTessellServiceInstances(id string, payload *model.AddDBServiceInstancesPayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(*payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/%s/service-instances", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) CloneTessellService(parentAvailabilityMachineId string, payload model.CloneTessellServicePayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/availability-machines/%s/clones", c.APIAddress, parentAvailabilityMachineId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) DeleteTessellService(id string, payload model.DeleteTessellServicePayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/services/%s", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) DeleteTessellServiceInstances(id string, payload *model.DeleteTessellServiceInstancePayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(*payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/services/%s/service-instances", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) GetTessellService(id string, d *schema.ResourceData) (*model.TessellServiceDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/%s", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}

	instanceNames := getInstanceNames(d)
	databaseNames := getDatabaseNames(d)

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellServiceDTO := model.TessellServiceDTO{}
	err = json.Unmarshal(body, &tessellServiceDTO)
	if err != nil {
		return nil, statusCode, err
	}

	// ordering support for arrays
	tessellServiceDTO.Instances = orderInstancesByName(tessellServiceDTO.Instances, instanceNames)
	tessellServiceDTO.Databases = orderDatabasesByName(tessellServiceDTO.Databases, databaseNames)

	return &tessellServiceDTO, statusCode, nil
}

func getInstanceNames(d *schema.ResourceData) *[]string {
	if d == nil {
		return nil
	}

	tfInstances := d.Get("instances")
	if tfInstances == nil {
		return nil
	}

	instanceMaps, _ := tfInstances.([]interface{})
	instanceNames := []string{}
	for _, instanceMap := range instanceMaps {
		inputInstance := instanceMap.(map[string]interface{})
		instanceNames = append(instanceNames, *helper.GetStringPointer(inputInstance["name"]))
	}
	return &instanceNames
}

func getDatabaseNames(d *schema.ResourceData) *[]string {
	if d == nil {
		return nil
	}

	tfDatabases := d.Get("databases")
	if tfDatabases == nil {
		return nil
	}

	databaseMaps, _ := tfDatabases.([]interface{})
	databaseNames := []string{}
	for _, databaseMap := range databaseMaps {
		inputDatabase := databaseMap.(map[string]interface{})
		databaseNames = append(databaseNames, *helper.GetStringPointer(inputDatabase["database_name"]))
	}
	return &databaseNames
}

func orderInstancesByName(instances *[]model.TessellServiceInstanceDTO, instanceNames *[]string) *[]model.TessellServiceInstanceDTO {
	if instanceNames == nil {
		return instances
	}
	instanceMap := make(map[string]model.TessellServiceInstanceDTO)
	for _, instance := range *instances {
		instanceMap[*instance.Name] = instance
	}

	orderedInstances := make([]model.TessellServiceInstanceDTO, 0, len(*instances))

	for _, name := range *instanceNames {
		if instance, exists := instanceMap[name]; exists {
			orderedInstances = append(orderedInstances, instance)
			delete(instanceMap, name)
		}
	}

	for _, instance := range *instances {
		if _, exists := instanceMap[*instance.Name]; exists {
			orderedInstances = append(orderedInstances, instance)
		}
	}

	return &orderedInstances
}

func orderDatabasesByName(databases *[]model.TessellDatabaseDTO, databaseNames *[]string) *[]model.TessellDatabaseDTO {
	if databaseNames == nil {
		return databases
	}
	databaseMap := make(map[string]model.TessellDatabaseDTO)
	for _, database := range *databases {
		databaseMap[*database.DatabaseName] = database
	}

	orderedDatabases := make([]model.TessellDatabaseDTO, 0, len(*databases))

	for _, name := range *databaseNames {
		if database, exists := databaseMap[name]; exists {
			orderedDatabases = append(orderedDatabases, database)
			delete(databaseMap, name)
		}
	}

	for _, database := range *databases {
		if _, exists := databaseMap[*database.DatabaseName]; exists {
			orderedDatabases = append(orderedDatabases, database)
		}
	}

	return &orderedDatabases
}

func (c *Client) GetTessellServices(name string, statuses []string, engineTypes []string, clonedFromServiceId string, clonedFromAvailabilityMachineId string, loadInstances bool, loadDatabases bool, owners []string, loadAcls bool) (*model.TessellServicesResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services", c.APIAddress), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("name", fmt.Sprintf("%v", name))
	q.Add("statuses", strings.Join(statuses, ","))
	q.Add("engine-types", strings.Join(engineTypes, ","))
	q.Add("cloned-from-service-id", fmt.Sprintf("%v", clonedFromServiceId))
	q.Add("cloned-from-availability-machine-id", fmt.Sprintf("%v", clonedFromAvailabilityMachineId))
	q.Add("load-instances", fmt.Sprintf("%v", loadInstances))
	q.Add("load-databases", fmt.Sprintf("%v", loadDatabases))
	q.Add("owners", strings.Join(owners, ","))
	q.Add("load-acls", fmt.Sprintf("%v", loadAcls))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellServicesResponse := model.TessellServicesResponse{}
	err = json.Unmarshal(body, &tessellServicesResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellServicesResponse, statusCode, nil
}

func (c *Client) ProvisionTessellService(payload model.ProvisionServicePayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services", c.APIAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) StartTessellService(id string, payload model.StartTessellServicePayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/%s/start", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) StopTessellService(id string, payload model.StopTessellServicePayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/%s/stop", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) SwitchoverTessellService(id string, payload *model.SwitchOverTessellServicePayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(*payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/%s/switchover", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) UpdateTessellService(id string, payload model.UpdateTessellServicePayload) (*model.TessellServiceDTO, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/%s", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellServiceDTO := model.TessellServiceDTO{}
	err = json.Unmarshal(body, &tessellServiceDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellServiceDTO, statusCode, nil
}

func (c *Client) UpdateTessellServiceCredentials(id string, payload model.ResetTessellServiceCredsPayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/%s/creds", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) DBServicePollForStatusCode(id string, statusCodeRequired int, timeout int, interval int) error {

	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	for {
		_, statusCode, err := c.GetTessellService(id, nil)
		if err != nil {
			if statusCode == statusCodeRequired {
				return nil
			}
			return fmt.Errorf("error while polling: %s", err.Error())
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out while polling")
		}
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}

func (c *Client) DBServicePollForStatus(id string, value string, timeout int, interval int) error {

	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	for {
		response, _, err := c.GetTessellService(id, nil)
		if err != nil {
			return fmt.Errorf("error while polling: %s", err.Error())
		}

		switch *response.Status {
		case value:
			return nil
		case "FAILED":
			return fmt.Errorf("received status FAILED while polling")
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out while polling")
		}
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}

func (c *Client) DBServicePollForUpdateInProgress(id string, referenceId string, timeout int, interval int) error {

	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	errorCountWhilePolling := 0

	for {
		response, _, err := c.GetTessellService(id, nil)
		if err != nil {
			errorCountWhilePolling += 1
			if errorCountWhilePolling > 3 {
				return fmt.Errorf("error while polling: %s", err.Error())
			} else {
				continue
			}
		}

		updateStillInProgress := false

		for _, resourceUpdateInfo := range *response.UpdatesInProgress {
			if *resourceUpdateInfo.ReferenceId == referenceId {
				updateStillInProgress = true
			}
		}

		if !updateStillInProgress {
			return nil
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out while polling")
		}
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}

func (c *Client) DBServicePollForInstanceAddition(id string, instanceName string, timeout int, interval int) error {
	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	errorCountWhilePolling := 0

	for {
		response, _, err := c.GetTessellService(id, nil)
		if err != nil {
			errorCountWhilePolling += 1
			if errorCountWhilePolling > 3 {
				return fmt.Errorf("error while polling: %s", err.Error())
			} else {
				continue
			}
		}

		updateStillInProgress := false

		for _, instance := range *response.Instances {
			if *instance.Name == instanceName {
				if *instance.Status == "CREATING" {
					updateStillInProgress = true
				} else if *instance.Status == "UP" {
					return nil
				} else {
					return fmt.Errorf("instance creation failed")
				}
				break
			}
		}

		if !updateStillInProgress {
			return nil
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out while polling")
		}
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}

func (c *Client) DBServicePollForInstanceSwitchover(id string, instanceId string, timeout int, interval int) error {
	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	errorCountWhilePolling := 0

	for {
		response, _, err := c.GetTessellService(id, nil)
		if err != nil {
			errorCountWhilePolling += 1
			if errorCountWhilePolling > 3 {
				return fmt.Errorf("error while polling: %s", err.Error())
			} else {
				continue
			}
		}

		switchoverStillInProgress := false

		if *response.Status == "SWITCHOVER" {
			switchoverStillInProgress = true
		} else if *response.Status == "DOWN" {
			return fmt.Errorf("instance switchover failed")
		} else {
			// check instance status
			for _, instance := range *response.Instances {
				if *instance.Id == instanceId {
					if *instance.Role == "primary" {
						return nil
					} else {
						switchoverStillInProgress = true
						break
					}
				}
			}
		}

		if !switchoverStillInProgress {
			return nil
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out while polling")
		}
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}

func (c *Client) GetTessellServiceInstance(id string, instanceId string) (*model.TessellServiceInstanceDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/%s/service-instances/%s", c.APIAddress, id, instanceId), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellServiceInstanceDTO := model.TessellServiceInstanceDTO{}
	err = json.Unmarshal(body, &tessellServiceInstanceDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellServiceInstanceDTO, statusCode, nil
}

func (c *Client) DBServicePollForInstanceDeletion(id string, instanceId string, timeout int, interval int) error {
	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	errorCountWhilePolling := 0

	for {
		response, statusCode, err := c.GetTessellServiceInstance(id, instanceId)
		if statusCode == 404 {
			return nil
		}
		if err != nil {
			errorCountWhilePolling += 1
			if errorCountWhilePolling > 3 {
				return fmt.Errorf("error while polling: %s", err.Error())
			} else {
				continue
			}
		}

		if *response.Status != "DELETING" {
			return fmt.Errorf("instance deletion failed")
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out while polling")
		}
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}
