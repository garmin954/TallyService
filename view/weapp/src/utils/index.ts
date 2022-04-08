import Taro from "@tarojs/taro";

/**
* @description 设置导航栏信息
*/
export const getNavBarInfo = () => {
  // 获取系统信息
  const {statusBarHeight=0, screenWidth} = Taro.getSystemInfoSync();
  // 胶囊按钮位置信息
  const menuButtonInfo = Taro.getMenuButtonBoundingClientRect();
  // 导航栏高度 = 状态栏到胶囊的间距（胶囊距上距离-状态栏高度） * 2 + 胶囊高度 + 状态栏高度
  const navBarHeight = (menuButtonInfo.top - statusBarHeight) * 2 + menuButtonInfo.height + statusBarHeight;
  const menuBotton = menuButtonInfo.top - statusBarHeight;
  const menuRight = screenWidth - menuButtonInfo.right;
  const menuHeight = menuButtonInfo.height;

  return {
    navBarHeight,
    menuBotton,
    menuRight,
    menuHeight
  }
}
