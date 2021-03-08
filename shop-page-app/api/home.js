import request from '../utils/request.js'

module.exports = {
  // 请求 banner 数据
  getBannerData: (params) => {
    return request({
      url: '/home/banner',
      method: 'GET',
      data: params
    })
  },
  // 请求推荐类目
  getCategoryData: (params) => {
    return request({
      url: '/home/category',
      method: 'GET',
      data: params
    })
  },
  // 请求优惠券
  getCouponData: (params) => {
    return request({
      url: '/home/coupon',
      method: 'GET',
      data: params
    })
  },
  // 请求推荐商品
  getGoodsData: (params) => {
    return request({
      url: '/home/goods',
      method: 'GET',
      data: params
    })
  },

}
