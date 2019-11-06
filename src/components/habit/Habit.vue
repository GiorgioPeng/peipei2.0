<template>
  <div>
    <div class="aircondition"></div>
    <div class="smoke"></div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Habit"
})
export default class Habit extends Vue {}
</script>
<style lang="stylus">
.aircondition {
  background-image: url('../../../public/habit/aircondition.png');
  background-size: 100% 100%;
  height: 300px;
  width: 80vw;
}

.smoke {
  background-image: url('../../../public/habit/smoke.png');
  background-size: 100% 100%;
  height: 300px;
  width: 80vw;
}
</style>