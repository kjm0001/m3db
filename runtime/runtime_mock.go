// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db/runtime/types.go

package runtime

import (
	"reflect"
	"time"

	"github.com/m3db/m3db/ratelimit"
	"github.com/m3db/m3x/close"

	"github.com/golang/mock/gomock"
)

// MockOptions is a mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsMockRecorder
}

// MockOptionsMockRecorder is the mock recorder for MockOptions
type MockOptionsMockRecorder struct {
	mock *MockOptions
}

// NewMockOptions creates a new mock instance
func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &MockOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockOptions) EXPECT() *MockOptionsMockRecorder {
	return _m.recorder
}

// Validate mocks base method
func (_m *MockOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockOptionsMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockOptions)(nil).Validate))
}

// SetPersistRateLimitOptions mocks base method
func (_m *MockOptions) SetPersistRateLimitOptions(value ratelimit.Options) Options {
	ret := _m.ctrl.Call(_m, "SetPersistRateLimitOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetPersistRateLimitOptions indicates an expected call of SetPersistRateLimitOptions
func (_mr *MockOptionsMockRecorder) SetPersistRateLimitOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetPersistRateLimitOptions", reflect.TypeOf((*MockOptions)(nil).SetPersistRateLimitOptions), arg0)
}

// PersistRateLimitOptions mocks base method
func (_m *MockOptions) PersistRateLimitOptions() ratelimit.Options {
	ret := _m.ctrl.Call(_m, "PersistRateLimitOptions")
	ret0, _ := ret[0].(ratelimit.Options)
	return ret0
}

// PersistRateLimitOptions indicates an expected call of PersistRateLimitOptions
func (_mr *MockOptionsMockRecorder) PersistRateLimitOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PersistRateLimitOptions", reflect.TypeOf((*MockOptions)(nil).PersistRateLimitOptions))
}

// SetWriteNewSeriesAsync mocks base method
func (_m *MockOptions) SetWriteNewSeriesAsync(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetWriteNewSeriesAsync", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWriteNewSeriesAsync indicates an expected call of SetWriteNewSeriesAsync
func (_mr *MockOptionsMockRecorder) SetWriteNewSeriesAsync(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetWriteNewSeriesAsync", reflect.TypeOf((*MockOptions)(nil).SetWriteNewSeriesAsync), arg0)
}

// WriteNewSeriesAsync mocks base method
func (_m *MockOptions) WriteNewSeriesAsync() bool {
	ret := _m.ctrl.Call(_m, "WriteNewSeriesAsync")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WriteNewSeriesAsync indicates an expected call of WriteNewSeriesAsync
func (_mr *MockOptionsMockRecorder) WriteNewSeriesAsync() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteNewSeriesAsync", reflect.TypeOf((*MockOptions)(nil).WriteNewSeriesAsync))
}

// SetWriteNewSeriesBackoffDuration mocks base method
func (_m *MockOptions) SetWriteNewSeriesBackoffDuration(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetWriteNewSeriesBackoffDuration", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWriteNewSeriesBackoffDuration indicates an expected call of SetWriteNewSeriesBackoffDuration
func (_mr *MockOptionsMockRecorder) SetWriteNewSeriesBackoffDuration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetWriteNewSeriesBackoffDuration", reflect.TypeOf((*MockOptions)(nil).SetWriteNewSeriesBackoffDuration), arg0)
}

// WriteNewSeriesBackoffDuration mocks base method
func (_m *MockOptions) WriteNewSeriesBackoffDuration() time.Duration {
	ret := _m.ctrl.Call(_m, "WriteNewSeriesBackoffDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// WriteNewSeriesBackoffDuration indicates an expected call of WriteNewSeriesBackoffDuration
func (_mr *MockOptionsMockRecorder) WriteNewSeriesBackoffDuration() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteNewSeriesBackoffDuration", reflect.TypeOf((*MockOptions)(nil).WriteNewSeriesBackoffDuration))
}

// SetWriteNewSeriesLimitPerShardPerSecond mocks base method
func (_m *MockOptions) SetWriteNewSeriesLimitPerShardPerSecond(value int) Options {
	ret := _m.ctrl.Call(_m, "SetWriteNewSeriesLimitPerShardPerSecond", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetWriteNewSeriesLimitPerShardPerSecond indicates an expected call of SetWriteNewSeriesLimitPerShardPerSecond
func (_mr *MockOptionsMockRecorder) SetWriteNewSeriesLimitPerShardPerSecond(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetWriteNewSeriesLimitPerShardPerSecond", reflect.TypeOf((*MockOptions)(nil).SetWriteNewSeriesLimitPerShardPerSecond), arg0)
}

// WriteNewSeriesLimitPerShardPerSecond mocks base method
func (_m *MockOptions) WriteNewSeriesLimitPerShardPerSecond() int {
	ret := _m.ctrl.Call(_m, "WriteNewSeriesLimitPerShardPerSecond")
	ret0, _ := ret[0].(int)
	return ret0
}

// WriteNewSeriesLimitPerShardPerSecond indicates an expected call of WriteNewSeriesLimitPerShardPerSecond
func (_mr *MockOptionsMockRecorder) WriteNewSeriesLimitPerShardPerSecond() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteNewSeriesLimitPerShardPerSecond", reflect.TypeOf((*MockOptions)(nil).WriteNewSeriesLimitPerShardPerSecond))
}

// SetTickSeriesBatchSize mocks base method
func (_m *MockOptions) SetTickSeriesBatchSize(value int) Options {
	ret := _m.ctrl.Call(_m, "SetTickSeriesBatchSize", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetTickSeriesBatchSize indicates an expected call of SetTickSeriesBatchSize
func (_mr *MockOptionsMockRecorder) SetTickSeriesBatchSize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetTickSeriesBatchSize", reflect.TypeOf((*MockOptions)(nil).SetTickSeriesBatchSize), arg0)
}

// TickSeriesBatchSize mocks base method
func (_m *MockOptions) TickSeriesBatchSize() int {
	ret := _m.ctrl.Call(_m, "TickSeriesBatchSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// TickSeriesBatchSize indicates an expected call of TickSeriesBatchSize
func (_mr *MockOptionsMockRecorder) TickSeriesBatchSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "TickSeriesBatchSize", reflect.TypeOf((*MockOptions)(nil).TickSeriesBatchSize))
}

// SetTickPerSeriesSleepDuration mocks base method
func (_m *MockOptions) SetTickPerSeriesSleepDuration(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetTickPerSeriesSleepDuration", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetTickPerSeriesSleepDuration indicates an expected call of SetTickPerSeriesSleepDuration
func (_mr *MockOptionsMockRecorder) SetTickPerSeriesSleepDuration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetTickPerSeriesSleepDuration", reflect.TypeOf((*MockOptions)(nil).SetTickPerSeriesSleepDuration), arg0)
}

// TickPerSeriesSleepDuration mocks base method
func (_m *MockOptions) TickPerSeriesSleepDuration() time.Duration {
	ret := _m.ctrl.Call(_m, "TickPerSeriesSleepDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TickPerSeriesSleepDuration indicates an expected call of TickPerSeriesSleepDuration
func (_mr *MockOptionsMockRecorder) TickPerSeriesSleepDuration() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "TickPerSeriesSleepDuration", reflect.TypeOf((*MockOptions)(nil).TickPerSeriesSleepDuration))
}

// SetTickMinimumInterval mocks base method
func (_m *MockOptions) SetTickMinimumInterval(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetTickMinimumInterval", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetTickMinimumInterval indicates an expected call of SetTickMinimumInterval
func (_mr *MockOptionsMockRecorder) SetTickMinimumInterval(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetTickMinimumInterval", reflect.TypeOf((*MockOptions)(nil).SetTickMinimumInterval), arg0)
}

// TickMinimumInterval mocks base method
func (_m *MockOptions) TickMinimumInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "TickMinimumInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TickMinimumInterval indicates an expected call of TickMinimumInterval
func (_mr *MockOptionsMockRecorder) TickMinimumInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "TickMinimumInterval", reflect.TypeOf((*MockOptions)(nil).TickMinimumInterval))
}

// SetMaxWiredBlocks mocks base method
func (_m *MockOptions) SetMaxWiredBlocks(value uint) Options {
	ret := _m.ctrl.Call(_m, "SetMaxWiredBlocks", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

// SetMaxWiredBlocks indicates an expected call of SetMaxWiredBlocks
func (_mr *MockOptionsMockRecorder) SetMaxWiredBlocks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetMaxWiredBlocks", reflect.TypeOf((*MockOptions)(nil).SetMaxWiredBlocks), arg0)
}

// MaxWiredBlocks mocks base method
func (_m *MockOptions) MaxWiredBlocks() uint {
	ret := _m.ctrl.Call(_m, "MaxWiredBlocks")
	ret0, _ := ret[0].(uint)
	return ret0
}

// MaxWiredBlocks indicates an expected call of MaxWiredBlocks
func (_mr *MockOptionsMockRecorder) MaxWiredBlocks() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MaxWiredBlocks", reflect.TypeOf((*MockOptions)(nil).MaxWiredBlocks))
}

// MockOptionsManager is a mock of OptionsManager interface
type MockOptionsManager struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsManagerMockRecorder
}

// MockOptionsManagerMockRecorder is the mock recorder for MockOptionsManager
type MockOptionsManagerMockRecorder struct {
	mock *MockOptionsManager
}

// NewMockOptionsManager creates a new mock instance
func NewMockOptionsManager(ctrl *gomock.Controller) *MockOptionsManager {
	mock := &MockOptionsManager{ctrl: ctrl}
	mock.recorder = &MockOptionsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockOptionsManager) EXPECT() *MockOptionsManagerMockRecorder {
	return _m.recorder
}

// Update mocks base method
func (_m *MockOptionsManager) Update(value Options) error {
	ret := _m.ctrl.Call(_m, "Update", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (_mr *MockOptionsManagerMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Update", reflect.TypeOf((*MockOptionsManager)(nil).Update), arg0)
}

// Get mocks base method
func (_m *MockOptionsManager) Get() Options {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Options)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockOptionsManagerMockRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockOptionsManager)(nil).Get))
}

// RegisterListener mocks base method
func (_m *MockOptionsManager) RegisterListener(l OptionsListener) close.SimpleCloser {
	ret := _m.ctrl.Call(_m, "RegisterListener", l)
	ret0, _ := ret[0].(close.SimpleCloser)
	return ret0
}

// RegisterListener indicates an expected call of RegisterListener
func (_mr *MockOptionsManagerMockRecorder) RegisterListener(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RegisterListener", reflect.TypeOf((*MockOptionsManager)(nil).RegisterListener), arg0)
}

// Close mocks base method
func (_m *MockOptionsManager) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockOptionsManagerMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockOptionsManager)(nil).Close))
}

// MockOptionsListener is a mock of OptionsListener interface
type MockOptionsListener struct {
	ctrl     *gomock.Controller
	recorder *MockOptionsListenerMockRecorder
}

// MockOptionsListenerMockRecorder is the mock recorder for MockOptionsListener
type MockOptionsListenerMockRecorder struct {
	mock *MockOptionsListener
}

// NewMockOptionsListener creates a new mock instance
func NewMockOptionsListener(ctrl *gomock.Controller) *MockOptionsListener {
	mock := &MockOptionsListener{ctrl: ctrl}
	mock.recorder = &MockOptionsListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockOptionsListener) EXPECT() *MockOptionsListenerMockRecorder {
	return _m.recorder
}

// SetRuntimeOptions mocks base method
func (_m *MockOptionsListener) SetRuntimeOptions(value Options) {
	_m.ctrl.Call(_m, "SetRuntimeOptions", value)
}

// SetRuntimeOptions indicates an expected call of SetRuntimeOptions
func (_mr *MockOptionsListenerMockRecorder) SetRuntimeOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRuntimeOptions", reflect.TypeOf((*MockOptionsListener)(nil).SetRuntimeOptions), arg0)
}
