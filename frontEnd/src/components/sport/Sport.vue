<template>
  <div>
    <div class="sportBg">
      <div class="sportChoice">
        <div class="run"></div>
        <div class="badminton"></div>
        <div class="basketball"></div>
        <div class="soccer"></div>
        <div class="bicycle"></div>
        <div class="swim"></div>
        <div class="weight"></div>
        <div class="pingpang"></div>
        <div class="yoga"></div>
      </div>
            <div class="sportChoiceInput">
        <div class="runInput"></div>
        <div class="badmintonInput"></div>
        <div class="basketballInput"></div>
        <div class="soccerInput"></div>
        <div class="bicycleInput"></div>
        <div class="swimInput"></div>
        <div class="weightInput"></div>
        <div class="pingpangInput"></div>
        <div class="yogaInput"></div>
      </div>
    </div>
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
.sportBg {
  background-image: url('../../../public/sport/division.png');
  background-size: 100% 100%;
  height: 80vw;
  width: 73.2vw;
  margin: 16vw auto;
}

.sportChoice {
  transform: translate(5.5vw,5vh);
  width: 95%;
  height: 80%;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(3, 1fr);
  grid-row-gap: 10px;
  grid-column-gap: 10px;
}

.run {
  background-image: url('../../../public/sport/run.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.badminton {
  background-image: url('../../../public/sport/badmiton.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.basketball {
  background-image: url('../../../public/sport/basketball.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.soccer {
  background-image: url('../../../public/sport/soccer.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.bicycle {
  background-image: url('../../../public/sport/bicycle.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.swim {
  background-image: url('../../../public/sport/swim.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.weight {
  background-image: url('../../../public/sport/fitness.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.pingpang {
  background-image: url('../../../public/sport/pingpang.png');
  background-size: 65% 85%;
  background-repeat:no-repeat;
}

.yoga {
  background-image: url('../../../public/sport/yoga.png');
  background-size: 65% 85%;
  background-repeat: no-repeat;
}
</style>
