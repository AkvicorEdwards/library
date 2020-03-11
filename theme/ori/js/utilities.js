function requireApi(url, elemId, func) {
    let httpRequest = new XMLHttpRequest();
    httpRequest.open('GET', url, true);
    console.log("API: " + url);
    console.log("ElemId: " + elemId);
    console.log("Function name: " + getFnName(func));
    httpRequest.send();
    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === 4 && httpRequest.status === 200) {
            let data = JSON.parse(httpRequest.responseText);
            document.getElementById(elemId).innerHTML = func(data);
        }
    };
}

function uploadArticleGetSelect(content) {
    let div = "";
    for (let i = 0; i < content.length; ++i) {
        div += `<option value="` + HTMLEncode(content[i].id) + `">` + HTMLEncode(content[i].category_name) + `</option>`;
    }
    return div
}

function articleCategoryGetList(content) {
    let div = "";
    for (let i = 0; i < content.length; ++i) {
        div += `<tr><td><a href="` + urlShort + HTMLEncode(content[i].id) + `">` + HTMLEncode(content[i].category_name) + `&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[`+content[i].num_in_category+`]</a></td></tr>`;
    }
    return div
}

function articleGetList(content) {
    let div = "";
    for (let i = 0; i < content.length; ++i) {
        div += `<tr><td><a href="` + urlShort + HTMLEncode(content[i].id) + `">` + HTMLEncode(content[i].title) + `&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[`+content[i].word_count+`]</a></td></tr>`;
    }
    return div
}

function articleGetContent(content) {
    let div = "<h1>" + content.title + "</h1>";
    for (let i = 0; i < content.whole.length; ++i) {
        div += `<p>`+HTMLEncode(content.whole[i]).replace(/ /g, "&nbsp;")+`</p>`;
    }
    return div
}

function getUrlRelativePath() {
    let url = document.location.toString();
    let arrUrl = url.split("//");
    let start = arrUrl[1].indexOf("/");
    let relUrl = arrUrl[1].substring(start);//stop省略，截取从start开始到结尾的所有字符
    if (relUrl.indexOf("?") !== -1) {
        relUrl = relUrl.split("?")[0];
    }
    console.log("getUrlRelativePath: " + relUrl);
    return relUrl;
}

function getFnName(callee) {
    let _callee = callee.toString().replace(/[\s\?]*/g, ""),
        comb = _callee.length >= 50 ? 50 : _callee.length;
    _callee = _callee.substring(0, comb);
    let name = _callee.match(/^function([^\(]+?)\(/);
    if (name && name[1]) {
        return name[1];
    }
    let caller = callee.caller, _caller = caller.toString().replace(/[\s\?]*/g, "");
    let last = _caller.indexOf(_callee), str = _caller.substring(last - 30, last);
    name = str.match(/var([^\=]+?)\=/);
    if (name && name[1]) {
        return name[1];
    }
    return "anonymous"
}

function HTMLEncode(html) {
    let temp = document.createElement("div");
    (temp.textContent != null) ? (temp.textContent = html) : (temp.innerText = html);
    let output = temp.innerHTML;
    temp = null;
    return output;
}

function loadScript(url, callback) {
    let script = document.createElement("script");
    script.type = "text/javascript";
    if (script.readyState) { //IE
        script.onreadystatechange = function () {
            if (script.readyState === "loaded" || script.readyState === "complete") {
                script.onreadystatechange = null;
                callback();
            }
        };
    } else { //Others
        script.onload = function () {
            callback();
        };
    }
    script.src = url;
    document.getElementsByTagName("head")[0].appendChild(script);
}

