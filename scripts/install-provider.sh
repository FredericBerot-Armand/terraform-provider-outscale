#!/bin/bash

set -e
project_dir=$(cd "$(dirname $0)" && pwd)
project_root=$(cd $project_dir/.. && pwd)
BUILD_DIR=$project_root

go build -o terraform-provider-outscale_v0.5.4
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/outscale-dev/outscale/0.5.4/linux_amd64
mv terraform-provider-outscale_v0.5.4 ~/.terraform.d/plugins/registry.terraform.io/outscale-dev/outscale/0.5.4/linux_amd64/