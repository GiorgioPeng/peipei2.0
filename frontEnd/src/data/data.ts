import { Vue } from "vue-property-decorator";
import baseUrl from '@/api/baseUrl'

interface IStudentData {
    school:string
    name:string
    studentId:string
    furturePlan:string
    major:string
    gender:string
    province:string
    weightOfProvince:string
    sour:string
    sweet:string
    spicy:string
    weightOfTaste:string

}
export default class Data extends Vue {
    public static sudentData: IStudentData = {
        school: '',
        name: '',
        studentId: '',
        furturePlan: '',
        major: '',
        gender: '',
        province: '',
        weightOfProvince: '',
        sour: '',
        sweet: '',
        spicy: '',
        weightOfTaste: '',
    }
    public static destinationUrl:string = baseUrl+''
    public static getJsonData(): any {
        return JSON.stringify(this.sudentData);
    }
}