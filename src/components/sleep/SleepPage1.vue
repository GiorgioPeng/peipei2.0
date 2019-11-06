<template>
  <div>
    <div class="line"></div>
    <div>
      <div class="settingHint"></div>
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
.day {
  background-image: url('../../../public/sleep/firstPage/day.png');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  background-position: left top;
  width: 200px;
  height: 200px;
}

.night {
  background-image: url('../../../public/sleep/firstPage/night.png');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  background-position: left top;
  width: 200px;
  height: 200px;
}
</style>