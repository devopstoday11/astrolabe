// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Astrolabe data protection framework API",
    "title": "Astrolabe API",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/astrolabe": {
      "get": {
        "description": "This returns the list of services that this Astrolabe server supports\n",
        "produces": [
          "application/json"
        ],
        "summary": "List available services",
        "operationId": "listServices",
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ServiceList"
            }
          }
        }
      }
    },
    "/astrolabe/tasks": {
      "get": {
        "description": "Lists running and recent tasks",
        "produces": [
          "application/json"
        ],
        "operationId": "listTasks",
        "responses": {
          "200": {
            "description": "List of recent task IDs",
            "schema": {
              "$ref": "#/definitions/TaskIDList"
            }
          }
        }
      }
    },
    "/astrolabe/tasks/nexus": {
      "get": {
        "description": "Provides a list of current task nexus",
        "produces": [
          "application/json"
        ],
        "operationId": "listTaskNexus",
        "responses": {
          "200": {
            "description": "Task nexus list",
            "schema": {
              "$ref": "#/definitions/TaskNexusList"
            }
          }
        }
      },
      "post": {
        "description": "Creates a new nexus for monitoring task completion",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "New task nexus",
            "schema": {
              "$ref": "#/definitions/TaskNexusID"
            }
          }
        }
      }
    },
    "/astrolabe/tasks/nexus/{taskNexusID}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "The nexus to wait on",
            "name": "taskNexusID",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "Time to wait (milliseconds) before returning if no tasks   complete",
            "name": "waitTime",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "description": "Last finished time seen by this client.  Tasks that have completed after this time tick will be returned, or if no tasks\nhave finished, the call will hang until waitTime has passed or a task finishes.  Starting time tick should\nbe the finished time of the last task that the caller saw completed on this nexus.  Use 0 to get all finished\ntasks (tasks that have finished and timed out of the server will not be shown)\n",
            "name": "lastFinishedNS",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/TaskNexusResponse"
            }
          }
        }
      }
    },
    "/astrolabe/tasks/{taskID}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Gets info about a running or recently completed task",
        "operationId": "getTaskInfo",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the task to retrieve info for",
            "name": "taskID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Info for running or recently completed task",
            "schema": {
              "$ref": "#/definitions/TaskInfo"
            }
          }
        }
      }
    },
    "/astrolabe/{service}": {
      "get": {
        "description": "List protected entities for the service.  Results will be returned in\ncanonical ID order (string sorted).  Fewer results may be returned than\nexpected, the ProtectedEntityList has a field specifying if the list has\nbeen truncated.\n",
        "produces": [
          "application/json"
        ],
        "operationId": "listProtectedEntities",
        "parameters": [
          {
            "type": "string",
            "description": "The service to list protected entities from",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "The maximum number of results to return (fewer results may be returned)",
            "name": "maxResults",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Results will be returned that come after this ID",
            "name": "idsAfter",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityList"
            }
          },
          "404": {
            "description": "Service or Protected Entity not found"
          }
        }
      },
      "post": {
        "description": "Copy a protected entity into the repository.  There is no option to\nembed data on this path, for a self-contained or partially\nself-contained object, use the restore from zip file option in the S3\nAPI REST API\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "copyProtectedEntity",
        "parameters": [
          {
            "type": "string",
            "description": "The service to copy the protected entity into",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "create",
              "create_new",
              "update"
            ],
            "type": "string",
            "description": "How to handle the copy.  create - a new protected entity with the\nProtected Entity ID will be created.  If the Protected Entity ID\nalready exists, the copy will fail.  create_new - A Protected Entity\nwith a new ID will be created with data and metadata from the source\nprotected entity.  Update - If a protected entity with the same ID\nexists it will be overwritten.  If there is no PE with that ID, one\nwill be created with the same ID. For complex Persistent Entities,\nthe mode will be applied to all of the component entities that are\npart of this operation as well.\n",
            "name": "mode",
            "in": "query",
            "required": true
          },
          {
            "description": "Info of ProtectedEntity to copy",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProtectedEntityInfo"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Create in progress",
            "schema": {
              "$ref": "#/definitions/CreateInProgressResponse"
            }
          }
        }
      }
    },
    "/astrolabe/{service}/{protectedEntityID}": {
      "get": {
        "description": "Get the info for a Protected Entity including name, data access and\ncomponents\n",
        "produces": [
          "application/json"
        ],
        "operationId": "getProtectedEntityInfo",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to retrieve info for",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityInfo"
            }
          }
        }
      },
      "delete": {
        "description": "Deletes a protected entity or snapshot of a protected entity (if the\nsnapshot ID is specified)\n",
        "produces": [
          "application/json"
        ],
        "operationId": "deleteProtectedEntity",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to retrieve info for",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityID"
            }
          }
        }
      }
    },
    "/astrolabe/{service}/{protectedEntityID}/snapshots": {
      "get": {
        "description": "Gets the list of snapshots for this protected entity\n",
        "produces": [
          "application/json"
        ],
        "operationId": "listSnapshots",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to retrieve info for",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "List succeeded",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityList"
            }
          },
          "404": {
            "description": "Service or Protected Entity not found"
          }
        }
      },
      "post": {
        "description": "Creates a new snapshot for this protected entity\n",
        "produces": [
          "application/json"
        ],
        "operationId": "createSnapshot",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to snapshot",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          },
          {
            "description": "Parameters for the snapshot.",
            "name": "params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/SnapshotParamList"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Snapshot created successfully, returns the new snapshot ID",
            "schema": {
              "$ref": "#/definitions/ProtectedEntitySnapshotID"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ComponentSpec": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "$ref": "#/definitions/ProtectedEntityID"
        },
        "server": {
          "type": "string"
        }
      }
    },
    "CreateInProgressResponse": {
      "type": "object",
      "properties": {
        "taskID": {
          "$ref": "#/definitions/TaskID"
        }
      }
    },
    "DataTransport": {
      "type": "object",
      "properties": {
        "params": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "transportType": {
          "type": "string"
        }
      }
    },
    "ProtectedEntityID": {
      "type": "string"
    },
    "ProtectedEntityInfo": {
      "type": "object",
      "required": [
        "id",
        "name",
        "metadataTransports",
        "dataTransports",
        "combinedTransports",
        "componentSpecs"
      ],
      "properties": {
        "combinedTransports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DataTransport"
          }
        },
        "componentSpecs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ComponentSpec"
          }
        },
        "dataTransports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DataTransport"
          }
        },
        "id": {
          "$ref": "#/definitions/ProtectedEntityID"
        },
        "metadataTransports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DataTransport"
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "ProtectedEntityList": {
      "type": "object",
      "properties": {
        "list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ProtectedEntityID"
          }
        },
        "truncated": {
          "type": "boolean"
        }
      }
    },
    "ProtectedEntitySnapshotID": {
      "type": "string"
    },
    "ServiceList": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "SnapshotPEParamItem": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "object"
        }
      }
    },
    "SnapshotPEParamList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/SnapshotPEParamItem"
      }
    },
    "SnapshotParamItem": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "$ref": "#/definitions/SnapshotPEParamList"
        }
      }
    },
    "SnapshotParamList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/SnapshotParamItem"
      }
    },
    "TaskID": {
      "type": "string"
    },
    "TaskIDList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/TaskID"
      }
    },
    "TaskInfo": {
      "type": "object",
      "required": [
        "id",
        "completed",
        "status",
        "startedTime",
        "startedTimeNS",
        "progress"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "details": {
          "type": "string"
        },
        "finishedTime": {
          "type": "string"
        },
        "finishedTimeNS": {
          "description": "Finished time in nanoseconds",
          "type": "integer"
        },
        "id": {
          "$ref": "#/definitions/TaskID"
        },
        "progress": {
          "type": "number",
          "maximum": 100
        },
        "result": {
          "type": "object"
        },
        "startedTime": {
          "type": "string"
        },
        "startedTimeNS": {
          "description": "Start time in nanoseconds",
          "type": "integer"
        },
        "status": {
          "type": "string",
          "enum": [
            "running",
            "success",
            "failed",
            "cancelled"
          ]
        }
      }
    },
    "TaskNexusID": {
      "type": "string"
    },
    "TaskNexusInfo": {
      "type": "object",
      "properties": {
        "associatedTasks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TaskID"
          }
        },
        "id": {
          "$ref": "#/definitions/TaskNexusID"
        }
      }
    },
    "TaskNexusList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/TaskNexusInfo"
      }
    },
    "TaskNexusResponse": {
      "type": "object",
      "properties": {
        "finished": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TaskInfo"
          }
        },
        "id": {
          "$ref": "#/definitions/TaskNexusID"
        }
      }
    }
  },
  "x-components": {}
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Astrolabe data protection framework API",
    "title": "Astrolabe API",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/astrolabe": {
      "get": {
        "description": "This returns the list of services that this Astrolabe server supports\n",
        "produces": [
          "application/json"
        ],
        "summary": "List available services",
        "operationId": "listServices",
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ServiceList"
            }
          }
        }
      }
    },
    "/astrolabe/tasks": {
      "get": {
        "description": "Lists running and recent tasks",
        "produces": [
          "application/json"
        ],
        "operationId": "listTasks",
        "responses": {
          "200": {
            "description": "List of recent task IDs",
            "schema": {
              "$ref": "#/definitions/TaskIDList"
            }
          }
        }
      }
    },
    "/astrolabe/tasks/nexus": {
      "get": {
        "description": "Provides a list of current task nexus",
        "produces": [
          "application/json"
        ],
        "operationId": "listTaskNexus",
        "responses": {
          "200": {
            "description": "Task nexus list",
            "schema": {
              "$ref": "#/definitions/TaskNexusList"
            }
          }
        }
      },
      "post": {
        "description": "Creates a new nexus for monitoring task completion",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "New task nexus",
            "schema": {
              "$ref": "#/definitions/TaskNexusID"
            }
          }
        }
      }
    },
    "/astrolabe/tasks/nexus/{taskNexusID}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "The nexus to wait on",
            "name": "taskNexusID",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "Time to wait (milliseconds) before returning if no tasks   complete",
            "name": "waitTime",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "description": "Last finished time seen by this client.  Tasks that have completed after this time tick will be returned, or if no tasks\nhave finished, the call will hang until waitTime has passed or a task finishes.  Starting time tick should\nbe the finished time of the last task that the caller saw completed on this nexus.  Use 0 to get all finished\ntasks (tasks that have finished and timed out of the server will not be shown)\n",
            "name": "lastFinishedNS",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/TaskNexusResponse"
            }
          }
        }
      }
    },
    "/astrolabe/tasks/{taskID}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Gets info about a running or recently completed task",
        "operationId": "getTaskInfo",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the task to retrieve info for",
            "name": "taskID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Info for running or recently completed task",
            "schema": {
              "$ref": "#/definitions/TaskInfo"
            }
          }
        }
      }
    },
    "/astrolabe/{service}": {
      "get": {
        "description": "List protected entities for the service.  Results will be returned in\ncanonical ID order (string sorted).  Fewer results may be returned than\nexpected, the ProtectedEntityList has a field specifying if the list has\nbeen truncated.\n",
        "produces": [
          "application/json"
        ],
        "operationId": "listProtectedEntities",
        "parameters": [
          {
            "type": "string",
            "description": "The service to list protected entities from",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "The maximum number of results to return (fewer results may be returned)",
            "name": "maxResults",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Results will be returned that come after this ID",
            "name": "idsAfter",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityList"
            }
          },
          "404": {
            "description": "Service or Protected Entity not found"
          }
        }
      },
      "post": {
        "description": "Copy a protected entity into the repository.  There is no option to\nembed data on this path, for a self-contained or partially\nself-contained object, use the restore from zip file option in the S3\nAPI REST API\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "copyProtectedEntity",
        "parameters": [
          {
            "type": "string",
            "description": "The service to copy the protected entity into",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "create",
              "create_new",
              "update"
            ],
            "type": "string",
            "description": "How to handle the copy.  create - a new protected entity with the\nProtected Entity ID will be created.  If the Protected Entity ID\nalready exists, the copy will fail.  create_new - A Protected Entity\nwith a new ID will be created with data and metadata from the source\nprotected entity.  Update - If a protected entity with the same ID\nexists it will be overwritten.  If there is no PE with that ID, one\nwill be created with the same ID. For complex Persistent Entities,\nthe mode will be applied to all of the component entities that are\npart of this operation as well.\n",
            "name": "mode",
            "in": "query",
            "required": true
          },
          {
            "description": "Info of ProtectedEntity to copy",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProtectedEntityInfo"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Create in progress",
            "schema": {
              "$ref": "#/definitions/CreateInProgressResponse"
            }
          }
        }
      }
    },
    "/astrolabe/{service}/{protectedEntityID}": {
      "get": {
        "description": "Get the info for a Protected Entity including name, data access and\ncomponents\n",
        "produces": [
          "application/json"
        ],
        "operationId": "getProtectedEntityInfo",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to retrieve info for",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityInfo"
            }
          }
        }
      },
      "delete": {
        "description": "Deletes a protected entity or snapshot of a protected entity (if the\nsnapshot ID is specified)\n",
        "produces": [
          "application/json"
        ],
        "operationId": "deleteProtectedEntity",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to retrieve info for",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityID"
            }
          }
        }
      }
    },
    "/astrolabe/{service}/{protectedEntityID}/snapshots": {
      "get": {
        "description": "Gets the list of snapshots for this protected entity\n",
        "produces": [
          "application/json"
        ],
        "operationId": "listSnapshots",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to retrieve info for",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "List succeeded",
            "schema": {
              "$ref": "#/definitions/ProtectedEntityList"
            }
          },
          "404": {
            "description": "Service or Protected Entity not found"
          }
        }
      },
      "post": {
        "description": "Creates a new snapshot for this protected entity\n",
        "produces": [
          "application/json"
        ],
        "operationId": "createSnapshot",
        "parameters": [
          {
            "type": "string",
            "description": "The service for the protected entity",
            "name": "service",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The protected entity ID to snapshot",
            "name": "protectedEntityID",
            "in": "path",
            "required": true
          },
          {
            "description": "Parameters for the snapshot.",
            "name": "params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/SnapshotParamList"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Snapshot created successfully, returns the new snapshot ID",
            "schema": {
              "$ref": "#/definitions/ProtectedEntitySnapshotID"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ComponentSpec": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "$ref": "#/definitions/ProtectedEntityID"
        },
        "server": {
          "type": "string"
        }
      }
    },
    "CreateInProgressResponse": {
      "type": "object",
      "properties": {
        "taskID": {
          "$ref": "#/definitions/TaskID"
        }
      }
    },
    "DataTransport": {
      "type": "object",
      "properties": {
        "params": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "transportType": {
          "type": "string"
        }
      }
    },
    "ProtectedEntityID": {
      "type": "string"
    },
    "ProtectedEntityInfo": {
      "type": "object",
      "required": [
        "id",
        "name",
        "metadataTransports",
        "dataTransports",
        "combinedTransports",
        "componentSpecs"
      ],
      "properties": {
        "combinedTransports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DataTransport"
          }
        },
        "componentSpecs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ComponentSpec"
          }
        },
        "dataTransports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DataTransport"
          }
        },
        "id": {
          "$ref": "#/definitions/ProtectedEntityID"
        },
        "metadataTransports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/DataTransport"
          }
        },
        "name": {
          "type": "string"
        }
      }
    },
    "ProtectedEntityList": {
      "type": "object",
      "properties": {
        "list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ProtectedEntityID"
          }
        },
        "truncated": {
          "type": "boolean"
        }
      }
    },
    "ProtectedEntitySnapshotID": {
      "type": "string"
    },
    "ServiceList": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "SnapshotPEParamItem": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "object"
        }
      }
    },
    "SnapshotPEParamList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/SnapshotPEParamItem"
      }
    },
    "SnapshotParamItem": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "$ref": "#/definitions/SnapshotPEParamList"
        }
      }
    },
    "SnapshotParamList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/SnapshotParamItem"
      }
    },
    "TaskID": {
      "type": "string"
    },
    "TaskIDList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/TaskID"
      }
    },
    "TaskInfo": {
      "type": "object",
      "required": [
        "id",
        "completed",
        "status",
        "startedTime",
        "startedTimeNS",
        "progress"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "details": {
          "type": "string"
        },
        "finishedTime": {
          "type": "string"
        },
        "finishedTimeNS": {
          "description": "Finished time in nanoseconds",
          "type": "integer"
        },
        "id": {
          "$ref": "#/definitions/TaskID"
        },
        "progress": {
          "type": "number",
          "maximum": 100,
          "minimum": 0
        },
        "result": {
          "type": "object"
        },
        "startedTime": {
          "type": "string"
        },
        "startedTimeNS": {
          "description": "Start time in nanoseconds",
          "type": "integer"
        },
        "status": {
          "type": "string",
          "enum": [
            "running",
            "success",
            "failed",
            "cancelled"
          ]
        }
      }
    },
    "TaskNexusID": {
      "type": "string"
    },
    "TaskNexusInfo": {
      "type": "object",
      "properties": {
        "associatedTasks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TaskID"
          }
        },
        "id": {
          "$ref": "#/definitions/TaskNexusID"
        }
      }
    },
    "TaskNexusList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/TaskNexusInfo"
      }
    },
    "TaskNexusResponse": {
      "type": "object",
      "properties": {
        "finished": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TaskInfo"
          }
        },
        "id": {
          "$ref": "#/definitions/TaskNexusID"
        }
      }
    }
  },
  "x-components": {}
}`))
}
