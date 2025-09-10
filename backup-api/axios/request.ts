import { useUserStore } from '@/stores/modules/user'
import axios from 'axios'
import { lc } from './storage'
import { changeConfig } from "./backup"

const instance = axios.create({
  timeout: 40000,
  baseURL: import.meta.env.VITE_GATEWAY,
})

instance.interceptors.request.use(
  config => {
    const userStore = useUserStore()
    const headers: any = {
      'sys-code': import.meta.env.VITE_SYS_CODE,
    }
    if (userStore.access_token) {
      headers.Authorization = 'Bearer ' + lc.get('access_token')
    }
    config.headers = { ...headers, ...config.headers }

    return changeConfig(config)
    // return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  res => {
    const userStore = useUserStore()
    if (userStore.access_token && new Date().getTime() > userStore.token_expired - 1000 * 60 * 30) {
      userStore.refreshToken()
    }

    // 存
    // saveResponse(res)
    // 取
    res.data = JSON.parse(res.data.data.response)

    if (res.status === 200 && res.request.responseType === 'blob') {
      //文件流
      return res
    }

    if (res.status === 200 && (res.data.success || res.data.StatusCode == 200)) {
      return res.data
    } else {
      if (res.data?.message) {
        window.$message.error(res.data.message, {
          duration: 5000,
        })
      }
      return Promise.reject(res.data)
    }
  },
  async error => {
    const userStore = useUserStore()

    if (error.response) {
      switch (error.response.status) {
        case 401:
          await userStore.loginOut()
          break
        case 404:
          window.$message.error('网络错误', {
            duration: 5000,
          })
          break

        default:
          window.$message.error(error.response.data.message, {
            duration: 5000,
          })
          break
      }
    }
    return Promise.reject(error)
  }
)

export default instance
