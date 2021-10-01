#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./enable-mcas.sh <namespace>"
	echo "./enable-mcas.sh lb-ns"
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
# Enable MCAS
enable_mcas() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then

		curl -sX POST ${c_URL_LADYBUG_NS}/mcas -H "${c_CT}" -H"${c_AUTH}";

	elif [ "$LADYBUG_CALL_METHOD" == "GRPC" ]; then

		#$APP_ROOT/src/grpc-api/cbadm/cbadm cluster create --config $APP_ROOT/src/grpc-api/cbadm/grpc_conf.yaml -i json -o json 
		echo "[ERROR] not yet implemented"; exit -1;

	else
		echo "[ERROR] missing LADYBUG_CALL_METHOD"; exit -1;
	fi
	
}


# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	enable_mcas;
fi
