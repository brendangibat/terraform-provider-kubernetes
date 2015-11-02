package main

import (
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
	vol := &api.Volume{
		Name:	userVolume["name"].(string),
	}
	populateVolumeSource(&vol.VolumeSource, userVolume)
	return vol
}

func populateVolumeSource(volumeSource *api.VolumeSource, userVolumeSource map[string]interface{}) *api.VolumeSource {
	if _, ok := userVolumeSource["host_path"]; ok {
		volumeSource.HostPath = buildHostPathVolumeSource(userVolumeSource["host_path"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["empty_dir"]; ok {
		volumeSource.EmptyDir = buildEmptyDirVolumeSource(userVolumeSource["empty_dir"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["gce_persistent_disk"]; ok {
		volumeSource.GCEPersistentDisk = buildGCEPersistentDiskVolumeSource(userVolumeSource["gce_persistent_disk"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["aws_elastic_block_store"]; ok {
		volumeSource.AWSElasticBlockStore = buildAWSElasticBlockStoreVolumeSource(userVolumeSource["aws_elastic_block_store"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["git_repo"]; ok {
		volumeSource.GitRepo = buildGitRepoVolumeSource(userVolumeSource["git_repo"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["secret"]; ok {
		volumeSource.Secret = buildSecretVolumeSource(userVolumeSource["secret"].(map[string]interface{}))
	}

	if _, ok := userVolumeSource["persistent_volume_claim"]; ok {
		volumeSource.PersistentVolumeClaim = buildPersistentVolumeClaimVolumeSource(userVolumeSource["persistent_volume_claim"].(map[string]interface{}))
	}

	return volumeSource
}

func buildHostPathVolumeSource(userHostPathVolumeSource map[string]interface{}) *api.HostPathVolumeSource {
	return &api.HostPathVolumeSource{
		Path: userHostPathVolumeSource["path"].(string),
	}
}

func buildEmptyDirVolumeSource(userEmptyDirVolumeSource map[string]interface{}) *api.EmptyDirVolumeSource {
	return &api.EmptyDirVolumeSource{
		Medium: userEmptyDirVolumeSource["medium"].(api.StorageMedium),
	}
}

func buildGCEPersistentDiskVolumeSource(userGCEPersistentDiskVolumeSource map[string]interface{}) *api.GCEPersistentDiskVolumeSource {
	gce := &api.GCEPersistentDiskVolumeSource{
		PDName: userGCEPersistentDiskVolumeSource["pd_name"].(string),
	}

	if _, ok := userGCEPersistentDiskVolumeSource["fs_type"]; ok {
		gce.FSType = userGCEPersistentDiskVolumeSource["fs_type"].(string)
	}

	if _, ok := userGCEPersistentDiskVolumeSource["partition"]; ok {
		gce.Partition = userGCEPersistentDiskVolumeSource["partition"].(int)
	}

	if _, ok := userGCEPersistentDiskVolumeSource["read_only"]; ok {
		gce.ReadOnly = userGCEPersistentDiskVolumeSource["read_only"].(bool)
	}

	return gce
}

func buildAWSElasticBlockStoreVolumeSource(userAWSElasticBlockStoreVolumeSource map[string]interface{}) *api.AWSElasticBlockStoreVolumeSource {
	awsEBS := &api.AWSElasticBlockStoreVolumeSource{
		VolumeID: userAWSElasticBlockStoreVolumeSource["volume_id"].(string),
	}

	if _, ok := userAWSElasticBlockStoreVolumeSource["fs_type"]; ok {
		awsEBS.FSType = userAWSElasticBlockStoreVolumeSource["fs_type"].(string)
	}

	if _, ok := userAWSElasticBlockStoreVolumeSource["partition"]; ok {
		awsEBS.Partition = userAWSElasticBlockStoreVolumeSource["partition"].(int)
	}

	if _, ok := userAWSElasticBlockStoreVolumeSource["read_only"]; ok {
		awsEBS.ReadOnly = userAWSElasticBlockStoreVolumeSource["read_only"].(bool)
	}

	return awsEBS
}

func buildGitRepoVolumeSource(userGitRepoVolumeSource map[string]interface{}) *api.GitRepoVolumeSource {
	return &api.GitRepoVolumeSource{
		Repository: userGitRepoVolumeSource["repository"].(string),
		Revision: userGitRepoVolumeSource["revision"].(string),
	}
}

func buildSecretVolumeSource(userSecretVolumeSource map[string]interface{}) *api.SecretVolumeSource {
	return &api.SecretVolumeSource{
		SecretName: userSecretVolumeSource["secret_name"].(string),
	}
}

func buildPersistentVolumeClaimVolumeSource(userPersistentVolumeClaimVolumeSource map[string]interface{}) *api.PersistentVolumeClaimVolumeSource {
	persistentClaim := &api.PersistentVolumeClaimVolumeSource{
		ClaimName: userPersistentVolumeClaimVolumeSource["claim_name"].(string),
	}

	if _, ok := userPersistentVolumeClaimVolumeSource["read_only"]; ok {
		persistentClaim.ReadOnly = userPersistentVolumeClaimVolumeSource["read_only"].(bool)
	}

	return persistentClaim
}
