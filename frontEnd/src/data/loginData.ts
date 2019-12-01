import { Vue } from 'vue-property-decorator'
import baseUrl from '@/api/baseUrl'
interface ILoginData{
    id:string;
    name:string;
    school:string
}
export default class LoginData extends Vue {
    public static loginData:ILoginData = {
        id:'',
        name:'',
        school:''
    }
    public static destinationUrl:string = baseUrl+'/students/?action=login'
    public static getLoginData():any{
        return JSON.stringify(this.loginData);
    }
}