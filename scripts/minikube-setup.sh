#!/usr/bin/env bash
set -euo pipefail

# Deploys the Student Management System to Minikube.
# Prerequisites: minikube, kubectl, docker (for image load).

MINIKUBE_PROFILE="${MINIKUBE_PROFILE:-minikube}"
IMAGE="indraplrg/technical-test:latest"

echo "==> Starting Minikube (profile: ${MINIKUBE_PROFILE})"
minikube start --profile "${MINIKUBE_PROFILE}" --driver=docker --memory=2048 --cpus=2

echo "==> Using Minikube Docker daemon"
eval "$(minikube --profile "${MINIKUBE_PROFILE}" docker-env)"

echo "==> Building image ${IMAGE}"
docker build -t "${IMAGE}" .

echo "==> Applying Kubernetes manifests"
kubectl apply -f k8s/postgres-configmap.yaml
kubectl apply -f k8s/postgres-secret.yaml
kubectl apply -f k8s/postgres-pvc.yaml
kubectl apply -f k8s/postgres-deployment.yaml
kubectl apply -f k8s/postgres-service.yaml
kubectl apply -f k8s/backend-configmap.yaml
kubectl apply -f k8s/backend-secret.yaml
kubectl apply -f k8s/backend-deployment.yaml
kubectl apply -f k8s/backend-service.yaml

echo "==> Waiting for rollout"
kubectl rollout status deployment/postgres --timeout=180s
kubectl rollout status deployment/backend --timeout=180s

echo "==> Service URL"
minikube --profile "${MINIKUBE_PROFILE}" service backend --url

echo "Done."