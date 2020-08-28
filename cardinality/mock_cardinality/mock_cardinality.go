// Code generated by MockGen. DO NOT EDIT.
// Source: cloud/services/prometheus-cardinality-exporter/cardinality (interfaces: PrometheusGaugeVec,PrometheusClient)

package mock_cardinality

import (
	gomock "github.com/golang/mock/gomock"
	prometheus "github.com/prometheus/client_golang/prometheus"
	http "net/http"
)

// MockPrometheusGaugeVec is a mock of PrometheusGaugeVec interface
type MockPrometheusGaugeVec struct {
	ctrl     *gomock.Controller
	recorder *MockPrometheusGaugeVecMockRecorder
}

// MockPrometheusGaugeVecMockRecorder is the mock recorder for MockPrometheusGaugeVec
type MockPrometheusGaugeVecMockRecorder struct {
	mock *MockPrometheusGaugeVec
}

// NewMockPrometheusGaugeVec creates a new mock instance
func NewMockPrometheusGaugeVec(ctrl *gomock.Controller) *MockPrometheusGaugeVec {
	mock := &MockPrometheusGaugeVec{ctrl: ctrl}
	mock.recorder = &MockPrometheusGaugeVecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockPrometheusGaugeVec) EXPECT() *MockPrometheusGaugeVecMockRecorder {
	return _m.recorder
}

// Collect mocks base method
func (_m *MockPrometheusGaugeVec) Collect(_param0 chan<- prometheus.Metric) {
	_m.ctrl.Call(_m, "Collect", _param0)
}

// Collect indicates an expected call of Collect
func (_mr *MockPrometheusGaugeVecMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Collect", arg0)
}

// Delete mocks base method
func (_m *MockPrometheusGaugeVec) Delete(_param0 prometheus.Labels) bool {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Delete indicates an expected call of Delete
func (_mr *MockPrometheusGaugeVecMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

// Describe mocks base method
func (_m *MockPrometheusGaugeVec) Describe(_param0 chan<- *prometheus.Desc) {
	_m.ctrl.Call(_m, "Describe", _param0)
}

// Describe indicates an expected call of Describe
func (_mr *MockPrometheusGaugeVecMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Describe", arg0)
}

// GetMetricWith mocks base method
func (_m *MockPrometheusGaugeVec) GetMetricWith(_param0 prometheus.Labels) (prometheus.Gauge, error) {
	ret := _m.ctrl.Call(_m, "GetMetricWith", _param0)
	ret0, _ := ret[0].(prometheus.Gauge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricWith indicates an expected call of GetMetricWith
func (_mr *MockPrometheusGaugeVecMockRecorder) GetMetricWith(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetricWith", arg0)
}

// MockPrometheusClient is a mock of PrometheusClient interface
type MockPrometheusClient struct {
	ctrl     *gomock.Controller
	recorder *MockPrometheusClientMockRecorder
}

// MockPrometheusClientMockRecorder is the mock recorder for MockPrometheusClient
type MockPrometheusClientMockRecorder struct {
	mock *MockPrometheusClient
}

// NewMockPrometheusClient creates a new mock instance
func NewMockPrometheusClient(ctrl *gomock.Controller) *MockPrometheusClient {
	mock := &MockPrometheusClient{ctrl: ctrl}
	mock.recorder = &MockPrometheusClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockPrometheusClient) EXPECT() *MockPrometheusClientMockRecorder {
	return _m.recorder
}

// Do mocks base method
func (_m *MockPrometheusClient) Do(_param0 *http.Request) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Do", _param0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (_mr *MockPrometheusClientMockRecorder) Do(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Do", arg0)
}
