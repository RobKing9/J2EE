// var server = "124.223.103.213:9090";
var server = "127.0.0.1:8080";

function LoadLikeCount(pageName) {
    var likeCount = document.getElementById("likeCount");
    //获取当前页面的点赞数
    var httpRequest = new XMLHttpRequest();
    // httpRequest.open('GET', 'http://124.223.103.213:9090/SeeLike?pageName=main', true);
    httpRequest.open('GET', "http://"+server+"/SeeLike?pageName="+pageName, true);
    httpRequest.send();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var res = httpRequest.responseText; //获取到json字符串，还需解析
            likeCount.innerText = res;
        }
    };
}

function onClickLike(pageName){
    var likeDiv = document.getElementById("likeDiv");
    var likeIcon = document.getElementById("likeIcon");

    if (likeDiv.classList.contains("unliked")) {
        likeDiv.classList.remove("unliked");
        likeDiv.classList.add("liked");
        likeIcon.classList.remove("unlikedIcon");
        likeIcon.classList.add("likedIcon");

        likeCount.innerText = parseInt(likeCount.innerText) + 1;

        var httpRequest = new XMLHttpRequest();
        httpRequest.open('GET', "http://"+server+"/like?pageName="+pageName, true);
        httpRequest.send();
    } else {
        likeDiv.classList.remove("liked");
        likeDiv.classList.add("unliked");
        likeIcon.classList.remove("likedIcon");
        likeIcon.classList.add("unlikedIcon");

        likeCount.innerText = parseInt(likeCount.innerText) - 1;

        var httpRequest = new XMLHttpRequest();
        httpRequest.open('GET', "http://"+server+"/unlike?pageName="+pageName, true);
        httpRequest.send();
    }
}