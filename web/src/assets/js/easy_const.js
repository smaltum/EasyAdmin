

const schemal = "http://"
const host = schemal + "localhost:8081"

const URL_BASE = host; //配置接口地址

//******************* USED URL ***********************\\

const SERVER_URL = {
    URL_LOGIN: "/api/v1/login", //登录
}


//******************* 菜单项 ***********************\\

const MENU = {
    MENU_ITEM: 1,
    MENU_SUB: 2,
}

export default {
    URL_BASE,
    SERVER_URL,
    MENU
}