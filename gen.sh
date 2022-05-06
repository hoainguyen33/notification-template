#!/bin/sh

gen --sqltype=mysql \
   	--connstr "root:uJKd+8WBrQcJ3RgM@tcp(localhost:3306)/getcare_messenger?charset=utf8&parseTime=True&loc=Local" \
   	--database getcare_messenger  \
   	--db \
   	--json \
   	--gorm \
   	--guregu \
   	--rest \
   	--out ./ \
   	--module getcare-messenger \
   	--makefile \
   	--json-fmt=snake \
   	--dao=repository \
   	--generate-dao \
   	--api=controller \
   	--no-overwrite \
   	--templateDir=templates \
   	--rest

mv controller/router.go route/api.go
rm -rf controller/http_utils.go