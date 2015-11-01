package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func buildVolumes(userVolumes []interface{}) []api.Volume {
	volumes := make([]api.Volume, len(userVolumes))
	if len(userVolumes) == 0 {
		return volumes
	}

	for index, userVolume := range userVolumes {
		volumes[index] = *buildVolume(userVolume.(map[string]interface{}))
	}

	return volumes
}

func buildVolume(userVolume map[string]interface{}) *api.Volume {
	return populateVolumeSource(&api.Volume{
			Name:	userVolume["name"].(string),
		}.(*api.VolumeSource), userVolume).(*api.Volume)
}

func populateVolumeSource(volumeSource *api.VolumeSource, userVolumeSource map[string]interface{}) *api.VolumeSource {
	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.HostPath = buildHostPathVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["empty_dir"]; ok {
		volumeSource.EmptyDir = buildEmptyDirVolumeSource(userVolumeSource["empty_dir"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.GCEPersistentDisk = buildGCEPersistentDiskVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.AWSElasticBlockStore = buildAWSElasticBlockStoreVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.GitRepo = buildGitRepoVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.Secret = buildSecretVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.NFS = buildHNFSVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.ISCSI = buildISCSIVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.Glusterfs = buildGlusterfsVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.PersistentVolumeClaim = buildPersistentVolumeClaimVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.RBD = buildRBDVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	return volume
}
