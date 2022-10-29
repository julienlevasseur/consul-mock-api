// Code generated by "mock-api-gen -type Consul -pkg mockconsul -endpoints agent.json -output ./agent.go"; DO NOT EDIT.

package mockconsul

import (
	"fmt"
	"io"

	"github.com/hashicorp/consul/api"
	mockapi "github.com/mkeeler/mock-http-api"
)

func (m *Consul) AgentCheckDeregister(checkID string, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/agent/check/deregister/%s", checkID))

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentCheckRegister(body *api.AgentCheckRegistration, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", "/v1/agent/check/register").WithBody(body)

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentCheckUpdateTLL(checkID string, body map[string]interface{}, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/agent/check/update/%s", checkID)).WithBody(body)

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentChecks(status int, reply map[string]*api.AgentCheck) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/checks")

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentConnectAuthorize(body *api.AgentAuthorizeParams, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("POST", "/v1/agent/connect/authorize").WithBody(body)

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentConnectCALeaf(serviceID string, status int, reply *api.LeafCert) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/agent/connect/ca/leaf/%s", serviceID))

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentConnectCARoots(status int, reply *api.CARootList) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/connect/ca/roots")

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentForceLeave(node string, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/agent/force-leave/%s", node))

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentHealthServiceByID(serviceID string, status int, reply *api.AgentServiceChecksInfo) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/agent/health/service/id/%s", serviceID))

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentHealthServiceByName(serviceName string, status int, reply []api.AgentServiceChecksInfo) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/agent/health/service/name/%s", serviceName))

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentHost(status int, reply map[string]interface{}) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/host")

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentJoin(address string, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/agent/join/%s", address))

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentLeave(status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", "/v1/agent/leave/%s")

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentMaintenance(status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", "/v1/agent/maintenance")

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentMembers(status int, reply []*api.AgentMember) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/members")

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentMetrics(status int, reply *api.MetricsInfo) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/metrics")

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentMonitor(status int, reply io.Reader) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/monitor")

	return m.WithStreamingReply(req, status, reply)
}

func (m *Consul) AgentReload(status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", "/v1/agent/reload")

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentSelf(status int, reply map[string]map[string]interface{}) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/self")

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentService(serviceID string, status int, reply *api.AgentService) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", fmt.Sprintf("/v1/agent/service/%s", serviceID))

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentServiceDeregister(status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", "/v1/agent/service/deregister")

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentServiceMaintenance(serviceID string, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/agent/service/maintenance/%s", serviceID))

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentServiceRegister(body *api.AgentServiceRegistration, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", "/v1/agent/service/register").WithBody(body)

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) AgentServices(status int, reply map[string]*api.AgentService) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("GET", "/v1/agent/services")

	return m.WithJSONReply(req, status, reply)
}

func (m *Consul) AgentTokenSet(tokenType string, body []byte, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", fmt.Sprintf("/v1/agent/token/%s", tokenType)).WithBody(body)

	return m.WithNoResponseBody(req, status)
}

func (m *Consul) SessionCreate(sessionEntry *api.SessionEntry, writeOptions *api.WriteOptions, status int) *mockapi.MockAPICall {
	req := mockapi.NewMockRequest("PUT", "/v1/session/create")

	return m.WithNoResponseBody(req, status)
}