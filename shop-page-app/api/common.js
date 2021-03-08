import request from '../utils/request.js'

module.exports = {
  // 请求常见问题
  getCommonQuestion: () => {
    return request({
      url: '/common/question',
      method: 'GET'
    })
  }
}