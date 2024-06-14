import mdItCustomAttrs from "markdown-it-custom-attrs";

export const CustomConfig = {
  head: [
    [
      "link",
      {
        rel: "stylesheet",
        href: "/plugins/bigimg/fancybox.css"//大图css
      }
    ],
    [
      "script",
      {
        src: "/plugins/bigimg/fancybox.umd.js"//大图js
      }
    ]
  ],
  markdown: {
    config: (md) => {
      // use more markdown-it plugins!
      md.use(mdItCustomAttrs, "image", {
        "data-fancybox": "gallery"
      });
    }
  }
};