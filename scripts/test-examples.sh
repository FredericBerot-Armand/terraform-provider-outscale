#!/bin/bash

set -e

project_dir=$(cd "$(dirname $0)" && pwd)
project_root=$(cd $project_dir/.. && pwd)
EXAMPLES_DIR=$project_root/examples

for f in $EXAMPLES_DIR/keypair
do
    if [ -d $f ]
    then
        cd $f
        echo $f
        terraform init
        terraform apply -auto-approve
        terraform destroy -auto-approve
        cd -
    fi
done

exit 0