/*
Copyright 2020 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"

	snapshotv1beta1api "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1beta1"

	v1 "github.com/j4m3s-s/velero/pkg/apis/velero/v1"
	persistence "github.com/j4m3s-s/velero/pkg/persistence"
	volume "github.com/j4m3s-s/velero/pkg/volume"
)

// BackupStore is an autogenerated mock type for the BackupStore type
type BackupStore struct {
	mock.Mock
}

// BackupExists provides a mock function with given fields: bucket, backupName
func (_m *BackupStore) BackupExists(bucket string, backupName string) (bool, error) {
	ret := _m.Called(bucket, backupName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(bucket, backupName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bucket, backupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBackup provides a mock function with given fields: name
func (_m *BackupStore) DeleteBackup(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRestore provides a mock function with given fields: name
func (_m *BackupStore) DeleteRestore(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBackupContents provides a mock function with given fields: name
func (_m *BackupStore) GetBackupContents(name string) (io.ReadCloser, error) {
	ret := _m.Called(name)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupMetadata provides a mock function with given fields: name
func (_m *BackupStore) GetBackupMetadata(name string) (*v1.Backup, error) {
	ret := _m.Called(name)

	var r0 *v1.Backup
	if rf, ok := ret.Get(0).(func(string) *v1.Backup); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Backup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupVolumeSnapshots provides a mock function with given fields: name
func (_m *BackupStore) GetBackupVolumeSnapshots(name string) ([]*volume.Snapshot, error) {
	ret := _m.Called(name)

	var r0 []*volume.Snapshot
	if rf, ok := ret.Get(0).(func(string) []*volume.Snapshot); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*volume.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDownloadURL provides a mock function with given fields: target
func (_m *BackupStore) GetDownloadURL(target v1.DownloadTarget) (string, error) {
	ret := _m.Called(target)

	var r0 string
	if rf, ok := ret.Get(0).(func(v1.DownloadTarget) string); ok {
		r0 = rf(target)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v1.DownloadTarget) error); ok {
		r1 = rf(target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPodVolumeBackups provides a mock function with given fields: name
func (_m *BackupStore) GetPodVolumeBackups(name string) ([]*v1.PodVolumeBackup, error) {
	ret := _m.Called(name)

	var r0 []*v1.PodVolumeBackup
	if rf, ok := ret.Get(0).(func(string) []*v1.PodVolumeBackup); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.PodVolumeBackup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsValid provides a mock function with given fields:
func (_m *BackupStore) IsValid() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListBackups provides a mock function with given fields:
func (_m *BackupStore) ListBackups() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// PutBackup provides a mock function with given fields: info
func (_m *BackupStore) PutBackup(info persistence.BackupInfo) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(persistence.BackupInfo) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRestoreLog provides a mock function with given fields: backup, restore, log
func (_m *BackupStore) PutRestoreLog(backup string, restore string, log io.Reader) error {
	ret := _m.Called(backup, restore, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) error); ok {
		r0 = rf(backup, restore, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRestoreResults provides a mock function with given fields: backup, restore, results
func (_m *BackupStore) PutRestoreResults(backup string, restore string, results io.Reader) error {
	ret := _m.Called(backup, restore, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) error); ok {
		r0 = rf(backup, restore, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *BackupStore) GetCSIVolumeSnapshots(backup string) ([]*snapshotv1beta1api.VolumeSnapshot, error) {
	panic("Not implemented")
	return nil, nil
}

func (_m *BackupStore) GetCSIVolumeSnapshotContents(backup string) ([]*snapshotv1beta1api.VolumeSnapshotContent, error) {
	panic("Not implemented")
	return nil, nil
}
