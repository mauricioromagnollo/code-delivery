<div align='justify'>

# Code Delivery

> Code Delivery is a case study on a delivery system, where it is possible to visualize the delivery vehicle in real time.

## Requirements

- View the delivery vehicle in real time;
- Multiple simultaneous deliverers;
- Simulator service that will send the real-time position of each delivery person;
- The data of each delivery, as well as positions, will be stored in Elasticsearch for further analysis;

## Non Functional Requirements

- To avoid information loss if the backend service is unavailable for a few moments, it will not be developed with REST. Will be developed using websockets and Apache Kafka will be used to send and receive data between systems;
- It is not the backend's responsibility to persist the data in Elasticsearch. For this, Kafka Connect will be used, which will also consume the simulator data and insert it into Elasticsearch;

## Techs

- Simulator: Golang
- Backend: Nest.js & MongoDB
- Frontend: React & Material UI
- Kafka & Kafka Connect
- Elasticsearch & Kibana
- Docker and Kubernetes
- Istio, Kiali, Prometheus & Grafana

</div>
