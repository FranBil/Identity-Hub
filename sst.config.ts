import { SSTConfig } from "sst";
import { PersonApiStack } from "./stacks/PersonStack";

export default {
  config(_input) {
    return {
      name: "Identity-Hub",
      region: "eu-west-1",
    };
  },
  stacks(app) {
    app.stack(PersonApiStack);
  }
} satisfies SSTConfig;
