// index.js
const defaultAvatarUrl = 'https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0'

Page({
  data: {
    loginCode:"",
    phoneCode:""
  },
  getPhoneNumber(e){
      wx.login({
        success: res => {
          this.setData({
            phoneCode:e.detail.code,
            loginCode:res.code
          })
        }
      })
 
  },
})
