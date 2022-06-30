<template>
  <div id="mainLayout">

    <!--  左侧  -->
    <div class="main-layout-left" :class="menuCollapseState ? 'main-layout-left-close' : 'main-layout-left-open'">

      <div class="main-layout-left-header"
        :class="menuCollapseState ? 'main-layout-left-header-close' : 'main-layout-left-header-open'">
        <img alt="" src="favicon.ico"
          :class="menuCollapseState ? 'main-layout-left-header-img-open-above' : 'main-layout-left-header-img-open-below'" />
        <span>EasyAdmin</span>
      </div>

      <el-menu class="main-layout-left-menu" background-color="#253b45" text-color="#eaeaea" @open="handleOpen"
        @close="handleClose" :collapse="menuCollapseState" @select="handleSelect">

        <!-- 菜单项 -->
        <div v-for="(menu, index) in menuList">

          <el-menu-item v-if="menu.type == 1" :index="menu.index">
            <i :class="menu.icon"></i>
            <span slot="title">{{ menu.label }}</span>
          </el-menu-item>

          <el-submenu v-if="menu.type == 2" :index="menu.index">
            <template slot="title">
              <i :class="menu.icon"></i>
              <span slot="title">{{ menu.label }}</span>
            </template>

            <!-- 子菜单 -->
            <el-menu-item-group>
              <el-menu-item v-for="(group, index) in menu.groups" :index="group.index">{{ group.label }}</el-menu-item>
            </el-menu-item-group>

          </el-submenu>
        </div>

      </el-menu>

    </div>

    <!--  右侧  -->
    <div class="main-layout-right">

      <el-header class="main-layout-right-header">

        <div class="main-layout-right-header-right">
          <i :class="menuCollapseState ? 'el-icon-d-arrow-left' : 'el-icon-d-arrow-right'" @click="switchCollapse"
            style="font-size: 18px;color: #333333;margin-left: 14px;">
            &nbsp;{{ menuCurrentLabel }}
          </i>

          <div>

            <el-dropdown trigger="click">

              <i class="el-icon-setting main-layout-right-header-right-i" />

              <el-dropdown-menu slot="dropdown" style="margin-top: 30px">
                <el-dropdown-item icon="el-icon-plus" @click.native="intentToPage('/index/manage/role')">
                  角色管理
                </el-dropdown-item>
                <el-dropdown-item icon="el-icon-plus" @click.native="intentToPage('/index/manage/user')">
                  用户管理
                </el-dropdown-item>
              </el-dropdown-menu>

            </el-dropdown>

            <el-dropdown trigger="click" size="small">

              <span class="el-dropdown-link" style="display: flex;align-items: center">
                <el-avatar shape="circle" size="medium" :src="loginUser.user.avatar" style="margin-right: 10px;" />
                {{ loginUser.user.name }}
                <i class="el-icon-arrow-down main-layout-right-header-right-i" />
              </span>
              <el-dropdown-menu slot="dropdown" style="margin-top: 20px">
                <el-dropdown-item icon="el-icon-plus">个人中心</el-dropdown-item>
                <el-dropdown-item icon="el-icon-circle-check" @click.native="logout">退出</el-dropdown-item>
              </el-dropdown-menu>

            </el-dropdown>
          </div>
        </div>

        <div style="height: 1px;background: #cdcdcd;" />
      </el-header>

      <!--   功能区   -->
      <div class="main-layout-right-content">
        <router-view />
      </div>

    </div>

  </div>

</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "Main",
  data() {
    return {
      // 菜单栏状态
      menuCollapseState: false,
      // 菜单栏
      menuList: [],
      // 当前页标签
      menuCurrentLabel: "仪表盘",
    };
  },
  computed: {
    ...mapGetters({
      loginUser: "getLoginUser"
    })
  },
  mounted() {
    this.menuList = this.loginUser.menu_list
    console.log(this.menuList)
  },
  methods: {
    // 登出
    logout() {
      this.$router.push({
        path: "/"
      })
    },
    handleOpen(key, keyPath) {
      console.log("handleOpen", key, keyPath, this.menuCollapseState);
    },
    handleClose(key, keyPath) {
      console.log("handleClose", key, keyPath, this.menuCollapseState);
    },
    handleSelect(index) {
      console.log(index)
    },
    // 开关menu
    switchCollapse() {
      this.menuCollapseState = !this.menuCollapseState;
    },
    // 跳转路由
    intentToPage(path) {
      this.$router.push({
        path: path
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/style/main.scss";
</style>
