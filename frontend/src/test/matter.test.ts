import { test } from "vitest";
import matter from "gray-matter";

test("matter", () => {
  console.log(
    matter(
      "---\n" +
        "{\n" +
        '    "layout": "home",\n' +
        '    "hero": {\n' +
        '        "name": "zh-vitepress simple template",\n' +
        '        "text": "zh-A VitePress Site",\n' +
        '        "tagline": "zh-My great project tagline",\n' +
        '        "actions": [\n' +
        "            {\n" +
        '                "text": "zh-Markdown Examples",\n' +
        '                "link": "/markdown-examples"\n' +
        "            },\n" +
        "            {\n" +
        '                "text": "zh-API Examples",\n' +
        '                "link": "/api-examples"\n' +
        "            },\n" +
        "            {\n" +
        '                "text": "zh-blog",\n' +
        '                "link": "http://xiaod.co"\n' +
        "            }\n" +
        "        ],\n" +
        '        "image": {\n' +
        '            "src": "",\n' +
        '            "alt": "",\n' +
        '            "width": "",\n' +
        '            "height": ""\n' +
        "        },\n" +
        '        "description": "",\n' +
        '        "features": [],\n' +
        '        "head": []\n' +
        "    },\n" +
        '    "features": [\n' +
        "        {\n" +
        '            "title": "zh-FeatureBBBBB",\n' +
        '            "details": "Lorem ipsum dolor sit amet, consectetur adipiscing elit"\n' +
        "        },\n" +
        "        {\n" +
        '            "title": "zh-Feature B",\n' +
        '            "details": "Lorem ipsum dolor sit amet, consectetur adipiscing elit"\n' +
        "        },\n" +
        "        {\n" +
        '            "title": "zh-Feature C",\n' +
        '            "details": "Lorem ipsum dolor sit amet, consectetur adipiscing elit"\n' +
        "        }\n" +
        "    ],\n" +
        '    "title": "index",\n' +
        '    "description": "",\n' +
        '    "navbar": true,\n' +
        '    "sideBar": true,\n' +
        '    "footer": false,\n' +
        '    "outline": [\n' +
        "        1,\n" +
        "        3\n" +
        "    ],\n" +
        '    "editLink": false,\n' +
        '    "lastUpdated": true,\n' +
        '    "aside": "right",\n' +
        '    "custom": {}\n' +
        "}\n" +
        "---",
    ),
    " 1",
  );
  console.log(
    matter(
      "---\n" +
        "{\n" +
        '    "layout": "home",\n' +
        '    "hero": {\n' +
        '        "name": "en-vitepress simple template",\n' +
        '        "text": "en-A VitePress Site",\n' +
        '        "tagline": "zh-My great project tagline",\n' +
        '        "actions": [\n' +
        "            {\n" +
        '                "text": "zh-Markdown Examples",\n' +
        '                "link": "/markdown-examples"\n' +
        "            },\n" +
        "            {\n" +
        '                "text": "zh-API Examples",\n' +
        '                "link": "/api-examples"\n' +
        "            },\n" +
        "            {\n" +
        '                "text": "zh-blog",\n' +
        '                "link": "http://xiaod.co"\n' +
        "            }\n" +
        "        ],\n" +
        '        "image": {\n' +
        '            "src": "",\n' +
        '            "alt": "",\n' +
        '            "width": "",\n' +
        '            "height": ""\n' +
        "        },\n" +
        '        "description": "",\n' +
        '        "features": [],\n' +
        '        "head": []\n' +
        "    },\n" +
        '    "features": [\n' +
        "        {\n" +
        '            "title": "zh-FeatureBBBBB",\n' +
        '            "details": "Lorem ipsum dolor sit amet, consectetur adipiscing elit"\n' +
        "        },\n" +
        "        {\n" +
        '            "title": "zh-Feature B",\n' +
        '            "details": "Lorem ipsum dolor sit amet, consectetur adipiscing elit"\n' +
        "        },\n" +
        "        {\n" +
        '            "title": "zh-Feature C",\n' +
        '            "details": "Lorem ipsum dolor sit amet, consectetur adipiscing elit"\n' +
        "        }\n" +
        "    ],\n" +
        '    "title": "index",\n' +
        '    "description": "",\n' +
        '    "navbar": true,\n' +
        '    "sideBar": true,\n' +
        '    "footer": false,\n' +
        '    "outline": [\n' +
        "        1,\n" +
        "        3\n" +
        "    ],\n" +
        '    "editLink": false,\n' +
        '    "lastUpdated": true,\n' +
        '    "aside": "right",\n' +
        '    "custom": {}\n' +
        "}\n" +
        "---",
    ),
    "2",
  );
});
