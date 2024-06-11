package vpsimpler

const CONFIGCONTENT = "import { CustomerConfig } from \"./customer\";\nimport { VpSimpleConfig } from \"./vpsimple\";\n\nexport default {\n  ...VpSimpleConfig,\n  ...CustomerConfig// customer config优先级更大\n};"
const CUSTOMERCONTENT = "export const CustomerConfig = {\n  // head: [],\n  // markdown: {}\n};"

const CONFIGFILE = "config.ts"
