#!/bin/bash

set -e
project_dir=$(cd "$(dirname $0)" && pwd)
project_root=$(cd $project_dir/.. && pwd)
BUILD_DIR=$project_root/tests/qa_provider_oapi

python3 --version || (echo "We need 'python3' intalled to run integration tests"; exit 1)
python3 -m venv .venv
source .venv/bin/activate
pip --version || (echo "We need 'pip' intalled to run integration tests"; exit 1)

cd $BUILD_DIR
pip install -r requirements.txt
pytest -v ./test_provider_oapi.py
rm -fr terraform.d || exit 0
