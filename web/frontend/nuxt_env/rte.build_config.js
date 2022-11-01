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
      {
        rel: 'stylesheet',
        href:
          'https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons'
      }
    ],
  },
  css: [
    // 'assets/style/theme.styl',
    'assets/style/scss/app.scss',
    'font-awesome/css/font-awesome.css',
    'roboto-fontface/css/roboto/roboto-fontface.css'
  ],
  plugins: [
    'plugins/vuetify',
    'plugins/axios',
    'plugins/vee-validate',
    { src: 'plugins/vue-perfect-scrollbar', ssr: false },
  ],
  router: {
    middleware: [
    ],
  },
}
