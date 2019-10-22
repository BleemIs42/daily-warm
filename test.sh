#!/usr/bin/env bash

echo -e "\
package main

// HTML for email template
const HTML = \`
$(cat email.html)
\`"  > html.go 

go run *.go > test.html


