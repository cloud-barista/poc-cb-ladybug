#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./uninstall-app-instance.sh <namespace> <app instance name>"
	echo "./uninstall-app-instance.sh lb-ns app-instance-01"
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

c_URL_LADYBUG_NS="${c_URL_LADYBUG}/ns/${v_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Namespace                  is '${v_NAMESPACE}'"
echo "- App instance name          is '${v_INSTANCE_NAME}'"


# ------------------------------------------------------------------------------
# Uninstall an app instance 
uninstall_app_instance() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then

		curl -sX DELETE ${c_URL_LADYBUG_NS}/apps/${v_INSTANCE_NAME} -H "${c_CT}" -H "${c_AUTH}"| jq;

	elif [ "$LADYBUG_CALL_METHOD" == "GRPC" ]; then

		echo "[ERROR] not yet implemented"; exit -1;
		#$APP_ROOT/src/grpc-api/cbadm/cbadm cluster delete --config $APP_ROOT/src/grpc-api/cbadm/grpc_conf.yaml -o json --ns ${v_NAMESPACE} --cluster ${v_CLUSTER_NAME}

	else
		echo "[ERROR] missing LADYBUG_CALL_METHOD"; exit -1;
	fi
	
}


# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	uninstall_app_instance;
fi
