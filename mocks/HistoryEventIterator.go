// Copyright (c) 2017-2021 Uber Technologies Inc.
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

// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	shared "go.uber.org/cadence/.gen/go/shared"
)

// HistoryEventIterator is an autogenerated mock type for the HistoryEventIterator type
type HistoryEventIterator struct {
	mock.Mock
}

// HasNext provides a mock function with given fields:
func (_m *HistoryEventIterator) HasNext() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *HistoryEventIterator) Next() (*shared.HistoryEvent, error) {
	ret := _m.Called()

	var r0 *shared.HistoryEvent
	if rf, ok := ret.Get(0).(func() *shared.HistoryEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.HistoryEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHistoryEventIterator interface {
	mock.TestingT
	Cleanup(func())
}

// NewHistoryEventIterator creates a new instance of HistoryEventIterator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHistoryEventIterator(t mockConstructorTestingTNewHistoryEventIterator) *HistoryEventIterator {
	mock := &HistoryEventIterator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
