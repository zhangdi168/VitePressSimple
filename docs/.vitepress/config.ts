import { CustomerConfig } from "./custom";
import { VpSimpleConfig } from "./vpsimple";

export default {
  ...VpSimpleConfig,
  ...CustomerConfig
};