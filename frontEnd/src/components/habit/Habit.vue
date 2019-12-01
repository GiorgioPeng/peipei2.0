<template>
  <div class="habitFrame">
    <div class="aircondition">
      <div></div>
      <div></div>
    </div>
    <div class="smoke">
      <div class="wetherSmoke">
        <div style="color:#ca6400">是</div>
        <div style="color:#a2a1a4">否</div>
      </div>
    </div>
    <div class="habitWeight">
      <el-slider
        v-model="smokeWeight"
        :step="1"
        :min="1"
        :max="5"
        :input-size="mini"
        style="transform:translate(0.5vw,-50px);width:45vw"
      ></el-slider>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Habit"
})
export default class Habit extends Vue {
  private smokeWeight: number = 1;
}
</script>
<style lang="stylus">
.habitWeight {
  .el-slider__bar {
    background-color: #ea6561 !important;
  }

  .el-slider__button {
    border: none !important;
    background-color: #ea6561 !important;
  }
}

.habitFrame {
  display: flex;
  margin-top: 15vw;
  margin-bottom: 4vw;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  width: 84vw;
  height: 700px;

  .aircondition {
    background-image: url('../../../public/habit/aircondition.png');
    background-size: 100% 100%;
    height: 45%;
    width: 80vw;
  }

  .smoke {
    background-image: url('../../../public/habit/smoke.png');
    background-size: 100% 100%;
    height: 45%;
    width: 80vw;
    .wetherSmoke{
      text-align:center;
      transform:translateY(118px);
    }
  }
}
</style>