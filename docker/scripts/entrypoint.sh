#!/bin/bash

ls -l /proj
java -jar /proj/build/libs/vertx-cellula-1.0-SNAPSHOT-fat.jar -cluster -conf /proj/src/main/conf/deployer-conf.json
