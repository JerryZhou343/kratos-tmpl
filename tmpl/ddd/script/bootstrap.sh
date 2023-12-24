#!/usr/bin/env bash

script_dir=$(dirname "$(readlink -f "$0")")

current_dir="$PWD"

config_file=""

# 如果 SITE_ENV 未设置值，设置为 "dev"
export SITE_ENV="${SITE_ENV:-dev}"

if [[ $SITE_ENV == "prod" ]]
then
    config_file="$script_dir/config/config_prod.yaml"
elif [[ $SITE_ENV == "test" ]]
then
    config_file="$script_dir/config/config_test.yaml"
else
    config_file="$script_dir/config/config_dev.yaml"
fi

echo "use config file:$config_file"

$script_dir/bin/server -conf $config_file