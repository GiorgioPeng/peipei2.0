<template>
  <div class="base">
    <div class="top">
      <div class="setting"></div>
      <div class="label"></div>
    </div>

  </div>
</template>
<script lang="ts">
import {Component, Vue,Prop} from "vue-property-decorator";

@Component({
    name:"dataCollect"
})
export default class extends Vue{
    private currentPage:string="";
    private currentLogo:string="";

};
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
</style>