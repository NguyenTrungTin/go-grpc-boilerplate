#!/usr/bin/env bash

set -e

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$(dirname "$CURRENT_DIR")"
GRPC_CERTS_DIR=$ROOT_DIR/certs/grpc

mkdir -p "$GRPC_CERTS_DIR"

C="VN"
ST="Lam Dong"
L="Da Lat"
O="Awesome"
CN="Company"
emailAddress="admin@company.com"

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 3650 -nodes -keyout "$GRPC_CERTS_DIR/ca-key.pem" -out "$GRPC_CERTS_DIR/ca-cert.pem" -subj "/C=$C/ST=$ST/L=$L/O=$O/CN=$CN/emailAddress=$emailAddress"

echo "CA's self-signed certificate"
openssl x509 -in "$GRPC_CERTS_DIR/ca-cert.pem" -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout "$GRPC_CERTS_DIR/server-key.pem" -out "$GRPC_CERTS_DIR/server-req.pem" -subj "/C=$C/ST=$ST/L=$L/O=$O/CN=$CN/emailAddress=$emailAddress"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in "$GRPC_CERTS_DIR/server-req.pem" -days 3650 -CA "$GRPC_CERTS_DIR/ca-cert.pem" -CAkey "$GRPC_CERTS_DIR/ca-key.pem" -CAcreateserial -out "$GRPC_CERTS_DIR/server-cert.pem" -extfile "$GRPC_CERTS_DIR/server-ext.cnf"

echo "Server's signed certificate"
openssl x509 -in "$GRPC_CERTS_DIR/server-cert.pem" -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout "$GRPC_CERTS_DIR/client-key.pem" -out "$GRPC_CERTS_DIR/client-req.pem" -subj "/C=$C/ST=$ST/L=$L/O=$O/CN=$CN/emailAddress=$emailAddress"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in "$GRPC_CERTS_DIR/client-req.pem" -days 3650 -CA "$GRPC_CERTS_DIR/ca-cert.pem" -CAkey "$GRPC_CERTS_DIR/ca-key.pem" -CAcreateserial -out "$GRPC_CERTS_DIR/client-cert.pem" -extfile "$GRPC_CERTS_DIR/client-ext.cnf"

echo "Client's signed certificate"
openssl x509 -in "$GRPC_CERTS_DIR/client-cert.pem" -noout -text
