<template>
  <div class="informationFrame">
    <div class="personalData">
      <div class="personalDataFrame">
        <div class="school"></div>
        <div>
          <div class="name"></div>
          <div class="studentId"></div>
        </div>
        <div class="futurePlan"></div>
        <div class="major"></div>
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
export default class Information extends Vue {}
</script>
<style lang="stylus">
.informationFrame {
  display: flex;
  margin-top: 8vw;
  margin-bottom: 4vw;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  width: 84vw;
  height: 90vh;
}

.personalData {
  background-image: url('../../../public/information/personalData.png');
  background-size: 100% 100%;
  height:35%;
  width: 80vw;

  .personalDataFrame {
    height: 80%;
    transform: translate(-20px, 30px);
    display: flex;
    flex-direction: column;
    align-items: flex-end;

    .school {
      background-image: url('../../../public/information/school.png');
      background-size: 100% 100%;
      width: 50%;
      height: 20%;
    }

    .futurePlan {
      background-image: url('../../../public/information/futurePlan.png');
      background-size: 100% 100%;
      width: 50%;
      height: 20%;
    }

    div {
      width: 50%;
      height: 20%;
      display: flex;
      flex-direction: row;

      .name {
        background-image: url('../../../public/information/name.png');
        background-size: 100% 100%;
        width: 50%;
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
      height: 20%;
    }

    .gender {
      width: 50%;
      height: 20%;
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
  width: 80vw;
}
</style>