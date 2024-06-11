import { CustomerConfig } from "./customer";
import { VpSimpleConfig } from "./vpsimple";

export default {
  ...VpSimpleConfig,
  ...CustomerConfig
};