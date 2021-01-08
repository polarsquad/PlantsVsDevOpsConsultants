# Plants vs DevOps Consultants
This repository reflects the full PlantSquad platform designed, written, and operated by Polar Squad.

## Blog
The related blog posts can be found here:

- https://medium.com/polarsquad/plants-vs-devops-consultants-part-1-introduction-polar-squad-15ca3215b462

## Platform
The platform plant-server code is generated from protoc-gen-go, to obtain this, you need golang installed.

Install proto for your OS here: https://github.com/protocolbuffers/protobuf/releases

In addition to that, please install the following go packages, make sure your go env is set up.
In addition, make sure the installed packages are added to your path. `export PATH="$(go env GOBIN):$PATH"`

```shell
go install google.golang.org/grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/protobuf
```