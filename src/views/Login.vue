<template>
  <div class="Login">
    <div class="logo">
      <img src="../../public/Login/logo.png" />
    </div>
    <div class="classForInputData">
      <input type="text" v-model.lazy="school" placeholder="学校" />
      <input type="text" v-model.lazy="name" placeholder="姓名" />
      <input type="text" v-model.lazy="studentId" placeholder="学号" />
    </div>
    <div class="loginButton" @click="loginFunction"></div>
    <div class="agree">
      <input type="checkbox" id="agree" />
      <label for="agree">{{agree}}</label>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";
// import VueRouter from "vue-router";
import router from "@/router/index";
// Vue.use(VueRouter);

@Component({
  name: "Login"
})
export default class extends Vue {
  private school: string = "";
  private name: string = "";
  private studentId: string = "";
  private agree: string = "同意佩佩读取您的个人信息";

  private loginFunction = () => {
    router.push({
      path: "/dataCollect"
    });
  };
}
</script>

<style lang="stylus">
* {
  margin: 0;
  padding: 0;
}

.Login {
  width: 100%;
  height: 100%;
  background-image: url('../../public/Login/background.png');
  background-size: 100% 100%;
  background-position: top left;
  background-repeat: no-repeat;
}

.logo {
  clear: both;
  display: block;
  margin: auto;
  padding-top: 19vh;
  width: 38.4vw;
  height: 38.4vw;
}

.logo img {
  height: 100%;
  width: 100%;
  object-fit: cover;
}

.classForInputData {
  display: flex;
  flex-direction: column;
  width: 70vw;
  margin: 10vh auto;
  margin-bottom: 0;
  height: 200px;
  justify-content: space-around;
}

.classForInputData input {
  width: 100%;
  height: 22%;
  border: none;
  outline: none;
  background-image: url('../../public/Login/input.png');
  background-size: 100% 100%;
  text-align: center;
  background-color: rgba(0, 0, 0, 0);
  color: rgba(77, 137, 138, 0.7);
}

.loginButton {
  margin: 3vh auto;
  background-image: url('../../public/Login/loginButton.png');
  background-position: 0 0;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  width: 47vw;
  height: 11vw;
}

.agree {
  display: flex;
  justify-content: center;
}

input['type="checkbox"'] {
  outline: none;
  background-color: white;
  border: 5px solid #5bbfbf;
}

label {
  font-size: 1vw;
  color: rgba(77, 137, 138, 0.7);
}
</style>