import request from '../utils/request.js'

module.exports = {
  // 请求商品详情数据
  getGoodsDetail: (params) => {
    return request({
      url: '/goods/' + params.id,
      method: 'GET'
    })
  }
}