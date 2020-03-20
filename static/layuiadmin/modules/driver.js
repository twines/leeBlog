layui.define(['table', 'form'], function (exports) {
  var $ = layui.$
      , table = layui.table
      , form = layui.form;

  table.render({
    elem: '#driver'
    , url: '/admin/driver/list' //模拟接口
    , cols: [[
      {type: 'checkbox', fixed: 'left'}
      , {field: 'id', width: 100, title: 'ID', sort: true}
      , {field: 'driverName', title: '司机姓名', minWidth: 100}
      , {field: 'driverMobile', title: '司机电话', minWidth: 100}
      , {field: 'driverIdNumber', title: '身份证号', minWidth: 100}
    ]]
    , page: true
    , limit: 30
    , height: 'full-220'
    , text: '对不起，加载出现异常！'
  });

  exports('driver', {})
});