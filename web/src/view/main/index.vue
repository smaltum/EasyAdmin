<template>
  <div id="mainLayout">

    <!--  左侧  -->
    <div :class="menuCollapseState?'main-layout-left-close':'main-layout-left-open'">

      <el-menu class="main-layout-left-menu" @open="handleOpen" @close="handleClose"
               :collapse="menuCollapseState">

        <li class="el-menu-item">

          <div :class="menuCollapseState?'main-layout-left-header-close':'main-layout-left-header-open'">
            <img alt="" src="favicon.ico" style="width: 40px;height: 40px;"/>
            <span>EasyAdmin</span>
          </div>

        </li>

        <el-menu-item index="0">
          <i class="el-icon-menu"></i>
          <span slot="title">仪表盘</span>
        </el-menu-item>

        <el-submenu index="1">
          <template slot="title">
            <i class="el-icon-setting"></i>
            <span slot="title">管理</span>
          </template>
          <el-menu-item-group>
            <span slot="title">分组一</span>
            <el-menu-item index="1-1">选项1</el-menu-item>
            <el-menu-item index="1-2">选项2</el-menu-item>
          </el-menu-item-group>
          <el-submenu index="1-4">
            <span slot="title">选项4</span>
            <el-menu-item index="1-4-1">选项1</el-menu-item>
          </el-submenu>
        </el-submenu>

      </el-menu>

    </div>

    <!--  右侧  -->
    <div class="main-layout-right">

      <el-header class="main-layout-right-header">

        <div class="main-layout-right-header-right">
          <i v-model="menuCollapseState" :class="menuCollapseState?'el-icon-d-arrow-left':'el-icon-d-arrow-right'"
             @click="switchCollapse" style="font-size: 18px;color: #333333;margin-left: 14px;">
            &nbsp;{{ menuCurrentLabel }}
          </i>

          <div>

            <el-dropdown trigger="click">

              <i class="el-icon-setting main-layout-right-header-right-i"/>

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
                <el-avatar shape="circle" size="medium" :src="loginUser.userAvtar"/>
                {{ loginUser.userName }}
                <i class="el-icon-arrow-down main-layout-right-header-right-i"/>
               </span>
              <el-dropdown-menu slot="dropdown" style="margin-top: 20px">
                <el-dropdown-item icon="el-icon-plus">个人中心</el-dropdown-item>
                <el-dropdown-item icon="el-icon-circle-check" @click.native="logout">退出</el-dropdown-item>
              </el-dropdown-menu>

            </el-dropdown>
          </div>
        </div>

        <div style="height: 1px;background: #cdcdcd;"/>
      </el-header>

      <!--   功能区   -->
      <div class="main-layout-right-content">
        <router-view/>
      </div>

    </div>

  </div>

</template>

<script>
export default {
  name: "Main",
  data() {
    return {
      // 菜单栏状态
      menuCollapseState: false,
      // 当前页标签
      menuCurrentLabel: "仪表盘",
      // 当前用户
      loginUser: {
        userName: '超级管理员', // 名称
        userType: 1, // 类型
        userAvtar: 'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png', // 头像
      },

    };
  },
  methods: {
    // 登出
    logout() {
      this.$router.push({
        path: "/"
      })
    },
    handleOpen(key, keyPath) {
      console.log(key, keyPath, this.menuCollapseState);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath, this.menuCollapseState);
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
