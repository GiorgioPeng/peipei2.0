<template>
  <div class="informationFrame">
    <div class="personalData">
      <div class="personalDataFrame">
        <input class="school" :value="school" placeholder="学校" />
        <div>
          <input class="name" :value="name" placeholder="姓名" />
          <input class="studentId" :value="id" placeholder="学号" />
        </div>
        <el-select v-model="futurePlan" class="major" placeholder="未来规划">
          <el-option v-for="item in planChoice" :key="item.value" :value="item.value"></el-option>
        </el-select>
        <el-select v-model="major" class="major" placeholder="专业">
          <el-option v-for="item in majorChoice" :key="item.value" :value="item.value"></el-option>
        </el-select>
        <div class="gender">
          <div :class="{male:!gender.male,maleChecked:gender.male}" @click="chooseGender('male')"></div>
          <div
            :class="{female:!gender.female,femaleChecked:gender.female}"
            @click="chooseGender('female')"
          ></div>
        </div>
      </div>
    </div>
    <div class="positionAndFood">
      <div class="provinceCollect">
        <el-select v-model="province" class="province" placeholder="省份">
          <el-option v-for="item in provinceChoice" :key="item.value" :value="item.value"></el-option>
        </el-select>
        <el-slider
          v-model="areaWeight"
          :step="1"
          :min="1"
          :max="5"
          style="transform:translate(20vw,133px);width:45vw"
        ></el-slider>
      </div>
      <div class="tasteCollect">
        <el-slider
          v-model="sour"
          :step="1"
          :min="1"
          :max="5"
          style="transform:translate(46vw,175px);width:20vw"
        ></el-slider>
        <el-slider
          v-model="sweet"
          :step="1"
          :min="1"
          :max="5"
          style="transform:translate(46vw,165px);width:20vw"
        ></el-slider>
        <el-slider
          v-model="spicy"
          :step="1"
          :min="1"
          :max="5"
          style="transform:translate(46vw,155px);width:20vw"
        ></el-slider>
        <el-slider
          class="tasteWeight"
          v-model="tasteWeight"
          :step="1"
          :min="1"
          :max="5"
          style="transform:translate(20vw,195px);width:45vw"
        ></el-slider>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import "element-ui/lib/theme-chalk/index.css";
import ElementUI from "element-ui";
Vue.use(ElementUI);
@Component({
  name: "Information"
})
export default class Information extends Vue {
  private school: string = "";
  private name: string = "";
  private id: string = "";
  private futurePlan: string = "";
  private planChoice: object[] = [
    {
      value: "搬砖"
    },
    {
      value: "偷电瓶"
    },
    {
      value: "当网管"
    },
    {
      value: "扫厕所"
    }
  ];
  private major: string = "";
  private majorChoice: object[] = [
    {
      value: "计科"
    },
    {
      value: "通信"
    },
    {
      value: "环境"
    },
    {
      value: "会计"
    },
    {
      value: "信管"
    }
  ];
  private gender: object = {
    male: false,
    female: false
  };
  private province: string = "";
  private provinceChoice: object[] = [
    { value: "北京" },
    { value: "天津" },
    { value: "上海" },
    { value: "重庆" },
    { value: "河北" },
    { value: "河南" },
    { value: "云南" },
    { value: "辽宁" },
    { value: "黑龙江" },
    { value: "湖南" },
    { value: "安徽" },
    { value: "山东" },
    { value: "新疆" },
    { value: "江苏" },
    { value: "浙江" },
    { value: "江西" },
    { value: "湖北" },
    { value: "广西" },
    { value: "甘肃" },
    { value: "山西" },
    { value: "内蒙古" },
    { value: "陕西" },
    { value: "吉林" },
    { value: "福建" },
    { value: "贵州" },
    { value: "广东" },
    { value: "青海" },
    { value: "西藏" },
    { value: "四川" },
    { value: "宁夏" },
    { value: "海南" },
    { value: "台湾" },
    { value: "香港" },
    { value: "澳门" }
  ];
  private areaWeight: number = 1;
  private sour: number = 1;
  private sweet: number = 1;
  private spicy: number = 1;
  private tasteWeight: number = 1;
  private chooseGender(genderClick: string) {
    if (genderClick === "male") {
      this.gender["male"] = true;
      this.gender["female"] = false;
    } else {
      this.gender["male"] = false;
      this.gender["female"] = true;
    }
  }
}
</script>
<style lang="stylus">
.el-input {
  width: 100% !important;
}

.el-input__inner {
  background-color: transparent !important;
  outline: none !important;
  border: none !important;
  text-align: center !important;
  padding: 0 !important;
  color: #707070 !important;
  height:31.5px !important;
}

.el-input__suffix {
  height: 30px;
  width: 22px;

  .el-input__suffix-inner i {
    background-image: url('../../../public/information/choice.png');
    background-size: 100% 100%;
    transform: rotate(0deg) !important;
  }
}

.provinceCollect {
  .el-slider__bar {
    background-color: #8fb6f0 !important;
  }

  .el-slider__button {
    border: none !important;
    background-color: #8fb6f0 !important;
  }
}

.tasteCollect {
  .el-slider__bar {
    background-color: #ea6561;
    background-image: linear-gradient(to right, green, yellow, red);
  }

  .el-slider__button {
    border: none !important;
    background-color: #ea6561 !important;
  }

  .tasteWeight {
    .el-slider__bar {
      background-color: #8fb6f0 !important;
      background-size: 0 0;
    }

    .el-slider__button {
      border: none !important;
      background-color: #8fb6f0 !important;
    }
  }
}

input {
  outline: none;
  border: none;
  text-align: center;
  color: #707070;
}

.informationFrame {
  display: flex;
  margin-top: 8vw;
  margin-bottom: 4vw;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  width: 84vw;
  height: 700px;
}

.personalData {
  background-image: url('../../../public/information/personalData.png');
  background-size: 100% 100%;
  height: 35%;
  width: 85vw;

  .personalDataFrame {
    height: 80%;
    transform: translate(-20px, 30px);
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    justify-content: space-between;

    .school {
      background-image: url('../../../public/information/school.png');
      background-size: 100% 100%;
      width: 50%;
      height: 16%;
    }

    .futurePlan {
      background-image: url('../../../public/information/futurePlan.png');
      background-size: 100% 100%;
      width: 50%;
      height: 16%;
    }

    div {
      width: 50%;
      height: 16%;
      display: flex;
      flex-direction: row;
      justify-content: space-between;

      .name {
        background-image: url('../../../public/information/name.png');
        background-size: 100% 100%;
        width: 40%;
        height: 100%;
      }

      .studentId {
        background-image: url('../../../public/information/studentId.png');
        background-size: 100% 100%;
        width: 50%;
        height: 100%;
      }
    }

    .major {
      background-image: url('../../../public/information/futurePlan.png');
      background-size: 100% 100%;
      width: 50%;
      height: 16%;
    }

    .gender {
      width: 50%;
      height: 16%;
      display: flex;
      flex-direction: row;
      justify-content: space-around;

      .male {
        background-image: url('../../../public/information/male.png');
        background-size: 100% 100%;
        height: 20px;
        width: 30px;
      }

      .maleChecked {
        background-image: url('../../../public/information/checkedMale.png');
        background-size: 100% 100%;
        height: 20px;
        width: 30px;
      }

      .female {
        background-image: url('../../../public/information/female.png');
        background-size: 100% 100%;
        height: 20px;
        width: 30px;
      }

      .femaleChecked {
        background-image: url('../../../public/information/checkedFemale.png');
        background-size: 100% 100%;
        height: 20px;
        width: 30px;
      }
    }
  }
}

.positionAndFood {
  background-image: url('../../../public/information/positionAndFood.png');
  background-size: 100% 100%;
  height: 65%;
  width: 85vw;

  .province {
    width: 50%;
    background-image: url('../../../public/information/futurePlan.png');
    background-size: 100% 100%;
    height: 30px;
    line-height: 30px;
    transform: translate(35vw, 70px);
  }
}
</style>