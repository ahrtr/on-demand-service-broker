// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	credhubcredhub_cli "github.com/cloudfoundry-incubator/credhub-cli/credhub"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/values"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/permissions"
	"github.com/pivotal-cf/on-demand-service-broker/credhub"
)

type FakeCredhubClient struct {
	GetByIdStub        func(id string) (credentials.Credential, error)
	getByIdMutex       sync.RWMutex
	getByIdArgsForCall []struct {
		id string
	}
	getByIdReturns struct {
		result1 credentials.Credential
		result2 error
	}
	getByIdReturnsOnCall map[int]struct {
		result1 credentials.Credential
		result2 error
	}
	GetLatestVersionStub        func(name string) (credentials.Credential, error)
	getLatestVersionMutex       sync.RWMutex
	getLatestVersionArgsForCall []struct {
		name string
	}
	getLatestVersionReturns struct {
		result1 credentials.Credential
		result2 error
	}
	getLatestVersionReturnsOnCall map[int]struct {
		result1 credentials.Credential
		result2 error
	}
	SetJSONStub        func(name string, value values.JSON, overwrite credhubcredhub_cli.Mode) (credentials.JSON, error)
	setJSONMutex       sync.RWMutex
	setJSONArgsForCall []struct {
		name      string
		value     values.JSON
		overwrite credhubcredhub_cli.Mode
	}
	setJSONReturns struct {
		result1 credentials.JSON
		result2 error
	}
	setJSONReturnsOnCall map[int]struct {
		result1 credentials.JSON
		result2 error
	}
	SetValueStub        func(name string, value values.Value, overwrite credhubcredhub_cli.Mode) (credentials.Value, error)
	setValueMutex       sync.RWMutex
	setValueArgsForCall []struct {
		name      string
		value     values.Value
		overwrite credhubcredhub_cli.Mode
	}
	setValueReturns struct {
		result1 credentials.Value
		result2 error
	}
	setValueReturnsOnCall map[int]struct {
		result1 credentials.Value
		result2 error
	}
	AddPermissionsStub        func(credName string, perms []permissions.Permission) ([]permissions.Permission, error)
	addPermissionsMutex       sync.RWMutex
	addPermissionsArgsForCall []struct {
		credName string
		perms    []permissions.Permission
	}
	addPermissionsReturns struct {
		result1 []permissions.Permission
		result2 error
	}
	addPermissionsReturnsOnCall map[int]struct {
		result1 []permissions.Permission
		result2 error
	}
	DeleteStub        func(name string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		name string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredhubClient) GetById(id string) (credentials.Credential, error) {
	fake.getByIdMutex.Lock()
	ret, specificReturn := fake.getByIdReturnsOnCall[len(fake.getByIdArgsForCall)]
	fake.getByIdArgsForCall = append(fake.getByIdArgsForCall, struct {
		id string
	}{id})
	fake.recordInvocation("GetById", []interface{}{id})
	fake.getByIdMutex.Unlock()
	if fake.GetByIdStub != nil {
		return fake.GetByIdStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByIdReturns.result1, fake.getByIdReturns.result2
}

func (fake *FakeCredhubClient) GetByIdCallCount() int {
	fake.getByIdMutex.RLock()
	defer fake.getByIdMutex.RUnlock()
	return len(fake.getByIdArgsForCall)
}

func (fake *FakeCredhubClient) GetByIdArgsForCall(i int) string {
	fake.getByIdMutex.RLock()
	defer fake.getByIdMutex.RUnlock()
	return fake.getByIdArgsForCall[i].id
}

func (fake *FakeCredhubClient) GetByIdReturns(result1 credentials.Credential, result2 error) {
	fake.GetByIdStub = nil
	fake.getByIdReturns = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) GetByIdReturnsOnCall(i int, result1 credentials.Credential, result2 error) {
	fake.GetByIdStub = nil
	if fake.getByIdReturnsOnCall == nil {
		fake.getByIdReturnsOnCall = make(map[int]struct {
			result1 credentials.Credential
			result2 error
		})
	}
	fake.getByIdReturnsOnCall[i] = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) GetLatestVersion(name string) (credentials.Credential, error) {
	fake.getLatestVersionMutex.Lock()
	ret, specificReturn := fake.getLatestVersionReturnsOnCall[len(fake.getLatestVersionArgsForCall)]
	fake.getLatestVersionArgsForCall = append(fake.getLatestVersionArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetLatestVersion", []interface{}{name})
	fake.getLatestVersionMutex.Unlock()
	if fake.GetLatestVersionStub != nil {
		return fake.GetLatestVersionStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getLatestVersionReturns.result1, fake.getLatestVersionReturns.result2
}

func (fake *FakeCredhubClient) GetLatestVersionCallCount() int {
	fake.getLatestVersionMutex.RLock()
	defer fake.getLatestVersionMutex.RUnlock()
	return len(fake.getLatestVersionArgsForCall)
}

func (fake *FakeCredhubClient) GetLatestVersionArgsForCall(i int) string {
	fake.getLatestVersionMutex.RLock()
	defer fake.getLatestVersionMutex.RUnlock()
	return fake.getLatestVersionArgsForCall[i].name
}

func (fake *FakeCredhubClient) GetLatestVersionReturns(result1 credentials.Credential, result2 error) {
	fake.GetLatestVersionStub = nil
	fake.getLatestVersionReturns = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) GetLatestVersionReturnsOnCall(i int, result1 credentials.Credential, result2 error) {
	fake.GetLatestVersionStub = nil
	if fake.getLatestVersionReturnsOnCall == nil {
		fake.getLatestVersionReturnsOnCall = make(map[int]struct {
			result1 credentials.Credential
			result2 error
		})
	}
	fake.getLatestVersionReturnsOnCall[i] = struct {
		result1 credentials.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetJSON(name string, value values.JSON, overwrite credhubcredhub_cli.Mode) (credentials.JSON, error) {
	fake.setJSONMutex.Lock()
	ret, specificReturn := fake.setJSONReturnsOnCall[len(fake.setJSONArgsForCall)]
	fake.setJSONArgsForCall = append(fake.setJSONArgsForCall, struct {
		name      string
		value     values.JSON
		overwrite credhubcredhub_cli.Mode
	}{name, value, overwrite})
	fake.recordInvocation("SetJSON", []interface{}{name, value, overwrite})
	fake.setJSONMutex.Unlock()
	if fake.SetJSONStub != nil {
		return fake.SetJSONStub(name, value, overwrite)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setJSONReturns.result1, fake.setJSONReturns.result2
}

func (fake *FakeCredhubClient) SetJSONCallCount() int {
	fake.setJSONMutex.RLock()
	defer fake.setJSONMutex.RUnlock()
	return len(fake.setJSONArgsForCall)
}

func (fake *FakeCredhubClient) SetJSONArgsForCall(i int) (string, values.JSON, credhubcredhub_cli.Mode) {
	fake.setJSONMutex.RLock()
	defer fake.setJSONMutex.RUnlock()
	return fake.setJSONArgsForCall[i].name, fake.setJSONArgsForCall[i].value, fake.setJSONArgsForCall[i].overwrite
}

func (fake *FakeCredhubClient) SetJSONReturns(result1 credentials.JSON, result2 error) {
	fake.SetJSONStub = nil
	fake.setJSONReturns = struct {
		result1 credentials.JSON
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetJSONReturnsOnCall(i int, result1 credentials.JSON, result2 error) {
	fake.SetJSONStub = nil
	if fake.setJSONReturnsOnCall == nil {
		fake.setJSONReturnsOnCall = make(map[int]struct {
			result1 credentials.JSON
			result2 error
		})
	}
	fake.setJSONReturnsOnCall[i] = struct {
		result1 credentials.JSON
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetValue(name string, value values.Value, overwrite credhubcredhub_cli.Mode) (credentials.Value, error) {
	fake.setValueMutex.Lock()
	ret, specificReturn := fake.setValueReturnsOnCall[len(fake.setValueArgsForCall)]
	fake.setValueArgsForCall = append(fake.setValueArgsForCall, struct {
		name      string
		value     values.Value
		overwrite credhubcredhub_cli.Mode
	}{name, value, overwrite})
	fake.recordInvocation("SetValue", []interface{}{name, value, overwrite})
	fake.setValueMutex.Unlock()
	if fake.SetValueStub != nil {
		return fake.SetValueStub(name, value, overwrite)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setValueReturns.result1, fake.setValueReturns.result2
}

func (fake *FakeCredhubClient) SetValueCallCount() int {
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	return len(fake.setValueArgsForCall)
}

func (fake *FakeCredhubClient) SetValueArgsForCall(i int) (string, values.Value, credhubcredhub_cli.Mode) {
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	return fake.setValueArgsForCall[i].name, fake.setValueArgsForCall[i].value, fake.setValueArgsForCall[i].overwrite
}

func (fake *FakeCredhubClient) SetValueReturns(result1 credentials.Value, result2 error) {
	fake.SetValueStub = nil
	fake.setValueReturns = struct {
		result1 credentials.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) SetValueReturnsOnCall(i int, result1 credentials.Value, result2 error) {
	fake.SetValueStub = nil
	if fake.setValueReturnsOnCall == nil {
		fake.setValueReturnsOnCall = make(map[int]struct {
			result1 credentials.Value
			result2 error
		})
	}
	fake.setValueReturnsOnCall[i] = struct {
		result1 credentials.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) AddPermissions(credName string, perms []permissions.Permission) ([]permissions.Permission, error) {
	var permsCopy []permissions.Permission
	if perms != nil {
		permsCopy = make([]permissions.Permission, len(perms))
		copy(permsCopy, perms)
	}
	fake.addPermissionsMutex.Lock()
	ret, specificReturn := fake.addPermissionsReturnsOnCall[len(fake.addPermissionsArgsForCall)]
	fake.addPermissionsArgsForCall = append(fake.addPermissionsArgsForCall, struct {
		credName string
		perms    []permissions.Permission
	}{credName, permsCopy})
	fake.recordInvocation("AddPermissions", []interface{}{credName, permsCopy})
	fake.addPermissionsMutex.Unlock()
	if fake.AddPermissionsStub != nil {
		return fake.AddPermissionsStub(credName, perms)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.addPermissionsReturns.result1, fake.addPermissionsReturns.result2
}

func (fake *FakeCredhubClient) AddPermissionsCallCount() int {
	fake.addPermissionsMutex.RLock()
	defer fake.addPermissionsMutex.RUnlock()
	return len(fake.addPermissionsArgsForCall)
}

func (fake *FakeCredhubClient) AddPermissionsArgsForCall(i int) (string, []permissions.Permission) {
	fake.addPermissionsMutex.RLock()
	defer fake.addPermissionsMutex.RUnlock()
	return fake.addPermissionsArgsForCall[i].credName, fake.addPermissionsArgsForCall[i].perms
}

func (fake *FakeCredhubClient) AddPermissionsReturns(result1 []permissions.Permission, result2 error) {
	fake.AddPermissionsStub = nil
	fake.addPermissionsReturns = struct {
		result1 []permissions.Permission
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) AddPermissionsReturnsOnCall(i int, result1 []permissions.Permission, result2 error) {
	fake.AddPermissionsStub = nil
	if fake.addPermissionsReturnsOnCall == nil {
		fake.addPermissionsReturnsOnCall = make(map[int]struct {
			result1 []permissions.Permission
			result2 error
		})
	}
	fake.addPermissionsReturnsOnCall[i] = struct {
		result1 []permissions.Permission
		result2 error
	}{result1, result2}
}

func (fake *FakeCredhubClient) Delete(name string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Delete", []interface{}{name})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeCredhubClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCredhubClient) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].name
}

func (fake *FakeCredhubClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredhubClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredhubClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getByIdMutex.RLock()
	defer fake.getByIdMutex.RUnlock()
	fake.getLatestVersionMutex.RLock()
	defer fake.getLatestVersionMutex.RUnlock()
	fake.setJSONMutex.RLock()
	defer fake.setJSONMutex.RUnlock()
	fake.setValueMutex.RLock()
	defer fake.setValueMutex.RUnlock()
	fake.addPermissionsMutex.RLock()
	defer fake.addPermissionsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredhubClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ credhub.CredhubClient = new(FakeCredhubClient)
