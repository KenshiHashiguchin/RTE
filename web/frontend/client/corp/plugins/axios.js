import https from 'https'

export default function ({ $axios, $cookies, redirect, store, error }) {
  $axios.onRequest((config) => {
    config.httpsAgent = new https.Agent({
      rejectUnauthorized: process.env.APP_ENV === 'production',
    })
    if (typeof $cookies.get('token') !== 'undefined') {
      config.headers.common.Authorization = 'Bearer ' + $cookies.get('token')
    }
    config.headers.common.Accept = 'application/json'
    if (process.env.name === 'local') {
      console.log('Making request to ' + config.url)
    }
    // config.timeout = 1000
    return config
  })

  $axios.onError((err) => {
    const code = parseInt(err.response && err.response.status)
    console.log(`ErrorCode ${code}`, err.message)
    if (err.response.status === 401) {
      // tokenが無効になっている場合、storeをログアウト状態にしてTOPにリダレクト
      if (store.state.auth.isAuthenticated) {
        store.dispatch('auth/logout')
        return redirect('/login?token=invalid')
      }
    }
    if (err.response.status === 500 || err.response.status === 503) {
      // 500 と 503エラーの場合 エラーベージを表示
      store.dispatch('master/browser_back/reset')
      return error({
        statusCode: err.response.status,
        message: err.response.statusText,
      })
    }

    if (err.response.status === 403 || err.response.status === 404) {
      // 403は404として見せる
      return error({
        statusCode: 404,
        message: err.response.statusText,
      })
    }
  })
}
