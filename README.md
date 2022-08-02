# Contract Engine Edge

This repository represents the Contract Engine, a Connector between KOSMoS Global and the Edge. The main task is to create Kubernetes custom resources (using the Custom Resource Definition CRD KosmosKubeEdge) based on the received active contracts from KOSMoS Global. These custom resources are then persisted into a Kubernetes cluster.   

## Quickstart
1) Build the connector via ```make``` or: 
```bash 
go build -o contractEngineEdge ./
```
2) Verify Tests are working via
```bash
go test ./...
```
3) Update Configuration File with parameters. See below [Configuration Section](#configuration)

4) Run contract Engine with 
```bash
./contractEngineEdge
```

## Under the Hood
The Contract Engine serves as a connector between KOSMoS Global and the Edge. For this purpose, the Contract Engine periodically retrieves KOSMoS Contracts from KOSMoS Global via a cron job. The following three steps are performed:
1. All contracts are obtained by sending a GET-HTTP request to
```KOSMoS_GLOBAL_HOST/contractEngine/allContractInstances``` . The response JSON looks like [mock_contract_list.json](/mock_contract_list.json)

2. Then the contracts are filtered by the attribute ```activated==1```.

3. For the activated contracts, KosmosKubeEdge (Kubernetes Custom Resource) objects are created. These are stored in the Kubernetes cluster.  




## Configuration

The Configuration of the application will be made through the file ```mock_contract_engine_parameters.toml```. Configuration for ```Keycloak,Harbor``` and ```KOSMoS_Global``` are needed. Configuration parameter for ```MQTT``` and ```database``` are optional, as they are not used as of now. The configuration parameters will be explained in the next four subsections

### Keycloak
The application uses Keycloak for authentification. Therefore it is required to setup Keycloak before. 

| parameter | default value | description |
| --------- | ----------- | -------- |
| keycloak_url | NULL | url where keycloak instance is running
|keycloak_username | edge-user | username used for authentication
| keycloak_pw | edge-pw | password used for authentication
| keycloak_client | kosmos-global | client that gets used for authentication
| keycloak_secret | NULL | secret used for authentication

### KOSMoS_Global
The Contract-Engine-Edge is dependend on a running KOSMoS_Global instance to get the current contracts. 



| parameter | default value | description |
| --------- | ----------- | -------- |
host | NULL | host Adress where KOSMoS_Global is running
user | NULL | username to log in
password | NULL | password of username
schema_pull_frequency | 120 | how often new image is pulled from harbor (in minutes)
contract_check_frequency | 10 | how often controller checks if contract is active (in minutes)
contract_publish_frequency | 30 | how often contract gets published on edge (in minutes)

### MQTT

| parameter | default value | description |
| --------- | ----------- | -------- |
host | mqtt-broker.kosmos-local | host where MQTT Broker is Running
port | 1883 | port where MQTT Broker is running
keepalive | 60 | maximum time interval that is permeitted to elapse between finishing transmitting one Packet and the point it starts sending the next.
user | NULL | username of MQTT-Client
password | NULL | password of MQTT-Client
blockchain_topic | "" | topics where blockchain messages are received
analysis_topic | "" | topics where analysis message are received


### Database

| parameter | default value | description |
| --------- | ----------- | -------- |
user | kosmos-local | username for db accesss
password | db_pass | password of username
database | kosmos_contracts | name of db
host | postgres.kosmos-local | host where db is running
port | 5432 | post where db is running

