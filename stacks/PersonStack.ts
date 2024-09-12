import { Api, EventBus, Function, StackContext, Table } from "sst/constructs";

export function PersonApiStack({ stack }: StackContext) {
    const personsTable = new Table(stack, "PersonsTable", {
        cdk: {
            table: {
                tableName: "PersonsTable",
            },
        },
        fields: {
            firstName: "string",
            lastName: "string",
            phoneNumber: "string",
            address: "string",
        },
        primaryIndex: { partitionKey: "lastName", sortKey: "phoneNumber" },
    });

    const bus = new EventBus(stack, "PersonEventBus", {
        rules: {
          personCreatedRule: {
            pattern: {
              source: ["com.example.identity_hub"],
              detailType: ["PersonCreated"],
            },
          },
        },
      });
      bus.attachPermissions(["lambda"])

    const listPersonFunction = new Function(stack, "GetAllPersons", {
        handler: "packages/lambda/list-persons/main.go",
        description: "Function for list Persons",
        runtime: "go1.x",
        environment: {
            TABLE_NAME: "PersonsTable"
        },
        bind: [personsTable]
    });

    const createPersonFunction = new Function(stack, "CreatePerson", {
        handler: "packages/lambda/create-person/main.go",
        description: "Function for create Person",
        runtime: "go1.x",
        environment: {
            TABLE_NAME: "PersonsTable"
        },
        bind: [personsTable, bus]
    });

    const api = new Api(stack, "Api", {
        routes: {
            "GET /v1/persons": listPersonFunction,
            "POST /v1/persons": createPersonFunction,
        }
    });

    stack.addOutputs({
        ApiEndpoint: api.url,
      });
}