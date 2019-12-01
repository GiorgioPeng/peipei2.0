import { AxiosResponse } from 'axios';
import axios from 'axios';


// get
export async function axiosGet(url: string, params: any = null): Promise<AxiosResponse<any>> {
    let response = await axios.get(url, { params });
    return response.data;
}

// post
export async function axiosPost(url: string, data: any): Promise<AxiosResponse<any>> {
    let response = await axios.post(url, data);
    return response.data;
}

// put
export async function axiosPut(url: string, data: any): Promise<AxiosResponse<any>> {
    let response = await axios.put(url, data);
    return response.data;
}
// patch
export async function axiosPatch(url: string, data: any): Promise<AxiosResponse<any>> {
    let response = await axios.patch(url, data);
    return response.data;
}

// delete
export async function axiosDelete(url: string, data: any): Promise<AxiosResponse<any>> {
    let response = await axios.delete(url, data);
    return response.data;
}