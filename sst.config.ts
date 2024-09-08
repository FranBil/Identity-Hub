import { SSTConfig } from "sst";
import { Api } from "sst/constructs";

export default {
  config(_input) {
    return {
      name: "Identity-Hub",
      region: "us-east-1",
    };
  },
  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go1.x",
    });
    app.stack(function Stack({ stack }) {
      const api = new Api(stack, "api", {
        routes: {
          "GET /": "packages/functions/main.go",
        },
      });
      stack.addOutputs({
        ApiEndpoint: api.url,
      });
    });
  },
} satisfies SSTConfig;
