#!/usr/bin/env bash

echo -e "\
package main

// HTML for email template
const HTML = \`
$(cat msg.txt)
\`"  > html.go 

cp .env.msg .env

go run *.go

rm .env
