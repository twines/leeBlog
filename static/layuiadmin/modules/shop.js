layui.define(['table', 'form'], function (exports) {
    var $ = layui.$
        , table = layui.table
        , form = layui.form;

    table.render({
        elem: '#shop'
        , url: '/admin/shop/list' //模拟接口
        , cols: [[
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 100, title: 'ID', sort: true}
            , {field: 'shopName', title: '商铺名称', minWidth: 100}
            , {field: 'shopAddress', title: '商铺地址', minWidth: 100}
            , {field: 'mobile', title: '手机'}
            , {field: 'phone', title: '固话'}
            , {field: 'email', title: '邮箱'}
            , {field: 'joinAt', title: '加入时间', sort: true}
            , {title: '操作', width: 150, align: 'center', fixed: 'right', toolbar: '#table-useradmin-webuser'}
        ]]
        , page: true
        , limit: 30
        , height: 'full-220'
        , text: '对不起，加载出现异常！'
    });

    exports('shop', {})
});