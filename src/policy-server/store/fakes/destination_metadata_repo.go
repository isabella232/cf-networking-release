// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/cf-networking-helpers/db"
)

type DestinationMetadataRepo struct {
	DeleteStub        func(tx db.Transaction, terminalGUID string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		tx           db.Transaction
		terminalGUID string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	UpsertStub        func(tx db.Transaction, terminalGUID, name, description string) error
	upsertMutex       sync.RWMutex
	upsertArgsForCall []struct {
		tx           db.Transaction
		terminalGUID string
		name         string
		description  string
	}
	upsertReturns struct {
		result1 error
	}
	upsertReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DestinationMetadataRepo) Delete(tx db.Transaction, terminalGUID string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		tx           db.Transaction
		terminalGUID string
	}{tx, terminalGUID})
	fake.recordInvocation("Delete", []interface{}{tx, terminalGUID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(tx, terminalGUID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *DestinationMetadataRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *DestinationMetadataRepo) DeleteArgsForCall(i int) (db.Transaction, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].tx, fake.deleteArgsForCall[i].terminalGUID
}

func (fake *DestinationMetadataRepo) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *DestinationMetadataRepo) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *DestinationMetadataRepo) Upsert(tx db.Transaction, terminalGUID string, name string, description string) error {
	fake.upsertMutex.Lock()
	ret, specificReturn := fake.upsertReturnsOnCall[len(fake.upsertArgsForCall)]
	fake.upsertArgsForCall = append(fake.upsertArgsForCall, struct {
		tx           db.Transaction
		terminalGUID string
		name         string
		description  string
	}{tx, terminalGUID, name, description})
	fake.recordInvocation("Upsert", []interface{}{tx, terminalGUID, name, description})
	fake.upsertMutex.Unlock()
	if fake.UpsertStub != nil {
		return fake.UpsertStub(tx, terminalGUID, name, description)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.upsertReturns.result1
}

func (fake *DestinationMetadataRepo) UpsertCallCount() int {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return len(fake.upsertArgsForCall)
}

func (fake *DestinationMetadataRepo) UpsertArgsForCall(i int) (db.Transaction, string, string, string) {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return fake.upsertArgsForCall[i].tx, fake.upsertArgsForCall[i].terminalGUID, fake.upsertArgsForCall[i].name, fake.upsertArgsForCall[i].description
}

func (fake *DestinationMetadataRepo) UpsertReturns(result1 error) {
	fake.UpsertStub = nil
	fake.upsertReturns = struct {
		result1 error
	}{result1}
}

func (fake *DestinationMetadataRepo) UpsertReturnsOnCall(i int, result1 error) {
	fake.UpsertStub = nil
	if fake.upsertReturnsOnCall == nil {
		fake.upsertReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.upsertReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DestinationMetadataRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DestinationMetadataRepo) recordInvocation(key string, args []interface{}) {
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
