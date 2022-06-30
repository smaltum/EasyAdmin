import axios from 'axios';
import easyConst from './easy_const.js'

//******************* OVER ***********************\\

axios.defaults.timeout = 15000; //响应时间
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'; //配置请求头
axios.defaults.baseURL = easyConst.URL_BASE;

//传参序列化(添加请求拦截器)
// axios.interceptors.request.use((config) => {
//   //添加token
// //   if (config.url !== SERVER_URL.URL_LOGIN) {
// //     let token = store.getters.getToken;
// //     if (token != undefined && token != '')
// //       config.headers['Token'] = token
// //   }
//   return config;
// }, (error) => {
//   return Promise.reject(error);
// });

//返回状态判断(添加响应拦截器)
axios.interceptors.response.use((res) => {

    console.log(res)

    switch (res.status) {
        case "401":
            break
        case "500":
            break
    }
    // return Promise.resolve(res);
    return res.data;
}, (error) => {
    return Promise.reject(error);
});

//返回一个Promise(发送get请求)
export function fetchGet(url, param) {
    return new Promise((resolve, reject) => {
        axios.get(url, {
            params: param
        })
            .then(response => {
                resolve(response)
            }, err => {
                reject(err)
            })
            .catch((error) => {
                reject(error)
            })
    })
}

//返回一个Promise(发送post请求)
export function fetchPost(url, params) {

    console.log(url);

    return new Promise((resolve, reject) => {
        axios.post(url, params)
            .then(response => {
                resolve(response);
            }, err => {
                reject(err);
            })
            .catch((error) => {
                reject(error)
            })
    })
}

//返回一个Promise(上传表单请求)
export function postForm(url, forms) {
    return new Promise((resolve, reject) => {
        var configs = {
            headers: { 'Content-Type': 'multipart/form-data' }
        };
        axios.post(url, forms, configs)
            .then(response => {
                resolve(response);
            }, err => {
                reject(err);
            })
            .catch((error) => {
                reject(error)
            })
    })
}

export default {
    fetchPost,
    fetchGet,
    postForm
}