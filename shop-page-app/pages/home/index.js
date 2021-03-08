import { 
  getBannerData,
  getCategoryData,
  getCouponData,
  getGoodsData
} from '../../api/home.js'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    // banner 广告
    banner_arr: [],
    // 推荐类目
    category_arr: [],
    // 今日优惠券
    coupon_arr: [],
    // 推荐商品
    goods_arr: []
  },
  /**
   * 拉取数据
   */
  fetchData: function() {
    this.fetchBannerData()
    this.fetchCategoryData()
    this.fetchCouponData()
    this.fetchGoodsData()
  },
  /**
   * 拉取 banner 数据
   */
  fetchBannerData() {
    getBannerData({}).then(res => {
      this.setData({
        banner_arr: res.data.data
      })
    }).catch(errors => {
      console.log(errors)
    })
  },
  /**
   * 拉取推荐分类
   */
  fetchCategoryData() {
    getCategoryData({}).then(res => {
      this.setData({
        category_arr: res.data.data
      })
    }).catch(errors => {
      console.log(errors)
    })
  },
  /**
   * 拉取优惠券
   */
  fetchCouponData() {
    getCouponData({}).then(res => {
      this.setData({
        coupon_arr: res.data.data
      })
    }).catch(errors => {
      console.log(errors)
    })
  },
  /**
   * 拉取推荐商品
   */
  fetchGoodsData() {
    getGoodsData({}).then(res => {
      this.setData({
        goods_arr: res.data.data
      })
    }).catch(errors => {
      console.log(errors)
    })
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