#!/bin/bash

docker-compose up -d

echo "Esperando a que el contenedor esté listo..."
while [ "$(docker inspect -f {{.State.Running}} zincsearch)" != "true" ]; do
  sleep 2
done

echo "Ejecutando script de carga de datos..."
./scripts/load_data.sh