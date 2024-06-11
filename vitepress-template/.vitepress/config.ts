import { CustomerConfig } from "./customer";
import { VpSimpleConfig } from "./vpsimple";


export default {
  ...VpSimpleConfig,
  ...CustomerConfig// customer config优先级更大
};
