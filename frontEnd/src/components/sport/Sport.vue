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
        <div
          v-for="(kind,key) in sportChoiceInput"
          :key="key"
          @click="handleClick(kind)"
          :class="{classForChecked:kind.checked,classForBasic:!kind.checked}"
        ></div>
      </div>
    </div>
  </div>
</template>
<script lang='ts'>
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Sport"
})
export default class Sport extends Vue {
  private sportChoiceInput: Object[] = [
    {
      kind: "runInput",
      checked: false
    },
    {
      kind: "badmintonInput",
      checked: false
    },
    {
      kind: "basketballInput",
      checked: false
    },
    {
      kind: "soccerballInput",
      checked: false
    },
    {
      kind: "bicycleInput",
      checked: false
    },
    {
      kind: "swimInput",
      checked: false
    },
    {
      kind: "fitnessInput",
      checked: false
    },
    {
      kind: "pingpangInput",
      checked: false
    },
    {
      kind: "yogaInput",
      checked: false
    }
  ];
  public handleClick(kind: any) {
    if (!kind.checked) kind.checked = true;
    else kind.checked = false;
  }
}
</script>
<style lang="stylus">
.sportBg {
  background-image: url('../../../public/sport/division.png');
  background-size: 100% 100%;
  height: 550px;
  width: 85vw;
  margin:50px auto;

  .sportChoice {
    transform: translate(1.5vw, 50px);
    width: 95%;
    height: 80%;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: repeat(3, 1fr);
    grid-row-gap: 10px;
    grid-column-gap: 10px;

    .run {
      background-image: url('../../../public/sport/run.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .badminton {
      background-image: url('../../../public/sport/badmiton.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .basketball {
      background-image: url('../../../public/sport/basketball.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .soccer {
      background-image: url('../../../public/sport/soccer.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .bicycle {
      background-image: url('../../../public/sport/bicycle.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .swim {
      background-image: url('../../../public/sport/swim.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .weight {
      background-image: url('../../../public/sport/fitness.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .pingpang {
      background-image: url('../../../public/sport/pingpang.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }

    .yoga {
      background-image: url('../../../public/sport/yoga.png');
      background-size: 65% 85%;
      background-repeat: no-repeat;
      background-position: center;
    }
  }

  .sportChoiceInput {
    transform: translate(1vw, -400px);
    width: 95%;
    height: 80%;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: repeat(3, 1fr);
    grid-row-gap: 10px;
    grid-column-gap: 10px;
  
    .classForBasic {
      background-image: url('../../../public/sport/check.png');
      background-size: 15% 15%;
      background-position: right bottom;
      background-repeat: no-repeat;
    }

    .classForChecked {
      background-image: url('../../../public/sport/checked.png');
      background-size: 15% 15%;
      background-position: right bottom;
      background-repeat: no-repeat;
    }
  }
}
</style>
