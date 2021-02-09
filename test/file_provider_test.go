package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFile1(t *testing.T) {
	url := "http://qa-test.dev.imtlab.io/payload1.json"
	jsonfile, err := GetFile(url)
	assert.Nil(t, err)
	assert.NotNil(t, jsonfile)
	assert.EqualValues(t, "imt", jsonfile.Customer.Name)
	assert.EqualValues(t, "la", jsonfile.Customer.Location)
	assert.EqualValues(t, "imt-la-1", jsonfile.Customer.Deployment)
	assert.EqualValues(t, "soda", jsonfile.Soda.App)
	assert.EqualValues(t, "21.1.1", jsonfile.Soda.Version)
	assert.EqualValues(t, "imt-la-1.cloudsoda.io", jsonfile.Soda.FQDN)
	assert.EqualValues(t, "soda", jsonfile.Soda.NameSpace)
	assert.EqualValues(t, "1000", jsonfile.Soda.Bandwidth)
	assert.EqualValues(t, true, jsonfile.Soda.Persistence.Enabled)
	assert.EqualValues(t, "fast", jsonfile.Soda.Persistence.Volume)
	assert.EqualValues(t, true, jsonfile.Soda.Backup.Enabled)
	assert.EqualValues(t, "NA", jsonfile.Soda.Backup.AwsAccessKey)
	assert.EqualValues(t, "NA", jsonfile.Soda.Backup.AwsAccessSecret)
	assert.EqualValues(t, "NA", jsonfile.Soda.Backup.AwsBucketName)
	assert.EqualValues(t, "10.146.177.223", jsonfile.MetalLb)
}

func TestGetFile2(t *testing.T) {
	url := "http://qa-test.dev.imtlab.io/payload1.json"
	jsonfile, err := GetFile(url)
	assert.Nil(t, err)
	assert.NotNil(t, jsonfile)
	assert.EqualValues(t, "imt", jsonfile.Customer.Name)
	assert.EqualValues(t, "la", jsonfile.Customer.Location)
	assert.EqualValues(t, "imt-la-1", jsonfile.Customer.Deployment)
	assert.EqualValues(t, "soda", jsonfile.Soda.App)
	assert.EqualValues(t, "21.1.1", jsonfile.Soda.Version)
	assert.EqualValues(t, "1000", jsonfile.Soda.Bandwidth)
	assert.EqualValues(t, true, jsonfile.Soda.Persistence.Enabled)
	assert.EqualValues(t, "local", jsonfile.Soda.Persistence.Volume)
	assert.EqualValues(t, "truee", jsonfile.Soda.Backup.Enabled)
	assert.EqualValues(t, "NA", jsonfile.Soda.Backup.AwsAccessKey)
	assert.EqualValues(t, "NA", jsonfile.Soda.Backup.AwsAccessSecret)
	assert.EqualValues(t, "NA", jsonfile.Soda.Backup.AwsBucketName)
	assert.EqualValues(t, "10.146.177.223", jsonfile.MetalLb)
}
