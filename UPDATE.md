# hacks to have fewer CRDs

Basic steps to build a custom provider-aws package without all the 900+ CRDs that come with the default provider.

High level steps gleaned from [here](https://github.com/crossplane/crossplane/issues/2869#issuecomment-1352141017)

[CRD scaling 1 pager](https://github.com/crossplane/crossplane/blob/master/design/one-pager-crd-scaling.md)

* start on master with latest from upstream
* fetch all tags
  ```
  git fetch --all
  ```

* check out the version you want and make a new branch
  ```
  git checkout v0.33.0
  git checkout -b platform-v0.33.0
  ```

* comment out unneeded apis from `config/externalname.go`
* comment out unneeded configure calls in `config/provider.go`
* ensure build submodule is checked out
  ```
  git submodule init 
  git submodule update
  ```
* make generate - watch for errors and fix - lol
* git add, commit, all the changes. tag and build
  ```
  export PROJECT_VERSION_TAG_GROUP=kompute
  git tag kompute/v0.33.0
  make build.all # should create oci packages under _output/xpkg/ for arm64 and amd64
  ```

* publish
  ```
  # load the tar ball into docker. Tag it and push it to ECR.
  # their publish process (under build/makelib/xpkg.mk) uses their upbound.io saas service somehow :shrug:
  BUILDIMG=$(docker load --input _output/xpkg/linux_amd64/provider-aws-v0.33.0.xpkg|cut -f3 -d":")
  docker tag ${BUILDIMG} 361053881171.dkr.ecr.ap-southeast-2.amazonaws.com/platform/upbound-aws:v0.33.0
  docker push 361053881171.dkr.ecr.ap-southeast-2.amazonaws.com/platform/upbound-aws:v0.33.0
  ```
