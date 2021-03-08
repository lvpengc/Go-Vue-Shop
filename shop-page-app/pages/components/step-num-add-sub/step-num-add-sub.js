// pages/components/step-num-add-sub/step-num-add-sub.js
Component({
  /**
   * 组件的属性列表
   */
  properties: {
    number: {
      type: Number,
      value: 1
    },
    data: {
      type: Object,
      value: {}
    }
  },

  /**
   * 组件的初始数据
   */
  data: {},

  /**
   * 组件的方法列表
   */
  methods: {
    numSub: function () {
      let num = this.data.number
      if (num > 1) {
        num--
        this.setData({
          number: num
        })
        this.triggerEvent("numberChange", {
          num: num,
          data: this.data.data
        })
      }
    },
    numAdd: function () {
      let num = this.data.number
      num++
      this.setData({
        number: num
      })
      this.triggerEvent("numberChange", {
        num: num,
        data: this.data.data
      })
    },
    numChange: function (e) {
      let num = e.detail.value
      this.setData({
        number: num
      })
      this.triggerEvent("numberChange", {
        num: num,
        data: this.data.data
      })
    }
  }
})
