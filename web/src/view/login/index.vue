<template>
  <div id="loginLayout">

    <div class="login-panel">

      <!--   输入框   -->
      <div class="login-panel-form">
        <div class="login-panel-form-title">
          <img class="login-panel-form-title-logo"
            :src="'https://pic1.zhimg.com/80/v2-bde536d2d0ed4eaeb6003e2bda5cdfdb_1440w.png?source=d16d100b'" alt>
          <p class="login-panel-form-title-p">EasyAdmin</p>
        </div>
        <el-form ref="loginFormData" :model="loginReq">

          <!--    用户名      -->
          <el-form-item prop="username">
            <el-input v-model="loginReq.username" placeholder="请输入用户名">
            </el-input>
          </el-form-item>

          <!--    密码      -->
          <el-form-item prop="password">
            <el-input v-model="loginReq.password" :type="'password'" placeholder="请输入密码">
            </el-input>
          </el-form-item>

          <!--    验证码      -->
          <el-form-item prop="captcha">
            <div class="login-panel-form-cap ">
              <el-input v-model="loginReq.captcha" placeholder="请输入验证码" style="width: 60%" />
              <div class="login-panel-form-cap-img">
                <img alt="请输入验证码">
              </div>
            </div>
          </el-form-item>

          <!--    登录按钮      -->
          <el-button type="primary" size="large" style="width: 100%;" @click="login">登&nbsp;&nbsp;录
          </el-button>

        </el-form>
      </div>

      <!--   占位   -->
      <div class="login_panel_right" />

    </div>
  </div>
</template>

<script>

import easyHttp from "../../assets/js/easy_http.js";
import easyConst from "../../assets/js/easy_const.js";
import { mapActions, mapGetters } from "vuex";

export default {
  name: "Login",
  data() {
    return {
      loginReq: {
        username: 'admin',
        password: 'admin',
        captcha: '',
      },
      loginResp: {
        login_time: "",
        user: {
          id: 0,
          type: 0,
          avatar: '',
          name: ''
        }
      }
    };
  },
  methods: {
    ...mapActions(["saveLoginUser"]),
    // 登录
    login() {
      let _this = this;
      easyHttp
        .fetchPost(easyConst.SERVER_URL.URL_LOGIN, this.loginReq)
        .then(res => {
          if (res.state == 0) {
            _this.$message({
              showClose: true,
              message: res.msg,
              type: 'error'
            });
          } else {
            _this.$message({
              showClose: true,
              message: res.msg,
              type: 'success'
            });
            _this.loginResp = res.data;
            _this.$router.push({
              path: '/index'
            })
            _this.saveLoginUser(_this.loginResp);
          }
        })
        .catch(err => {

        });
    }

  }
};
</script>

<style lang="scss" scoped>
@import "@/style/login.scss";
</style>
