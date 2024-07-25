
#!/bin/bash -x
# Hinzufügen von Einträgen zur /etc/hosts Datei mit sudo
echo "192.168.49.2 test-app.com" | sudo tee -a /etc/hosts
echo "127.0.0.1 minikube.data.gov.au" | sudo tee -a /etc/hosts

# Anwenden der cert-manager YAML-Datei
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.15.1/cert-manager.yaml

# Warten, bis die cert-manager Pods laufen
echo "Warten auf den cert-manager..."
kubectl wait --for=condition=available --timeout=300s deployment --all -n cert-manager

# Anwenden der selfsigned-issuer YAML-Datei
kubectl apply -f https://gist.githubusercontent.com/t83714/51440e2ed212991655959f45d8d037cc/raw/7b16949f95e2dd61e522e247749d77bc697fd63c/selfsigned-issuer.yaml

# Minikube Ingress Addon aktivieren
minikube addons enable ingress

kubectl apply -n sdvdemo-certificate -f https://github.com/markusfelsner/sdvdemo-certificate/sdvdemo-certificate.yaml

