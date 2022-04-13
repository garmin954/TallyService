import request from "@/utils/request";

export default {
  wxLogin: (data: {
    code: string
  }) => {
    return request({ method: "POST", data, url: '/api/wxlogin' })
  },
  wxRegister: (data: {
    encryptedData: string
    iv: string
    code: string
  }) => {
    return request({ method: "POST", data, url: '/api/wxregister' })
  }
}

