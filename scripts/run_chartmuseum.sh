#!/bin/sh

# ref: https://github.com/helm/chartmuseum

# ------------------------------------------------------------------------------
# const
c_CM_PORT=38080

# download and install chartmuseum
curl https://raw.githubusercontent.com/helm/chartmuseum/main/scripts/get-chartmuseum | bash 

# run chartmuseum
chartmuseum --debug --port=${c_CM_PORT}\
  --storage="local" \
  --storage-local-rootdir="./chartstorage"
