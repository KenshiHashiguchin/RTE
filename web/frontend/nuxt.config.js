require('dotenv').config()
const environment = process.env.APP_ENV || 'local'
const appType = process.env.APP_TYPE || 'user'
console.log('appType')
console.log(appType)
console.log(environment)
const envSet = require(`./nuxt_env/env.${appType}_${environment}.js`)
const buildConfSet = require(`./nuxt_env/rte.build_config.js`)
const VuetifyLoaderPlugin = require('vuetify-loader/lib/plugin')

export default {
  env: {
    name: environment,
    userAppUrl: process.env.APP_USER_URL,
    merchantAppUrl: process.env.APP_MERCHANT_URL,
  },
  srcDir: `client/${appType}/`,

  render: {
    compressor: (req, res, next) => {
      next()
    },
  },

  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: buildConfSet.head || {},

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: buildConfSet.css || [],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: buildConfSet.plugins || [],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    // '@nuxtjs/eslint-module',
    // https://go.nuxtjs.dev/stylelint
    // '@nuxtjs/stylelint-module',
    '@nuxtjs/vuetify',
  ],
  vuetify: {
    customVariables: ['~/assets/style/scss/variables.scss'],
    treeShake: true
  },
  moment: {
    locales: ['ja'],
  },

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
    '@nuxtjs/axios',
    [
      'cookie-universal-nuxt',
      {
        parseJSON: false,
      },
    ],
  ],
  loading: {color: '#3adced'},
  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
    transpile: [
      'vee-validate/dist/rules'
    ],
    extractCSS: true,
    devtools: process.env.APP_DEBUG === 'true',
    html: {
      minify: {
        collapseBooleanAttributes: true,
        decodeEntities: true,
        minifyCSS: true,
        minifyJS: true,
        processConditionalComments: true,
        removeEmptyAttributes: true,
        removeRedundantAttributes: true,
        trimCustomFragments: true,
        useShortDoctype: true,
        removeComments: true,
        preserveLineBreaks: false,
        collapseWhitespace: true,
      },
    },
    // transpile: ['vuetify/lib'],
    // plugins: [new VuetifyLoaderPlugin()],
    // loaders: {
    //   stylus: {
    //     import: ["~/assets/style/variables.styl"]
    //   }
    // },
    // build.templates local設定
    // templates: [
    //   {
    //     src:
    //       environment === 'local'
    //         ? `client/${appType}/local/app.html`
    //         : `client/${appType}/_app.html`,
    //     dst: 'views/app.template.html',
    //   },
    // ],
  },
  axios: envSet.axios || {},
  router: buildConfSet.router || {},
  watchers: {
    webpack: {
      poll: process.env.WATCH_POLL,
    },
  },
}
