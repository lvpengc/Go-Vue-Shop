import { ENV_DATA, ENV } from './config.js'
import { trim } from './functions.js'

/**
 * 封装 request
 */
export default ({
  url,
  data = {},
  method = 'GET'
}) => {
  return new Promise((resolve, reject) => {
    wx.request({
      url: trim(ENV_DATA.BASE_URL, '/') + '/' + trim(url, '/'),
      data,
      method,
      header: {
        'Content-Type': 'application/json',
        'X-Token': wx.getStorageSync('token')
      },
      success: res => {
        resolve(res);
      },
      fail: err => {
        reject(err)
      }
    })
  })
}