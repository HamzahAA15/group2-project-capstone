// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_certificateRepository is a generated GoMock package.
package mock_certificateRepository

import (
	reflect "reflect"
	certificateEntities "sirclo/project-capstone/entities/certificateEntities"

	gomock "github.com/golang/mock/gomock"
)

// MockCertificateInterface is a mock of CertificateInterface interface.
type MockCertificateInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateInterfaceMockRecorder
}

// MockCertificateInterfaceMockRecorder is the mock recorder for MockCertificateInterface.
type MockCertificateInterfaceMockRecorder struct {
	mock *MockCertificateInterface
}

// NewMockCertificateInterface creates a new mock instance.
func NewMockCertificateInterface(ctrl *gomock.Controller) *MockCertificateInterface {
	mock := &MockCertificateInterface{ctrl: ctrl}
	mock.recorder = &MockCertificateInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateInterface) EXPECT() *MockCertificateInterfaceMockRecorder {
	return m.recorder
}

// CountVaccineIsPending mocks base method.
func (m *MockCertificateInterface) CountVaccineIsPending(userID string, dossage int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountVaccineIsPending", userID, dossage)
	ret0, _ := ret[0].(int)
	return ret0
}

// CountVaccineIsPending indicates an expected call of CountVaccineIsPending.
func (mr *MockCertificateInterfaceMockRecorder) CountVaccineIsPending(userID, dossage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountVaccineIsPending", reflect.TypeOf((*MockCertificateInterface)(nil).CountVaccineIsPending), userID, dossage)
}

// GetCertificate mocks base method.
func (m *MockCertificateInterface) GetCertificate(id string) (certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificate", id)
	ret0, _ := ret[0].(certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificate indicates an expected call of GetCertificate.
func (mr *MockCertificateInterfaceMockRecorder) GetCertificate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockCertificateInterface)(nil).GetCertificate), id)
}

// GetCertificateUser mocks base method.
func (m *MockCertificateInterface) GetCertificateUser(userID string) ([]certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateUser", userID)
	ret0, _ := ret[0].([]certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateUser indicates an expected call of GetCertificateUser.
func (mr *MockCertificateInterfaceMockRecorder) GetCertificateUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateUser", reflect.TypeOf((*MockCertificateInterface)(nil).GetCertificateUser), userID)
}

// GetCertificates mocks base method.
func (m *MockCertificateInterface) GetCertificates(orderBy string) ([]certificateEntities.Certificates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificates", orderBy)
	ret0, _ := ret[0].([]certificateEntities.Certificates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificates indicates an expected call of GetCertificates.
func (mr *MockCertificateInterfaceMockRecorder) GetCertificates(orderBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificates", reflect.TypeOf((*MockCertificateInterface)(nil).GetCertificates), orderBy)
}

// GetVaccineDose mocks base method.
func (m *MockCertificateInterface) GetVaccineDose(userID, status string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVaccineDose", userID, status)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetVaccineDose indicates an expected call of GetVaccineDose.
func (mr *MockCertificateInterfaceMockRecorder) GetVaccineDose(userID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaccineDose", reflect.TypeOf((*MockCertificateInterface)(nil).GetVaccineDose), userID, status)
}

// UploadCertificateVaccine mocks base method.
func (m *MockCertificateInterface) UploadCertificateVaccine(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadCertificateVaccine", certificate)
	ret0, _ := ret[0].(certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadCertificateVaccine indicates an expected call of UploadCertificateVaccine.
func (mr *MockCertificateInterfaceMockRecorder) UploadCertificateVaccine(certificate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadCertificateVaccine", reflect.TypeOf((*MockCertificateInterface)(nil).UploadCertificateVaccine), certificate)
}

// VerifyCertificate mocks base method.
func (m *MockCertificateInterface) VerifyCertificate(certificate certificateEntities.Certificate) (certificateEntities.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyCertificate", certificate)
	ret0, _ := ret[0].(certificateEntities.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyCertificate indicates an expected call of VerifyCertificate.
func (mr *MockCertificateInterfaceMockRecorder) VerifyCertificate(certificate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyCertificate", reflect.TypeOf((*MockCertificateInterface)(nil).VerifyCertificate), certificate)
}
