import mock from 'mockjs'
import easyConst from './easy_const.js'

mock.setup({
  timeout: '200-600'
});

const getImage = () => {
  const urlArr = [
    'http://images.gongzuoshouji.cn/teacher/2021-08-09/84751fbcf153487f868a21a77048d19c.jpg',
    'http://images.gongzuoshouji.cn/teacher/20210804/196b0fa5d1726432492ad5edae6a565b.jpg',
    'http://images.gongzuoshouji.cn/teacher/20210730/a8e65bfaf7575a051a45b8dedd528af4.jpg',
  ]
  const index = mock.Random.integer(0, urlArr.length - 1)
  return urlArr[index]
}


// 登录
mock.mock(easyConst.URL_BASE + easyConst.SERVER_URL.URL_LOGIN, {
  "state": 1,
  "msg": "登录成功",
  "data": {
    "last_login_time": "@datetime",
    "user": {
      "id|1000000-9999999": 1000000,
      "type": 1,
      "avatar": getImage(),
      "name": "@cname",
    },
    "menu_list": [
      {
        "type": 1, // 1 el-menu-item 2 el-submenu
        "index": "1",
        "icon": 'el-icon-menu',
        "label": '仪表盘'
      },
      {
        "type": 2,
        "index": "2",
        "icon": 'el-icon-setting',
        "label": '系统管理',
        "groups": [
          {
            "label": '用户',
            "index": "2-1",
            "router": "/index/manage/user"
          },
          {
            "label": '角色',
            "index": "2-2",
            "router": "/index/manage/role"
          },
          {
            "label": '菜单',
            "index": "2-3",
            "router": "/index/manage/menu"
          },
          {
            "label": '日志',
            "index": "2-4",
            "router": "/index/manage/log"
          }
        ]
      }
    ]
  }
});