// This file was generated by counterfeiter
package fakes

import (
	"confab"
	"sync"
)

type FakeconsulRPCClient struct {
	StatsStub        func() (map[string]map[string]string, error)
	statsMutex       sync.RWMutex
	statsArgsForCall []struct{}
	statsReturns     struct {
		result1 map[string]map[string]string
		result2 error
	}
	ListKeysStub        func(token string) (confab.KeyringResponse, error)
	listKeysMutex       sync.RWMutex
	listKeysArgsForCall []struct {
		token string
	}
	listKeysReturns struct {
		result1 confab.KeyringResponse
		result2 error
	}
	InstallKeyStub        func(key, token string) (confab.KeyringResponse, error)
	installKeyMutex       sync.RWMutex
	installKeyArgsForCall []struct {
		key   string
		token string
	}
	installKeyReturns struct {
		result1 confab.KeyringResponse
		result2 error
	}
	UseKeyStub        func(key, token string) (confab.KeyringResponse, error)
	useKeyMutex       sync.RWMutex
	useKeyArgsForCall []struct {
		key   string
		token string
	}
	useKeyReturns struct {
		result1 confab.KeyringResponse
		result2 error
	}
	RemoveKeyStub        func(key, token string) (confab.KeyringResponse, error)
	removeKeyMutex       sync.RWMutex
	removeKeyArgsForCall []struct {
		key   string
		token string
	}
	removeKeyReturns struct {
		result1 confab.KeyringResponse
		result2 error
	}
	LeaveStub        func() error
	leaveMutex       sync.RWMutex
	leaveArgsForCall []struct{}
	leaveReturns     struct {
		result1 error
	}
}

func (fake *FakeconsulRPCClient) Stats() (map[string]map[string]string, error) {
	fake.statsMutex.Lock()
	fake.statsArgsForCall = append(fake.statsArgsForCall, struct{}{})
	fake.statsMutex.Unlock()
	if fake.StatsStub != nil {
		return fake.StatsStub()
	} else {
		return fake.statsReturns.result1, fake.statsReturns.result2
	}
}

func (fake *FakeconsulRPCClient) StatsCallCount() int {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return len(fake.statsArgsForCall)
}

func (fake *FakeconsulRPCClient) StatsReturns(result1 map[string]map[string]string, result2 error) {
	fake.StatsStub = nil
	fake.statsReturns = struct {
		result1 map[string]map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeconsulRPCClient) ListKeys(token string) (confab.KeyringResponse, error) {
	fake.listKeysMutex.Lock()
	fake.listKeysArgsForCall = append(fake.listKeysArgsForCall, struct {
		token string
	}{token})
	fake.listKeysMutex.Unlock()
	if fake.ListKeysStub != nil {
		return fake.ListKeysStub(token)
	} else {
		return fake.listKeysReturns.result1, fake.listKeysReturns.result2
	}
}

func (fake *FakeconsulRPCClient) ListKeysCallCount() int {
	fake.listKeysMutex.RLock()
	defer fake.listKeysMutex.RUnlock()
	return len(fake.listKeysArgsForCall)
}

func (fake *FakeconsulRPCClient) ListKeysArgsForCall(i int) string {
	fake.listKeysMutex.RLock()
	defer fake.listKeysMutex.RUnlock()
	return fake.listKeysArgsForCall[i].token
}

func (fake *FakeconsulRPCClient) ListKeysReturns(result1 confab.KeyringResponse, result2 error) {
	fake.ListKeysStub = nil
	fake.listKeysReturns = struct {
		result1 confab.KeyringResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeconsulRPCClient) InstallKey(key string, token string) (confab.KeyringResponse, error) {
	fake.installKeyMutex.Lock()
	fake.installKeyArgsForCall = append(fake.installKeyArgsForCall, struct {
		key   string
		token string
	}{key, token})
	fake.installKeyMutex.Unlock()
	if fake.InstallKeyStub != nil {
		return fake.InstallKeyStub(key, token)
	} else {
		return fake.installKeyReturns.result1, fake.installKeyReturns.result2
	}
}

func (fake *FakeconsulRPCClient) InstallKeyCallCount() int {
	fake.installKeyMutex.RLock()
	defer fake.installKeyMutex.RUnlock()
	return len(fake.installKeyArgsForCall)
}

func (fake *FakeconsulRPCClient) InstallKeyArgsForCall(i int) (string, string) {
	fake.installKeyMutex.RLock()
	defer fake.installKeyMutex.RUnlock()
	return fake.installKeyArgsForCall[i].key, fake.installKeyArgsForCall[i].token
}

func (fake *FakeconsulRPCClient) InstallKeyReturns(result1 confab.KeyringResponse, result2 error) {
	fake.InstallKeyStub = nil
	fake.installKeyReturns = struct {
		result1 confab.KeyringResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeconsulRPCClient) UseKey(key string, token string) (confab.KeyringResponse, error) {
	fake.useKeyMutex.Lock()
	fake.useKeyArgsForCall = append(fake.useKeyArgsForCall, struct {
		key   string
		token string
	}{key, token})
	fake.useKeyMutex.Unlock()
	if fake.UseKeyStub != nil {
		return fake.UseKeyStub(key, token)
	} else {
		return fake.useKeyReturns.result1, fake.useKeyReturns.result2
	}
}

func (fake *FakeconsulRPCClient) UseKeyCallCount() int {
	fake.useKeyMutex.RLock()
	defer fake.useKeyMutex.RUnlock()
	return len(fake.useKeyArgsForCall)
}

func (fake *FakeconsulRPCClient) UseKeyArgsForCall(i int) (string, string) {
	fake.useKeyMutex.RLock()
	defer fake.useKeyMutex.RUnlock()
	return fake.useKeyArgsForCall[i].key, fake.useKeyArgsForCall[i].token
}

func (fake *FakeconsulRPCClient) UseKeyReturns(result1 confab.KeyringResponse, result2 error) {
	fake.UseKeyStub = nil
	fake.useKeyReturns = struct {
		result1 confab.KeyringResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeconsulRPCClient) RemoveKey(key string, token string) (confab.KeyringResponse, error) {
	fake.removeKeyMutex.Lock()
	fake.removeKeyArgsForCall = append(fake.removeKeyArgsForCall, struct {
		key   string
		token string
	}{key, token})
	fake.removeKeyMutex.Unlock()
	if fake.RemoveKeyStub != nil {
		return fake.RemoveKeyStub(key, token)
	} else {
		return fake.removeKeyReturns.result1, fake.removeKeyReturns.result2
	}
}

func (fake *FakeconsulRPCClient) RemoveKeyCallCount() int {
	fake.removeKeyMutex.RLock()
	defer fake.removeKeyMutex.RUnlock()
	return len(fake.removeKeyArgsForCall)
}

func (fake *FakeconsulRPCClient) RemoveKeyArgsForCall(i int) (string, string) {
	fake.removeKeyMutex.RLock()
	defer fake.removeKeyMutex.RUnlock()
	return fake.removeKeyArgsForCall[i].key, fake.removeKeyArgsForCall[i].token
}

func (fake *FakeconsulRPCClient) RemoveKeyReturns(result1 confab.KeyringResponse, result2 error) {
	fake.RemoveKeyStub = nil
	fake.removeKeyReturns = struct {
		result1 confab.KeyringResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeconsulRPCClient) Leave() error {
	fake.leaveMutex.Lock()
	fake.leaveArgsForCall = append(fake.leaveArgsForCall, struct{}{})
	fake.leaveMutex.Unlock()
	if fake.LeaveStub != nil {
		return fake.LeaveStub()
	} else {
		return fake.leaveReturns.result1
	}
}

func (fake *FakeconsulRPCClient) LeaveCallCount() int {
	fake.leaveMutex.RLock()
	defer fake.leaveMutex.RUnlock()
	return len(fake.leaveArgsForCall)
}

func (fake *FakeconsulRPCClient) LeaveReturns(result1 error) {
	fake.LeaveStub = nil
	fake.leaveReturns = struct {
		result1 error
	}{result1}
}

// var _ confab.consulRPCClient = new(FakeconsulRPCClient)
