import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/__docusaurus/debug',
    component: ComponentCreator('/__docusaurus/debug', '4eb'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/config',
    component: ComponentCreator('/__docusaurus/debug/config', 'fe5'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/content',
    component: ComponentCreator('/__docusaurus/debug/content', '2ea'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/globalData',
    component: ComponentCreator('/__docusaurus/debug/globalData', 'b87'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/metadata',
    component: ComponentCreator('/__docusaurus/debug/metadata', '6dc'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/registry',
    component: ComponentCreator('/__docusaurus/debug/registry', '680'),
    exact: true
  },
  {
    path: '/__docusaurus/debug/routes',
    component: ComponentCreator('/__docusaurus/debug/routes', '79d'),
    exact: true
  },
  {
    path: '/markdown-page',
    component: ComponentCreator('/markdown-page', '83b'),
    exact: true
  },
  {
    path: '/docs',
    component: ComponentCreator('/docs', 'bae'),
    routes: [
      {
        path: '/docs',
        component: ComponentCreator('/docs', 'b49'),
        routes: [
          {
            path: '/docs',
            component: ComponentCreator('/docs', '953'),
            routes: [
              {
                path: '/docs/category/all-books',
                component: ComponentCreator('/docs/category/all-books', '206'),
                exact: true,
                sidebar: "bookSidebar"
              },
              {
                path: '/docs/intro',
                component: ComponentCreator('/docs/intro', 'a02'),
                exact: true,
                sidebar: "bookSidebar"
              }
            ]
          }
        ]
      }
    ]
  },
  {
    path: '/',
    component: ComponentCreator('/', '801'),
    exact: true
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
