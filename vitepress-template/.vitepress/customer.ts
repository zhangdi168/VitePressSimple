import mdItCustomAttrs from "markdown-it-custom-attrs";

export const CustomerConfig = {
  head: [
    [
      "link",
      {
        rel: "stylesheet",
        href: "/bigimg/fancybox.css"//大图css
      }
    ],
    [
      "script",
      {
        src: "/bigimg/fancybox.umd.js"//大图js
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