#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./get-package.sh <namespace> <package name> <version>"
	echo "./get-package.sh lb-ns my-package 1.0.0"
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

# 2. Package Name
if [ "$#" -gt 1 ]; then v_PACKAGE_NAME="$2"; else	v_PACKAGE_NAME="${PACKAGE_NAME}"; fi
if [ "${v_PACKAGE_NAME}" == "" ]; then 
	read -e -p "Package name ? : " v_PACKAGE_NAME
fi
if [ "${v_PACKAGE_NAME}" == "" ]; then echo "[ERROR] missing <package name>"; exit -1; fi

# 3. Package Version
if [ "$#" -gt 2 ]; then v_PACKAGE_VERSION="$3"; else	v_PACKAGE_VERSION="${PACKAGE_VERSION}"; fi
if [ "${v_PACKAGE_VERSION}" == "" ]; then 
	read -e -p "Package version ? : " v_PACKAGE_VERSION
fi
if [ "${v_PACKAGE_VERSION}" == "" ]; then echo "[ERROR] missing <package version>"; exit -1; fi


c_URL_LADYBUG_NS="${c_URL_LADYBUG}/ns/${v_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Namespace                  is '${v_NAMESPACE}'"
echo "- Package name               is '${v_PACKAGE_NAME}'"
echo "- Package version            is '${v_PACKAGE_VERSION}'"


# ------------------------------------------------------------------------------
# get a package
get_package_with_version() {

	if [ "$LADYBUG_CALL_METHOD" == "REST" ]; then
		
		curl -sX GET ${c_URL_LADYBUG_NS}/packages/${v_PACKAGE_NAME}/${v_PACKAGE_VERSION} -H "${c_CT}" -H "${c_AUTH}" | jq;

	elif [ "$LADYBUG_CALL_METHOD" == "GRPC" ]; then

		#$APP_ROOT/src/grpc-api/cbadm/cbadm package get --config $APP_ROOT/src/grpc-api/cbadm/grpc_conf.yaml -o json --ns ${v_NAMESPACE} --package ${v_PACKAGE_NAME}
		echo "[ERROR] not yet implementd"; exit -1;

		
	else
		echo "[ERROR] missing LADYBUG_CALL_METHOD"; exit -1;
	fi
	
}


# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	get_package_with_version;
fi
