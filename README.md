#Policy engine microservice

## prerequsites for developement and testing
 - Go 1.15
 - Docker/podman (podman is preferred) 

## Steps to build policy-engine microservice
- Execute below commands inside policy-engine directory.
 
`Note Use either steps 1,2 and 3 OR 1,4 and 5`  
`Use step 2 & 3, if you want to push images to public repository quay.io`  
`Use step 4 & 5 if you want to push image to existing image registry`

1) go build
2) podman build -f Dockerfile -t quay.io/abdulrahuman_s/policy-engine:v1.0.0 .
3) podman push quay.io/abdulrahuman_s/policy-engine:v1.0.0

4) podman build -f Dockerfile -t default-route-openshift-image-registry.apps-crc.testing/policy-engine:v1.0.0 .
5) podman push default-route-openshift-image-registry.apps-crc.testing/policy-engine:v1.0.0

## Deploy image to Openshift
**nfmf-operator** is used to deploy **policy-engine** microservice in openshift cluster. Hence refer [nfmf-operator/README.md](../nfmf-operator/README.md)