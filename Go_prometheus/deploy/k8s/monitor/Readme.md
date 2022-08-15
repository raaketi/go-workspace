Prometheus Monitoring for application using Servicemonitor ,it can be used to list all the namespaces and lookup for port name "http-web".
  endpoints:
  - port: http-prom-metrics
    path: /metrics
in kubernetes service (based on match expressions or Match labels available in Servicemonitor) and check if they provide /metrics endpoint and start scraping the metrics .

  selector:
    matchExpressions:
      - {key: app.kubernetes.io/managed-by, operator: Exists}


Alternative way to scrape the metrics of any application 

for this approach we dont require a servicemonitors

install Prometheus Stack using helm chart 

helm install prometheus prometheus-community/kube-prometheus-stack -n prometheus

it will install prometheus stack with default scrape config where we might not find the scrape config for generic application service endpoints, so we add it using upgrading the scrape config using values.yaml

prom_custom_values.yaml can be found in current directory 

helm upgrade \
  --namespace prometheus \
  -f prom_custom_values.yaml \
  prometheus prometheus-community/kube-prometheus-stack


and Make sure you have added a prometheus annotation to the service of the applications

prometheus.io/scrape: "true"

and also we are scraping the endpoints which have *metrics as the name of the port inthe service 

if you want to check the promehteus helm chart we have used ,run below command to get the tar file in the local directory 


helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm pull prometheus-community/kube-prometheus-stack
