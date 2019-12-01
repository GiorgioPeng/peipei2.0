<template>
  <div class="base">
    <div class="top">
      <!-- 设置按钮 -->
      <div class="topBgImg"></div>
      <div class="setting"></div>
      <!-- 顶部表示页面的部分 -->
      <div class="label"></div>
      <!-- 最上方的那几个掉下来的部分 -->
      <div class="typeLogo">
        <div
          :class="{ checkedInformation : isCheckedInformation , noCheckedInformation : !isCheckedInformation }"
          @click="toggle(0)"
        ></div>
        <div
          :class="{ checkedSleep : isCheckedSleep , noCheckedSleep : !isCheckedSleep }"
          @click="toggle(1)"
        ></div>
        <div
          :class="{ checkedHabit : isCheckedHabit , noCheckedHabit : !isCheckedHabit }"
          @click="toggle(2)"
        ></div>
        <div
          :class="{ checkedSport : isCheckedSport , noCheckedSport : !isCheckedSport}"
          @click="toggle(3)"
        ></div>
        <div
          :class="{ checkedEntertainment : isCheckedEntertainment , noCheckedEntertainment : !isCheckedEntertainment}"
          @click="toggle(4)"
        ></div>
      </div>
    </div>
    <!-- 主要内容 -->
    <div class="main">
      <div v-if="pageSet[currentPage] === 'Information'">
        <Information></Information>
      </div>
      <div v-else-if="pageSet[currentPage] === 'Sleep'">
        <Sleep></Sleep>
      </div>
      <div v-else-if="pageSet[currentPage] === 'Habit'">
        <Habit></Habit>
      </div>
      <div v-else-if="pageSet[currentPage] === 'Sport'">
        <Sport></Sport>
      </div>
      <div v-else>
        <Entertainment></Entertainment>
      </div>
    </div>

    <!-- 页面切换及完成 -->
    <div v-if="pageSet[currentPage] !== 'Entertainment'" @click="changePage" class="nextOne"></div>
    <div v-else class="finish"></div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import pushData from "../api/pushData";
import Information from "@/components/information/Information.vue";
import Habit from "@/components/habit/Habit.vue";
import Sport from "@/components/sport/Sport.vue";
import Sleep from "@/components/sleep/SleepPage1.vue";
import Entertainment from "@/components/entertainment/EntertainmentPage1.vue";
@Component({
  name: "dataCollect",
  components: {
    Information,
    Sleep,
    Habit,
    Sport,
    Entertainment
  }
})
export default class extends Vue {
  private currentPage: number = 0;
  private pageSet: string[] = [
    "Information",
    "Sleep",
    "Habit",
    "Sport",
    "Entertainment"
  ];
  private isChecked: boolean[] = [true, false, false, false, false];
  // 之后用于替换下面五个属性
  private isCheckedInformation: boolean = true;
  private isCheckedSleep: boolean = true;
  private isCheckedHabit: boolean = true;
  private isCheckedSport: boolean = true;
  private isCheckedEntertainment: boolean = true;
  public changePage() {
    this.currentPage++;
    let res = pushData();
    console.log(res);
  }
  public toggle(label: number) {
    this.isChecked = this.isChecked.map(item => (item = false));
    this.isChecked[label] = true;
    this.currentPage = label;
    console.log(this.isChecked);
  }
}
</script>
<style lang="stylus">
html,body{
  margin:0;
  padding:0;
}
.base {
  margin: 0;
  padding: 0;
  height: 1100px;
  width: 100%;
  background-image: url('../../public/base/backgroundBase.png');
  background-repeat: no-repeat;
  background-size: 100% 100%;
  overflow: hidden;
}

.top {
  width: 100%;
  height: 150px;

  .topBgImg {
    background-image: url('../../public/base/topBase.png');
    background-position: top left;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 150px;
    z-index: 3;
  }

  .setting {
    background-image: url('../../public/base/settingBase.png');
    background-position: top left;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    width: 8vw;
    height: 8vw;
    position: absolute;
    top: 10vw;
    left: 5vw;
    z-index: 4;
  }

  .label {
    position: absolute;
    top: 8vw;
    left: 30vw;
    z-index: 4;
    background-image: url('../../public/base/sportTitle.png');
    background-size: 100% 100%;
    background-position: left top;
    width: 40vw;
    height: 12vw;
  }
}

.main {
  transform: translate(8vw, 35px);
  width: 84vw;
}

.checkedInformation {
  position: absolute;
  left: 15vw;
  top: 75px;
  background-image: url('../../public/base/checkedInformation.png');
  background-size: 100% 100%;
  width: 12.5vw;
  height: 27vw;
  z-index: 1;
}

.checkedSleep {
  position: absolute;
  top: 70px;
  left: 30vw;
  background-image: url('../../public/base/checkedSleep.png');
  background-size: 100% 100%;
  width: 12.5vw;
  height: 27vw;
  z-index: 1;
}

.checkedHabit {
  position: absolute;
  top: 80px;
  left: 44.7vw;
  background-image: url('../../public/base/checkedHabit.png');
  background-size: 100% 100%;
  width: 12.5vw;
  height: 27vw;
  z-index: 1;
}

.checkedSport {
  position: absolute;
  top: 100px;
  left: 59.5vw;
  background-image: url('../../public/base/checkedSport.png');
  background-size: 100% 100%;
  width: 12.5vw;
  height: 27vw;
  z-index: 1;
}

.checkedEntertainment {
  position: absolute;
  top: 110px;
  left: 74vw;
  background-image: url('../../public/base/checkedEntertainment.png');
  background-size: 100% 100%;
  width: 12.5vw;
  height: 27vw;
  z-index: 1;
}

.nextOne {
  background-image: url('../../public/base/nextBase.png');
  background-size: 100% 100%;
  width: 44.5vw;
  height: 11.5vw;
  transform: translate(28vw, 80px);
}
</style>