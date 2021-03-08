// pages/category/index.js
import { 
  getCategoryData,
  getGoodsData
} from '../../api/category.js'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    onFresh: false,
    //窗口高度
    winHeight: "",
    scrollLeft: 0,
    category_arr: [],
    active_tab: undefined,
    goods_arr: [],
    active_category_icon_url: '',
    active_category_name: '',
    active_category_desc: '',
  },
  /**
   * 拉取分类列表
   */
  async fetchCategoryData() {
    await getCategoryData({}).then(res => {
      this.setData({
        category_arr: res.data.data
      })
    }).catch(errors => {
      console.log(errors)
    })
  },
  fetchGoodsData() {
    // 只加载尚未加载的分类商品数据
    if (this.data.goods_arr.length == 0 || !this.data.goods_arr[this.data.active_tab] || this.data.goods_arr[this.data.active_tab].length == 0) {
      this.setData({
        onFresh: true
      })
      getGoodsData({
        category_id: this.data.active_tab
      }).then(res => {
        let tmp = this.data.goods_arr
        tmp[this.data.active_tab] = res.data.data
        this.setData({
          goods_arr: tmp,
          onFresh: false
        })
      }).catch(errors => {
        console.log(errors)
      })

    }
  },
  initCategoryInfo: function() {
    for (const index in this.data.category_arr) {
      if (this.data.category_arr[index].id == this.data.active_tab) {
        this.setData({
          active_category_icon_url: this.data.category_arr[index].icon_url,
          active_category_name: this.data.category_arr[index].name,
          active_category_desc: this.data.category_arr[index].desc,
        })
        break
      }
    }
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: async function (options) {
    await this.fetchCategoryData()
    if (options.id) {
      this.setData({
        active_tab: options.id,
      })
      this.initCategoryInfo()
      this.fetchGoodsData()
    }
    //  高度自适应
    wx.getSystemInfo({
      success: (res) => {
        const clientHeight = res.windowHeight,
          clientWidth = res.windowWidth,
          rpxR = 750 / clientWidth;
        this.setData({
          winHeight: clientHeight * rpxR - 180
        });
      }
    });
  },
  // 点击分类切换
  switchTab: function (e) {
    const cur = e.currentTarget.dataset.index
    if (cur != this.data.active_tab) {
      this.setData({
        active_tab: cur,
      })
      this.initCategoryInfo()
      this.fetchGoodsData()
    }
  },
  //判断当前滚动超过一屏时，设置tab标题滚动条。
  checkTab: function () {
    if (this.data.active_tab > 3) {
      this.setData({
        scrollLeft: 300
      })
    } else {
      this.setData({
        scrollLeft: 0
      })
    }
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
    this.data.goods_arr[this.data.active_tab] = []
    this.fetchGoodsData()
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