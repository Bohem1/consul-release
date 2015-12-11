# consul-release
---

This is a [BOSH](http://bosh.io) release for [consul](https://github.com/hashicorp/consul).

* [CI](https://mega.ci.cf-app.com/pipelines/consul)
* [Roadmap](https://www.pivotaltracker.com/n/projects/1382120)

###Contents

1. [Deploying](#deploying)
2. [Running Tests](#running-tests)

## Deploying

In order to deploy consul-release you must follow the standard steps for deploying software with BOSH.

We assume you have already deployed and targeted a BOSH director. For more instructions on how to do that please see the [BOSH documentation](http://bosh.io/docs).

### 1. Uploading a stemcell

Find the "BOSH Lite Warden" stemcell you wish to use. [bosh.io](https://bosh.io/stemcells) provides a resource to find and download stemcells.  Then run `bosh upload release STEMCELL_URL_OR_PATH_TO_DOWNLOADED_STEMCELL`.

### 2. Creating a release

From within the consul-release director run `bosh create release --force` to create a development release.

### 3. Uploading a release

Once you've created a development release run `bosh upload release` to upload your development release to the director.

### 4. Generating a deployment manifest

We provide a set of scripts and templates to generate a simple deployment manifest. You should use these as a starting point for creating your own manifest, but they should not be considered comprehensive or production-ready.

In order to automatically generate a manifest you must have installed [spiff](https://github.com/cloudfoundry-incubator/spiff).  Once installed, manifests can be generated using `./scripts/generate_consul_deployment_manifest [STUB LIST]` with the provided stubs:

1. director_uuid_stub
	
	The director_uuid_stub provides the uuid for the currently targeted BOSH director.
	```yaml
	---
	director_uuid: DIRECTOR_UUID
	```
2. instance_count_stub

	The instance count stub provides the ability to overwrite the number of instances of consul to deploy. The minimal deployment of consul is shown below:
	```yaml
	---
	instance_count_overrides:
	  consul_z1:
	    instances: 1
	  consul_z2:
	    instances: 0
	```

	NOTE: at no time should you deploy only 2 instances of consul.
3. persistent_disk_stub

	The persistent disk stub allows you to override the size of the persistent disk used in each instance of the consul job. If you wish to use the default settings provide a stub with only an empty hash:
	```yaml
	---
	persistent_disk_overrides: {}
	```
	
	To override disk sizes the format is as follows
	```yaml
	---
	persistent_disk_overrides:
	  consul_z1: 1234
	  consul_z2: 1234	
	```
	
4. iaas_settings

	The IaaS settings stub contains IaaS-specific settings, including networks, cloud properties, and compilation properties. Please see the BOSH documentation for setting up networks and subnets on your IaaS of choice. We currently allow for three network configurations on your IaaS: consul1, consul2, and compilation. You must also specify the stemcell to deploy against as well as the version (or latest).
	
We provide [default stubs for a BOSH-Lite deployment](https://github.com/cloudfoundry-incubator/consul-release/blob/master/manifest-generation/bosh-lite-stubs).  Specifically:

* instance_count_stub: [manifest-generation/bosh-lite-stubs/instance-count-overrides.yml](manifest-generation/bosh-lite-stubs/instance-count-overrides.yml)
* persistent_disk_stub: [manifest-generation/bosh-lite-stubs/persistent-disk-overrides.yml](manifest-generation/bosh-lite-stubs/persistent-disk-overrides.yml)
* iaas_settings: [manifest-generation/bosh-lite-stubs/iaas-settings-consul.yml](manifest-generation/bosh-lite-stubs/iaas-settings-consul.yml)

[Optional]

1. If you wish to override the name of the release and the deployment (default: consul) you can provide a release_name_stub with the following format:
	
	```yaml
	---
	name_overrides:
	  release_name: NAME
	  deployment_name: NAME
	```

Output the result of the above command to a file: `./scripts/generate_consul_deployment_manifest [STUB LIST] > OUTPUT_MANIFEST_PATH`.

### 5. Deploy.

Run `bosh -d OUTPUT_MANIFEST_PATH deploy`.

## Confab Tests

Run the `confab` tests by executing the `src/confab/scripts/test` executable.

## Acceptance Tests

The acceptance tests deploy a new consul cluster and exercise a variety of features, including scaling the number of nodes, as well as destructive testing to verify resilience.

### Prerequisites

The following should be installed on the local machine:

- Consul
- Golang (>= 1.5)

If using homebrew, these can be installed with:

```
brew install consul go
```

### Network setup

#### BOSH-Lite

Make sure you’ve run `bin/add-route`.
This will setup some routing rules to give the tests access to the consul VMs.

#### AWS

You will want to run your tests from a VM within the same subnet as determined in your iaas-settings stub.
This assumes you are using a private subnet within a VPC.

### Environment setup

This repository assumes that it is the root of your `GOPATH`. You can set this up by doing the following:

```shell
source .envrc
```

Or if you have `direnv` installed:

```shell
direnv allow
```

### Running the CONSATS

#### Running locally

Run all the tests with:

```
CONSATS_CONFIG=[config_file.json] ./scripts/test
```

Run a specific set of tests with:

```
CONSATS_CONFIG=[config_file.json] ./scripts/test <some test packages>
```

The `CONSATS_CONFIG` environment variable points to a configuration file which specifies the endpoint of the BOSH director.
When specifying location of the CONSATS_CONFIG, it must be an absolute path on the filesystem.

See below for more information on the contents of this configuration file.

### CONSATS config

An example config json for BOSH-lite would look like:

```json
cat > integration_config.json << EOF
{
  "bosh_target": "192.168.50.4",
  "bosh_username": "admin",
  "bosh_password": "admin"
}
EOF
export CONSATS_CONFIG=$PWD/integration_config.json
```

The full set of config parameters is explained below:
* `bosh_target` (required) Public BOSH IP address that will be used to host test environment.
* `bosh_username` (required) Username for the BOSH director login.
* `bosh_password` (required) Password for the BOSH director login.

#### Running as BOSH errand

The `acceptance-tests` BOSH errand assumes that the BOSH director has already uploaded the correct versions of the dependent releases.
The required releases are:
* [bosh-warden-cpi-release](http://bosh.io/releases/github.com/cppforlife/bosh-warden-cpi-release?version=28)
* [turbulence-release](http://bosh.io/releases/github.com/cppforlife/turbulence-release?version=0.4)
* [consul-release](http://bosh.io/releases/github.com/cloudfoundry-incubator/consul-release) or `bosh create release && bosh upload release`
