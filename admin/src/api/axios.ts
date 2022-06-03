import axios from "axios";
import router from '../router';
import { ElMessage } from "element-plus";
import qs from 'qs'
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 5000
})

service.interceptors.response.use(
    
    response =>{
        debugger
        return response
    },error => {
        debugger
        return Promise.reject(new Error(error.response.data))
    }
)

service.interceptors.request.use(
    config => {
        return config
    },
    error =>{
        return Promise.reject(new Error(error))
    }
)

interface BZResponse<U> {
    /** 是否成功状态 */
    success: boolean;
    /** 响应数据 */
    obj: U;
    /** 错误提示 */
    errorMsg?: string;
    /** 错误提示 */
    code?: number;
}

function checkCode(res: any){
    if (res.status === -404) {
        ElMessage.error('未请求到数据');
    }
    if (res.data && !res.data.success) {
        ElMessage.error('未请求到数据');
    }
    return res;
}
//公用报错提示
function checkStatus(response: any){
    if ( response && (/^(200|304|400)$/.test(response.status)) ) {
        // 未登录跳转
        if(response.data.msg && response.data.msg === 'token 认证失败'){
            router.push('/managelogin'); //未登录
            return false;
        }else if (response.data.code == "200") {
            return response.data;
            ElMessage.error(response.data.msg);
        }else {// 请求出错提示
            ElMessage.error("请求出错");
            return false;
        }
    }else{
        return false;
    }
    
}

export default {
    service,
    getJson<U>(data: any, url: string) {
        let token: any = "Bearer "+window.localStorage.getItem("token")
        return axios({
            method: "get",
            url: url,
            data: qs.stringify(data),
            headers: {
                "Authorization": token,
                'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
            }
        })
        .then((res: any) => {
            // 正常返回数据
            return checkStatus(res);
        })
        .catch(err => {
            return checkCode(err);
        });
    },
    postJson<U>(data: any, url: string) {
        let token: any = "Bearer "+window.localStorage.getItem("token")
        return axios({
            method: "post",
            url: url,
            data: qs.stringify(data),
            headers: {
                "Authorization": token,
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        })
        .then((res: any) => {
            // 正常返回数据
            return checkStatus(res);
        })
        .catch(err => {
            return checkCode(err);
        });
    },
    deleteJson<U>(data: any, url: string) {
        let token: any = "Bearer "+window.localStorage.getItem("token")
        return axios({
            method: "DELETE",
            url: url,
            data: qs.stringify(data),
            headers: {
                "Authorization": token,
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        })
        .then((res: any) => {
            // 正常返回数据
            return checkStatus(res);
        })
        .catch(err => {
            return checkCode(err);
        });
    }
}    
