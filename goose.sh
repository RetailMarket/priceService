#!/usr/bin/env bash
goose -env=test -pgschema=workflow $1
echo 'goose completed...'
