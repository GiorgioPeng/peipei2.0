<template>
  <div>
    <div class="line"></div>
    <div class="sleepPageFirstFrame">
      <div class="settingSleepTime"></div>
      <div class="hint"></div>
      <div class="day"></div>
      <div class="night"></div>
    </div>
    <div class="weightPart">
      <div class="weightHint"></div>
      <div class="weight"></div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Sleep"
})
export default class Sleep extends Vue {}
</script>
<style lang="stylus">
.line {
  width: 5.5vw;
  height: 45vh;
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
  height: 70vh;

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
  }

  .night {
    background-image: url('../../../public/sleep/firstPage/night.png');
    background-size: 100% 100%;
    background-repeat: no-repeat;
    background-position: left top;
    width: 60vw;
    height: 30%;
    margin-left: 6vw;
  }
}

.weightPart {
  height: 20vh;

  .weightHint {
    height: 5.8vw;
    width: 24.7vw;
    background-image: url('../../../public/sleep/firstPage/weightHint.png');
    background-size: 100% 100%;
    margin-left: 4.6vw;
  }
}
</style>