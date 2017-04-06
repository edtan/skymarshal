// This file was generated by counterfeiter
package providerfakes

import (
	"encoding/json"
	"sync"

	"github.com/concourse/atc/auth/provider"
	"github.com/jessevdk/go-flags"
)

type FakeTeamProvider struct {
	ProviderConstructorStub        func(provider.AuthConfig, string) (provider.Provider, bool)
	providerConstructorMutex       sync.RWMutex
	providerConstructorArgsForCall []struct {
		arg1 provider.AuthConfig
		arg2 string
	}
	providerConstructorReturns struct {
		result1 provider.Provider
		result2 bool
	}
	providerConstructorReturnsOnCall map[int]struct {
		result1 provider.Provider
		result2 bool
	}
	AddAuthGroupStub        func(*flags.Group) provider.AuthConfig
	addAuthGroupMutex       sync.RWMutex
	addAuthGroupArgsForCall []struct {
		arg1 *flags.Group
	}
	addAuthGroupReturns struct {
		result1 provider.AuthConfig
	}
	addAuthGroupReturnsOnCall map[int]struct {
		result1 provider.AuthConfig
	}
	UnmarshalConfigStub        func(*json.RawMessage) (provider.AuthConfig, error)
	unmarshalConfigMutex       sync.RWMutex
	unmarshalConfigArgsForCall []struct {
		arg1 *json.RawMessage
	}
	unmarshalConfigReturns struct {
		result1 provider.AuthConfig
		result2 error
	}
	unmarshalConfigReturnsOnCall map[int]struct {
		result1 provider.AuthConfig
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamProvider) ProviderConstructor(arg1 provider.AuthConfig, arg2 string) (provider.Provider, bool) {
	fake.providerConstructorMutex.Lock()
	ret, specificReturn := fake.providerConstructorReturnsOnCall[len(fake.providerConstructorArgsForCall)]
	fake.providerConstructorArgsForCall = append(fake.providerConstructorArgsForCall, struct {
		arg1 provider.AuthConfig
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ProviderConstructor", []interface{}{arg1, arg2})
	fake.providerConstructorMutex.Unlock()
	if fake.ProviderConstructorStub != nil {
		return fake.ProviderConstructorStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.providerConstructorReturns.result1, fake.providerConstructorReturns.result2
}

func (fake *FakeTeamProvider) ProviderConstructorCallCount() int {
	fake.providerConstructorMutex.RLock()
	defer fake.providerConstructorMutex.RUnlock()
	return len(fake.providerConstructorArgsForCall)
}

func (fake *FakeTeamProvider) ProviderConstructorArgsForCall(i int) (provider.AuthConfig, string) {
	fake.providerConstructorMutex.RLock()
	defer fake.providerConstructorMutex.RUnlock()
	return fake.providerConstructorArgsForCall[i].arg1, fake.providerConstructorArgsForCall[i].arg2
}

func (fake *FakeTeamProvider) ProviderConstructorReturns(result1 provider.Provider, result2 bool) {
	fake.ProviderConstructorStub = nil
	fake.providerConstructorReturns = struct {
		result1 provider.Provider
		result2 bool
	}{result1, result2}
}

func (fake *FakeTeamProvider) ProviderConstructorReturnsOnCall(i int, result1 provider.Provider, result2 bool) {
	fake.ProviderConstructorStub = nil
	if fake.providerConstructorReturnsOnCall == nil {
		fake.providerConstructorReturnsOnCall = make(map[int]struct {
			result1 provider.Provider
			result2 bool
		})
	}
	fake.providerConstructorReturnsOnCall[i] = struct {
		result1 provider.Provider
		result2 bool
	}{result1, result2}
}

func (fake *FakeTeamProvider) AddAuthGroup(arg1 *flags.Group) provider.AuthConfig {
	fake.addAuthGroupMutex.Lock()
	ret, specificReturn := fake.addAuthGroupReturnsOnCall[len(fake.addAuthGroupArgsForCall)]
	fake.addAuthGroupArgsForCall = append(fake.addAuthGroupArgsForCall, struct {
		arg1 *flags.Group
	}{arg1})
	fake.recordInvocation("AddAuthGroup", []interface{}{arg1})
	fake.addAuthGroupMutex.Unlock()
	if fake.AddAuthGroupStub != nil {
		return fake.AddAuthGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addAuthGroupReturns.result1
}

func (fake *FakeTeamProvider) AddAuthGroupCallCount() int {
	fake.addAuthGroupMutex.RLock()
	defer fake.addAuthGroupMutex.RUnlock()
	return len(fake.addAuthGroupArgsForCall)
}

func (fake *FakeTeamProvider) AddAuthGroupArgsForCall(i int) *flags.Group {
	fake.addAuthGroupMutex.RLock()
	defer fake.addAuthGroupMutex.RUnlock()
	return fake.addAuthGroupArgsForCall[i].arg1
}

func (fake *FakeTeamProvider) AddAuthGroupReturns(result1 provider.AuthConfig) {
	fake.AddAuthGroupStub = nil
	fake.addAuthGroupReturns = struct {
		result1 provider.AuthConfig
	}{result1}
}

func (fake *FakeTeamProvider) AddAuthGroupReturnsOnCall(i int, result1 provider.AuthConfig) {
	fake.AddAuthGroupStub = nil
	if fake.addAuthGroupReturnsOnCall == nil {
		fake.addAuthGroupReturnsOnCall = make(map[int]struct {
			result1 provider.AuthConfig
		})
	}
	fake.addAuthGroupReturnsOnCall[i] = struct {
		result1 provider.AuthConfig
	}{result1}
}

func (fake *FakeTeamProvider) UnmarshalConfig(arg1 *json.RawMessage) (provider.AuthConfig, error) {
	fake.unmarshalConfigMutex.Lock()
	ret, specificReturn := fake.unmarshalConfigReturnsOnCall[len(fake.unmarshalConfigArgsForCall)]
	fake.unmarshalConfigArgsForCall = append(fake.unmarshalConfigArgsForCall, struct {
		arg1 *json.RawMessage
	}{arg1})
	fake.recordInvocation("UnmarshalConfig", []interface{}{arg1})
	fake.unmarshalConfigMutex.Unlock()
	if fake.UnmarshalConfigStub != nil {
		return fake.UnmarshalConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unmarshalConfigReturns.result1, fake.unmarshalConfigReturns.result2
}

func (fake *FakeTeamProvider) UnmarshalConfigCallCount() int {
	fake.unmarshalConfigMutex.RLock()
	defer fake.unmarshalConfigMutex.RUnlock()
	return len(fake.unmarshalConfigArgsForCall)
}

func (fake *FakeTeamProvider) UnmarshalConfigArgsForCall(i int) *json.RawMessage {
	fake.unmarshalConfigMutex.RLock()
	defer fake.unmarshalConfigMutex.RUnlock()
	return fake.unmarshalConfigArgsForCall[i].arg1
}

func (fake *FakeTeamProvider) UnmarshalConfigReturns(result1 provider.AuthConfig, result2 error) {
	fake.UnmarshalConfigStub = nil
	fake.unmarshalConfigReturns = struct {
		result1 provider.AuthConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamProvider) UnmarshalConfigReturnsOnCall(i int, result1 provider.AuthConfig, result2 error) {
	fake.UnmarshalConfigStub = nil
	if fake.unmarshalConfigReturnsOnCall == nil {
		fake.unmarshalConfigReturnsOnCall = make(map[int]struct {
			result1 provider.AuthConfig
			result2 error
		})
	}
	fake.unmarshalConfigReturnsOnCall[i] = struct {
		result1 provider.AuthConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.providerConstructorMutex.RLock()
	defer fake.providerConstructorMutex.RUnlock()
	fake.addAuthGroupMutex.RLock()
	defer fake.addAuthGroupMutex.RUnlock()
	fake.unmarshalConfigMutex.RLock()
	defer fake.unmarshalConfigMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTeamProvider) recordInvocation(key string, args []interface{}) {
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

var _ provider.TeamProvider = new(FakeTeamProvider)
