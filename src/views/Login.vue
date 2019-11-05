<template>
  <div class="Login">
    <div class="classForInputData">
      <input type="text" :value="school" />
      <input type="text" :value="name" />
      <input type="text" :value="studentId" />
    </div>
    <input type="button" value="登录" />
    <input type="checkbox" id="agree" />
    <label for="agree">{{agree}}</label>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";

@Component({
  name: "Login"
})
export default class extends Vue {
  private school: string = "学校";
  private name: string = "姓名";
  private studentId: string = "学号";
  private agree: string = "同意佩佩读取您的个人信息";
}
</script>

<style lang="stylus">
.classForInputData {
  display: flex;
  flex-direction: column;
  width: 70vw;
  margin: 100px auto;
  height: 200px;
  justify-content: space-around;
}

.classForInputData input {
  width: 100%;
  height: 30px;
  border-radius: 10%;
  border: thin soild rgba(77, 137, 138, 0.7);
  outline: none;
}
</style>