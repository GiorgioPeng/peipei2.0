<template>
  <div>
    <div class="gamePhoto">
      <div class="biggerGamePhoto"></div>
      <div class="smallerGamePhoto"></div>
    </div>
    <div class="frame">
      <div class="biggerFrame">
        <div style="text-align:center">网络游戏</div>
        <div style="text-align:center">手游</div>
        <div class="smallGameChoice" v-for="game in gameChoice" :key="game.type">
          <div style="padding-left:15px">
            <span style="color:#707070">{{game.type}}</span>
            <br />
            <span style="color:#939497">{{game.game[0]===0?'请点击右侧图标进入选择':game.game[0]}}</span>
          </div>
          <div class="rightIcon"></div>
        </div>
        <div class="siderBar">
          <div class="circle"></div>
        </div>
      </div>
      <div class="smallerFrame"></div>
    </div>
    <div class="pages" style="display:flex;justify-content:center;width:100%;height:10px">
      <div
        v-for="type in types"
        :key="type.name"
        :class="{uncheckedType:!type.checked,checkedType:type.checked}"
      ></div>
    </div>
    <div class="entertainmentWeight"></div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Entertainment"
})
export default class Entertainment extends Vue {
  private types: object[] = [
    {
      name: "type1",
      checked: false
    },
    {
      name: "type2",
      checked: false
    },
    {
      name: "type3",
      checked: false
    },
    {
      name: "type4",
      checked: false
    },
    {
      name: "type5",
      checked: false
    }
  ];
  private gameChoice: object[] = [
    {
      type: "王者类",
      game: [0, "王者荣耀"]
    },
    {
      type: "射击类",
      game: [0, "绝地求生"]
    },
    {
      type: "棋牌类",
      game: [0, ""]
    }
  ];
  private current: boolean = false; // use to identify wether focus on
}
</script>
<style lang="stylus">
.gamePhoto {
  width: 100%;
  height: 18vh;
  margin: 0;

  .biggerGamePhoto {
    background-image: url('../../../public/entertainment/firstPage/PUBG.png');
    background-size: 100% 100%;
    width: 60%;
    height: 100%;
    margin: 10vh auto;
    margin-bottom: 1vh;
  }

  .smallerGamePhoto {
    background-image: url('../../../public/entertainment/firstPage/backGame.png');
    background-size: 100% 100%;
    width: 40%;
    height: 70%;
    transform: translate(70vw, -16vh);
  }
}

.frame {
  width: 80%;
  height: 50vw;
  margin: 1vh auto;
  margin-bottom: -8px;

  .biggerFrame {
    background-image: url('../../../public/entertainment/firstPage/frame1.png');
    background-size: 100% 100%;
    width: 100%;
    height: 95%;

    .smallGameChoice {
      display: flex;
      justify-content: space-around;
      align-items: center;
      background-image: url('../../../public/entertainment/firstPage/smallNotChoose.png');
      background-position: left;
      background-size: 10% 60%;
      background-repeat: no-repeat;
      padding-left: 3px;
      margin: 0 12px;

      .rightIcon {
        background-image: url('../../../public/entertainment/firstPage/rightIcon.png');
        background-size: 100% 100%;
        height: 15px;
        width: 15px;
      }
    }

    .bigGameChoice {
      display: flex;
      justify-content: space-around;
      align-items: center;
      background-image: url('../../../public/entertainment/firstPage/bigNotChoose.png');
      background-position: left;
      background-size: 10% 60%;
      background-repeat: no-repeat;
      padding-left: 3px;
      margin-left: 12px;

      .rightIcon {
        background-image: url('../../../public/entertainment/firstPage/rightIcon.png');
        background-size: 100% 100%;
        height: 15px;
        width: 15px;
      }
    }

    .siderBar {
      transform: translate(64vw, -20vh);
      background-image: url('../../../public/entertainment/firstPage/siderbar.png');
      background-size: 100% 100%;
      height: 85%;
      width: 2%;

      .circle {
        background-image: url('../../../public/entertainment/firstPage/smallChoose.png');
        background-size: 100% 100%;
        background-repeat: no-repeat;
        margin-left: -3px;
        height: 13px;
        width: 13px;
      }
    }
  }
}

.uncheckedType {
  margin-left: 5px;
  background-image: url('../../../public/entertainment/firstPage/uncheckedType.png');
  background-size: 100% 100%;
  width: 10px;
  height: 10px;
}

.entertainmentWeight {
  background-image: url('../../../public/entertainment/firstPage/weightHint.png');
  background-size: 30% 35%;
  background-repeat: no-repeat;
  background-position: left top;
  margin-top: 5px;
  width: 100%;
  height: 8vh;
}
</style>