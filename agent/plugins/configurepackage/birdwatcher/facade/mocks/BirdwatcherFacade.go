// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import request "github.com/aws/aws-sdk-go/aws/request"
import ssm "github.com/aws/aws-sdk-go/service/ssm"

// BirdwatcherFacade is an autogenerated mock type for the BirdwatcherFacade type
type BirdwatcherFacade struct {
	mock.Mock
}

// DescribeDocument provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) DescribeDocument(_a0 *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ssm.DescribeDocumentOutput
	if rf, ok := ret.Get(0).(func(*ssm.DescribeDocumentInput) *ssm.DescribeDocumentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.DescribeDocumentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ssm.DescribeDocumentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDocumentRequest provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) DescribeDocumentRequest(_a0 *ssm.DescribeDocumentInput) (*request.Request, *ssm.DescribeDocumentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ssm.DescribeDocumentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ssm.DescribeDocumentOutput
	if rf, ok := ret.Get(1).(func(*ssm.DescribeDocumentInput) *ssm.DescribeDocumentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ssm.DescribeDocumentOutput)
		}
	}

	return r0, r1
}

// GetDocument provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) GetDocument(_a0 *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ssm.GetDocumentOutput
	if rf, ok := ret.Get(0).(func(*ssm.GetDocumentInput) *ssm.GetDocumentOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.GetDocumentOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ssm.GetDocumentInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDocumentRequest provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) GetDocumentRequest(_a0 *ssm.GetDocumentInput) (*request.Request, *ssm.GetDocumentOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ssm.GetDocumentInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ssm.GetDocumentOutput
	if rf, ok := ret.Get(1).(func(*ssm.GetDocumentInput) *ssm.GetDocumentOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ssm.GetDocumentOutput)
		}
	}

	return r0, r1
}

// GetManifest provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) GetManifest(_a0 *ssm.GetManifestInput) (*ssm.GetManifestOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ssm.GetManifestOutput
	if rf, ok := ret.Get(0).(func(*ssm.GetManifestInput) *ssm.GetManifestOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.GetManifestOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ssm.GetManifestInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManifestRequest provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) GetManifestRequest(_a0 *ssm.GetManifestInput) (*request.Request, *ssm.GetManifestOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ssm.GetManifestInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ssm.GetManifestOutput
	if rf, ok := ret.Get(1).(func(*ssm.GetManifestInput) *ssm.GetManifestOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ssm.GetManifestOutput)
		}
	}

	return r0, r1
}

// PutConfigurePackageResult provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) PutConfigurePackageResult(_a0 *ssm.PutConfigurePackageResultInput) (*ssm.PutConfigurePackageResultOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ssm.PutConfigurePackageResultOutput
	if rf, ok := ret.Get(0).(func(*ssm.PutConfigurePackageResultInput) *ssm.PutConfigurePackageResultOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.PutConfigurePackageResultOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ssm.PutConfigurePackageResultInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutConfigurePackageResultRequest provides a mock function with given fields: _a0
func (_m *BirdwatcherFacade) PutConfigurePackageResultRequest(_a0 *ssm.PutConfigurePackageResultInput) (*request.Request, *ssm.PutConfigurePackageResultOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ssm.PutConfigurePackageResultInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ssm.PutConfigurePackageResultOutput
	if rf, ok := ret.Get(1).(func(*ssm.PutConfigurePackageResultInput) *ssm.PutConfigurePackageResultOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ssm.PutConfigurePackageResultOutput)
		}
	}

	return r0, r1
}