import Data from '@/data/data'
import { axiosPut } from './baseApi'
export default async function pushData(){
    let jsonData:any = Data.getJsonData();
    let url:string = Data.destinationUrl;
    let respondData:any = await axiosPut(url,jsonData)
    return respondData;
}