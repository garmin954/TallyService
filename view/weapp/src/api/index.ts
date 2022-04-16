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
  },

  fetchLedgerUser: (data: {}) => {
    return request({ method: "GET", data, url: '/api/ledger/user' })
  },

  fetchClassify: (data: {}) => {
    return request({ method: "GET", data, url: '/api/classify' })
  },
}

