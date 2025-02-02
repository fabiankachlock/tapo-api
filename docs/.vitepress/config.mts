import { defineConfig } from "vitepress";

const advanced = [
  { text: "Overview", link: "/advanced/overview" },
  { text: "API client", link: "/advanced/api-client" },
  { text: "Custom Request", link: "/advanced/custom-request" },
  { text: "Custom Response", link: "/advanced/custom-response" },
  { text: "Custom Protocol", link: "/advanced/custom-protocol" },
];

const devices = [
  {
    text: "Supported devices",
    link: "/devices/supported-devices",
  },
  { text: "Light", link: "/devices/light" },
  { text: "Colored Light", link: "/devices/colored-light" },
  { text: "RGB Light Strip", link: "/devices/rgb-light-strip" },
  { text: "RGBIC Light Strip", link: "/devices/rgbic-light-strip" },
  { text: "Plug", link: "/devices/plug" },
  {
    text: "Energy Monitoring Plug",
    link: "/devices/energy-monitoring-plug",
  },
  { text: "Power Strip", link: "/devices/power-strip" },
  { text: "Hub", link: "/devices/hub" },
];

const childDevices = [
  { text: "Button", link: "/devices/children/button" },
  {
    text: "Motion Sensor",
    link: "/devices/children/motion-sensor",
  },
  {
    text: "Contact Sensor",
    link: "/devices/children/contact-sensor",
  },
  {
    text: "Water Leak Sensor",
    link: "/devices/children/water-leak-sensor",
  },
  {
    text: "Temperatur And Humidity Sensor",
    link: "/devices/children/temp-hum-sensor",
  },
];

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "tapo-api",
  description: "Unofficial Tapo API Client written in Go.",
  lang: "en-US",

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
      {
        text: "Devices",
        items: [
          ...devices,
          {
            text: "Child Devices",
            items: [...childDevices],
          },
        ],
      },
      {
        text: "Advanced",
        items: advanced,
      },
    ],

    sidebar: [
      {
        text: "Introduction",
        link: "/introduction",
      },
      {
        text: "Devices",
        items: [
          ...devices,
          {
            text: "Child Devices",
            items: [...childDevices],
          },
        ],
      },
      {
        text: "Advanced usage",
        items: advanced,
      },
    ],
    editLink: {
      pattern:
        "https://github.com/fabiankachlock/tapo-api/edit/main/docs/:path",
      text: "Edit this page on GitHub",
    },
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
