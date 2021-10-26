#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./install-app-instance.sh <namespace> <app instance name> <app package name>"
	echo "./install-app-instance.sh lb-ns app-instance-01 app-package-01"
	exit 0; 
fi

source ./conf.env

# ------------------------------------------------------------------------------
# const


# -----------------------------------------------------------------
# parameter

# 1. namespace
if [ "$#" -gt 0 ]; then v_NAMESPACE="$1"; else	v_NAMESPACE="${NAMESPACE}"; fi
if [ "${v_NAMESPACE}" == "" ]; then 
	read -e -p "Namespace ? : " v_NAMESPACE
fi
if [ "${v_NAMESPACE}" == "" ]; then echo "[ERROR] missing <namespace>"; exit -1; fi

# 2. App Instance Name
if [ "$#" -gt 1 ]; then v_INSTANCE_NAME="$2"; else	v_INSTANCE_NAME="${INSTANCE_NAME}"; fi
if [ "${v_INSTANCE_NAME}" == "" ]; then 
	read -e -p "App instance name  ? : "  v_INSTANCE_NAME
fi
if [ "${v_INSTANCE_NAME}" == "" ]; then echo "[ERROR] missing <instance name>"; exit -1; fi

# 3. App Package Name
if [ "$#" -gt 2 ]; then v_PACKAGE_NAME="$3"; else	v_PACKAGE_NAME="${PACKAGE_NAME}"; fi
if [ "${v_PACKAGE_NAME}" == "" ]; then 
	read -e -p "App package name  ? : "  v_PACKAGE_NAME
fi
if [ "${v_PACKAGE_NAME}" == "" ]; then echo "[ERROR] missing <app package name>"; exit -1; fi


c_URL_LADYBUG_NS="${c_URL_LADYBUG}/ns/${v_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Namespace                  is '${v_NAMESPACE}'"
echo "- App instance name          is '${v_INSTANCE_NAME}'"
echo "- App package name           is '${v_PACKAGE_NAME}'"


# ------------------------------------------------------------------------------
# Install an app instance 
install_app_instance() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then

		resp=$(curl -sX POST ${c_URL_LADYBUG_NS}/apps -H "${c_CT}" -H "${c_AUTH}" -d @- <<EOF
		{
			"instance": "${v_INSTANCE_NAME}",
			"package": "${v_PACKAGE_NAME}",
			"wait": true,
			"force": false,
			"upgradeCRDs": false
		}
EOF
		); echo ${resp} | jq

	elif [ "$LADYBUG_CALL_METHOD" == "GRPC" ]; then

		echo "[ERROR] not yet implemented"; exit -1;
#		$APP_ROOT/src/grpc-api/cbadm/cbadm cluster create --config $APP_ROOT/src/grpc-api/cbadm/grpc_conf.yaml -i json -o json -d \
#		'{
#			"namespace":  "'${v_NAMESPACE}'",
#			"ReqInfo": {
#					"name": "'${v_CLUSTER_NAME}'",
#					"config": {
#						"kubernetes": {
#							"networkCni": "kilo",
#							"podCidr": "10.244.0.0/16",
#							"serviceCidr": "10.96.0.0/12",
#							"serviceDnsDomain": "cluster.local"
#						}
#					},
#					"controlPlane": [
#						{
#							"connection": "config-aws-ap-northeast-1",
#							"count": 1,
#							"spec": "t2.medium"
#						}
#					],
#					"worker": [
#						{
#							"connection": "config-gcp-asia-northeast3",
#							"count": 1,
#							"spec": "n1-standard-2"
#						}
#					]
#				}
#		}'	
#
	else
		echo "[ERROR] missing LADYBUG_CALL_METHOD"; exit -1;
	fi
	
}


# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	install_app_instance;
fi
