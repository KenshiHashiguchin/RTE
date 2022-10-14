module.exports = {
  axios: {
    debug: process.env.APP_DEBUG === 'true',
    baseURL: 'https://nginx:8443',
    browserBaseURL: process.env.APP_MERCHANT_URL,
  },
}
