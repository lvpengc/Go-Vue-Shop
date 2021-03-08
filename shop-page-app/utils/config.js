// 环境变量
const ENV = 'DEV'
// 环境信息
const ENV_DATA = {
  DEV: {
    BASE_URL: 'http://localhost:8080/' 
  },
  STAGE: {
    BASE_URL: 'http://localhost:8080'
  },
  PROD: {
    BASE_URL: 'http://localhost:8080'
  }
}
module.exports = {
  ENV,
  ENV_DATA: ENV_DATA[ENV]
}