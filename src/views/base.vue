<template>
  <div class="base">
    <div class="top">
      <div class="setting"></div>
      <div class="label"></div>
      <div class="typeLogo">
        <div
          :class="{ checkedInformation : isCheckedInformation , noCheckedInformation : !isCheckedInformation }"
        ></div>
        <div :class="{ checkedSleep : isCheckedSleep , noCheckedSleep : !isCheckedSleep }"></div>
        <div :class="{ checkedHabit : isCheckedHabit , noCheckedHabit : !isCheckedHabit }"></div>
        <div :class="{ checkedSport : isCheckedSport , noCheckedSport : !isCheckedSport}"></div>
        <div :class="isCheckedEntertainment"></div>
      </div>
    </div>

    <div class="main">
      <div v-if="currentPage === 'Information'">
        <Information></Information>
      </div>
      <div v-else-if="currentPage === 'Sleep'">
        <Sleep></Sleep>
      </div>
      <div v-else-if="currentPage === 'Habit'">
        <Habit></Habit>
      </div>
      <div v-else-if="currentPage === 'Sport'">
        <Sport></Sport>
      </div>
      <div v-else>
        <Entertainment></Entertainment>
      </div>
    </div>

    <div v-if="currentPage !== 'Entertainment'" class="nextOne"></div>
    <div v-else class="finish"></div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
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
  private currentPage: string = "Sport";
  private pageSet: string[] = [
    "Information",
    "Sleep",
    "Habit",
    "Sport",
    "Entertainment"
  ];
}
</script>
<style lang="stylus">
html, body {
  height: 100%;
  width: 100%;
  padding: 0;
  margin: 0;
}

.base {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  background-image: url('../../public/base/backgroundBase.png');
  background-position: top left;
  background-repeat: no-repeat;
  background-size: 100% 100%;
}

.top {
  background-image: url('../../public/base/topBase.png');
  background-position: top left;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 20vh;
}

.setting {
  background-image: url('../../public/base/settingBase.png');
  background-position: top left;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  width: 8vw;
  height: 8vw;
  position: fixed;
  top: 10vw;
  left: 5vw;
  z-index: 2;
}

.label {
  position: fixed;
  top: 8vw;
  left: 30vw;
  z-index: 2;
  background-image: url('../../public/base/sportTitle.png');
  background-size: 100% 100%;
  background-position: left top;
  width: 40vw;
  height: 12vw;
}

.main {
  position: fixed;
  top: 20vh;
  left: 8vw;
  width: 84vw;
  height: 64.5vh;
  overflow: hidden;
  background-color: rgba(10, 10, 10, 0.1);
}
</style>