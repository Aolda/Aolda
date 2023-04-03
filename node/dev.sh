#!/bin/bash

function aloda-dev() {
    # 현재 폴더 경로 저장
    CURRENT_DIR=$(pwd)

    # 명령어를 실행할 폴더로 이동
    cd /Users/ichanju/Desktop/아올다/Aolda/contract

    if [ "$1" == "compile" ]; then
        # compile 명령어 실행
        sudo npx hardhat compile
    elif [ "$1" == "run" ]; then
        # run 명령어 실행
        output=$(sudo npx hardhat run --network dev ./scripts/deploy.ts)
        echo "$output"
        contract_address=$( echo "$output" | grep "aoldaClient" | awk '{print $2}')
        if [ -n "$contract_address" ]; then
            sed -i '' '/^CONTRACT_ADDRESS=/d' /Users/ichanju/Desktop/아올다/Aolda/node/.env | sed -e :a -e '/^\n*$/{$d;N;ba' -e '}'
            echo "CONTRACT_ADDRESS=$contract_address" >> /Users/ichanju/Desktop/아올다/Aolda/node/.env
        fi
    elif [ "$1" == "rm" ]; then
        # rm 명령어 실행
        sudo rm -rf cache
        sudo rm -rf artifacts 
    else
        echo "Invalid command"
        exit 1
    fi

    # 원래 폴더로 이동
    cd $CURRENT_DIR
}

# 인자로 폴더 경로와 명령어를 받아서 실행
aloda-dev $1
