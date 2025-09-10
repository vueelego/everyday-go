import axios from "axios";

export type BackupParam = {
  api: string
  query?: string
  body?: string
  response?: string
}

export function saveResponse(res: any) {
  let data: BackupParam = {
    api: (res.config.baseURL || '') + res.config.url,
    query: res.config.params ? JSON.stringify(res.config.params) : '',
    body: res.config.params ? JSON.stringify(res.config.data) : '',
    response: typeof res.data == "object" ? JSON.stringify(res.data) : res.data ?? ''
  }

  axios.post("http://localhost:6001/v1/save", data, {
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(res => {
    console.log("save success " + data.api, data.query, data.query, data.response)
  })
}

export function changeConfig(config: any) {
  let param: BackupParam = {
    api: config.baseURL + config.url,
    // query: config.params,
    // body: config.data ? JSON.stringify(config.data) : "",
  }

  config.data = param

  config.baseURL = null
  config.url = "http://localhost:6001/v1/get"
  config.method = "post"
  return config
}

export function getResponse(data: BackupParam) {
  axios.post("http://localhost:6001/v1/get", data, {
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(res => {
    console.log("get res", res)
  })
}