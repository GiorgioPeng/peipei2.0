import LoginData from '@/data/loginData'
import {axiosPost} from './baseApi'
export default async function login(){
    let loginData:any = LoginData.getLoginData();
    let url:string = LoginData.destinationUrl;
    let respondData:any = axiosPost(url,loginData);
    return respondData.then(()=>console.log('chenggong')).catch((res:any)=>console.log(res));
    // status  状态码
    // token 登陆凭证
    // msg 登陆信息
}