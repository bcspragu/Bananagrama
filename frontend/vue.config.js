module.exports = {
  devServer: {
    disableHostCheck: true,
    host: '0.0.0.0',
    port: 8081,
    proxy: {
      '/api/': {
          target: 'http://localhost:8080',
      },
    }
  }
}
