import Cookies from 'js-cookie'

const setVal = function (key, value) {
  return Cookies.set(key, value)
};

const getVal = function (key) {
  return Cookies.get(key)
};

const removeVal = function (key) {
  return Cookies.remove(key)
};

export default {
  setVal,
  getVal,
  removeVal
}
