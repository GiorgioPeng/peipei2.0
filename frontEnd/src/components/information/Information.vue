<template>
  <div class="informationFrame">
    <div class="personalData">
      <div class="personalDataFrame">
        <input class="school" :value='school'/>
        <div>
          <input class="name" :value="name"/>
          <input class="studentId" :value="id"/>
        </div>
        <div class="futurePlan">未来规划</div>
        <div class="major">专业</div>
        <div class="gender">
          <div class="male"></div>
          <div class="female"></div>
        </div>
      </div>
    </div>
    <div class="positionAndFood"></div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
@Component({
  name: "Information"
})
export default class Information extends Vue {
  private school:string='学校'
  private name:string = '姓名'
  private id:string = '学号'
  private futurePlan:string = '未来规划'
  private major:string = '专业'
}
</script>
<style lang="stylus">
input{
  outline:none;
  border:none;
  text-align:center;
  color:#707070
}
.informationFrame {
  display: flex;
  margin-top: 8vw;
  margin-bottom: 4vw;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  width: 84vw;
  height: 700px;
}

.personalData {
  background-image: url('../../../public/information/personalData.png');
  background-size: 100% 100%;
  height:35%;
  width: 85vw;

  .personalDataFrame {
    height: 80%;
    transform: translate(-20px, 30px);
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    justify-content:space-between
    .school {
      background-image: url('../../../public/information/school.png');
      background-size: 100% 100%;
      width: 50%;
      height: 16%;
    }

    .futurePlan {
      background-image: url('../../../public/information/futurePlan.png');
      background-size: 100% 100%;
      width: 50%;
      height: 16%;
    }

    div {
      width: 50%;
      height: 16%;
      display: flex;
      flex-direction: row;
      justify-content:space-between
      .name {
        background-image: url('../../../public/information/name.png');
        background-size: 100% 100%;
        width: 40%;
        height: 100%;
      }

      .studentId {
        background-image: url('../../../public/information/studentId.png');
        background-size: 100% 100%;
        width: 50%;
        height: 100%;
      }
    }

    .major {
      background-image: url('../../../public/information/futurePlan.png');
      background-size: 100% 100%;
      width: 50%;
      height: 16%;
    }

    .gender {
      width: 50%;
      height: 16%;
      display: flex;
      flex-direction: row;
      justify-content: space-around;

      .male {
        background-image: url('../../../public/information/male.png');
        background-size: 100% 100%;
        height: 20px;
        width: 20px;
      }

      .female {
        background-image: url('../../../public/information/female.png');
        background-size: 100% 100%;
        height: 20px;
        width: 20px;
      }
    }
  }
}

.positionAndFood {
  background-image: url('../../../public/information/positionAndFood.png');
  background-size: 100% 100%;
  height: 65%;
  width: 85vw;
}
</style>