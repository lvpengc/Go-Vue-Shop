import request from '../utils/request.js'

module.exports = {
  // 请求推荐类目
  getCategoryData: (params) => {
    return request({
      url: '/category',
      method: 'GET',
      data: params
    })
  },
  // 请求推荐类目
  getGoodsData: (params) => {
    return request({
      url: '/category/' + params.category_id,
      method: 'GET'
    })
  },

}
