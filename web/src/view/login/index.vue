<template>
  <div id="loginLayout">

    <div class="login-panel">

      <!--   占位   -->
      <div class="login_panel_right"/>

      <!--   输入框   -->
      <div class="login-panel-form">
        <div class="login-panel-form-title">
          <img
            class="login-panel-form-title-logo"
            :src="'https://pic1.zhimg.com/80/v2-bde536d2d0ed4eaeb6003e2bda5cdfdb_1440w.png?source=d16d100b'"
            alt>
          <p class="login-panel-form-title-p">EasyAdmin</p>
        </div>
        <el-form
          ref="loginFormData"
          :model="loginData"
        >

          <!--    用户名      -->
          <el-form-item prop="username">
            <el-input
              v-model="loginData.username"
              placeholder="请输入用户名"
            >
            </el-input>
          </el-form-item>

          <!--    密码      -->
          <el-form-item prop="password">
            <el-input
              v-model="loginData.password"
              :type="'password'"
              placeholder="请输入密码"
            >
            </el-input>
          </el-form-item>

          <!--    验证码      -->
          <el-form-item prop="captcha">
            <div class="login-panel-form-cap ">
              <el-input
                v-model="loginData.captcha"
                placeholder="请输入验证码"
                style="width: 60%"
              />
              <div class="login-panel-form-cap-img">
                <img
                  alt="请输入验证码"
                >
              </div>
            </div>
          </el-form-item>

          <!--    登录按钮      -->
          <el-button
            type="primary"
            size="large"
            style="width: 100%;"
            @click="login"
          >登&nbsp;&nbsp;录
          </el-button>

        </el-form>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      loginData: {
        username: 'admin',
        password: 'admin',
        captcha: '',
      }
    };
  },
  methods: {

    // 登录
    login() {
      this.$router.push({
        path: "/index"
      })
    }

  }
};
</script>

<style lang="scss" scoped>
@import "@/style/login.scss";
</style>
