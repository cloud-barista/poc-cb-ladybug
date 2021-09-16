#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./create-cluster.sh <namespace> <clsuter name>"
	echo "./create-cluster.sh lb-ns my-cluster"
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

# 2. Cluster Name
if [ "$#" -gt 1 ]; then v_CLUSTER_NAME="$2"; else	v_CLUSTER_NAME="${CLUSTER_NAME}"; fi
if [ "${v_CLUSTER_NAME}" == "" ]; then 
	read -e -p "Cluster name  ? : "  v_CLUSTER_NAME
fi
if [ "${v_CLUSTER_NAME}" == "" ]; then echo "[ERROR] missing <cluster name>"; exit -1; fi


c_URL_LADYBUG_NS="${c_URL_LADYBUG}/ns/${v_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Namespace                  is '${v_NAMESPACE}'"
echo "- Cluster name               is '${v_CLUSTER_NAME}'"


# ------------------------------------------------------------------------------
# Create a cluster
create_cluster() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then

		resp=$(curl -sX POST ${c_URL_LADYBUG_NS}/clusters -H "${c_CT}" -H "${c_AUTH}" -d @- <<EOF
		{
			"name": "${v_CLUSTER_NAME}",
			"config": {
				"kubernetes": {
					"networkCni": "kilo",
					"podCidr": "10.244.0.0/16",
					"serviceCidr": "10.96.0.0/12",
					"serviceDnsDomain": "cluster.local"
				}
			},
			"controlPlane": [
				{
					"connection": "config-aws-ap-northeast-2",
					"count": 1,
					"spec": "t2.medium"
				}
			],
			"worker": [
				{
					"connection": "config-aws-ap-northeast-1",
					"count": 1,
					"spec": "t2.small"
				},
				{
					"connection": "config-gcp-asia-northeast3",
					"count": 1,
					"spec": "n1-standard-2"
				}
			]
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
	create_cluster;
fi
