<template>
  <div>
    <div class="gamePhoto">
      <div class="biggerGamePhoto"></div>
      <div class="smallerGamePhoto"></div>
    </div>
    <div class="frame">
      <div class="biggerFrame"></div>
      <div class="smallerFrame"></div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Entertainment"
})
export default class Entertainment extends Vue {}
</script>
<style lang="stylus">
.gamePhoto {
  width: 80vw;
  height: 50vw;
}

.biggerGamePhoto {
  background-image: url('../../../public/entertainment/firstPage/PUBG.png');
  background-size: 100% 100%;
  width: 100%;
  height: 100%;
}

.frame {
  width: 80vw;
  height: 50vw;
}

.biggerFrame {
  background-image: url('../../../public/entertainment/firstPage/frame1.png');
  background-size: 100% 100%;
  width: 100%;
  height: 100%;
}
</style>