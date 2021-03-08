// pages/goods/detail/index.js
import {
  getGoodsDetail
} from '../../../api/goods.js'
import {
  getCommonQuestion
} from '../../../api/common.js'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    // 商品数据
    goods_data: {},
    // 是否已选择规格
    is_choose: false,
    // 选择的规格文案
    choose_stock: '',
    // 最近一条评论
    recent_comment: undefined,
    // 常见问题
    common_question: []
  },
  fetchGoodsDetail: function(params) {
    getGoodsDetail(params).then(res => {
      this.setData({
        goods_data: res.data.data,
        // 取最近一条评论用于默认展示
        recent_comment: res.data.data.comment[0]
      })
    }).catch(errors => {
      console.log(errors)
    })
  },
  fetchCommonQuestion() {
    getCommonQuestion().then(res => {
      this.setData({
        common_question: res.data.data
      })
    }).catch(errors => {
      console.log(errors)
    })
  },
  //图片点击事件
  preview: function(event){
    var src = event.currentTarget.dataset.src;//获取data-src
    var imgList = event.currentTarget.dataset.list;//获取data-list
    //图片预览
    wx.previewImage({
      current: src, // 当前显示图片的http链接
      urls: imgList // 需要预览的图片http链接列表
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    if (options.id) {
      this.fetchGoodsDetail({
        id: options.id
      })
    }
    this.fetchCommonQuestion()
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