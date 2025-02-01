import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "tapo-api",
  description: "Unofficial Tapo API Client written in Go.",
  lang: "en",

  head: [
    [
      "link",
      {
        rel: "icon",
        type: "image/png",
        href: "/static/gopher.png",
      },
    ],
    [
      "script",
      {},
      `
      // Add a listener that avoids the Vitepress Router if there is a hijack attribute
      window.addEventListener('click', (e) => {
          const link = e.target.closest('a');
          if (!link) return;

          if (link.getAttribute('hijack') == 'false') {
            e.stopImmediatePropagation();
            window.location = link.getAttribute('href')
            return true
          }
        },
        { capture: true }
      )
      `,
    ],
  ],

  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    logo: "/static/gopher.png",
    nav: [
      { text: "Home", link: "/" },
      { text: "Examples", link: "/markdown-examples" },
    ],

    sidebar: [
      {
        text: "Examples",
        items: [
          { text: "Markdown Examples", link: "/markdown-examples" },
          { text: "Runtime API Examples", link: "/api-examples" },
        ],
      },
    ],
    socialLinks: [
      { icon: "github", link: "https://github.com/fabiankachlock/tapo-api" },
    ],
    search: {
      provider: "local",
    },
    footer: {
      copyright: "Copyright © 2024-present Fabian Kachlock",
      message:
        "Made with <3 using <a href='https://vitepress.dev/' hijack='false' >vitepress</a>",
    },
  },
});
