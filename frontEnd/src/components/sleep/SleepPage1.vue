<template>
  <div v-if="page">
    <div class="line"></div>
    <div class="sleepPageFirstFrame">
      <div class="settingSleepTime"></div>
      <div class="hint"></div>
      <div class="day" @click="page=false">{{this.am}}</div>
      <div class="night" @click="page=false">{{this.pm}}</div>
    </div>
    <div class="weightPart">
      <div class="weightHint"></div>
      <div class="weight">
        <el-slider
          v-model="sleepWeight"
          :step="1"
          :min="1"
          :max="5"
          style="transform:translate(8vw,10px);width:70vw"
        ></el-slider>
      </div>
    </div>
  </div>
  <SleepPage2 v-else></SleepPage2>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import SleepPage2 from "./SleepPage2.vue";
@Component({
  name: "Sleep",
  components: { SleepPage2 }
})
export default class Sleep extends Vue {
  private page: boolean = true;
  private sleepWeight: number = 1;
  private am:string="点击添加平均睡眠时间段"
  private pm:string="点击添加平均睡眠时间段"
}
</script>
<style lang="stylus">
.weight {
  .el-slider__bar {
    background-color: #56bebf !important;
  }

  .el-slider__button {
    border: none !important;
    background-color: #56bebf !important;
  }
}

.line {
  width: 5.5vw;
  height: 400px;
  background-image: url('../../../public/sleep/firstPage/line.png');
  background-size: 100% 100%;
  position: absolute;
  left: 8.6vw;
  top: 18vw;
}

.sleepPageFirstFrame {
  display: flex;
  margin-top: 15vw;
  margin-bottom: 4vw;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  width: 84vw;
  height: 550px;

  .hint {
    background-image: url('../../../public/sleep/firstPage/settingHint.png');
    background-size: 100% 100%;
    width: 44.5vw;
    height: 6vw;
    margin-left: -6vw;
  }

  .day {
    background-image: url('../../../public/sleep/firstPage/day.png');
    background-size: 100% 100%;
    background-repeat: no-repeat;
    background-position: left top;
    width: 60vw;
    height: 30%;
    margin-left: 6vw;
    color:white;
    font-size:16px;
    text-align:center;
    line-height:280px;
  }

  .night {
    background-image: url('../../../public/sleep/firstPage/night.png');
    background-size: 100% 100%;
    background-repeat: no-repeat;
    background-position: left top;
    width: 60vw;
    height: 30%;
    margin-left: 6vw;
    color:white;
    font-size:16px;
    text-align:center;
    line-height:280px;
  }
}

.weightPart {
  height: 100px;

  .weightHint {
    height: 5.8vw;
    width: 24.7vw;
    background-image: url('../../../public/sleep/firstPage/weightHint.png');
    background-size: 100% 100%;
    margin-left: 4.6vw;
  }
}
</style>