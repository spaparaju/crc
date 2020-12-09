package bundle

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const reference = `{
  "version": "1.0",
  "type": "snc",
  "buildInfo": {
    "buildTime": "2020-10-26T04:48:26+00:00",
    "openshiftInstallerVersion": "./openshift-install v4.6.0\nbuilt from commit ebdbda57fc18d3b73e69f0f2cc499ddfca7e6593\nrelease image registry.svc.ci.openshift.org/origin/release:4.5",
    "sncVersion": "git4.1.14-137-g14e7"
  },
  "clusterInfo": {
    "openshiftVersion": "4.6.1",
    "clusterName": "crc",
    "baseDomain": "testing",
    "appsDomain": "apps-crc.testing",
    "sshPrivateKeyFile": "id_rsa_crc",
    "kubeConfig": "kubeconfig",
    "kubeadminPasswordFile": "kubeadmin-password"
  },
  "nodes": [
    {
      "kind": [
        "master",
        "worker"
      ],
      "hostname": "crc-h66l2-master-0",
      "diskImage": "crc.qcow2",
      "internalIP": "192.168.126.11"
    }
  ],
  "storage": {
    "diskImages": [
      {
        "name": "crc.qcow2",
        "format": "qcow2",
        "size": "12268273664",
        "sha256sum": "245a0e5acd4f09000a9a5f37d731082ed1cf3fdcad1b5320cbe9b153c9fd82a4"
      }
    ]
  },
  "driverInfo": {
    "name": "libvirt"
  }
}`

func TestUnmarshalMarshal(t *testing.T) {
	var bundle CrcBundleInfo
	assert.NoError(t, json.Unmarshal([]byte(reference), &bundle))
	assert.Equal(t, CrcBundleInfo{
		Version: "1.0",
		Type:    "snc",
		BuildInfo: BuildInfo{
			BuildTime:                 "2020-10-26T04:48:26+00:00",
			OpenshiftInstallerVersion: "./openshift-install v4.6.0\nbuilt from commit ebdbda57fc18d3b73e69f0f2cc499ddfca7e6593\nrelease image registry.svc.ci.openshift.org/origin/release:4.5",
			SncVersion:                "git4.1.14-137-g14e7",
		},
		ClusterInfo: ClusterInfo{
			OpenShiftVersion:      "4.6.1",
			ClusterName:           "crc",
			BaseDomain:            "testing",
			AppsDomain:            "apps-crc.testing",
			SSHPrivateKeyFile:     "id_rsa_crc",
			KubeConfig:            "kubeconfig",
			KubeadminPasswordFile: "kubeadmin-password",
		}, Nodes: []Node{
			{
				Kind:       []string{"master", "worker"},
				Hostname:   "crc-h66l2-master-0",
				DiskImage:  "crc.qcow2",
				InternalIP: "192.168.126.11",
			},
		},
		Storage: Storage{
			DiskImages: []DiskImage{
				{
					Name:     "crc.qcow2",
					Format:   "qcow2",
					Size:     "12268273664",
					Checksum: "245a0e5acd4f09000a9a5f37d731082ed1cf3fdcad1b5320cbe9b153c9fd82a4",
				},
			},
		},
		DriverInfo: DriverInfo{
			Name: "libvirt",
		},
		cachedPath: ""}, bundle)
	bin, err := json.Marshal(bundle)
	assert.NoError(t, err)
	assert.JSONEq(t, string(bin), reference)
}
