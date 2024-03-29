package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"terraform-provider-tessell/internal/model"
)

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

func (c *Client) GetTessellService(id string) (*model.TessellServiceDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/%s", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}

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

func (c *Client) ProvisionTessellService(payload model.ProvisionTessellServicePayload) (*model.TaskSummary, int, error) {
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
		_, statusCode, err := c.GetTessellService(id)
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
		response, _, err := c.GetTessellService(id)
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
		response, _, err := c.GetTessellService(id)
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
