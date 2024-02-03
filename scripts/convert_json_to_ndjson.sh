#!/bin/bash

# Creando el directorio destino
mkdir -p ./tmp/dataEmails
mkdir -p ./tmp/dataEmails/ndjson

# Directorio que contiene los archivos JSON
dir="./tmp/dataEmails"
dirTo="./tmp/dataEmails/ndjson"

# Iterar sobre todos los archivos JSON en el directorio
for file in "$dir"/*.json; do
    # Obtener el nombre base del archivo (sin la extensión)
    base=$(basename "$file" .json)
    echo "Procesando $base..."

    cat "$file" | jq -c '.[]' > "$dirTo/$base.ndjson"
done