// @ts-check
const {themes} = require('prism-react-renderer');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'behavior3go',
  tagline: 'Golang 行为树 / Behavior Tree for Go',
  favicon: 'img/favicon.ico',

  url: 'https://henrytien.github.io',
  baseUrl: '/behavior-tree/',

  organizationName: 'henrytien',
  projectName: 'behavior-tree',

  onBrokenLinks: 'throw',

  trailingSlash: false,

  markdown: {
    hooks: {
      onBrokenMarkdownLinks: 'warn',
    },
  },

  headTags: [
    {
      tagName: 'meta',
      attributes: {
        name: 'keywords',
        content: 'behavior tree, behavior3, golang, go, AI, game AI, 行为树, behavior3go',
      },
    },
    {
      tagName: 'meta',
      attributes: {
        name: 'author',
        content: 'henrytien',
      },
    },
  ],

  i18n: {
    defaultLocale: 'zh-Hans',
    locales: ['zh-Hans', 'en'],
    localeConfigs: {
      'zh-Hans': {label: '中文'},
      en: {label: 'English'},
    },
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: './sidebars.js',
          editUrl: 'https://github.com/henrytien/behavior-tree/edit/master/website/',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
        sitemap: {
          changefreq: 'weekly',
          priority: 0.5,
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      metadata: [
        {name: 'description', content: 'behavior3go — a Go implementation of the behavior3 behavior tree, compatible with the official online editor.'},
      ],
      navbar: {
        title: 'behavior3go',
        items: [
          {
            type: 'docSidebar',
            sidebarId: 'mainSidebar',
            position: 'left',
            label: '文档 / Docs',
          },
          {
            type: 'localeDropdown',
            position: 'right',
          },
          {
            href: 'https://github.com/henrytien/behavior-tree',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {label: '简介 / Intro', to: '/docs/intro'},
              {label: '节点参考 / Nodes', to: '/docs/nodes/composites'},
            ],
          },
          {
            title: 'More',
            items: [
              {label: 'GitHub', href: 'https://github.com/henrytien/behavior-tree'},
              {label: '在线编辑器 / Editor', href: 'https://henrytien.github.io/behavior-tree-editor/'},
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} henrytien. Built with Docusaurus.`,
      },
      prism: {
        theme: themes.github,
        darkTheme: themes.dracula,
        additionalLanguages: ['go', 'json', 'bash'],
      },
    }),
};

module.exports = config;
