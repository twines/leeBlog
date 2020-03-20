layui.define(['table', 'form'], function (exports) {
    var $ = layui.$
        , table = layui.table
        , form = layui.form;

    table.render({
        elem: '#task'
        , url: '/admin/task/list' //模拟接口
        , cols: [[
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 100, title: 'ID', sort: true}
            , {field: 'driverName', title: '商铺名', minWidth: 100}
            , {field: 'driverMobile', title: '任务描述', minWidth: 100}
            , {field: 'driverMobile', title: '开始时间', minWidth: 100}
            , {field: 'driverMobile', title: '结束时间', minWidth: 100}
            , {field: 'driverIdNumber', title: '状态', minWidth: 100}
        ]]
        , page: true
        , limit: 30
        , height: 'full-220'
        , text: '对不起，加载出现异常！'
    });

    exports('task', {})
});