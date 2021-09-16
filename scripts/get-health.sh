#!/bin/bash
# -----------------------------------------------------------------
source ./conf.env

# ------------------------------------------------------------------------------
# const


# -----------------------------------------------------------------
# parameter


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Ladybug URL is '${c_URL_LADYBUG}'"


# ------------------------------------------------------------------------------
# get Node
get_health() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then
		
		curl -sX GET ${c_URL_LADYBUG}/health -H "${c_CT}" -H "${c_AUTH}";

	elif [ "$LADYBUG_CALL_METHOD" == "GRPC" ]; then

		#$APP_ROOT/src/grpc-api/cbadm/cbadm node get --config $APP_ROOT/src/grpc-api/cbadm/grpc_conf.yaml -o json --ns ${v_NAMESPACE} --cluster ${v_CLUSTER_NAME} --node ${v_NODE_NAME}	
		echo "[ERROR] not yet implemented"; exit -1;
		
	else
		echo "[ERROR] missing LADYBUG_CALL_METHOD"; exit -1;
	fi
	
}


# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	get_health;
fi
