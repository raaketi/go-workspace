Prometheus Monitoring for application using Servicemonitor ,it can be used to list all the namespaces and lookup for port name "http-web".
  endpoints:
  - port: http-web
    path: /metrics
in kubernetes service (based on match expressions or Match labels available in Servicemonitor) and check if they provide /metrics endpoint and start scraping the metrics .

  selector:
    matchExpressions:
      - {key: app.kubernetes.io/managed-by, operator: Exists}