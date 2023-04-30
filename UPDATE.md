# hacks to have fewer CRDs

Basic steps to build a custom provider-aws package without all the 900+ CRDs that come with the default provider.

High level steps gleaned from [here](https://github.com/crossplane/crossplane/issues/2869#issuecomment-1352141017)

[CRD scaling 1 pager](https://github.com/crossplane/crossplane/blob/master/design/one-pager-crd-scaling.md)

* start on master with latest from upstream
* fetch all tags
  ```
  git fetch --all
  ```

* check out the release branch you want and make a new branch
  ```
  git checkout release-0.34
  git checkout -b platform-v0.34.0
  ```

* grab the previous modified versions of required files
  ```
  git checkout platform-v0.33.0 config/externalname.go
  git checkout platform-v0.33.0 config/provider.go
  git reset config/externalname.go config/provider.go
  ```

* check the diffs of these files
  - ensure any missing upstream changes are added to these files
  ```
  git icdiff
  ```

* ensure build submodule is checked out
  ```
  git submodule init 
  git submodule update
  ```

* comment out unneeded apis from `config/externalname.go`
* comment out unneeded configure calls in `config/provider.go`

* remove any api directories that are no longer needed.
  ```
  # remove all api folders
  ls -1 apis/ |grep -v -e .go -e v1alpha1 -e v1beta1|xargs -I % rm -rf apis/%
  # add back only those wanted
  grep Configure, config/provider.go|grep -v '//'|cut -f1 -d'.'|xargs -I % git checkout apis/%
  ```

* make generate - watch for errors and fix - lol
  ```
  make generate
  ```

* git add, commit, all the changes. tag and build
  ```
  git add package internal apis examples-generated config
  git commit -m "kompute release v0.34.0"
  export PROJECT_VERSION_TAG_GROUP=kompute
  git tag kompute/v0.34.0
  make build.all # should create oci packages under _output/xpkg/ for arm64 and amd64
  ```

* publish
  ```
  # load the tar ball into docker. Tag it and push it to ECR.
  # their publish process (under build/makelib/xpkg.mk) uses their upbound.io saas service somehow :shrug:
  BUILDIMG=$(docker load --input _output/xpkg/linux_amd64/provider-aws-v0.34.0.xpkg|cut -f3 -d":")
  docker tag ${BUILDIMG} 361053881171.dkr.ecr.ap-southeast-2.amazonaws.com/platform/upbound-aws:v0.34.0
  docker push 361053881171.dkr.ecr.ap-southeast-2.amazonaws.com/platform/upbound-aws:v0.34.0
  ```


```
(⎈ ap-k8s-developer-consumer-830726149330@omega-usea1-v1:consumer)
work/provider-aws [platform-v0.33.0●] » cp -Rp config /tmp/                 
(⎈ ap-k8s-developer-consumer-830726149330@omega-usea1-v1:consumer)
work/provider-aws [platform-v0.33.0●] » cp UPDATE.md /tmp/ 
(⎈ ap-k8s-developer-consumer-830726149330@omega-usea1-v1:consumer)
work/provider-aws [platform-v0.33.0●] » gco release-0.34    
M	build
Switched to branch 'release-0.34'
Your branch is up to date with 'origin/release-0.34'.
(⎈ ap-k8s-developer-consumer-830726149330@omega-usea1-v1:consumer)
work/provider-aws [release-0.34●] » gco -b platform-v0.34.0
Switched to a new branch 'platform-v0.34.0'
(⎈ ap-k8s-developer-consumer-830726149330@omega-usea1-v1:consumer)
work/provider-aws [platform-v0.34.0●] » cp /tmp/config/externalname.go config/externalname.go 
(⎈ ap-k8s-developer-consumer-830726149330@omega-usea1-v1:consumer)
work/provider-aws [platform-v0.34.0●] » cp /tmp/config/provider.go config/provider.go 
```