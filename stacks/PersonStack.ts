import { Api, StackContext } from "sst/constructs";

export function PersonStack({ stack}: StackContext) {
    stack.setDefaultFunctionProps({
        runtime: "go1.x",
        timeout: "30 seconds",
        memorySize: "128 MB",
        logRetention: "one_day"  
    })

    const api = new Api(stack, "Api", {
        routes: {
            "GET /persons": "packages/functions/main.go"
        }
    });
    stack.addOutputs({
        ApiEndpoint: api.url,
      });
}