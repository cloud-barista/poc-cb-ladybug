#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 2 ]; then 
	echo "./upload-package.sh <namespace> <package file>"
	echo "./upload-package.sh lb-ns my-package-1.0.0.tgz"
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

if [ "$#" -gt 1 ]; then v_PACKAGE_FILE="$2"; else	v_PACKAGE_FILE="${PACKAGE_FILE}"; fi
if [ "${v_PACKAGE_FILE}" == "" ]; then 
	read -e -p "Package file ? : " v_PACKAGE_FILE
fi
if [ "${v_PACKAGE_FILE}" == "" ]; then echo "[ERROR] missing <package file>"; exit -1; fi


c_URL_LADYBUG_NS="${c_URL_LADYBUG}/ns/${v_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Namespace                  is '${v_NAMESPACE}'"
echo "- Package file               is '${v_PACKAGE_FILE}'"


# ------------------------------------------------------------------------------
# Create a cluster
upload_package() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then

		resp=$(curl -sX POST ${c_URL_LADYBUG_NS}/packages -H"${c_AUTH}" -F "package=@${v_PACKAGE_FILE}" -w "%{response_code}") ;
		echo ${resp} | jq

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
	upload_package;
fi
