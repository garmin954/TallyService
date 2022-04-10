import { useGlobalIconFont } from './utils/iconfont/helper';
export default {
  usingComponents: Object.assign(useGlobalIconFont()),
  pages: [
    "pages/index/index",
    "pages/login/index",
    "pages/mine/index",
    "pages/statistics/index",
    "pages/ledger/index"

  ],
  window: {
    backgroundTextStyle: "light",
    navigationBarBackgroundColor: "#fff",
    navigationBarTitleText: "WeChat",
    navigationBarTextStyle: "black",
  },
  tabBar: {
    custom: false,
    color: '#333333',
    selectedColor: '#667eea',
    backgroundColor: '#fff',
    borderStyle: 'white',
    list:[
      {
        pagePath: "pages/index/index",
        text: "明细"
      },
      {
        pagePath: "pages/statistics/index",
        text: "统计"
      },
      {
        pagePath: "pages/mine/index",
        text: "我的"
      }
    ]
  },

};
