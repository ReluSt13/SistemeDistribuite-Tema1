#!/bin/bash

for i in {1..15}
do
    go run ./client/command_client.go -exercise=Exercitiul$i -client=Client$i &
done