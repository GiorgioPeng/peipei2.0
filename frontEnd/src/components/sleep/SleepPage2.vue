<template>
  <div class="timeSetting">
    <div class="dayOrNight">
      <div :class="{getUpChecked:timeFrame=='AM',getUp:timeFrame=='PM'}" @click="timeFrame='AM'"></div>
      <div :class="{goToBedChecked:timeFrame=='PM',goToBed:timeFrame=='AM'}" @click="timeFrame='PM'"></div>
    </div>
    <div style="display:flex;justify-content:center;align-items:center">
      <div class="back"></div>
      <div class="time">
        <div class="mainDisplay">
          <div>{{timeFrame=='AM'?am[0]:pm[0]}}</div>
          <div>{{timeFrame=='AM'?am[1]:pm[1]}}</div>
          <div></div>
          <div>{{timeFrame=='AM'?am[2]:pm[2]}}</div>
          <div>{{timeFrame=='AM'?am[3]:pm[3]}}</div>
          <div>{{timeFrame=='AM'?am[4]:pm[4]}}</div>
          <div>{{timeFrame=='AM'?am[5]:pm[5]}}</div>
          <div></div>
          <div>{{timeFrame=='AM'?am[6]:pm[6]}}</div>
          <div>{{timeFrame=='AM'?am[7]:pm[7]}}</div>
        </div>
        <div class="detailChoice">
          <div class="amOfPm">{{this.timeFrame}}</div>
          <div class="timeChoice">
            <div v-if="timeFrame=='AM'">
              <div v-for="am in amChoiceDetail" :key="am" @click="chooseTime(am,'am')">{{am}}</div>
            </div>
            <div v-else>
              <div v-for="pm in pmChoiceDetail" :key="pm" @click="chooseTime(pm,'pm')">{{pm}}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="save"></div>
  </div>
</template>
<script lang='ts'>
import { Vue, Component } from "vue-property-decorator";
@Component({
  name: "SleepPage2"
})
export default class SleepPage2 extends Vue {
  private am: string[] = ["0", "8", "0", "1", "0", "9", "0", "0"];
  private pm: string[] = ["2", "2", "0", "1", "2", "3", "0", "0"];
  private timeFrame: string = "AM";
  private amChoiceDetail: string[] = [
    "07:00之前",
    "07:01-08:00",
    "08:01-09:00",
    "09:01-10:00",
    "10:01之后"
  ];
  private pmChoiceDetail: string[] = [
    "21:00之前",
    "21:01-22:00",
    "22:01-23:00",
    "23:01-24:00",
    "24:01之后"
  ];
  public chooseTime(time: string, time2: string) {
    let i = 0;
    let tempTime = time;
    if (time2 == "am") {
      if (time.indexOf("之前") != -1) {
        time = time.slice(0, time.length - 2);
        time = "XX:XX-" + time;
      } else if (time.indexOf("之后") != -1) {
        time = time.slice(0, time.length - 2);
        time += "-XX:XX";
      }
      time.split("-").map(res => {
        res.split(":").map(res => {
          res.split("").map(res => (this.am[i++] = res));
        });
      });
      for (let k in this.amChoiceDetail) {
        if (this.amChoiceDetail[k] == tempTime)
          this.amChoiceDetail[k] = this.amChoiceDetail[2];
      }
      this.amChoiceDetail[2] = tempTime;

      this.$forceUpdate();
      console.log(this.am);
    }
    if (time2 == "pm") {
      if (time.indexOf("之前") != -1) {
        time = time.slice(0, time.length - 2);
        time = "XX:XX-" + time;
      } else if (time.indexOf("之后") != -1) {
        time = time.slice(0, time.length - 2);
        time += "-XX:XX";
      }
      time.split("-").map(res => {
        res.split(":").map(res => {
          res.split("").map(res => (this.pm[i++] = res));
        });
      });
      for (let k in this.pmChoiceDetail) {
        if (this.pmChoiceDetail[k] == tempTime)
          this.pmChoiceDetail[k] = this.pmChoiceDetail[2];
      }
      this.pmChoiceDetail[2] = tempTime;

      this.$forceUpdate();
      console.log(this.pm);
    }
  }
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

    .getUpChecked {
      background-image: url('../../../public/sleep/secondPage/getUpTimeChecked.png');
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

    .goToBedChecked {
      background-image: url('../../../public/sleep/secondPage/sleepTimeChecked.png');
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
    transform: translateX(10px);
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
      display: grid;
      grid-template-columns: 45px 45px 22px 45px 45px;
      grid-column-gap: 5px;
      grid-row-gap: 10px;

      div {
        font-size: 60px;
        font-weight: bold;
        display: flex;
        justify-content: center;
        align-items: center;
        color: white;
        background-image: url('../../../public/sleep/secondPage/shadow.png');
        background-size: 100% 50%;
        background-position: center top;
        background-repeat: no-repeat;
      }

      div:nth-child(3) {
        background-image: none;
      }

      div:nth-child(8) {
        background-image: none;
      }
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
          text-align: center;

          div:nth-child(3) {
            color: #56bebf;
            background-image: url('../../../public/sleep/secondPage/greySlot.png');
            background-size: 100% 100%;
          }
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