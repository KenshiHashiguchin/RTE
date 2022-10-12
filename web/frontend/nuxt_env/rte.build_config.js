module.exports = {
  head: {
    title: 'rte',
    htmlAttrs: {
      lang: 'ja',
    },
    meta: [
      {
        name: 'theme-color',
        content: '#0057A7',
      },
      {
        name: 'theme-color',
        media: '(prefers-color-scheme: light)',
        content: '#0057A7',
      },
      {
        name: 'theme-color',
        media: '(prefers-color-scheme: dark)',
        content: '#002c45',
      },
    ],
    link: [
      { rel: 'icon', href: '/favicon.ico', sizes: 'any' },
      { rel: 'icon', href: '/icon.svg', type: 'image/svg+xml' },
      {
        rel: 'apple-touch-icon',
        href: '/apple-touch-icon.png',
      },
    ],
  },
  css: [
    '@/assets/style/scss/app.scss',
  ],
  plugins: [
    'plugins/axios',
  ],
  router: {
    middleware: [
    ],
  },
}
