Email.onchange = function(){
    var Email = this.value;
    var reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
    if(!reg.test(Email)){
        alert("邮箱格式不正确，请重新输入！");
        return false;
    }
}

username.onchange = function(){
    var  username = this.value;
    var reg = /^\S{1,10}$/;
    if(!reg.test( username)){
        alert("用户名长度不能超过10位！");
        return false;
    }
}

