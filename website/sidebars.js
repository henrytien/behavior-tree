/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  mainSidebar: [
    'intro',
    'getting-started',
    'concepts',
    'design-notes',
    {
      type: 'category',
      label: '节点参考 / Node Reference',
      items: [
        'nodes/composites',
        'nodes/decorators',
        'nodes/actions',
        'nodes/conditions',
        'nodes/subtree',
      ],
    },
    'examples',
    'faq',
  ],
};

module.exports = sidebars;
