package vpsimpler

const INDEXTSCONTENT = "import { CustomerConfig } from \"./customer\";\nimport { VpSimpleConfig } from \"./vpsimple\";\nimport { defineConfig } from \"vitepress\";\n\nexport default defineConfig({\n  ...VpSimpleConfig,\n  ...CustomerConfig// customer config优先级更大\n});"
const CUSTOMERCONTENT = "export const CustomerConfig = {\n  // head: [],\n  // markdown: {}\n};"

const CONFIGFILE = "config.ts"
