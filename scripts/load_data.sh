#!/bin/bash

mkdir -p ./tmp

curl -L http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz -o ./tmp/enron_mail_20110402.tgz
tar -xzf ./tmp/enron_mail_20110402.tgz -C ./tmp

echo "Build indexer"
cd scripts && go build indexer.go && cd ..
echo "Run indexer"
./scripts/indexer ./tmp/enron_mail_20110402/maildir

./scripts/convert_json_to_ndjson.sh
cd tmp/dataEmails/ndjson
dir="./"
for file in "$dir"/*.ndjson; do
    base=$(basename "$file")
    curl http://localhost:4080/api/_bulk -i -u admin:admin --data-binary "@$base"
done