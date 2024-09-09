import { Api, Function, StackContext, Table } from "sst/constructs";

export function PersonApiStack({ stack }: StackContext) {
    const personsTable = new Table(stack, "PersonsTable", {
        cdk: {
            table: {
                tableName: "PersonsTable",
            },
        },
        fields: {
            id: "string",
            firstName: "string",
            lastName: "string",
            phoneNumber: "string",
            address: "string",
        },
        primaryIndex: { partitionKey: "id" },
    });

    const listPersonFunction = new Function(stack, "LetPersons", {
        handler: "packages/functions/list-persons/main.main",
        runtime: "go1.x",
        environment: {
            TABLE_NAME: "PersonTable"
        },
    });

    const createPersonFunction = new Function(stack, "CreatePerson", {
        handler: "packages/functions/create-person/main.main",
        runtime: "go1.x",
        environment: {
            TABLE_NAME: "PersonTable"
        },
    });

    // personsTable.grantRead(listPersonFunction);
    // personsTable.grantReadWrite(createPersonFunction);

    const api = new Api(stack, "Api", {
        routes: {
            "GET /persons": listPersonFunction,
            "POST /persons": createPersonFunction,
        }
    });
    stack.addOutputs({
        ApiEndpoint: api.url,
      });
}