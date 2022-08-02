package main

import (
	"context"
	"fmt"

	kosmosv1 "github.com/kosmos-industrie40/setting-up-kosmos-edge/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func createObject(c client.Writer, id int, activated_contracts kosmosv1.KosmosKubeEdgeSpec) error {

	// Using a typed object.
	contract := &kosmosv1.KosmosKubeEdge{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      fmt.Sprintf("kosmos-contract-%d", id), //%d, contract_ID ergänzen
		},
		TypeMeta: metav1.TypeMeta{
			Kind: "kosmosv1.KosmosKubeEdge",
		},
		Spec: kosmosv1.KosmosKubeEdgeSpec{
			Body: kosmosv1.BodyProperties{
				Machine: activated_contracts.Body.Machine,
				// Sensors:                     getSensors(activated_contracts.Body.Sensors),
				Sensors: activated_contracts.Body.Sensors,
				// RequiredTechnicalContainers: getRequiredTechnicalContainers(activated_contracts.Body.RequiredTechnicalContainers),
				RequiredTechnicalContainers: activated_contracts.Body.RequiredTechnicalContainers,
				Analysis: kosmosv1.AnalysisBody{
					Enable: activated_contracts.Body.Analysis.Enable,
					// Systems: getSystemsAnalysis(activated_contracts.Body.Analysis.Systems),
					Systems: activated_contracts.Body.Analysis.Systems,
				},
				BlockchainConnection: kosmosv1.BlockchainConnectionBody{
					Uri: activated_contracts.Body.BlockchainConnection.Uri,
					// ContainerList: getContainers(activated_contracts.Body.BlockchainConnection.ContainerList),
					ContainerList: activated_contracts.Body.BlockchainConnection.ContainerList,
					Sensors:       activated_contracts.Body.BlockchainConnection.Sensors,
				},
				// MachineConnection:  getMachineConnection(activated_contracts.Body.MachineConnection),
				MachineConnection:  activated_contracts.Body.MachineConnection,
				KosmosLocalSystems: activated_contracts.Body.KosmosLocalSystems,
				CheckSignature:     activated_contracts.Body.CheckSignature,
				Metadata:           activated_contracts.Body.Metadata,
				Contract: kosmosv1.ContractBody{
					Valid: kosmosv1.ValidContract{
						Start: activated_contracts.Body.Contract.Valid.Start,
						End:   activated_contracts.Body.Contract.Valid.End,
					},
					CreationTime:   activated_contracts.Body.Contract.CreationTime,
					Id:             activated_contracts.Body.Contract.Id,
					ParentContract: activated_contracts.Body.Contract.ParentContract,
					Version:        activated_contracts.Body.Contract.Version,
					// Permissions:    getPermissionContracts(activated_contracts.Body.Contract.Permissions),
					Permissions: activated_contracts.Body.Contract.Permissions,
					Partners:    activated_contracts.Body.Contract.Partners,
				},
			},
			Signature: kosmosv1.SignatureProperties{
				Meta: kosmosv1.MetaSignature{
					Algorithm: activated_contracts.Signature.Meta.Algorithm,
					Date:      activated_contracts.Signature.Meta.Date,
				},
				Signature: activated_contracts.Signature.Signature,
			},
		},
	}
	// c is a created client.
	err := c.Create(context.Background(), contract)

	if err != nil {
		return err
	}

	return nil
}

// func getPermissionContracts(permContract kosmosv1.PermissionsContract) kosmosv1.PermissionsContract {
// 	var permContractsRead = []string{}
// 	var permContractsWrite = []string{}
// 	var permContracts = kosmosv1.PermissionsContract{}

// 	for _, read := range permContract.Read {
// 		permContractsRead = append(permContractsRead, read)
// 	}

// 	for _, write := range permContract.Write {
// 		permContractsWrite = append(permContractsRead, write)
// 	}

// 	permContracts = kosmosv1.PermissionsContract{
// 		Read:  permContractsRead,
// 		Write: permContractsWrite,
// 	}

// 	return permContracts
// }

// func getRequiredTechnicalContainers(reqTechCon []kosmosv1.RequiredTechnicalContainersBody) []kosmosv1.RequiredTechnicalContainersBody {
// 	var reqTechConS = []kosmosv1.RequiredTechnicalContainersBody{}
// 	for _, reqTechConElements := range reqTechCon {
// 		reqTechConS = append(reqTechConS, kosmosv1.RequiredTechnicalContainersBody{
// 			System:     reqTechConElements.System,
// 			Containers: getContainers(reqTechConElements.Containers),
// 		})
// 	}
// 	return reqTechConS
// }

// func getSystemsAnalysis(sysAna []kosmosv1.SystemsAnalysis) []kosmosv1.SystemsAnalysis {
// 	var sysAnas = []kosmosv1.SystemsAnalysis{}
// 	for _, sysAnaElements := range sysAna {
// 		sysAnas = append(sysAnas, kosmosv1.SystemsAnalysis{
// 			System: sysAnaElements.System,
// 			Enable: sysAnaElements.Enable,
// 			Connection: kosmosv1.ConnectionSystems{
// 				Url:       sysAnaElements.Connection.Url,
// 				UserMgmt:  sysAnaElements.Connection.UserMgmt,
// 				Interval:  sysAnaElements.Connection.Interval,
// 				Container: getContainer(sysAnaElements.Connection.Container),
// 			},
// 			Pipelines: getPipelineDefinitions(sysAnaElements.Pipelines),
// 		})
// 	}
// 	return sysAnas
// }

// func getPipelineDefinitions(pipeline []kosmosv1.PipelinesDefinitions) []kosmosv1.PipelinesDefinitions {
// 	var pipelines = []kosmosv1.PipelinesDefinitions{}
// 	for _, pipelineElements := range pipeline {
// 		pipelines = append(pipelines, kosmosv1.PipelinesDefinitions{
// 			ML_Trigger: kosmosv1.ML_TriggerPipelines{
// 				Type:       pipelineElements.ML_Trigger.Type,
// 				Definition: pipelineElements.ML_Trigger.Definition,
// 			},
// 			Sensors:  pipelineElements.Sensors,
// 			Pipeline: getPipelines(pipelineElements.Pipeline),
// 		})
// 	}
// 	return pipelines
// }

// func getPipelines(pipeline []kosmosv1.PipelinePipelines) []kosmosv1.PipelinePipelines {
// 	var pipelines = []kosmosv1.PipelinePipelines{}
// 	for _, pipelineElements := range pipeline {
// 		// pipelines = append(pipelines, kosmosv1.PipelinePipelines{
// 		// 	Container:     getContainer(pipelineElements.Container),
// 		// 	PersistOutput: pipelineElements.PersistOutput,
// 		// })
// 		pipelineA := kosmosv1.PipelinePipelines{
// 			Container:     getContainer(pipelineElements.Container),
// 			PersistOutput: pipelineElements.PersistOutput,
// 		}
// 		if pipelineElements.From != nil {
// 			pipelineA.From = &kosmosv1.ModelDefinitions{
// 				Url: pipelineElements.From.Url,
// 				Tag: pipelineElements.From.Tag,
// 			}
// 		}
// 		if pipelineElements.To != nil {
// 			pipelineA.To = &kosmosv1.ModelDefinitions{
// 				Url: pipelineElements.To.Url,
// 				Tag: pipelineElements.To.Tag,
// 			}
// 		}

// 		pipelines = append(pipelines, pipelineA)
// 	}
// 	return pipelines
// }

// func getMachineConnection(machCon []kosmosv1.MachineConnectionBody) []kosmosv1.MachineConnectionBody {
// 	var machCons = []kosmosv1.MachineConnectionBody{}
// 	for _, machConElements := range machCon {
// 		machCons = append(machCons, kosmosv1.MachineConnectionBody{
// 			Connector: getContainer(machConElements.Connector),
// 			ConnectionData: kosmosv1.ConnectionData{
// 				Uri:         machConElements.ConnectionData.Uri,
// 				Credentials: machConElements.ConnectionData.Credentials,
// 				MachineID:   machConElements.ConnectionData.MachineID,
// 			},
// 			SubscribeData: machConElements.SubscribeData,
// 		})
// 	}
// 	return machCons
// }

// func getContainers(containers []kosmosv1.ContainerDefinitions) []kosmosv1.ContainerDefinitions {
// 	var cons = []kosmosv1.ContainerDefinitions{}
// 	for _, containersElements := range containers {
// 		cons = append(cons, getContainer(containersElements))
// 	}
// 	return cons
// }

// func getContainer(containers kosmosv1.ContainerDefinitions) kosmosv1.ContainerDefinitions {
// 	var cons = kosmosv1.ContainerDefinitions{}
// 	cons = kosmosv1.ContainerDefinitions{
// 		Url:         containers.Url,
// 		Tag:         containers.Tag,
// 		Arguments:   containers.Arguments,
// 		Environment: containers.Environment,
// 		Ports:       getPorts(containers.Ports),
// 	}
// 	return cons
// }

// func getPorts(port []kosmosv1.PortsContainer) []kosmosv1.PortsContainer {
// 	var ports = []kosmosv1.PortsContainer{}
// 	for _, portElements := range port {
// 		ports = append(ports, kosmosv1.PortsContainer{
// 			Dst:   portElements.Dst,
// 			Src:   portElements.Src,
// 			Label: portElements.Label,
// 		})
// 	}
// 	return ports
// }

// func getSensors(sensor []kosmosv1.SensorsBody) []kosmosv1.SensorsBody {
// 	var sensors = []kosmosv1.SensorsBody{}
// 	for _, sensorElements := range sensor {
// 		sensors = append(sensors, kosmosv1.SensorsBody{
// 			Name:            sensorElements.Name,
// 			StorageDuration: getStorageDurationSensors(sensorElements.StorageDuration),
// 			Meta:            sensorElements.Meta,
// 		})
// 	}
// 	return sensor
// }

// func getStorageDurationSensors(stoDurSen []kosmosv1.StorageDurationSensors) []kosmosv1.StorageDurationSensors {
// 	var stoDurSens = []kosmosv1.StorageDurationSensors{}
// 	for _, stoDurSenElements := range stoDurSen {
// 		stoDurSens = append(stoDurSens, kosmosv1.StorageDurationSensors{
// 			SystemName: stoDurSenElements.SystemName,
// 			Duration:   stoDurSenElements.Duration,
// 		})
// 	}
// 	return stoDurSens
// }
