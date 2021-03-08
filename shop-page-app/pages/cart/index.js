// pages/cart/index.js
import {
  getCartGoods,
  updateCartGoodsStock,
  updateCartGoodsNum,
  delCartGoods,
  switchCartGoodsItem,
  switchCartGoodsAll,
  checkoutCartGoods
} from '../../api/cart.js'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    // 是否编辑
    is_edit: false,
    // 全部总额
    total_fee: 0,
    // 商品列表
    cart_goods_arr: [],
    // 是否全选
    is_select_all: false,
    // 进入编辑前被选中的id数组
    select_arr: []
  },
  // 切换数量
  goodsNumChange: function (e) {
    // 更新商品数量
    const tmp = {}
    this.setData({
      cart_goods_arr: this.data.cart_goods_arr.map(item => {
        if (item.id == e.detail.data.id) {
          item.number = e.detail.num
          Object.assign(tmp, item)
        }
        return item
      })
    })
    // 重新计算总额
    // 对当前id切换指定stock下的数量，如id不存在或stock已更改则返回错误，数量以当前为准覆盖
    updateCartGoodsNum({
      id: e.detail.data.id,
      num: e.detail.num,
      goods_stock_id: tmp.goods_stock_id
    }).then(res => {
      console.log('数量更新成功')
    }).catch(errors => {
      console.log(errors)
    })
    // 刷新列表
    this.fetchCartData()
  },
  // 商品勾选
  changeItem: function(e) {
    console.log(e.detail)
    // 更新商品勾选
    const tmp = {}
    this.setData({
      cart_goods_arr: this.data.cart_goods_arr.map(item => {
        if (e.detail.value.includes(item.id + '')) {
          console.log(123, item)
          // 如果item未选中但之前已选中，则此是操作对象
          if (!item.is_checked) {
            Object.assign(tmp, item)
          }
          item.is_checked = true
        } else {
          console.log(456, item)
          // 如果item已选中但之前未选中，则此是操作对象
          if (item.is_checked) {
            Object.assign(tmp, item)
          }
          item.is_checked = false
        }
        return item
      })
    })
    // 判断是否全选
    this.checkIfSelectAll()
    if (!this.data.is_edit) {
      // 重新计算总额
      // 切换指定id的checked状态，若id不存在或stock已更改则返回错误
      switchCartGoodsItem({
        id: tmp.id,
        goods_stock_id: tmp.goods_stock_id,
        is_checked: !tmp.is_checked
      }).then(res => {
        console.log('item状态更新成功')
      }).catch(errors => {
        console.log(errors)
      })
      // 刷新列表
      this.fetchCartData()
    }
  },
  // 全选
  checkedAll: function(e) {
    // 全选或取消全选
    if (this.data.is_select_all) {
      this.setData({
        cart_goods_arr: this.data.cart_goods_arr.map(item => {
          item.is_checked = false
          return item
        })
      })
    } else {
      this.setData({
        cart_goods_arr: this.data.cart_goods_arr.map(item => {
          item.is_checked = true
          return item
        })
      })
    }
    // 判断是否全选
    this.checkIfSelectAll()
    if (!this.data.is_edit) {
      // 重新计算总额
      // 切换所有的checked状态，若id不存在或stock已更改则直接忽略
      switchCartGoodsAll({
        is_checked: this.data.is_select_all
      }).then(res => {
        console.log('所有状态更新成功')
      }).catch(errors => {
        console.log(errors)
      })
      // 刷新列表
      this.fetchCartData()
    }
  },
  // 获取页面数据
  fetchData() {
    this.fetchCartData()
  },
  // 获取购物车数据
  fetchCartData: function() {
    getCartGoods().then(res => {
      this.setData({
        cart_goods_arr: res.data.data.goods_list,
        total_fee: res.data.data.total_fee
      })
      // 判断是否全选
      this.checkIfSelectAll()
    }).catch(errors => {
      console.log(errors)
    })
  },
  // 判断是否全选 
  checkIfSelectAll: function() {
    // 任意未选中，则flag为false
    let flag = true
    this.data.cart_goods_arr.forEach(item => {
      if (item.is_checked == false) {
        flag = false
      }
    })
    // 更新是否全选
    this.setData({
      is_select_all: flag && (this.data.cart_goods_arr.length > 0)
    })
  }
  ,
  // 编辑模式
  editCart: function() {
    if (this.data.is_edit) { // 取消或完成编辑
      // 再恢复之前的选中状态
      this.setData({
        cart_goods_arr: this.data.cart_goods_arr.map(item => {
          if (this.data.select_arr.includes(item.id)) {
            item.is_checked = true
          }
          return item
        }),
        select_arr: [],
        is_edit: false
      })
      this.checkIfSelectAll()
      // 刷新列表
      this.fetchCartData()
    } else { // 进入编辑模式
      let tmp = []
      // 记录当前被选中的商品
      this.data.cart_goods_arr.forEach(item => {
        if (item.is_checked == true) {
          tmp.push(item.id)
        }
      })
      // 取消所有checkbox的选中
      this.setData({
        cart_goods_arr: this.data.cart_goods_arr.map(item => {
          item.is_checked = false
          return item
        }),
        is_select_all: false,
        select_arr: tmp,
        is_edit: true
      })
    }
  },
  // 删除所选商品
  deleteCart: function() {
    let tmp = []
    // 获取所有选中要删除的商品
    this.data.cart_goods_arr.forEach(item => {
      if (item.is_checked == true) {
        tmp.push(item.id)
      }
    })
    // 未选中则提示
    if (tmp.length == 0) {
      wx.showToast({
        title: '您还没有选择商品哦！',
        icon: 'none',
      })
    } else {
      let _that = this
      wx.showModal({
        title: '操作确认',
        content: '确认删除所选商品？',
        confirmColor: '#09bb07',
        success(res) {
          if (res.confirm) {
            delCartGoods({
              id: tmp
            }).then(res => {
              // 实时更新页面数据，去掉已删除的
              _that.setData({
                cart_goods_arr: _that.data.cart_goods_arr.filter(item => {
                  return !tmp.includes(item.id)
                })
              })
              wx.showToast({
                title: '操作成功',
              })
            }).catch(errors => {
              console.log(errors)
            })
          }
        }
      })
    }
  },
  checkoutOrder: function() {
    let tmp = []
    // 获取所有选中要删除的商品
    this.data.cart_goods_arr.forEach(item => {
      if (item.is_checked == true) {
        tmp.push(item.id)
      }
    })
    // 未选中则提示
    if (tmp.length <= 0) {
      wx.showToast({
        title: '您还没有选择商品哦！',
        icon: 'none',
      })
    } else {
      checkoutCartGoods().then(res => {
        // 订单创建成功需跳转页面
        console.log('订单创建成功需跳转页面')
      }).catch(errors => {
        console.log(errors)
      })
    }
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.fetchData()
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})