import Taro from '@tarojs/taro';
import interceptors from './interceptors';

// 请求连接前缀
export const baseUrl = "http://192.168.31.180:8080";
// export const baseUrl = "http://192.168.1.21:8080";

// 输出日志信息
export const noConsole = false;

interceptors.forEach((interceptorItem) => Taro.addInterceptor(interceptorItem));

interface OptionsType {
  method: 'GET' | 'POST' | 'PUT';
  data: any;
  url: string;
  noLoading?: boolean;
}
export default (options: OptionsType = { method: 'GET', data: {}, url: '', noLoading: false }) => {
  if (!options.noLoading) {
    Taro.showLoading({
      title: '加载中'
    });
  }
  if (!noConsole) {
    console.log(`%c【请求】%c ${new Date().toLocaleString()}【 URL=${options.url} 】PARA：`,  `color:#077fa2`,``);
    console.log( options.data )
    console.log(`%c <<<`, `border-right:5px solid #077fa2;padding-right:5px; color:#077fa2`);
  }
  for (const key in options.data) {
    if (options.data.hasOwnProperty(key) && (options.data[key] === undefined || options.data[key] == null)) {
      delete options.data[key];
    }
  }
  return Taro.request({
    url: baseUrl + options.url,
    data: {
      ...options.data
    },
    header: {
      'Content-Type': 'application/json',
      "Authorization": `Bearer ${Taro.getStorageSync("TOKEN")}`
    },
    method: options.method
  }).then((res) => {
    // 拦截token
    if(res?.data?.token && res?.data?.expired){
      Taro.setStorageSync("TOKEN", res.data.token)
      Taro.setStorageSync("TOKEN_EXPIRED", res.data.expired)
    }

    setTimeout(() => {
      Taro.hideLoading();
    }, 100);
    if (!noConsole) {
      console.log(`%c【返回】%c ${new Date().toLocaleString('zh', { hour12: false })}【${options.url} 】 响应数据：`, `color:#78ad5c`,``);
      console.log( res, )
      console.log(`%c <<<`, `border-right:5px solid #78ad5c;padding-right:5px; color:#78ad5c`);
    }

    return Promise.resolve(res)
  });
};
