<template>
  <div>
    <div class="sport"></div>
  </div>
</template>
<script>
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Sport"
})
export default class Sport extends Vue {}
</script>
<style lang="stylus">
.sport {
  background-image: url('../../../public/sport/division.png');
  background-size: 100% 100%;
  height: 300px;
  width: 80vw;
}
</style>