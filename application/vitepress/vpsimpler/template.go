package vpsimpler

const ConfigContent = "import { CustomConfig } from \"./custom\";\nimport { VpSimpleConfig } from \"./vpsimple\";\n\nexport default {\n  ...VpSimpleConfig,\n  ...CustomConfig// custom config优先级更大\n};"
const CustomContent = "export const CustomConfig = {\n  // head: [],\n  // markdown: {}\n};"

// ConfigFile1 旧版本的配置文件路径
const ConfigFile1 = ".vitepress/config.mts"
const ConfigFile1Rename = ".vitepress/config_bak.mts"

const ConfigFile = ".vitepress/config/index.ts"
const VpsimpleFile = ".vitepress/config/vpsimple.ts"
const CustomerFile = ".vitepress/config/custom.ts"
