import { CustomerConfig } from "./customer";
import { VpSimpleConfig } from "./vpsimple";


export default {
  ...VpSimpleConfig,//来自软件vpsimple自动写入的配置
  ...CustomerConfig// customer config优先级更大
};
