## Core Openreplay application configuration folder

  This folder contains configuration for core openreplay apps. All applications share common helm chart named *openreplay* which can be overridden by `<applicaion>.yaml` file.
  
  **Below is a sample template.**
  
  ```yaml
  namespace: app        # In which namespace alert runs.
  image:
    repository: 998611063711.dkr.ecr.eu-central-1.amazonaws.com/alert # Which image to use
    pullPolicy: IfNotPresent
    tag: "latest"       # Overrides the image tag whose default is the chart appVersion.

  imagePullSecrets:     
    # If needed credentials to pull the image.
    # aws-registry is created using script docker_registry.sh.
    - name: aws-registry 

  service:
    type: ClusterIP
    port: 9000

  resources:
    limits:
      cpu: 256m
      memory: 512Mi
    requests:
      cpu: 100m
      memory: 128Mi

  # env vars for the application
  env:
    ALERT_NOTIFICATION_STRING: https://parrot.openreplay.io/alerts/notifications
    CLICKHOUSE_STRING: tcp://clickhouse.db.svc.cluster.local:9000/default
    POSTGRES_STRING: postgres://postgresql.db.svc.cluster.local:5432
  ```