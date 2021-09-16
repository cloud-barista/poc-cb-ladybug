#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then
	echo "./get-all-clusters.sh <namespace>"
	echo "./get-all-clusters.sh lb-ns"
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

c_URL_LADYBUG_NS="${c_URL_LADYBUG}/ns/${v_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Namespace                  is '${v_NAMESPACE}'"


# ------------------------------------------------------------------------------
# get all clusters
get_all_clusters() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then

		curl -sX GET ${c_URL_LADYBUG_NS}/clusters -H "${c_CT}" -H "${c_AUTH}" | jq;

	elif [ "$LADYBUG_CALL_METHOD" == "GRPC" ]; then

		#$APP_ROOT/src/grpc-api/cbadm/cbadm cluster get --config $APP_ROOT/src/grpc-api/cbadm/grpc_conf.yaml -o json --ns ${v_NAMESPACE} --cluster ${v_CLUSTER_NAME}
		echo "[ERROR] not yet implementd"; exit -1;

	else
		echo "[ERROR] missing LADYBUG_CALL_METHOD"; exit -1;
	fi

}


# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then
	echo ""
	echo "------------------------------------------------------------------------------"
	get_all_clusters
fi
