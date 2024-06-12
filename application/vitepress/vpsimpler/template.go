package vpsimpler

const ConfigContent = "import { CustomerConfig } from \"./customer\";\nimport { VpSimpleConfig } from \"./vpsimple\";\n\nexport default {\n  ...VpSimpleConfig,\n  ...CustomerConfig// customer config优先级更大\n};"
const CustomerContent = "export const CustomerConfig = {\n  // head: [],\n  // markdown: {}\n};"

// ConfigFile1 旧版本的配置文件路径
const ConfigFile1 = ".vitepress/config.mts"
const ConfigFile1Rename = ".vitepress/config_bak.mts"

const ConfigFile = ".vitepress/config/index.ts"
const VpsimpleFile = ".vitepress/config/vpsimple.ts"
const CustomerFile = ".vitepress/config/customer.ts"
