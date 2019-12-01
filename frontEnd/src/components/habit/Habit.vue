<template>
  <div class="habitFrame">
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
  }
}
</style>