<template>
  <div class="timeSetting">
    <div class="dayOrNight">
      <div class="getUp"></div>
      <div class="goToBed"></div>
    </div>
    <div style="display:flex;justify-content:center;align-items:center">
      <div class="back"></div>
      <div class="time">
        <div class="mainDisplay"></div>
        <div class="detailChoice">
          <div class="amOfPm">{{this.timeFrame}}</div>
          <div class="timeChoice">
            <div v-if="timeFrame=='AM'">
              <div v-for="am in amChoiceDetail" :key="am">{{am}}</div>
            </div>
            <div v-else v-for="pm in pmChoiceDetail" :key="pm">{{pm}}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="save"></div>
  </div>
</template>
<script lang='ts'>
import { Vue, Component } from "vue-property-decorator";
export default class SleepPage2 extends Vue {
  private timeFrame: string = "AM";
  private amChoiceDetail: string[] = [
    "7:00之前",
    "7:01-8:00",
    "8:01-9:00",
    "9:01-10:00",
    "10:01之后"
  ];
  private pmChoiceDetail: string[] = [
    "21:00之前",
    "21:01-22:00",
    "22:01-23:00",
    "23:01-24:00",
    "24:01之后"
  ];
}
</script>
<style lang="stylus">
.timeSetting {
  margin: 50px auto;
  width: 85vw;
  height: 700px;

  .dayOrNight {
    width: 100%;
    height: 50px;
    display: flex;
    justify-content: space-around;
    background-image: url('../../../public/sleep/secondPage/line.png');
    background-size: 1% 70%;
    background-repeat: no-repeat;
    background-position: center center;
    align-items: center;

    .getUp {
      background-image: url('../../../public/sleep/secondPage/getUpTime.png');
      background-size: 100% 100%;
      width: 100px;
      height: 25px;
    }

    .goToBed {
      background-image: url('../../../public/sleep/secondPage/sleepTime.png');
      background-size: 100% 100%;
      width: 100px;
      height: 25px;
    }
  }

  .back {
    background-image: url('../../../public/sleep/secondPage/back.png');
    background-size: 100% 100%;
    height: 30px;
    width: 30px;
    transform:translateX(10px);
  }

  .time {
    height: 600px;
    flex: 1;

    .mainDisplay {
      margin: 30px auto;
      background-image: url('../../../public/sleep/secondPage/backgroundOfTime.png');
      background-size: 100% 100%;
      width: 70%;
      height: 30%;
    }

    .detailChoice {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 390px;

      .amOfPm {
        background-image: url('../../../public/sleep/secondPage/blueBottomLine.png');
        background-size: 100% 10%;
        background-repeat: no-repeat;
        background-position: center bottom;
        line-height: 10px;
        text-align: center;
        font-size: 30px;
        height: 30px;
        width: 50px;
        color: #ceccd1;
      }

      .timeChoice {
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        height: 100%;

        div {
          font-size: 30px;
          padding: 15px;
          color: #ceccd1;
        }
      }
    }
  }

  .save {
    background-image: url('../../../public/sleep/secondPage/save.png');
    background-size: 100% 100%;
    height: 50px;
    width: 190px;
  }
}
</style>