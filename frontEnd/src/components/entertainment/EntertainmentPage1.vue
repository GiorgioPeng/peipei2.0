<template>
  <div>
    <div class="gamePhoto">
      <div class="biggerGamePhoto"></div>
      <div class="smallerGamePhoto"></div>
    </div>
    <div class="frame">
      <div class="biggerFrame">
        <div style="text-align:center;color:#707070;font-weight:bold">网络游戏</div>
        <div style="text-align:center;color:#949397;font-size:10px">手游</div>
        <div class="choice">
          <div
            :class="{smallGameChoice:!game.choose,smallGameChoiceChecked:game.choose}"
            v-for="game in gameChoice"
            :key="game.type"
            @click="game.choose = !game.choose"
          >
            <div style="padding-left:20px">
              <span style="color:#707070">{{game.type}}</span>
              <br />
              <span
                style="color:#939497;font-size:10px"
              >{{game.game[0]===0?'请点击右侧图标进入选择':game.game[0]}}</span>
            </div>
            <div class="rightIcon"></div>
          </div>
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
    <div class="entertainmentWeight">
      <el-slider
        v-model="entertainmentWeight"
        :step="1"
        :min="1"
        :max="5"
        style="transform:translate(8vw,30px);width:70vw"
      ></el-slider>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Entertainment"
})
export default class Entertainment extends Vue {
  private entertainmentWeight: number = 1;
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
  ]; //下面5个球
  private gameChoice: object[] = [
    {
      type: "MOBA类",
      game: [0],
      choose: false
    },
    {
      type: "射击类",
      game: [0],
      choose: false
    },
    {
      type: "棋牌类",
      game: [0],
      choose: false
    },
    {
      type: "动作类",
      game: [0],
      choose: false
    },
    {
      type: "策略类",
      game: [0],
      choose: false
    },
    {
      type: "休闲类",
      game: [0],
      choose: false
    },
    {
      type: "儿童类",
      game: [0],
      choose: false
    },
    {
      type: "益智类",
      game: [0],
      choose: false
    },
    {
      type: "女生类",
      game: [0],
      choose: false
    },
    {
      type: "冒险类",
      game: [0],
      choose: false
    },
    {
      type: "体育类",
      game: [0],
      choose: false
    }
  ];
  private current: boolean = false; // use to identify wether focus on
}
</script>
<style lang="stylus">
.entertainmentWeight {
  .el-slider__bar {
    background-color: #56bebf !important;
  }

  .el-slider__button {
    border: none !important;
    background-color: #56bebf !important;
  }
}

.gamePhoto {
  width: 100%;
  height: 150px;
  margin: 0;

  .biggerGamePhoto {
    background-image: url('../../../public/entertainment/firstPage/PUBG.png');
    background-size: 100% 100%;
    width: 60%;
    height: 100%;
    margin: 50px auto;
    margin-bottom: 5px;
  }

  .smallerGamePhoto {
    background-image: url('../../../public/entertainment/firstPage/backGame.png');
    background-size: 100% 100%;
    width: 40%;
    height: 70%;
    transform: translate(70vw, -130px);
  }
}

.frame {
  width: 80%;
  height: 400px;
  margin: 20px auto;
  margin-bottom: -8px;

  .biggerFrame {
    background-image: url('../../../public/entertainment/firstPage/frame1.png');
    background-size: 100% 100%;
    width: 100%;
    height: 95%;
    padding-top: 10px;

    .choice {
      height: 320px;
      overflow: scroll;

      .smallGameChoice {
        display: flex;
        justify-content: space-around;
        align-items: center;
        background-image: url('../../../public/entertainment/firstPage/smallNotChoose.png');
        background-position: 10px center;
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

      .smallGameChoiceChecked {
        display: flex;
        justify-content: space-around;
        align-items: center;
        background-image: url('../../../public/entertainment/firstPage/smallChoose.png');
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
  height: 80px;
}

::-webkit-scrollbar {
  width: 10px;
  height: 0px;
}

/* 定义滚动条的轨道颜色、内阴影及圆角 */
::-webkit-scrollbar-track {
  background-image: url('../../../public/entertainment/firstPage/siderbar.png');
  background-size: 100% 100%;
}

/* 定义滑块颜色、内阴影及圆角 */
::-webkit-scrollbar-thumb {
  background-image: url('../../../public/entertainment/firstPage/smallChoose.png');
  background-size: 100% 100%;
}
</style>