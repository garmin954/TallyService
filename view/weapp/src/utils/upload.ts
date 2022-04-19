import Taro from "@tarojs/taro"
import OSS from "ali-oss/dist/aliyun-oss-sdk"


export const uploadOss =async (path) => {
  const uploadTask = Taro.uploadFile({
    url: 'https://garmina.oss-cn-hongkong.aliyuncs.com', //仅为示例，非真实的接口地址
    filePath: path,
    name: 'file',
    formData:{

    },
    success: function (res){
      var data = res.data
      //do something
    }
  })
  uploadTask.progress((res) => {
    console.log('上传进度', res.progress)
    console.log('已经上传的数据长度', res.totalBytesSent)
    console.log('预期需要上传的数据总长度', res.totalBytesExpectedToSend)
  })
  uploadTask.abort() // 取消上传任务
}

