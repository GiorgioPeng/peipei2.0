<template>
  <div>
    <div class="personalData"></div>
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
.personalData {
  background-image: url('../../../public/information/personalData.png');
  background-size: 100% 100%;
  height: 300px;
  width: 80vw;
}

.positionAndFood {
  background-image: url('../../../public/information/positionAndFood.png');
  background-size: 100% 100%;
  height: 300px;
  width: 80vw;
}
</style>