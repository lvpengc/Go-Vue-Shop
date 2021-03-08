// pages/components/step-num-add-sub/step-num-add-sub.js
Component({
  /**
   * 组件的属性列表
   */
  properties: {
    // 是否可操作，1-可，2-不可
    type: {
      type: Number,
      value: 1
    },
    // 星星分值
    score: {
      type: Number,
      value: 0
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
    changeScore: function(e) {
      console.log(this.data, e)
      if (this.data.type == 1) {
        console.log("haha")
        this.setData({
          score: parseInt(e.currentTarget.dataset.value)
        })
        console.log(this.data.score)
        this.triggerEvent("markStar", {
          num: parseInt(e.currentTarget.dataset.value),
        })
      }
    }
  }
})
