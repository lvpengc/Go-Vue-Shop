import request from '../utils/request.js'

module.exports = {
  // 请求购物车数据
  getCartGoods: (params) => {
    return request({
      url: '/cart',
      method: 'GET',
      data: params
    })
  },
  // 添加商品到购物车
  addCartGoods: (params) => {
    return request({
      url: '/cart',
      method: 'POST',
      data: params
    })
  },
  // 更新购物车商品规格的选择
  updateCartGoodsStock: (params) => {
    return request({
      url: '/cart/' + params.id + '/stock',
      method: 'PUT',
      data: params
    })
  },
  // 更新购物车商品数量的选择
  updateCartGoodsNum: (params) => {
    return request({
      url: '/cart/' + params.id + '/num',
      method: 'PUT',
      data: params
    })
  },
  // 删除指定商品
  delCartGoods: (params) => {
    return request({
      url: '/cart/',
      method: 'DELETE',
      data: params
    })
  },
  // 选择或取消指定id的商品
  switchCartGoodsItem: (params) => {
    return request({
      url: '/cart/' + params.id + '/switch-checked',
      method: 'PUT',
      data: params
    })
  },
  // 选择或取消指定所有的商品
  switchCartGoodsAll: (params) => {
    return request({
      url: '/cart/switch-checked',
      method: 'PUT',
      data: params
    })
  },
  // 结算选中商品
  checkoutCartGoods: (params) => {
    return request({
      url: '/cart/checkout',
      method: 'POST',
      data: params
    })
  }
}
