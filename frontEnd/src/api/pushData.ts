import Data from '@/data/data'
import { axiosPut } from './baseApi'
export default async function pushData(){
    let jsonData = Data.getJsonData();
    let url = Data.destinationUrl;
    let respondData = await axiosPut(url,jsonData)
    return respondData;
}